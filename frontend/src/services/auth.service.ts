"use server";

import { cookies } from "next/headers";

interface authRequest_t {
  id?: number;
  username?: string;
  email?: string;
  password?: string;
  firstName?: string;
  lastName?: string;
}

const handleLogin = async (req: authRequest_t) => {
  const url: string = `${process.env.NEXT_PUBLIC_API_URL}/v1/users/login`;

  const reqOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(req),
  };

  const response = await fetch(url, reqOptions).catch((err) => {
    console.error(err);
  });

  // console.log(response);

  if (!response || response.status !== 200) {
    return null;
  }

  return await response.json();
};

const handleRegister = async (req: authRequest_t) => {
  const url: string = `${process.env.NEXT_PUBLIC_API_URL}/v1/users/create`;

  const reqOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(req),
  };

  const response = await fetch(url, reqOptions).catch((err) => {
    console.error(err);
  });

  if (!response || response.status !== 200) {
    return null;
  }

  return await response.json();
};

// const refetchUserData = async () => {
//   let userData = await getUserData();
//
//   const url: string = `${process.env.NEXT_PUBLIC_API_URL}/v1/users/id/${userData.id}`;
//
//   const reqOptions = {
//     method: "GET",
//     headers: {
//       "Content-Type": "application/json",
//     },
//   };
//
//   const response = await fetch(url, reqOptions).catch((err) => {
//     console.error(err);
//   });
//
//   if (!response || response.status !== 200) {
//     return false;
//   }
//
//   let data = await response.json();
//
//   // setCookie("session", JSON.stringify(data));
// };

// const handleProfileUpdate = async (req: authRequest_t) => {
//   let userData = await getUserData();
//
//   req.id = userData.id;
//
//   const url: string = `${process.env.NEXT_PUBLIC_API_URL}/v1/users/update`;
//
//   const reqOptions = {
//     method: "PUT",
//     headers: {
//       "Content-Type": "application/json",
//     },
//     body: JSON.stringify(req),
//   };
//
//   const response = await fetch(url, reqOptions).catch((err) => {
//     console.error(err);
//   });
//
//   if (!response || response.status !== 200) {
//     return false;
//   }
//
//   let data = await response.json();
//
//   setCookie("session", JSON.stringify(data));
// };

// const setCookie = (name: string, value: string) => {
//   cookies().set(name, value, {
//     path: "/",
//     maxAge: 60 * 60 * 24 * 7, // 1 week
//     httpOnly: true,
//   });
// };

// const getUserData = async () => {
//   const data = await cookies().get("session");
//
//   let sessionData = null;
//
//   if (data?.value) {
//     sessionData = JSON.parse(data.value);
//
//     // console.log({ sessionData });
//
//     return sessionData;
//   }
//
//   return "";
// };

export {
  handleLogin,
  handleRegister,
  // getUserData,
  // refetchUserData,
  // handleProfileUpdate,
};

export type { authRequest_t };
