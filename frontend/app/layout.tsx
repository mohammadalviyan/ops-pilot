import type { Metadata } from "next";
import { AntdRegistry } from "@ant-design/nextjs-registry";

import Providers from "./providers";
import "./globals.css";
import "../styles/glass.css";

export const metadata: Metadata = {
  title: "OpsPilot",
  description: "RPA operations command center",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <AntdRegistry>
          <Providers>{children}</Providers>
        </AntdRegistry>
      </body>
    </html>
  );
}
