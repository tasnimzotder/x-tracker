//
// Created by Tasnim Zotder on 07/11/2023.
//

#ifndef X_TRACKER_ALPHA_COMM_H
#define X_TRACKER_ALPHA_COMM_H

#include <Arduino.h>
#include <ArduinoJson.h>
#include <MQTTClient.h>
#include <WiFi.h>
#include <WiFiClientSecure.h>

#include "esp_log.h"
#include "secrets.h"

struct Coordinates {
    double latitude;
    double longitude;
};

struct MQTTUploadData {
    const char *deviceID;
    bool fallDetected;
    bool panicButtonPressed;
    bool batteryLow;
    bool isSafe;
    uint32_t captureTime;
    uint32_t captureDate;
    Coordinates coordinates;
};

struct DeviceActivity {
    int deviceID;
    bool fallDetected;
    bool panicButtonPressed;
};

class Communications {
private:
    char _ssid[32]{};
    char _password[32]{};
    char _server[32]{};
    int _port = 8883;

    uint32_t _deviceID{};

    WiFiClientSecure _net = WiFiClientSecure();
    MQTTClient _mqttClient = MQTTClient(256);

    const char *TAG = "COMM";

public:
    Communications();

    void setup();

    //    void setDeviceID(uint32_t deviceID);

    esp_err_t connectToWiFi();

    esp_err_t connectToMQTT();

    esp_err_t publishMQTTData(MQTTUploadData mqttUploadData);

    esp_err_t publishMQTTActivity(DeviceActivity deviceActivity);

    static void messageHandler(String &topic, String &payload);
};

#endif  // X_TRACKER_ALPHA_COMM_H
