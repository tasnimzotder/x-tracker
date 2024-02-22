package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/tasnimzotder/x-tracker/constants"
	"github.com/tasnimzotder/x-tracker/models"
	"github.com/tasnimzotder/x-tracker/utils"
)

type getLastLocationsRequest struct {
	DeviceID int64 `json:"device_id" binding:"required"`
	Limit    int   `json:"limit" binding:"required"`
}

func GetLocationsByDeviceID(s *Server, DeviceID int64) ([]models.Location, error) {
	var locations []models.Location

	// todo: get locations from influxdb

	// sort locations by timestamp
	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Timestamp > locations[j].Timestamp
	})

	return locations, nil
}

func (s *Server) WriteDataToInfluxDB(
	payload []byte,
) {
	org := constants.INFLUX_DB_ORG
	bucket := constants.INFLUX_DB_BUCKET
	writeAPI := s.InfluxdbClient.WriteAPIBlocking(org, bucket)

	//type Data struct {
	//	DeviceID      int     `json:"device_id"`
	//	ClientID      string  `json:"client_id"`
	//	Timestamp     int64   `json:"timestamp"`
	//	Lat           float64 `json:"lat"`
	//	Lng           float64 `json:"lng"`
	//	BatteryStatus int     `json:"battery_status"`
	//}

	var data models.Location

	err := json.Unmarshal(payload, &data)
	if err != nil {
		log.Fatal(err)
	}

	deviceID := fmt.Sprintf("%d", data.DeviceID)

	// check if the location falls within a geofence
	geofence, err := s.Queries.GetGeofencesByDevice(context.Background(), int64(data.DeviceID))
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
			data.Long,
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
		"long":           data.Long,
		"battery_status": data.BatteryStatus,
		// "timestamp":      time.UnixMilli(data.Timestamp),
	}

	point := write.NewPoint("edge_location", tags, fields, time.UnixMilli(data.Timestamp))
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}

	//	todo: remove this
	// send msg to kafka
	//
	//topic := "location"
	//key := deviceID
	//value := payload
	//
	//err = utils.SendKafkaMessage(s.kafkaProducer, topic, key, string(value))
	//if err != nil {
	//	log.Printf("Error: %v", err)
	//}
}
