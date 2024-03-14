"use client";

// import Map from "react-map-gl";

import React from "react";
import "mapbox-gl/dist/mapbox-gl.css";
import mapboxgl from "mapbox-gl";
import DevicesMap from "@/components/maps/DevicesMap";

const TOKEN = process.env.NEXT_PUBLIC_MAPBOX_TOKEN as string;
mapboxgl.accessToken = TOKEN;

export default function Home() {
  return (
    <div>
      <h1>Map</h1>
      <DevicesMap />
    </div>
  );
}
