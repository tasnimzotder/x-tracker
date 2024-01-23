//
// Created by Tasnim Zotder on 07/11/2023.
//

#ifndef X_TRACKER_ALPHA_GPS_H
#define X_TRACKER_ALPHA_GPS_H

#include <Arduino.h>
#include <HardwareSerial.h>
#include <TinyGPSPlus.h>

#include "esp_log.h"

/**
 * @brief GPS data structure
 * @details This structure contains the GPS data
 * @param latitude - latitude in double
 * @param longitude - longitude in double
 * @param currentLocation - true if the GPS data is updated
 * @param date - date in raw format
 * @param time - time in raw format
 */
struct GPSData {
    double latitude;
    double longitude;
    bool currentLocation;
    //    double altitude;
    //    int satellites;
    uint32_t date;
    uint32_t time;
};

class GPS {
private:
    uint8_t _pin_txd2{};
    uint8_t _pin_rxd2{};
    int _baud_rate{};

    TinyGPSPlus _gps;
    GPSData _gpsData{
        .latitude = 0,
        .longitude = 0,
        .currentLocation = false,
        .date = 0,
        .time = 0};

    // define tag for logging
    const char *TAG = "GPS";

public:
    GPS();

    /**
     * @brief Setup the GPS module
     * @param pin_txd2
     * @param pin_rxd2
     * @param baud_rate
     */
    void setup(uint8_t pin_txd2, uint8_t pin_rxd2, int baud_rate);

    /**
     * @brief Get the GPS data
     * @return GPSData
     */
    GPSData getGPSData();
};

#endif  // X_TRACKER_ALPHA_GPS_H
