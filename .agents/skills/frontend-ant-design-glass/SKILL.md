# Frontend Ant Design Glass Skill

Use this skill when implementing OpsPilot frontend UI using Next.js and Ant Design.

---

## Goal

Build frontend pages that feel like a polished RPA operations dashboard using Ant Design components with OpsPilot Glass Style.

The UI should feel:

- Professional.
- Clean.
- Modern.
- Operational.
- Data-first.
- Slightly glassmorphism.
- Human-designed.
- Not generic AI-generated SaaS.

---

## Required Docs

Before coding, read:

- `AGENTS.md`
- `docs/PRD.md`
- `docs/ARCHITECTURE.md`
- `docs/API_CONTRACT.md`
- `docs/UI_SYSTEM.md`
- `docs/DECISIONS.md`
- Active task file in `/tasks`

If available, also read:

- `frontend/config/antd-theme.ts`
- `frontend/styles/glass.css`
- Existing reusable frontend components.

---

## Ant Design Rules

Use Ant Design as the primary UI component library.

Prefer Ant Design for:

- `Layout`
- `Menu`
- `Card`
- `Button`
- `Form`
- `Input`
- `Select`
- `DatePicker`
- `Table`
- `Tag`
- `Badge`
- `Dropdown`
- `Modal`
- `Drawer`
- `Tabs`
- `Timeline`
- `Statistic`
- `Progress`
- `Alert`
- `Empty`
- `Skeleton`
- `App` message, modal, and notification context.

Do not build custom components from scratch when Ant Design already provides a good base component.

Do not use shadcn/ui, Material UI, Bootstrap, Chakra UI, or another UI component library unless the active task explicitly changes the frontend decision and updates `docs/DECISIONS.md`.

---

## Next.js Rules

If using Next.js App Router:

- Use `@ant-design/nextjs-registry` in `frontend/app/layout.tsx`.
- Put `ConfigProvider` in `frontend/app/providers.tsx`.
- `frontend/app/providers.tsx` must be a client component.
- Use `"use client"` only where needed.
- Keep server and client components separated where practical.
- Do not make every component a client component by default.

Required files for Ant Design setup:

```txt
frontend/app/layout.tsx
frontend/app/providers.tsx
frontend/config/antd-theme.ts
frontend/styles/glass.css
```

---

## Glass Style Rules

Use glass effect carefully.

Use glass style for:

- App shell.
- Sidebar.
- Topbar.
- Dashboard stat cards.
- Filter panels.
- Important summary cards.
- Empty state panels.

Do not overuse glass effect for:

- Every table cell.
- Every button.
- Every form field.
- Dense forms.
- Long text areas.
- Error or critical messages.

Glass style should improve hierarchy and depth, not reduce readability.

---

## Design Rules

Avoid:

- Plain default Ant Design look.
- Too many identical cards.
- Random gradients.
- Excessive shadows.
- Unclear contrast.
- Template-looking dashboard.
- Placeholder text like `Lorem ipsum`.
- Generic copy like `Manage your data`.
- One-off styling repeated across pages.

Prefer:

- Clear operational wording.
- Useful helper text.
- Status tags.
- Priority tags.
- Context under numbers.
- Empty, loading, and error states.
- Realistic seed or demo data when backend is not ready.
- Reusable components before full page implementation.

---

## Required Reusable Components

Create or reuse these components when relevant.

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
  MetricGroup

components/feedback/
  EmptyState
  LoadingState
  ErrorState

components/forms/
  FormSection
  FieldError

components/charts/
  ChartCard
  TicketStatusChart
  RobotRunChart
  TicketTrendChart
```

---

## Component Guidance

### GlassCard

- Wrap Ant Design `Card`.
- Apply `op-card-soft`, `op-glass`, or similar local class.
- Keep Ant Design `Card` API as much as possible.
- Avoid heavy custom styling inside feature pages.

### GlassStatCard

- Use for dashboard metrics.
- Must include label, value, and context.
- Optional trend or severity indicator is allowed.
- Avoid cards that only show label and number.

### StatusTag

- Wrap Ant Design `Tag`.
- Map OpsPilot statuses to consistent visual direction.
- Do not manually style status tags in every page.

### PriorityTag

- Wrap Ant Design `Tag`.
- Map `LOW`, `MEDIUM`, `HIGH`, and `CRITICAL` consistently.
- Do not manually style priority tags in every page.

### DataTable

- Wrap Ant Design `Table`.
- Include loading and empty state.
- Keep pagination behavior consistent.
- Prefer reusable filter integration.
- Avoid custom table implementations per module.

### PageHeader

- Include title, description, and optional primary action.
- Avoid generic page descriptions.
- The description should explain what the page helps the user do.

### EmptyState

- Use human, operational copy.
- Avoid `Data not found`.
- Good examples:
  - `No tickets need attention right now.`
  - `No robot run logs recorded yet.`
  - `No assets are expiring soon. Everything looks safe for now.`

---

## Frontend Task Flow

Before coding:

1. Read the active task.
2. Identify whether this is setup, layout, component, page, or integration work.
3. Check existing reusable components.
4. Summarize UX direction.
5. List components that will be created or reused.
6. Confirm whether client components are needed.

During coding:

1. Use Ant Design components first.
2. Apply OpsPilot Glass Style through wrapper components and CSS classes.
3. Keep feature logic under `features/*`.
4. Keep shared UI under `components/*`.
5. Keep theme values in `config/antd-theme.ts`.
6. Keep glass classes in `styles/glass.css`.
7. Do not implement out-of-scope features.
8. Do not introduce another UI library.

After coding:

1. Run relevant validation.
2. Update the active task status.
3. Update `backlog/STATUS.md`.
4. Update `docs/UI_SYSTEM.md` if reusable UI rules changed.
5. Update `docs/DECISIONS.md` if a frontend decision changed.
6. Summarize changed files and UX decisions.

---

## Hard Rules

- Do not add another UI library unless explicitly requested.
- Do not mix shadcn/ui with Ant Design.
- Do not override Ant Design internal class names globally unless unavoidable.
- Prefer theme tokens, wrapper components, and local CSS classes.
- Do not implement features outside the active task.
- Do not create full feature pages before base layout and reusable components exist.
- Do not leave placeholder copy like `Lorem ipsum`.
- Do not create many one-off components if a reusable component would fit.
- Do not scatter color, radius, shadow, or glass style values across feature pages.
- Do not make the UI look like plain default Ant Design.

---

## Final Response Format

After implementation, summarize:

- Ant Design components used.
- Reusable components created or reused.
- UX decisions made.
- Glass style usage.
- Any remaining mock data.
- Validation or checks run.
- Files updated.
