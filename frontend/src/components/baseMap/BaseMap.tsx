"use client";

import "mapbox-gl/dist/mapbox-gl.css";
import Map, { Marker } from "react-map-gl";
import { MdLocationPin } from "react-icons/md";
import React, { useRef, useEffect, useState } from "react";
import mapboxgl from "mapbox-gl";
import styles from "./BaseMap.module.scss";
import { Card } from "@mantine/core";

const TOKEN = process.env.NEXT_PUBLIC_MAPBOX_TOKEN as string;

interface BaseMapProps {
  data: GeoJSON.FeatureCollection<GeoJSON.LineString>;
}

const BaseMap = (props: BaseMapProps) => {
  const mapContainer = useRef(null);
  const map = useRef<mapboxgl.Map>(null);
  const [lng, setLng] = useState(
    props.data.features[0].geometry.coordinates[0][0],
  );
  const [lat, setLat] = useState(
    props.data.features[0].geometry.coordinates[0][1],
  );
  const [zoom, setZoom] = useState(16);

  useEffect(() => {
    if (map.current) return; // initialize map only once

    // @ts-ignore
    map.current = new mapboxgl.Map({
      accessToken: TOKEN,
      // @ts-ignore
      container: mapContainer.current,
      style: "mapbox://styles/mapbox/streets-v12", // stylesheet location
      // style: "mapbox://styles/mapbox/dark-v10", // stylesheet location
      center: [lng, lat],
      zoom: zoom,
    }) as mapboxgl.Map;

    map.current.addControl(
      new mapboxgl.GeolocateControl({
        positionOptions: {
          enableHighAccuracy: true,
        },
        trackUserLocation: true,
      }),
    );

    // @ts-ignore
    map.current.on("move", () => {
      // @ts-ignore
      setLng(map.current.getCenter().lng.toFixed(4));
      // @ts-ignore
      setLat(map.current.getCenter().lat.toFixed(4));
      // @ts-ignore
      setZoom(map.current.getZoom().toFixed(2));
    });
  }, [zoom, props.data, lat, lng]);

  useEffect(() => {
    //  add layer from geojson
    if (!map.current) return;

    if (map.current.getLayer("route")) {
      // @ts-ignore
      map.current.getSource("route").setData(props.data);
    }
  }, [props.data]);

  useEffect(() => {
    //  add layer from geojson
    if (!map.current) return;

    if (map.current.getLayer("route")) {
      // @ts-ignore
      map.current.getSource("route").setData(props.data);
    }

    // Add marker to the latest location
    if (props.data && props.data.features.length > 0) {
      const latestLocation =
        props.data.features[props.data.features.length - 1];
      const marker = new mapboxgl.Marker()
        .setLngLat([lng, lat])
        .addTo(map.current)
        .setPopup(
          new mapboxgl.Popup({ offset: 25 }).setHTML(
            `<h3>Latest location</h3>`,
          ),
        );
    }
  }, [props.data]);

  return (
    <div className={styles.container}>
      <div className={styles.sidebar}>
        Longitude: {lng} | Latitude: {lat} | Zoom: {zoom}
      </div>

      <Card shadow="sm" padding={0} radius={"md"} withBorder>
        <div
          ref={mapContainer}
          style={{
            height: "500px",
            width: "700px",
          }}
        ></div>
      </Card>
    </div>
  );
};

export default BaseMap;
