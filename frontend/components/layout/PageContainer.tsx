import type { ReactNode } from "react";

type PageContainerProps = {
  children: ReactNode;
  narrow?: boolean;
};

export function PageContainer({ children, narrow = false }: PageContainerProps) {
  return (
    <section
      className={`op-page-container${narrow ? " op-page-container-narrow" : ""}`}
    >
      {children}
    </section>
  );
}
