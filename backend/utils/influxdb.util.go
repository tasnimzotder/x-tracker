package utils

//func WriteLocationDataToInfluxDB(
//	s *api.Server,
//	// payload []byte,
//	data models.Location,
//) {
//	org := constants.INFLUX_DB_ORG
//	bucket := constants.INFLUX_DB_BUCKET
//	writeAPI := s.InfluxdbClient.WriteAPIBlocking(org, bucket)
//
//	//type Data struct {
//	//	DeviceID      int     `json:"device_id"`
//	//	ClientID      string  `json:"client_id"`
//	//	Timestamp     int64   `json:"timestamp"`
//	//	Lat           float64 `json:"lat"`
//	//	Lng           float64 `json:"lng"`
//	//	BatteryStatus int     `json:"battery_status"`
//	//}
//
//	//var data models.Location
//
//	//err := json.Unmarshal(payload, &data)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	deviceID := fmt.Sprintf("%d", data.DeviceID)
//
//	// check if the location falls within a geofence
//	geofence, err := s.Queries.GetGeofencesByDevice(context.Background(), int64(data.DeviceID))
//	if err != nil {
//		log.Printf("Error: %v", err)
//	}
//
//	for _, fence := range geofence {
//
//		fenceData := Geofence{
//			CenterLat:  fence.CenterLat,
//			CenterLong: fence.CenterLong,
//			Radius:     fence.Radius,
//			Rule:       fence.Rule,
//		}
//
//		// log.Printf("Geofence: %v", fence)
//		status, distance := IsLocationInGeofence(
//			data.Lat,
//			data.Long,
//			fenceData,
//		)
//
//		if status {
//			log.Printf("Device %v is within the geofence: %v with distance: %v", deviceID, fence.GeofenceName, distance)
//		}
//	}
//
//	// write data to influxdb
//	tags := map[string]string{
//		"device_id": deviceID,
//	}
//
//	fields := map[string]interface{}{
//		"lat":            data.Lat,
//		"long":           data.Long,
//		"battery_status": data.BatteryStatus,
//		// "timestamp":      time.UnixMilli(data.Timestamp),
//	}
//
//	point := write.NewPoint("edge_location", tags, fields, time.UnixMilli(data.Timestamp))
//	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
//		log.Fatal(err)
//	}
//
//	//	todo: remove this
//	// send msg to kafka
//
//	//topic := "location"
//	//key := deviceID
//	//value := payload
//	//
//	//err = utils.SendKafkaMessage(s.kafkaProducer, topic, key, string(value))
//	//if err != nil {
//	//	log.Printf("Error: %v", err)
//	//}
//}
