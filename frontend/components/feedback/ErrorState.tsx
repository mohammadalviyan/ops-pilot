"use client";

import { Alert } from "antd";

type ErrorStateProps = {
  title?: string;
  description: string;
};

export function ErrorState({
  title = "Unable to load this workspace",
  description,
}: ErrorStateProps) {
  return (
    <div className="op-error-state">
      <Alert type="error" showIcon message={title} description={description} />
    </div>
  );
}
