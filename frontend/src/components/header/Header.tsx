"use client";

import { Box, Button, Group } from "@mantine/core";
import styles from "./Header.module.scss";
import { useRouter, usePathname } from "next/navigation";
import Link from "next/link";
import { useEffect, useState } from "react";
import { getUserData, handleLogout } from "@/services/auth.service";

const Header = () => {
  const [userData, setUserData] = useState<any>(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const router = useRouter();

  useEffect(() => {
    const fetchData = async () => {
      let userData = await getUserData();

      if (userData == "") {
        setIsLoggedIn(false);
        setUserData(null);
      } else {
        setIsLoggedIn(true);
        setUserData(userData);
      }
    };

    fetchData();
  }, [usePathname().toString()]);

  return (
    <Box pb={15}>
      <header className={styles.header}>
        <Group justify={"space-between"}>
          <Link href={"/"}>
            <h1 className={styles.brandName}>xTracker</h1>
          </Link>

          <Group gap={10}>
            <Link href={"devices"}>Devices</Link>
          </Group>

          <Group>{userData && <p>Welcome {userData.username}</p>}</Group>

          <Group>
            {!isLoggedIn ? (
              <Button onClick={() => router.push("/auth")}>Log in</Button>
            ) : (
              <Button onClick={() => handleLogout()}>Logout</Button>
            )}
          </Group>
        </Group>
      </header>
    </Box>
  );
};

export default Header;
