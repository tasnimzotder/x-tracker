import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

import "@mantine/core/styles.css";

import { MantineProvider, ColorSchemeScript } from "@mantine/core";
import Header from "@/components/header/Header";
import Footer from "@/components/footer/Footer";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <head>
        <ColorSchemeScript />
      </head>
      <body className={inter.className}>
        <MantineProvider>
          <Header />
          {children}
          <Footer />
        </MantineProvider>
      </body>
    </html>
  );
}
