"use client";

import { Empty } from "antd";

type EmptyStateProps = {
  title: string;
  description: string;
};

export function EmptyState({ title, description }: EmptyStateProps) {
  return (
    <div className="op-empty-state op-glass-strong">
      <Empty description={false}>
        <h2 className="op-topbar-title">{title}</h2>
        <p className="op-page-description" style={{ marginInline: "auto" }}>
          {description}
        </p>
      </Empty>
    </div>
  );
}
