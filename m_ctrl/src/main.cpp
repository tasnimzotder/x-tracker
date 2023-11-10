//
// Created by Tasnim Zotder on 07/11/2023.
//

#include <Arduino.h>
#include "gps/gps.h"
#include "comm/comm.h"
#include "esp_log.h"

#define TXD2 17
#define RXD2 16

#define PANIC_BUTTON 4

int GPSBaud = 9600;

int publishDelay = 1000 * 10;
int64_t lastPublish = 0;

GPS gps;
Communications comm;

const char *TAG = "MAIN";

bool panicButtonPressed = false;

// function prototypes
void displayInfo(GPSData gpsData);

// esp task
void task1(void *pvParameters) {
    while (true) {
//        read panic button (pull-up)
        int panicButtonState = digitalRead(PANIC_BUTTON);

        if (panicButtonState == LOW && !panicButtonPressed) {
            panicButtonPressed = true;
            ESP_LOGI(TAG, "Panic button pressed");

            DeviceActivity deviceActivity = {
                    .deviceID = (int)AWS_IOT_THING_NAME,
                    .fallDetected = false,
                    .panicButtonPressed = true
            };

            esp_err_t err = comm.publishMQTTActivity(deviceActivity);
            if (err != ESP_OK) {
                ESP_LOGE(TAG, "Error publishing MQTT activity");
                return;
            }

            ESP_LOGI(TAG, "MQTT activity published");
        } else if (panicButtonState == HIGH && panicButtonPressed) {
            panicButtonPressed = false;
            ESP_LOGI(TAG, "Panic button released");
        }


        vTaskDelay(
                100 / portTICK_PERIOD_MS); // wait for one second (portTICK_PERIOD_MS is a constant defined by FreeRTOS
    }
}

void task2(void *pvParameters) {
    while (true) {
        int64_t currentMillis = millis();

        if (currentMillis - lastPublish < publishDelay) {
//            return;
            continue;
        }

        GPSData gpsData = gps.getGPSData();

        if (gpsData.currentLocation) {
            displayInfo(gpsData);

            // upload data to AWS IoT
            MQTTUploadData mqttUploadData = {
                    .deviceID = AWS_IOT_THING_NAME,
                    .fallDetected = false,
                    .panicButtonPressed = panicButtonPressed,
                    .batteryLow = false,
                    .isSafe = true,
                    .captureTime = gpsData.time,
                    .captureDate = gpsData.date,
                    .coordinates = {
                            .latitude = gpsData.latitude,
                            .longitude = gpsData.longitude
                    }
            };

            esp_err_t err = comm.publishMQTTData(mqttUploadData);
            if (err != ESP_OK) {
                ESP_LOGE(TAG, "Error publishing MQTT data");
//                return;
                continue;
            }

            ESP_LOGI(TAG, "MQTT data published");

            lastPublish = currentMillis;
        }

        vTaskDelay(
                1000 / portTICK_PERIOD_MS); // wait for one second (portTICK_PERIOD_MS is a constant defined by FreeRTOS
    }
}


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

    //     set up panic button
    pinMode(PANIC_BUTTON, INPUT_PULLUP);

    //    create task
    xTaskCreate(
            task1,   /* Task function. */
            "Task1",     /* String with name of task. */
            4096,            /* Stack size in bytes. */
            nullptr,             /* Parameter passed as input of the task */
            1,                /* Priority of the task. */
            nullptr);            /* Task handle. */

    xTaskCreate(
            task2,   /* Task function. */
            "Task2",     /* String with name of task. */
            4096,            /* Stack size in bytes. */
            nullptr,             /* Parameter passed as input of the task */
            1,                /* Priority of the task. */
            nullptr);            /* Task handle. */

    delay(5000);
}


void loop() {


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