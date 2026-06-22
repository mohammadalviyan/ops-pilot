"use client";

import { BellOutlined } from "@ant-design/icons";
import { Avatar, Space, Typography } from "antd";

const { Text } = Typography;

export function Topbar() {
  return (
    <header className="op-topbar op-glass">
      <Space align="center" style={{ width: "100%", justifyContent: "space-between" }}>
        <div>
          <h2 className="op-topbar-title">RPA Operations Workspace</h2>
          <p className="op-topbar-caption">Shell foundation ready for MVP workflows.</p>
        </div>
        <Space size={16}>
          <BellOutlined aria-label="Notifications placeholder" />
          <Space size={10}>
            <Avatar>OP</Avatar>
            <Text type="secondary">User profile pending auth</Text>
          </Space>
        </Space>
      </Space>
    </header>
  );
}
