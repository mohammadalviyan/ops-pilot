"use client";

import { Skeleton, Space } from "antd";

type LoadingStateProps = {
  label?: string;
};

export function LoadingState({ label = "Preparing workspace..." }: LoadingStateProps) {
  return (
    <div className="op-loading-state op-glass-strong" aria-busy="true">
      <Space direction="vertical" size={18} style={{ width: "100%" }}>
        <p className="op-topbar-title">{label}</p>
        <Skeleton active paragraph={{ rows: 3 }} title={false} />
      </Space>
    </div>
  );
}
