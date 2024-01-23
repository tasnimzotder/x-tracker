"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

interface authRequest_t {
  username: string;
  email?: string;
  password: string;
}

const handleLogin = async (req: authRequest_t): Promise<boolean> => {
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

  console.log(response);

  if (!response || response.status !== 200) {
    return false;
  }

  let data = await response.json();
  // console.log(data);

  // const sessionData = {
  //   username: data.username,
  //   email: data.email,
  //   id: data.id,
  // };

  cookies().set("session", JSON.stringify(data), {
    path: "/",
    maxAge: 60 * 60 * 24 * 7, // 1 week
    httpOnly: true,
  });

  // console.log(JSON.parse(<string>cookies().get("session")?.value).username);

  redirect("/");

  // save data

  return true;
};

const handleLogout = () => {
  cookies().set("session", "", {
    path: "/",
    maxAge: -1,
    httpOnly: true,
  });

  redirect("/auth");
};

const getUserData = async () => {
  const data = await cookies().get("session");

  let sessionData = null;

  if (data?.value) {
    sessionData = JSON.parse(data.value);

    console.log({ sessionData });

    return sessionData;
  }

  return "";
};

export { handleLogin, getUserData, handleLogout };

export type { authRequest_t };
