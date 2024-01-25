package api

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	"github.com/gin-gonic/gin"
	"github.com/tasnimzotder/x-tracker/interfaces"
	"github.com/tasnimzotder/x-tracker/utils"
	"net/http"
	"strconv"
)

type getLastLocationsRequest struct {
	DeviceID int64 `json:"device_id" binding:"required"`
	Limit    int   `json:"limit" binding:"required"`
}

func (s *Server) getLastLocations(ctx *gin.Context) {
	var req getLastLocationsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// get timestream data
	querySvc := timestreamquery.NewFromConfig(s.AWS_Config)

	//queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table WHERE time > ago(%dh) ORDER BY time DESC`, req.Limit)
	//queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table ORDER BY time DESC LIMIT %d`, req.Limit)
	queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table WHERE deviceID = '%s' ORDER BY time DESC LIMIT %d`, strconv.FormatInt(req.DeviceID, 10), req.Limit)

	queryInput := &timestreamquery.QueryInput{
		QueryString: &queryPtr,
	}

	queryOutput, err := querySvc.Query(context.Background(), queryInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	data := queryOutput.Rows

	type GeoJSONFeatureCollection struct {
		Type     string                                `json:"type"`
		Features []interfaces.GeoJSONFeatureLineString `json:"features"`
	}

	var geoJSON GeoJSONFeatureCollection
	geoJSON.Type = "FeatureCollection"

	var feature interfaces.GeoJSONFeatureLineString
	feature.Type = "Feature"
	feature.Geometry.Type = "LineString"

	for _, row := range data {
		coordinates := []float32{
			utils.ConvStrToFloat(*row.Data[2].ScalarValue),
			utils.ConvStrToFloat(*row.Data[1].ScalarValue),
		}

		feature.Geometry.Coordinates = append(
			feature.Geometry.Coordinates, coordinates,
		)
		feature.Properties.DeviceID = *row.Data[0].ScalarValue
		feature.Properties.Time = append(feature.Properties.Time, *row.Data[3].ScalarValue)
	}

	geoJSON.Features = append(geoJSON.Features, feature)
	ctx.JSON(http.StatusOK, geoJSON)
}
