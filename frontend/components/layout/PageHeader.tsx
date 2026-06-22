import type { ReactNode } from "react";

type PageHeaderProps = {
  title: string;
  description: string;
  action?: ReactNode;
};

export function PageHeader({ title, description, action }: PageHeaderProps) {
  return (
    <header className="op-page-header">
      <div>
        <h1 className="op-page-title">{title}</h1>
        <p className="op-page-description">{description}</p>
      </div>
      {action}
    </header>
  );
}
