import { EmptyState } from "../../components/feedback/EmptyState";
import { PageContainer } from "../../components/layout/PageContainer";
import { PageHeader } from "../../components/layout/PageHeader";

export default function LoginPage() {
  return (
    <main className="op-auth-surface">
      <PageContainer narrow>
        <PageHeader
          title="OpsPilot Login"
          description="Authentication will be implemented in the dedicated auth task."
        />
        <EmptyState
          title="Login is not enabled yet"
          description="This placeholder keeps the route available without adding auth logic."
        />
      </PageContainer>
    </main>
  );
}
