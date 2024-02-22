package api

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/tasnimzotder/x-tracker/utils"
	"log"
	"net/http"
)

func (s *Server) sseLocationUpdates(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Important for CORS, adjust if needed

	DeviceID := ctx.Param("device_id")
	if DeviceID == "" {
		err := ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("DeviceID is required"))
		if err != nil {
			log.Printf("Error aborting streaming: %v", err)
		}

		return
	}

	// check if device exists
	_, err := s.Queries.GetDevice(ctx, int64(utils.ConvStrToFloat(DeviceID)))
	if err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("Device not found"))
		if err != nil {
			log.Printf("Error aborting streaming: %v", err)
		}

		return
	}

	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		err := ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("streaming unsupported"))
		if err != nil {
			log.Printf("Error aborting streaming: %v", err)
			return
		}

		return
	}

	// kafka
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf("Error creating Kafka consumer: %v", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("location_updates", 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Error creating Kafka partition consumer: %v", err)
		return
	}
	defer partitionConsumer.Close()

	// Logic to get updates (replace with your Kafka consumer or data source)
	for {
		select {
		case <-ctx.Request.Context().Done():
			log.Printf("Client disconnected from SSE")
			return
		case msg := <-partitionConsumer.Messages():
			if string(msg.Key) != DeviceID {
				continue
			}

			locationData := fmt.Sprintf("%s", string(msg.Value))

			//log.Printf("Sending SSE: %s", msg.Key)
			_, err := fmt.Fprintf(ctx.Writer, "data: %s\n\n", locationData)
			if err != nil {
				log.Printf("Error writing to stream: %v", err)
				return
			}

			flusher.Flush()
		}
	}
}
