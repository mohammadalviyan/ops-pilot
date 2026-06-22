# Architecture — OpsPilot

## 1. Overview

OpsPilot is an internal RPA Operations, Ticketing, Monitoring, and Reporting Dashboard.

The system is built as a monorepo with separate backend and frontend applications.

```txt
opspilot/
  backend/
  frontend/
  docs/
  backlog/
  tasks/
  docker-compose.yml
  AGENTS.md
```

OpsPilot should be developed using a task-based AI workflow. The agent must work from one active task at a time, follow `AGENTS.md`, and update project documentation after completing each task.

---

## 2. Architecture Goals

- Keep MVP simple and locally runnable.
- Use modular backend structure.
- Keep frontend feature-based and reusable.
- Use Ant Design as the primary frontend component system.
- Apply OpsPilot Glass Style consistently without sacrificing readability.
- Make API contract stable for FE/BE collaboration.
- Avoid over-engineering before MVP is complete.
- Make the project friendly for AI coding agents by using clear docs, task files, reusable skills, and status tracking.
- Ensure every completed task updates the relevant project docs so the agent does not lose context.

---

## 3. Tech Stack Decisions

| Area | Decision |
|---|---|
| Backend | Go |
| HTTP Router | Gin |
| Database | PostgreSQL |
| DB Driver | pgx |
| Migration | golang-migrate |
| Password Hashing | bcrypt |
| Auth | JWT-based auth |
| Logger | slog or zerolog |
| Frontend | Next.js + React + TypeScript |
| Frontend Router | Next.js App Router |
| Component Base | Ant Design |
| UI Style | OpsPilot Glass Style |
| Theme System | Ant Design ConfigProvider theme tokens |
| Styling | Ant Design theme + local CSS classes |
| Icons | @ant-design/icons |
| Forms | Ant Design Form, with optional zod validation when needed |
| Validation | Backend validation is required; frontend validation may use Ant Design rules or zod |
| Table | Ant Design Table |
| Chart | @ant-design/charts or Recharts if explicitly chosen by task |
| API Client | Centralized fetch wrapper in `frontend/lib/api.ts` |
| Local Runtime | Docker Compose |

### Frontend UI Decision

OpsPilot frontend uses **Next.js + Ant Design + OpsPilot Glass Style**.

Ant Design is selected because OpsPilot is an internal operations dashboard that needs strong support for:

- Dashboard layout.
- Tables.
- Forms.
- Filters.
- Status tags.
- Timeline.
- Modal/drawer interactions.
- Reports and monitoring screens.

OpsPilot Glass Style is applied on top of Ant Design using:

- `ConfigProvider` theme tokens.
- Local CSS classes.
- Reusable wrapper components.
- Consistent layout and data-display components.

Do not use shadcn/ui, Material UI, Bootstrap, Chakra UI, or another UI component library unless a future task explicitly changes this decision and records it in `docs/DECISIONS.md`.

---

## 4. Documentation Source of Truth

The agent must use these docs as source of truth:

```txt
docs/
  PRD.md              Product requirements and MVP scope.
  ARCHITECTURE.md     Technical architecture and project conventions.
  API_CONTRACT.md     API endpoint, request, response, and error contract.
  DB_SCHEMA.md        Database schema, enum, indexes, and migration notes.
  UI_SYSTEM.md        Frontend visual direction, UX rules, and component rules.
  DECISIONS.md        Accepted technical/product decisions.
  AI_WORKFLOW.md      Agent workflow, task lifecycle, and prompt usage.
```

Rules:

- Do not duplicate large sections from one doc into another.
- Do not rewrite all docs unless the active task explicitly requests it.
- Update only the docs affected by the completed task.
- If implementation changes an API, update `docs/API_CONTRACT.md`.
- If implementation changes database schema, update `docs/DB_SCHEMA.md`.
- If implementation changes reusable UI/component strategy, update `docs/UI_SYSTEM.md`.
- If implementation makes a technical decision, update `docs/DECISIONS.md`.

---

## 5. Backend Structure

```txt
backend/
  cmd/
    api/
      main.go
  internal/
    config/
    database/
    middleware/
    module/
      auth/
        handler.go
        service.go
        repository.go
        model.go
        dto.go
        routes.go
      user/
      ticket/
      robot/
      asset/
      dashboard/
      report/
      knowledge/
    pkg/
      response/
      validator/
      logger/
      jwt/
      password/
      pagination/
  migrations/
  docs/
  go.mod
  go.sum
```

---

## 6. Backend Layering Rules

### Handler

Responsibilities:

- Parse request.
- Validate request DTO.
- Call service.
- Return standard response.

Must not:

- Contain SQL query.
- Contain business rules.
- Directly access database.

### Service

Responsibilities:

- Business logic.
- Status transition logic.
- SLA calculation.
- Activity generation.
- Permission checks.

Must not:

- Know HTTP details.
- Write raw response object.

### Repository

Responsibilities:

- Database query.
- Transaction execution.
- Data persistence.

Must not:

- Contain HTTP logic.
- Contain UI or presentation logic.

---

## 7. Frontend Structure

Frontend uses Next.js App Router and Ant Design.

```txt
frontend/
  app/
    layout.tsx
    providers.tsx
    globals.css
    login/
      page.tsx
    dashboard/
      page.tsx
    tickets/
      page.tsx
      create/
        page.tsx
      [id]/
        page.tsx
    robots/
      page.tsx
      [id]/
        page.tsx
    robot-runs/
      page.tsx
    assets/
      page.tsx
    reports/
      page.tsx
    knowledge-base/
      page.tsx

  components/
    layout/
      AppShell.tsx
      Sidebar.tsx
      Topbar.tsx
      PageContainer.tsx
      PageHeader.tsx

    glass/
      GlassCard.tsx
      GlassPanel.tsx
      GlassStatCard.tsx
      GlassPageHeader.tsx

    data-display/
      StatusTag.tsx
      PriorityTag.tsx
      DataTable.tsx
      FilterToolbar.tsx
      ActivityTimeline.tsx
      MetricGroup.tsx

    feedback/
      EmptyState.tsx
      LoadingState.tsx
      ErrorState.tsx

    forms/
      FormSection.tsx
      FieldError.tsx

    charts/
      ChartCard.tsx
      TicketStatusChart.tsx
      RobotRunChart.tsx
      TicketTrendChart.tsx

  features/
    auth/
      components/
      hooks/
      services/
      types.ts
    dashboard/
      components/
      hooks/
      services/
      types.ts
    tickets/
      components/
      hooks/
      services/
      types.ts
    robots/
      components/
      hooks/
      services/
      types.ts
    assets/
      components/
      hooks/
      services/
      types.ts
    reports/
      components/
      hooks/
      services/
      types.ts
    knowledge-base/
      components/
      hooks/
      services/
      types.ts

  config/
    antd-theme.ts
    navigation.ts

  lib/
    api.ts
    auth.ts
    constants.ts
    date.ts
    format.ts
    utils.ts

  hooks/
    useDebounce.ts
    usePagination.ts

  types/
    api.ts
    common.ts
    enums.ts

  styles/
    glass.css

  public/
```

### Frontend Structure Rules

- `app/*/page.tsx` should compose feature components, not contain heavy UI logic.
- `features/*` contains feature-specific UI, hooks, services, and types.
- `components/*` contains reusable cross-feature components.
- `components/glass/*` contains OpsPilot glass wrapper components.
- `components/data-display/*` contains reusable table, badge/tag, timeline, and metric components.
- `config/antd-theme.ts` is the only place for Ant Design theme tokens.
- `styles/glass.css` is the only place for shared glass style classes.
- API calls must go through `lib/api.ts` or feature service files that use `lib/api.ts`.

---

## 8. Ant Design Setup Rules

If using Next.js App Router, install:

```bash
npm install antd @ant-design/icons @ant-design/nextjs-registry
```

If charts are needed:

```bash
npm install @ant-design/charts
```

Required setup:

### `frontend/app/layout.tsx`

- Import global CSS.
- Import `styles/glass.css`.
- Use `AntdRegistry` from `@ant-design/nextjs-registry`.
- Wrap children with `Providers`.

### `frontend/app/providers.tsx`

- Must be a client component.
- Use `ConfigProvider`.
- Use Ant Design `App` component for message, modal, and notification context.
- Import theme from `config/antd-theme.ts`.

### `frontend/config/antd-theme.ts`

- Store all Ant Design theme tokens.
- Define primary, success, warning, error, background, border radius, and component tokens.
- Do not scatter theme values across random page components.

### `frontend/styles/glass.css`

- Store shared glass classes.
- Define background gradients.
- Define `op-glass`, `op-glass-strong`, `op-card-soft`, and related helper classes.
- Keep contrast readable.

---

## 9. OpsPilot Glass Style Rules

OpsPilot uses a subtle glass dashboard style.

The UI should feel like:

- RPA operations command center.
- Internal monitoring cockpit.
- Professional support dashboard.
- Clean but not empty.
- Modern but not decorative.
- Data-first and readable.

Use glass effect for:

- App shell.
- Sidebar.
- Topbar.
- Dashboard stat cards.
- Filter panels.
- Important summary cards.

Avoid glass effect for:

- Every table cell.
- Every button.
- Dense forms.
- Long text areas.
- Low-contrast content.
- Critical error messages.

Glass style must improve visual hierarchy, not reduce readability.

---

## 10. Frontend Component Strategy

The frontend must avoid one-off UI implementation.

Before creating a page, create or reuse base components.

Required base components:

```txt
components/layout/
  AppShell
  Sidebar
  Topbar
  PageContainer
  PageHeader

components/glass/
  GlassCard
  GlassPanel
  GlassStatCard
  GlassPageHeader

components/data-display/
  StatusTag
  PriorityTag
  DataTable
  FilterToolbar
  ActivityTimeline

components/feedback/
  EmptyState
  LoadingState
  ErrorState

components/forms/
  FormSection
  FieldError
```

Rules:

- Use Ant Design components as the base.
- Create wrapper components only when it improves consistency.
- Do not manually style status badges in every page.
- Do not repeat table/filter logic in every module.
- Each page must support loading, empty, and error state when relevant.
- Do not build custom table, modal, dropdown, date picker, form, or menu from scratch if Ant Design already provides it.
- Do not mix another component library into the frontend.

---

## 11. Frontend Page Rules

Every page should include:

- Clear page title.
- Short helper description.
- Primary action on top-right if needed.
- Main content area.
- Loading state.
- Error state.
- Empty state when applicable.
- Consistent status/priority tags.
- Responsive behavior for desktop and tablet.

Pages should answer real operational questions:

- What needs attention today?
- Which robots failed?
- Which tickets are waiting?
- Which license or asset is risky?
- What should the support team do next?

Avoid:

- Generic SaaS copy.
- Placeholder text like `Lorem ipsum`.
- Repetitive equal-size cards without hierarchy.
- Large empty white sections.
- Random gradients.
- Default Ant Design look without OpsPilot identity.

---

## 12. Authentication Flow

MVP auth uses email/password and JWT.

```txt
User submits login form
  ↓
POST /api/v1/auth/login
  ↓
Backend validates credentials
  ↓
Backend returns token and user profile
  ↓
Frontend stores auth state
  ↓
Protected routes require authenticated user
```

Recommended MVP storage:

- Use httpOnly cookie if implemented early.
- If cookie auth delays MVP, use local storage temporarily and document the trade-off in `docs/DECISIONS.md`.

Frontend auth rules:

- Auth state should be managed in `features/auth`.
- Auth API calls should use centralized API client.
- Protected route behavior must be documented when implemented.
- Do not expose JWT secret or sensitive config to frontend.

---

## 13. Database Overview

Core tables:

- users
- roles
- departments
- applications
- ticket_categories
- tickets
- ticket_comments
- ticket_activities
- robots
- robot_run_logs
- assets
- escalations
- knowledge_base
- sla_rules

Database details must be maintained in `docs/DB_SCHEMA.md`.

Rules:

- Use migrations for schema changes.
- Do not modify database schema without adding migration files.
- Keep enum values aligned with PRD and API contract.
- Add indexes for commonly filtered fields when implementing list endpoints.

---

## 14. API Design Principles

- Prefix all endpoints with `/api/v1`.
- Use consistent success/error/pagination response.
- Use query params for filtering list endpoints.
- Use PATCH for partial status/assignment changes.
- Soft delete where appropriate.
- Backend should generate `ticket_no`.
- Keep API request/response documented in `docs/API_CONTRACT.md`.

Standard response format:

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

Paginated response format:

```json
{
  "success": true,
  "message": "Success",
  "data": [],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total_data": 100,
    "total_page": 10
  }
}
```

---

## 15. Local Development

Minimum local services:

```txt
PostgreSQL
Backend API
Frontend App
```

Optional:

```txt
pgAdmin
Redis
```

Rules:

- Use Docker Compose for local development.
- Keep `.env.example` updated.
- Never commit real `.env`.
- Update `README.md` if setup commands change.
- Prefer Makefile or npm scripts for repeatable commands.

---

## 16. Observability for MVP

For MVP, keep observability simple:

- Structured logs.
- Request logging middleware.
- Error response consistency.
- Health check endpoint.

Future:

- OpenTelemetry.
- Sentry.
- Loki/Grafana.

---

## 17. Security Notes

- Password must be hashed.
- JWT secret must come from environment variable.
- Never commit `.env`.
- Validate all input.
- Restrict requester access to own tickets where role rules are implemented.
- Internal comments must not be visible to requester.
- Do not expose backend secrets in frontend environment variables.
- Frontend environment variables must use `NEXT_PUBLIC_` only for safe public values.

---

## 18. Agent Documentation Rules

Every coding task must update project documentation when relevant.

After completing a task, the agent must update:

- Active task file in `/tasks`.
- `backlog/STATUS.md`.
- `backlog/BACKLOG.md` if task status changes.
- `docs/DECISIONS.md` if a technical decision was made.
- `docs/API_CONTRACT.md` if API request/response changes.
- `docs/DB_SCHEMA.md` if database schema changes.
- `docs/UI_SYSTEM.md` if UI system or reusable components change.
- `README.md` if setup command or environment variable changes.

The agent must not rewrite all documentation unless the active task explicitly asks for it.

Documentation updates must be minimal, accurate, and related to the task.

---

## 19. Task Execution Workflow

The agent must work on one task at a time.

Required flow:

1. Read `AGENTS.md`.
2. Read the active task file.
3. Read only relevant docs.
4. Summarize task scope.
5. Create a short implementation plan.
6. Implement only the active task.
7. Run relevant validation.
8. Update task status.
9. Update `backlog/STATUS.md`.
10. Summarize changed files and next task recommendation.

The agent must not implement features outside the active task.

### If the active task is frontend-related

The agent must read:

- `docs/UI_SYSTEM.md`
- `.agents/skills/frontend-ant-design-glass/SKILL.md`
- `docs/API_CONTRACT.md` when API integration is involved

### If the active task is backend-related

The agent must read:

- `docs/API_CONTRACT.md`
- `docs/DB_SCHEMA.md`
- `.agents/skills/backend-go/SKILL.md`

### If the active task is database-related

The agent must read:

- `docs/DB_SCHEMA.md`
- `docs/DECISIONS.md`
- migration files

---

## 20. Frontend Task Policy

For frontend tasks, the agent must prioritize this order:

1. Setup Ant Design provider and theme.
2. Setup OpsPilot glass style.
3. Create reusable layout components.
4. Create reusable data-display components.
5. Create page shell.
6. Integrate API.
7. Add loading, empty, and error states.
8. Polish UX copy.

Do not start from full page implementation if base components do not exist yet.

Recommended frontend build order:

1. Ant Design provider, theme, and glass CSS.
2. AppShell, Sidebar, Topbar, PageContainer, PageHeader.
3. GlassCard, GlassPanel, GlassStatCard.
4. StatusTag, PriorityTag, DataTable, FilterToolbar.
5. Login page.
6. Dashboard shell.
7. Ticket list page.
8. Ticket detail page.
9. Ticket create form.
10. Robot list/detail pages.
11. Asset list page.
12. Reports page.

---

## 21. MVP Build Sequence

1. Project foundation.
2. Docker Compose and environment setup.
3. Backend health check and config.
4. Database migration and seed data.
5. Frontend Ant Design provider, theme, and glass CSS.
6. Frontend app shell and reusable base components.
7. Auth backend.
8. Login frontend.
9. Ticketing core backend.
10. Ticket list and ticket create frontend.
11. Ticket detail, comments, activities.
12. Dashboard summary and chart.
13. Robot monitoring.
14. Asset monitoring.
15. Reports.
16. Knowledge base.

---

## 22. Non-MVP Constraints

Do not implement before MVP core is stable:

- Direct UiPath Orchestrator integration.
- Microsoft Teams integration.
- WhatsApp integration.
- Real-time websocket monitoring.
- AI root cause analysis.
- Complex approval workflow.
- SSO enterprise.
- Multi-tenant company logic.
- Kubernetes deployment.

Future features may be prepared through clean architecture, but must not block MVP delivery.
