"use client";

import { App, ConfigProvider } from "antd";
import type { ReactNode } from "react";

import { antdTheme } from "../config/antd-theme";

type ProvidersProps = {
  children: ReactNode;
};

export default function Providers({ children }: ProvidersProps) {
  return (
    <ConfigProvider theme={antdTheme}>
      <App>{children}</App>
    </ConfigProvider>
  );
}
