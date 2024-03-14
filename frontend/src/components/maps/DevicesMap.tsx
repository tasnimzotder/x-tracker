import React from "react";
import { useAuth } from "@/contexts/authContext";
import Map, { Marker, NavigationControl } from "react-map-gl";
import mapboxgl from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";
import DevicePinMap from "@/components/maps/DevicePinMap";

const TOKEN = process.env.NEXT_PUBLIC_MAPBOX_TOKEN as string;
mapboxgl.accessToken = TOKEN;

const DevicesMap = () => {
  const [viewport, setViewport] = React.useState({
    latitude: 28.679,
    longitude: 77.0697,
    zoom: 10,
  });
  const [marker, setMarker] = React.useState({
    latitude: 28.679,
    longitude: 77.0697,
  });

  const { recentLocation } = useAuth();

  React.useEffect(() => {
    if (recentLocation) {
      setViewport({
        latitude: recentLocation.lat,
        longitude: recentLocation.long,
        zoom: 10,
      });

      setMarker({
        latitude: recentLocation.lat,
        longitude: recentLocation.long,
      });
    }
  }, [recentLocation]);

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        // height: "100vh",
      }}
    >
      <Map
        mapboxAccessToken={TOKEN}
        mapLib={mapboxgl}
        interactive={true}
        mapStyle={"mapbox://styles/mapbox/streets-v12"}
        style={{
          width: "80%",
          height: "700px",
          borderRadius: "10px",
          boxShadow: "0 0 10px rgba(0, 0, 0, 0.5)",
        }}
        {...viewport}
      >
        <Marker
          longitude={marker.longitude}
          latitude={marker.latitude}
          anchor="bottom"
          draggable={false}
        >
          <DevicePinMap
            DeviceID={recentLocation?.device_id}
            DeviceType={"edge"}
          />
        </Marker>
        <NavigationControl />
      </Map>
    </div>
  );
};

export default DevicesMap;
