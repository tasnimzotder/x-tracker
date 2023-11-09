package api

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"github.com/gin-gonic/gin"
	"math/rand"
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
	queryPtr := fmt.Sprintf(`SELECT DISTINCT deviceID, latitude, longitude, time FROM "xtrackerDB".xtracker_table ORDER BY time DESC LIMIT %d`, req.Limit)

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

	type getLastLocationsResponse struct {
		DeviceID string  `json:"device_id"`
		Lat      float32 `json:"lat"`
		Lon      float32 `json:"lon"`
		Time     string  `json:"time"`
	}

	var response []getLastLocationsResponse

	for _, row := range data {
		response = append(response, getLastLocationsResponse{
			DeviceID: *row.Data[0].ScalarValue,
			//Lat:      *row.Data[1].ScalarValue,
			//Lon:      *row.Data[2].ScalarValue,
			Lat:  addRandomValues(*row.Data[1].ScalarValue),
			Lon:  addRandomValues(*row.Data[2].ScalarValue),
			Time: *row.Data[3].ScalarValue,
		})

	}

	ctx.JSON(http.StatusOK, response)
}

func addRandomValues(initialValue string) float32 {
	//	convert to float
	val, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		panic(err)
	}

	//	add random value
	val += (float64(1) - float64(2)*rand.Float64()) / float64(1000)

	//	convert back to string
	//return strconv.FormatFloat(val, 'f', 6, 64)
	//return initialValue

	return float32(val)
}
