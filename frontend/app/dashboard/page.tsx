import { EmptyState } from "../../components/feedback/EmptyState";
import { AppShell } from "../../components/layout/AppShell";
import { PageContainer } from "../../components/layout/PageContainer";
import { PageHeader } from "../../components/layout/PageHeader";

export default function DashboardPage() {
  return (
    <AppShell>
      <PageContainer>
        <PageHeader
          title="Operations Dashboard"
          description="Review the RPA operations workspace before live dashboard data is connected."
        />
        <EmptyState
          title="Dashboard data is not connected yet"
          description="The app shell is ready. Summary metrics and charts will be added in a later dashboard task."
        />
      </PageContainer>
    </AppShell>
  );
}
