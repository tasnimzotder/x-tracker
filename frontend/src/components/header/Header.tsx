import { Box, Button, Group } from "@mantine/core";
import styles from "./Header.module.scss";

const Header = () => {
  return (
    <Box pb={15}>
      <header className={styles.header}>
        <Group justify={"space-between"}>
          <h1 className={styles.brandName}>xTracker</h1>

          <Group gap={10}>
            <a>Home</a>
            <a>About us</a>
            <a>Contact us</a>
          </Group>

          <Group>
            <Button>Log in</Button>
          </Group>
        </Group>
      </header>
    </Box>
  );
};

export default Header;
