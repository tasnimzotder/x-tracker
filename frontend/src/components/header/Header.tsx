"use client";

import { Avatar, Box, Button, Group } from "@mantine/core";
import styles from "./Header.module.scss";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { useAuth } from "@/contexts/authContext";

const Header = () => {
  const { loggedIn, userData, logout } = useAuth();

  const router = useRouter();

  return (
    <Box pb={15}>
      <header className={styles.header}>
        <Group justify={"space-between"}>
          <Group gap={20}>
            <Link href={"/"}>
              <h1 className={styles.brandName}>xTracker</h1>
            </Link>

            <Link href={"devices"}>Devices</Link>
          </Group>

          <Group>
            {userData && (
              <p>
                <Link href={"/profile"}>
                  <Avatar
                    src={
                      "https://avatars.githubusercontent.com/u/44049528mm?v=4"
                    }
                    alt={"user profile picture"}
                  />
                </Link>
              </p>
            )}

            {!loggedIn ? (
              <Button onClick={() => router.push("/auth")}>Log in</Button>
            ) : (
              <Button onClick={() => logout()}>Logout</Button>
            )}
          </Group>
        </Group>
      </header>
    </Box>
  );
};

export default Header;
