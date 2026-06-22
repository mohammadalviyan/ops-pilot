# TASK-005 — Frontend Ant Design Foundation and App Shell

## Status

done

## Priority

P0

## Goal

Set up the frontend foundation for Next.js App Router, Ant Design, and OpsPilot Glass Style, then create the reusable app shell components needed before feature pages are built.

## Scope

- Install and document required frontend dependencies:
  - `antd`
  - `@ant-design/icons`
  - `@ant-design/nextjs-registry`
- Configure Next.js App Router Ant Design setup:
  - Use `AntdRegistry` in `frontend/app/layout.tsx`.
  - Add `frontend/app/providers.tsx` as a client component.
  - Wrap the app with Ant Design `ConfigProvider`.
  - Use Ant Design `App` for message, modal, and notification context.
- Add `frontend/config/antd-theme.ts` for Ant Design theme tokens.
- Add `frontend/styles/glass.css` for shared OpsPilot Glass Style classes.
- Create reusable layout components:
  - `frontend/components/layout/AppShell.tsx`
  - `frontend/components/layout/Sidebar.tsx`
  - `frontend/components/layout/Topbar.tsx`
  - `frontend/components/layout/PageContainer.tsx`
  - `frontend/components/layout/PageHeader.tsx`
- Create basic feedback components if reasonable for this foundation pass:
  - `frontend/components/feedback/EmptyState.tsx`
  - `frontend/components/feedback/LoadingState.tsx`
  - `frontend/components/feedback/ErrorState.tsx`
- Add a minimal `/dashboard` page shell only to verify layout rendering.
- Add a minimal `/login` placeholder only if needed to keep routes referenced by navigation from breaking.
- Keep styling aligned with `docs/UI_SYSTEM.md` and `.agents/skills/frontend-ant-design-glass/SKILL.md`.
- Update README setup commands if new frontend install or run commands are introduced.

## Out of Scope

- Do not implement dashboard data, metrics, charts, or real operational content.
- Do not implement authentication, protected routes, sessions, JWT handling, or login submission.
- Do not implement API integration or a shared API client.
- Do not implement ticket CRUD, robot monitoring, asset monitoring, reports, or knowledge base pages.
- Do not add chart libraries in this task.
- Do not use Tailwind, shadcn/ui, Material UI, Bootstrap, Chakra UI, or another primary UI library.
- Do not create advanced glass/data-display/form/chart components unless required to make the shell coherent.

## Acceptance Criteria

- Frontend dependencies for Ant Design and Next.js App Router registry are installed.
- Next.js App Router renders Ant Design without hydration/style issues.
- `frontend/app/layout.tsx` imports global styles and `frontend/styles/glass.css`.
- `frontend/app/layout.tsx` uses `AntdRegistry`.
- `frontend/app/providers.tsx` uses `ConfigProvider` and Ant Design `App`.
- `frontend/config/antd-theme.ts` centralizes primary theme tokens.
- `frontend/styles/glass.css` centralizes shared glass style classes.
- `AppShell`, `Sidebar`, `Topbar`, `PageContainer`, and `PageHeader` exist and are reusable.
- Basic `EmptyState`, `LoadingState`, and `ErrorState` exist if included in scope.
- `/dashboard` displays inside the app shell with no real data integration.
- `/login` placeholder exists if navigation or routing needs it.
- The frontend runs locally.
- Relevant lint/build/type checks run or skipped with a clear reason.
- README is updated if setup commands changed.

## Validation Plan

- Run the relevant frontend install command.
- Run the frontend development server or build command to verify the app starts.
- Run lint/type/build checks available in the frontend package.
- Manually verify `/dashboard` renders the app shell.
- Manually verify Ant Design styles and OpsPilot glass CSS are loaded.
- Confirm no API calls, auth logic, or dashboard data were added.

## Validation Results

- `npm install` completed and generated `frontend/package-lock.json`.
- `npm run lint` passed.
- `npm run typecheck` passed.
- `npm run build` passed.
- `npm run dev -- --hostname 127.0.0.1 --port 3000` started successfully with approval for localhost binding.
- `curl -I http://127.0.0.1:3000/dashboard` returned `200 OK`.
- Confirmed the task adds only shell, theme, glass style, and feedback foundation; no auth logic, dashboard business data, or backend API integration was added.

## Relevant Docs

- `AGENTS.md`
- `docs/ARCHITECTURE.md`
- `docs/DECISIONS.md`
- `docs/UI_SYSTEM.md`
- `.agents/skills/frontend-ant-design-glass/SKILL.md`
- `backlog/STATUS.md`

## Agent Instructions

- Work only on this task.
- Before coding, summarize the task goal, affected files, UX direction, and validation plan.
- Keep the implementation small and reviewable.
- Use Ant Design components first.
- Keep theme values in `frontend/config/antd-theme.ts`.
- Keep shared glass classes in `frontend/styles/glass.css`.
- Do not implement out-of-scope features.
- After coding, run relevant checks.
- Update this task status.
- Update `backlog/STATUS.md`.
- Update `backlog/BACKLOG.md` if task status changes.
- Update `docs/UI_SYSTEM.md` only if reusable UI rules changed.
- Update `docs/DECISIONS.md` only if a technical decision changed.
