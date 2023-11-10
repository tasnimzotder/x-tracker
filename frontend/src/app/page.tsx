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
  const [locations, setLocations] = useState<
    GeoJSON.FeatureCollection<GeoJSON.LineString>
  >({
    type: "FeatureCollection",
    features: [
      {
        type: "Feature",
        properties: {
          time: "",
          device_id: "",
        },
        geometry: {
          type: "LineString",
          coordinates: [],
        },
      },
    ],
  });

  const handleClick = async () => {
    const res = await getLastLocations("test", 1);
    setLocations(res);
  };

  return (
    <main className={styles.main}>
      <div>xTracker</div>

      <button onClick={handleClick}>Click me</button>

      {locations.features[0].geometry.coordinates.length > 0 && (
        <div>
          <BaseMap data={locations} />
        </div>
      )}
    </main>
  );
}
