package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tasnimzotder/x-tracker/constants"
	"log"
	"sort"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/tasnimzotder/x-tracker/utils"
)

type getLastLocationsRequest struct {
	DeviceID int64 `json:"device_id" binding:"required"`
	Limit    int   `json:"limit" binding:"required"`
}

// note: keep the struct keys as snake case for proper json marshalling for dynamodb
type Location struct {
	Device_ID int64 `json:"device_id"`
	//Client_ID           string `json:"client_id"`
	Lat                 string `json:"lat"`
	Lng                 string `json:"lng"`
	Timestamp           int64  `json:"timestamp"`
	Processed_Timestamp int64  `json:"processed_timestamp"`
	Battery_Status      int    `json:"battery_status"`
}

func GetLocationsByDeviceID(s *Server, DeviceID int64) ([]Location, error) {
	var locations []Location

	// get location from influxdb
	//queryAPI := s.influxdbClient.QueryAPI(constants.INFLUX_DB_ORG)
	//query := `from(bucket: "locations")
	//			|> range(start: -10000m`

	// currTimestampMilli := utils.GetCurrentTimeMilli()
	// prevTimestampMilli := currTimestampMilli - (60 * 60 * 1000)

	// log.Println(prevTimestampMilli)

	// keyEx := expression.Key("device_id").Equal(expression.Value(DeviceID)).And(expression.Key("timestamp").GreaterThanEqual(expression.Value(prevTimestampMilli)))

	// expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()

	// if err != nil {
	// 	return nil, err
	// } else {
	// 	queryPaginator := dynamodb.NewQueryPaginator(s.DynamoDB.DynamoDBClient, &dynamodb.QueryInput{
	// 		TableName: &s.DynamoDB.TableName,
	// 		// todo: use index
	// 		IndexName:                 aws.String("device_id-timestamp-index"),
	// 		ExpressionAttributeNames:  expr.Names(),
	// 		ExpressionAttributeValues: expr.Values(),
	// 		KeyConditionExpression:    expr.KeyCondition(),
	// 		// Limit:                     aws.Int32(1),
	// 	})

	// 	for queryPaginator.HasMorePages() {
	// 		response, err = queryPaginator.NextPage(context.TODO())
	// 		if err != nil {
	// 			log.Printf("Error: %v", err)
	// 			return nil, err
	// 		} else {
	// 			var locationPage []Location

	// 			err = attributevalue.UnmarshalListOfMaps(response.Items, &locationPage)
	// 			if err != nil {
	// 				return nil, err
	// 			} else {
	// 				locations = append(locations, locationPage...)
	// 			}
	// 		}
	// 	}
	// }

	// sort locations by timestamp
	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Timestamp > locations[j].Timestamp
	})

	return locations, nil
}

func (s *Server) WriteDataToInfluxDB(
	topic string,
	payload []byte,
) {
	org := constants.INFLUX_DB_ORG
	bucket := constants.INFLUX_DB_BUCKET
	writeAPI := s.influxdbClient.WriteAPIBlocking(org, bucket)

	type Data struct {
		DeviceID      int     `json:"device_id"`
		ClientID      string  `json:"client_id"`
		Timestamp     int64   `json:"timestamp"`
		Lat           float64 `json:"lat"`
		Lng           float64 `json:"lng"`
		BatteryStatus int     `json:"battery_status"`
	}

	var data Data

	err := json.Unmarshal(payload, &data)
	if err != nil {
		log.Fatal(err)
	}

	deviceID := fmt.Sprintf("%d", data.DeviceID)

	// check if the location falls within a geofence
	geofence, err := s.queries.GetGeofencesByDevice(context.Background(), int64(data.DeviceID))
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, fence := range geofence {

		fenceData := utils.Geofence{
			CenterLat:  fence.CenterLat,
			CenterLong: fence.CenterLong,
			Radius:     fence.Radius,
			Rule:       fence.Rule,
		}

		// log.Printf("Geofence: %v", fence)
		status, distance := utils.IsLocationInGeofence(
			data.Lat,
			data.Lng,
			fenceData,
		)

		if status {
			log.Printf("Device %v is within the geofence: %v with distance: %v", deviceID, fence.GeofenceName, distance)
		}
	}

	// write data to influxdb
	tags := map[string]string{
		"device_id": deviceID,
	}

	fields := map[string]interface{}{
		"lat":            data.Lat,
		"lng":            data.Lng,
		"battery_status": data.BatteryStatus,
		// "timestamp":      time.UnixMilli(data.Timestamp),
	}

	point := write.NewPoint("edge_location", tags, fields, time.UnixMilli(data.Timestamp))
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}
