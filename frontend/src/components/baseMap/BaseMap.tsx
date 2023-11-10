"use client";

import "mapbox-gl/dist/mapbox-gl.css";
import Map, { Marker } from "react-map-gl";
import { MdLocationPin } from "react-icons/md";
import React, { useRef, useEffect, useState } from "react";
import mapboxgl from "mapbox-gl";
import styles from "./BaseMap.module.scss";
import { Card } from "@mantine/core";

const TOKEN: string =
  "pk.eyJ1IjoidGFzbmltem90ZGVyIiwiYSI6ImNsb3BrNnk1eTBiYngyanM4dHZqZXA4MWMifQ.Gvrhme69_5PTIp3zGd85TQ";

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

    map.current.on("load", () => {
      // @ts-ignore
      map.current.addSource("route", {
        type: "geojson",
        data: props.data,
      });

      // @ts-ignore
      map.current.addLayer({
        id: "route",
        type: "line",
        source: "route",
        layout: {
          "line-join": "round",
          "line-cap": "round",
        },
        paint: {
          "line-color": "#888",
          "line-width": 2,
        },
      });
    });
  }, []);

  useEffect(() => {
    //  add layer from geojson
    if (!map.current) return;

    if (map.current.getLayer("route")) {
      // @ts-ignore
      map.current.getSource("route").setData(props.data);
    }
  }, [props.data]);

  let marker: mapboxgl.Marker;

  // useEffect(() => {
  //   // @ts-ignore
  //   map.current.setCenter([props.lon, props.lat]);

  //   setLng(props.lon);
  //   setLat(props.lat);
  //   setZoom(16);
  // }, [props.lat, props.lon]);

  // useEffect(() => {
  //   // @ts-ignore
  //   map.current.setCenter([lng, lat]);
  // }, [lat, lng, zoom]);

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
