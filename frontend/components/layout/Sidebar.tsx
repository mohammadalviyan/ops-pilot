"use client";

import {
  AppstoreOutlined,
  BookOutlined,
  DashboardOutlined,
  DatabaseOutlined,
  FileTextOutlined,
  MonitorOutlined,
  RobotOutlined,
  ToolOutlined,
} from "@ant-design/icons";
import { Menu } from "antd";
import Link from "next/link";
import { usePathname } from "next/navigation";

const navigationItems = [
  {
    key: "/dashboard",
    icon: <DashboardOutlined />,
    label: <Link href="/dashboard">Dashboard</Link>,
  },
  {
    key: "/tickets",
    icon: <FileTextOutlined />,
    label: "Tickets",
    disabled: true,
  },
  {
    key: "/robots",
    icon: <RobotOutlined />,
    label: "Robots",
    disabled: true,
  },
  {
    key: "/robot-runs",
    icon: <MonitorOutlined />,
    label: "Robot Runs",
    disabled: true,
  },
  {
    key: "/assets",
    icon: <DatabaseOutlined />,
    label: "Assets",
    disabled: true,
  },
  {
    key: "/reports",
    icon: <AppstoreOutlined />,
    label: "Reports",
    disabled: true,
  },
  {
    key: "/knowledge-base",
    icon: <BookOutlined />,
    label: "Knowledge Base",
    disabled: true,
  },
  {
    key: "/settings",
    icon: <ToolOutlined />,
    label: "Settings",
    disabled: true,
  },
];

function selectedKey(pathname: string) {
  const match = navigationItems.find((item) => pathname.startsWith(item.key));
  return match?.key ?? "/dashboard";
}

export function Sidebar() {
  const pathname = usePathname();

  return (
    <nav aria-label="Primary navigation">
      <div className="op-brand">
        <div className="op-brand-mark">OP</div>
        <div className="op-brand-text">
          <span className="op-brand-name">OpsPilot</span>
          <span className="op-brand-caption">RPA command center</span>
        </div>
      </div>
      <Menu
        mode="inline"
        selectedKeys={[selectedKey(pathname)]}
        items={navigationItems}
        style={{ borderInlineEnd: 0, background: "transparent" }}
      />
    </nav>
  );
}
