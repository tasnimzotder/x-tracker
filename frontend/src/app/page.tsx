"use client";

import { useState } from "react";
import styles from "./page.module.scss";
import {
  getLastLocations,
  locationService_t,
} from "@/services/location.service";
import { formatTimeToLocaleString } from "@/utils/time.util";
import BaseMap from "@/components/baseMap/BaseMap";

export default function Home() {
  const [locations, setLocations] = useState<locationService_t[]>([
    {
      lat: 28.613,
      lon: 77.2295,
      time: "",
      device_id: "0",
    },
  ]);

  const handleClick = async () => {
    const res = await getLastLocations("test", 10);
    setLocations(res);
  };

  return (
    <main className={styles.main}>
      <div>xTracker</div>

      <button onClick={handleClick}>Click me</button>

      {locations.length > 0 && (
        <div>
          <div>
            {locations[0]?.time !== "" && (
              <span>
                Latest location: {formatTimeToLocaleString(locations[0]?.time)}
              </span>
            )}
            <br />
            <span>Latitude: {locations[0]?.lat}</span> <br />
            <span>Longitude: {locations[0]?.lon}</span> <br />
            <span>Device ID: {locations[0]?.device_id}</span> <br />
          </div>

          <BaseMap lat={locations[0]?.lat} lon={locations[0]?.lon} />
        </div>
      )}
    </main>
  );
}
