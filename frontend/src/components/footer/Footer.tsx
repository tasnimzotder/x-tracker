import { Container, Group, Anchor } from "@mantine/core";
// import { MantineLogo } from "@mantine/ds";
import classes from "./Footer.module.scss";

const links = [
  { link: "#", label: "Contact" },
  { link: "#", label: "Privacy" },
  { link: "#", label: "Blog" },
  { link: "#", label: "Careers" },
];

export function Footer() {
  const items = links.map((link) => (
    <Anchor<"a">
      c="dimmed"
      key={link.label}
      href={link.link}
      // onClick={(event) => event.preventDefault()}
      size="sm"
    >
      {link.label}
    </Anchor>
  ));

  return (
    <div className={classes.footer}>
      <Container className={classes.inner}>
        {/*<MantineLogo size={28} />*/}
        <div>xTracker</div>
        <Group className={classes.links}>{items}</Group>
      </Container>
    </div>
  );
}

export default Footer;
