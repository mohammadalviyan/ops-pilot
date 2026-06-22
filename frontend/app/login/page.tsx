import { LoginForm } from "../../features/auth/components/LoginForm";
import { PageContainer } from "../../components/layout/PageContainer";
import { PageHeader } from "../../components/layout/PageHeader";

export default function LoginPage() {
  return (
    <main className="op-auth-surface">
      <PageContainer narrow>
        <PageHeader
          title="OpsPilot Login"
          description="Access the RPA operations workspace for tickets, robot monitoring, assets, and reporting."
        />
        <LoginForm />
      </PageContainer>
    </main>
  );
}
