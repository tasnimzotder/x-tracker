//
// Created by Tasnim Zotder on 07/11/2023.
//

#include <Arduino.h>
#include "gps/gps.h"
#include "comm/comm.h"
#include "esp_log.h"

#define TXD2 17
#define RXD2 16

int GPSBaud = 9600;

int publishDelay = 5000;
unsigned long lastPublish = 0;

GPS gps;
Communications comm;

const char *TAG = "MAIN";


void setup() {
    esp_err_t err;

    Serial.begin(115200);
    Serial.println("Starting up");

    Serial2.begin(GPSBaud, SERIAL_8N1, RXD2, TXD2);

    gps.setup(TXD2, RXD2, GPSBaud);
    comm.setup();

    err = comm.connectToWiFi();
    if (err != ESP_OK) {
        ESP_LOGE(TAG, "Error connecting to WiFi");
        return;
    }

    err = comm.connectToMQTT();
    if (err != ESP_OK) {
        ESP_LOGE(TAG, "Error connecting to MQTT");
        return;
    }

    Serial.println("x-tracker-alpha");

    delay(5000);
}

// function prototypes
void displayInfo(GPSData gpsData);


void loop() {
    unsigned long currentMillis = millis();

    if (currentMillis - lastPublish >= publishDelay) {
        lastPublish = currentMillis;
    } else {
        return;
    }

    GPSData gpsData = gps.getGPSData();

    if (gpsData.currentLocation) {
        displayInfo(gpsData);

        // upload data to AWS IoT
        MQTTUploadData mqttUploadData = {
                .latitude = gpsData.latitude,
                .longitude = gpsData.longitude,
                .date = gpsData.date,
                .time = gpsData.time,
                .deviceID = "x-tracker-alpha"
        };

        esp_err_t err = comm.publishMQTTData(mqttUploadData);
        if (err != ESP_OK) {
            ESP_LOGE(TAG, "Error publishing MQTT data");
            return;
        } else {
            ESP_LOGI(TAG, "MQTT data published");
        }
    }

    delay(100);
}


void displayInfo(GPSData gpsData) {
    Serial.print("Latitude: ");
    Serial.println(gpsData.latitude, 6);
    Serial.print("Longitude: ");
    Serial.println(gpsData.longitude, 6);

    uint32_t rawTime = gpsData.time;

    uint32_t rawTimeHour = rawTime / 1000000;
    uint32_t rawTimeMinute = (rawTime % 1000000) / 10000;
    uint32_t rawTimeSecond = (rawTime % 10000) / 100;

    Serial.print("Time: ");
    Serial.print(rawTimeHour);
    Serial.print(":");
    Serial.print(rawTimeMinute);
    Serial.print(":");
    Serial.println(rawTimeSecond);
}