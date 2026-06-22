# Decisions — OpsPilot

This file records technical and product decisions so AI agents do not re-litigate the same choices repeatedly.

---

## ADR-001 — Use Monorepo

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Use one repository with separate `backend/` and `frontend/` folders.

### Reason

OpsPilot is a portfolio/internal product with tight FE-BE iteration. A monorepo keeps docs, API contract, tasks, and implementation close together.

### Consequence

Root-level documentation and AI instructions must clearly mention which files belong to backend or frontend.

---

## ADR-002 — Use Go for Backend

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Backend uses Go.

### Reason

This aligns with the project goal and the developer's backend focus.

---

## ADR-003 — Use Gin for HTTP Router

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Use Gin for REST API routing.

### Reason

Gin is widely used in Go backend projects and fits the desired REST API style.

### Alternative

Chi can also work, but the project should choose one router early to avoid inconsistent patterns.

---

## ADR-004 — Use PostgreSQL

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Use PostgreSQL as the main database.

### Reason

The domain is relational: users, tickets, robots, assets, comments, activities, and reports.

---

## ADR-005 — Use Migration-based Schema Management

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Use `golang-migrate` or equivalent migration tooling.

### Reason

AI agents must not change database schema ad hoc. Every schema change must be traceable.

---

## ADR-006 — Use JWT Auth for MVP

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Use JWT-based authentication with hashed password for MVP.

### Reason

It is simple enough for MVP and already aligned with PRD.

### Future

SSO may be added after MVP.

---

## ADR-007 — Use Next.js + TypeScript for Frontend

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Frontend uses Next.js, React, and TypeScript.

### Reason

Dashboard app benefits from routing, reusable components, and type safety.

---

## ADR-008 — Use Tailwind + shadcn/ui

**Status:** Superseded  
**Date:** 2026-06-19  
**Superseded By:** ADR-011 — Use Ant Design with OpsPilot Glass Style

### Previous Decision

Use Tailwind CSS and shadcn/ui for UI components.

### Reason

This was initially proposed to speed up dashboard implementation while keeping a modern, clean UI.

### Superseded Reason

OpsPilot has shifted toward a more enterprise-dashboard component strategy. Ant Design is a better fit for ticketing, monitoring, tables, forms, filters, timeline, and reporting screens.

### Consequence

Do not use Tailwind + shadcn/ui as the main frontend component strategy unless a future ADR explicitly changes the decision again.

---

## ADR-009 — Execute One Ticket Per Agent Session

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

AI agent should execute only one `tasks/TASK-xxx.md` per session.

### Reason

This prevents context drift, scope creep, and unnecessary token usage.

---

## ADR-010 — Keep Future Integrations Out of MVP

**Status:** Accepted  
**Date:** 2026-06-19

### Decision

Do not implement direct UiPath, Teams, WhatsApp, websocket real-time monitoring, SSO, or AI RCA in the first MVP.

### Reason

The current goal is a stable operational dashboard with ticketing, monitoring, and reporting foundation.

---

## ADR-011 — Use Ant Design with OpsPilot Glass Style

**Status:** Accepted  
**Date:** 2026-06-22

### Decision

OpsPilot frontend will use Ant Design as the primary UI component library.

The visual identity will use **OpsPilot Glass Style**, implemented through:

- Ant Design `ConfigProvider` theme tokens.
- Local CSS classes.
- Reusable wrapper components.
- Consistent layout and data-display components.

### Reason

OpsPilot is an internal RPA operations dashboard with many enterprise-style UI needs:

- Dashboard summary cards.
- Data tables.
- Filters.
- Forms.
- Status tags.
- Priority tags.
- Timeline/activity logs.
- Modal/drawer interactions.
- Reports and monitoring screens.

Ant Design provides strong enterprise components out of the box. OpsPilot Glass Style prevents the UI from looking too plain, too default, or too generic.

### Consequences

- Do not use shadcn/ui, Material UI, Bootstrap, Chakra UI, or another component library as the primary UI system.
- Frontend tasks must reuse Ant Design components first.
- Custom components should wrap Ant Design, not replace it.
- UI consistency must be controlled through `frontend/config/antd-theme.ts` and `frontend/styles/glass.css`.
- Shared UI wrappers should live in `frontend/components/*`.
- Feature-specific UI should live in `frontend/features/*`.
- The agent must read `docs/UI_SYSTEM.md` and `.agents/skills/frontend-ant-design-glass/SKILL.md` before implementing visible frontend UI tasks.

### Implementation Notes

Required packages for Next.js App Router:

```bash
npm install antd @ant-design/icons @ant-design/nextjs-registry
```

Optional chart package:

```bash
npm install @ant-design/charts
```

Required frontend files:

```txt
frontend/app/layout.tsx
frontend/app/providers.tsx
frontend/config/antd-theme.ts
frontend/styles/glass.css
```

---

## ADR-012 — Use Reusable Frontend Base Components Before Building Feature Pages

**Status:** Accepted  
**Date:** 2026-06-22

### Decision

Before implementing full feature pages, the frontend must establish reusable base components.

Required base components include:

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

### Reason

AI agents often generate repetitive and generic UI when they start directly from feature pages. Reusable base components make the UI more consistent, human-designed, and maintainable.

### Consequences

- Frontend page tasks should reuse existing base components.
- If a required base component does not exist, the task should create it first.
- Do not manually style status badges, priority badges, tables, and empty states repeatedly in each page.
- Do not build a full page before the required shell and reusable components exist.

---

## ADR-013 — Use Temporary Local Storage for MVP Frontend Auth State

**Status:** Accepted  
**Date:** 2026-06-23

### Decision

For the MVP frontend auth flow, store the login access token and user profile in browser `localStorage`.

### Reason

The backend currently returns a bearer token from `POST /api/v1/auth/login`. Using `localStorage` keeps the first frontend login task small and allows the app to validate the basic auth flow before introducing cookie/session infrastructure.

### Consequences

- This is a temporary MVP approach, not the final preferred security posture.
- Frontend auth code must not store plaintext passwords.
- Frontend auth code must keep storage access centralized in `frontend/lib/auth.ts`.
- A future task may migrate to httpOnly cookie auth; if it does, update this decision and the API/client behavior together.
