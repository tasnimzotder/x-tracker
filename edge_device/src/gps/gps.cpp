//
// Created by Tasnim Zotder on 07/11/2023.
//

#include "gps.h"

GPS::GPS() = default;

void GPS::setup(uint8_t pin_txd2, uint8_t pin_rxd2, int baud_rate) {
    _pin_txd2 = pin_txd2;
    _pin_rxd2 = pin_rxd2;
    _baud_rate = baud_rate;

    //    Serial.println("Setting up GPS");
    ESP_LOGI(TAG, "Setting up GPS");

    Serial2.begin(
        this->_baud_rate,
        SERIAL_8N1,
        this->_pin_rxd2,
        this->_pin_txd2);
    this->_gps = TinyGPSPlus();

    //    Serial.println("GPS setup complete");
    ESP_LOGI(TAG, "GPS setup complete");

    //    delay(3000);
    vTaskDelay(3000 / portTICK_PERIOD_MS);
}

GPSData GPS::getGPSData() {
    while (Serial2.available() > 0) {
        //        if (this->_gps.encode(Serial2.read())) {
        //            ESP_LOGI("GPS", "GPS data received");
        //        }
        _gps.encode(Serial2.read());

        if (this->_gps.location.isUpdated()) {
            if (_gpsData.time != this->_gps.time.value()) {
                _gpsData.currentLocation = true;
            } else {
                _gpsData.currentLocation = false;
            }

            _gpsData.latitude = this->_gps.location.lat();
            _gpsData.longitude = this->_gps.location.lng();

            _gpsData.date = this->_gps.date.value();
            _gpsData.time = this->_gps.time.value();

            return this->_gpsData;
        }

        _gpsData.currentLocation = false;

        //        ESP_LOGW(TAG, "GPS data not received");

        //        break;
    }

    return this->_gpsData;
}
