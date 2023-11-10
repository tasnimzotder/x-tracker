package api

import (
	"backend/interfaces"
	"fmt"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getLastLocationsRequest struct {
	DeviceID string `json:"device_id" binding:"required"`
	Limit    int32  `json:"limit" binding:"required,min=1"`
}

func (s *Server) getLastLocations(ctx *gin.Context) {
	var req getLastLocationsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// get timestream data
	querySvc := timestreamquery.New(s.Session)

	//queryPtr := "SELECT * FROM \"gps\" WHERE \"device_id\" = '" + req.DeviceID + "' ORDER BY time DESC LIMIT " + string(req.Limit)
	//queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table ORDER BY time DESC LIMIT %d`, req.Limit)
	queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table WHERE time > ago(%dh) ORDER BY time DESC`, req.Limit)

	queryInput := &timestreamquery.QueryInput{
		QueryString: &queryPtr,
	}

	queryOutput, err := querySvc.Query(queryInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return the query result
	data := queryOutput.Rows

	// geojson
	//var geoJSON GeoJSONFeatureCollection
	//
	//geoJSON.Type = "FeatureCollection"
	//
	//for _, row := range data {
	//	var feature GeoJSONFeature
	//	feature.Type = "Feature"
	//	feature.Geometry.Type = "Point"
	//	feature.Geometry.Coordinates = []float32{addRandomValues(*row.Data[2].ScalarValue), addRandomValues(*row.Data[1].ScalarValue)}
	//	feature.Properties.DeviceID = *row.Data[0].ScalarValue
	//	feature.Properties.Time = *row.Data[3].ScalarValue
	//
	//	geoJSON.Features = append(geoJSON.Features, feature)
	//}

	type GeoJSONFeatureCollection struct {
		Type     string                                `json:"type"`
		Features []interfaces.GeoJSONFeatureLineString `json:"features"`
	}

	// return a line geojson
	var geoJSON GeoJSONFeatureCollection

	geoJSON.Type = "FeatureCollection"

	var feature interfaces.GeoJSONFeatureLineString
	feature.Type = "Feature"
	feature.Geometry.Type = "LineString"

	for _, row := range data {
		coordinate := []float32{convStrToFloat(*row.Data[2].ScalarValue), convStrToFloat(*row.Data[1].ScalarValue)}
		feature.Geometry.Coordinates = append(feature.Geometry.Coordinates, coordinate)

		feature.Properties.DeviceID = *row.Data[0].ScalarValue
		feature.Properties.Time = append(feature.Properties.Time, *row.Data[3].ScalarValue)
	}

	geoJSON.Features = append(geoJSON.Features, feature)

	ctx.JSON(http.StatusOK, geoJSON)
}

type GetLastLocationResponseType struct {
	DeviceID  string `json:"device_id"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Time      string `json:"time"`
}

// Todo: improve this function
func (s *Server) GetLastLocation() (GetLastLocationResponseType, error) {
	var rsp GetLastLocationResponseType

	// get timestream data
	querySvc := timestreamquery.New(s.Session)
	//queryPtr := "SELECT * FROM \"gps\" WHERE \"device_id\" = '" + req.DeviceID + "' ORDER BY time DESC LIMIT " + string(req.Limit)
	//queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table ORDER BY time DESC LIMIT %d`, req.Limit)
	queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table ORDER BY time DESC LIMIT 1`)

	queryInput := &timestreamquery.QueryInput{
		QueryString: &queryPtr,
	}

	queryOutput, err := querySvc.Query(queryInput)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return rsp, err
	}

	// return the query result
	data := queryOutput.Rows

	rsp.DeviceID = *data[0].Data[0].ScalarValue
	rsp.Longitude = *data[0].Data[2].ScalarValue
	rsp.Latitude = *data[0].Data[1].ScalarValue

	//ctx.JSON(http.StatusOK, rsp)
	return rsp, nil
}

func convStrToFloat(value string) float32 {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return float32(val)
}
