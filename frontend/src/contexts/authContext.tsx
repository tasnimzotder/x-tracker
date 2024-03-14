"use client";

import React, {
  createContext,
  useContext,
  useEffect,
  useRef,
  useState,
} from "react";
import {
  authRequest_t,
  handleLogin,
  handleRegister,
} from "@/services/auth.service";
import { deleteCookie, getCookie, setCookie } from "cookies-next";
import { useRouter } from "next/navigation";
import { getDeviceList } from "@/services/device.service";
import { getRecentLocation } from "@/services/location.service";

const AuthContext = createContext({
  locations: [] as Array<any>,
  loggedIn: false as boolean,
  userData: null as any,
  devices: [] as Array<any>,
  recentLocation: null as any,
  loadLocations: false as boolean,
  setLoadLocations: (val: boolean) => {},
  login: (req: authRequest_t) => {},
  register: (req: authRequest_t) => {},
  logout: () => {},
});
const useAuth = () => useContext(AuthContext);

const AuthContextProvider = ({ children }: { children: React.ReactNode }) => {
  // const [ws, setWs] = useState<WebSocket | any>();
  const [locations, setLocations] = useState<Array<any>>([]);
  const [loggedIn, setLoggedIn] = useState<boolean>(false);
  const [userData, setUserData] = useState<any>(null);
  const [devices, setDevices] = useState<Array<any>>([]);
  const [loadLocations, setLoadLocations] = useState<boolean>(false);

  const [recentLocation, setRecentLocation] = useState();

  const cookie_name = "session";
  const router = useRouter();

  // let ws: WebSocket | any = null;
  const ws = useRef<WebSocket | null>(null);

  const updateDeviceList = async () => {
    const res = await getDeviceList(userData.id);

    // todo: remove this
    await getRecentLocation(1);

    if (res != null) {
      setDevices(res);
    }
  };

  const updateRecentLocation = async (device_id: number) => {
    const eventSource = new EventSource(
      `${process.env.NEXT_PUBLIC_API_URL}/v1/sse/location_updates/${device_id}`,
    );
    // const eventSource = new EventSource(
    //   `http://localhost:8080/v1/sse/location_updates/1`,
    // );

    eventSource.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log("Message received");
      console.log({ data });

      setRecentLocation(data);

      // setLocations((prev) => [...prev, data]);
    };

    eventSource.onerror = (error) => {
      console.error("EventSource failed:", error);
    };
  };

  useEffect(() => {
    let cookie = getCookie(cookie_name);

    console.log({ cookie });

    if (cookie) {
      setLoggedIn(true);
      setUserData(JSON.parse(cookie as string));

      if (userData) {
        updateDeviceList();
      }
    }

    // updateRecentLocation(1);
  }, []);

  useEffect(() => {
    if (loggedIn && userData) {
      updateDeviceList();
    }
  }, [loggedIn, userData]);

  useEffect(() => {
    console.log(devices);

    if (devices.length > 0) {
      updateRecentLocation(devices[0].id);
    }
  }, [devices]);

  const resetAuth = () => {
    deleteCookie(cookie_name);

    setLoggedIn(false);
    setUserData(null);
    setDevices([]);
    setLocations([]);
    setLoadLocations(false);

    if (ws.current) {
      // ws.current.close();
      ws.current?.send("disconnect");
      // setWs(null);
      ws.current = null;
      console.log("WebSocket closed");
    }
  };

  const login = async (req: authRequest_t) => {
    const res = await handleLogin(req);

    if (res != null) {
      setLoggedIn(true);
      setCookie(cookie_name, JSON.stringify(res), {
        maxAge: 60 * 60 * 24 * 7, // 1 week
      });
      setUserData(res);

      router.push("/profile");
    } else {
      resetAuth();
    }
  };

  const register = async (req: authRequest_t) => {
    const res = await handleRegister(req);

    if (res != null) {
      setLoggedIn(true);
      setCookie(cookie_name, JSON.stringify(res), {
        maxAge: 60 * 60 * 24 * 7, // 1 week
      });
      setUserData(res);

      router.push("/profile");
    } else {
      resetAuth();
    }
  };

  const fetchUserData = () => {};

  const logout = () => {
    resetAuth();

    router.push("/auth");
  };

  useEffect(() => {
    console.log({ loggedIn });
    console.log({ userData });
  }, [loggedIn, userData]);

  const contextValues = {
    // ws,
    locations,
    loggedIn,
    userData,
    devices,
    recentLocation,
    loadLocations,
    setLoadLocations,
    login,
    register,
    logout,
    fetchUserData,
  };

  return (
    <AuthContext.Provider value={contextValues}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthContextProvider, useAuth };
