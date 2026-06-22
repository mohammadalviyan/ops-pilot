"use client";

import type { ReactNode } from "react";
import { Layout } from "antd";

import { Sidebar } from "./Sidebar";
import { Topbar } from "./Topbar";

const { Content, Sider } = Layout;

type AppShellProps = {
  children: ReactNode;
};

export function AppShell({ children }: AppShellProps) {
  return (
    <div className="op-app-shell">
      <Layout className="op-shell-layout">
        <Sider className="op-sidebar op-glass" width={260}>
          <Sidebar />
        </Sider>
        <Layout className="op-shell-layout">
          <Topbar />
          <Content className="op-content">{children}</Content>
        </Layout>
      </Layout>
    </div>
  );
}
