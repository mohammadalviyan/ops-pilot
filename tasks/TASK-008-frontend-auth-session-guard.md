# TASK-008 — Frontend Auth Session and Route Guard

## Status

todo

## Priority

P0

## Goal

Implement frontend auth session handling and route protection for the OpsPilot dashboard shell using the login result from TASK-007 and the backend `GET /api/v1/auth/me` endpoint.

## Scope

- Add frontend auth session utilities for reading, writing, and clearing the MVP access token and user profile.
- Add an auth provider or route-level client guard appropriate for the current Next.js App Router setup.
- Call `GET /api/v1/auth/me` to validate the current bearer token.
- Protect dashboard app routes such as `/dashboard`.
- Redirect unauthenticated users from protected routes to `/login`.
- Redirect authenticated users away from `/login` to `/dashboard`.
- Add logout behavior that calls `POST /api/v1/auth/logout`, clears local frontend session state, and returns the user to `/login`.
- Update `Topbar` to show the authenticated user name/role and provide a logout action.
- Keep route guard behavior simple and MVP-focused.
- Use existing Ant Design and OpsPilot Glass Style components for loading/error states.

## Out of Scope

- Do not implement backend auth changes.
- Do not implement role-based permission or per-page authorization.
- Do not implement refresh tokens.
- Do not implement httpOnly cookie auth unless explicitly chosen and documented as part of this task.
- Do not implement ticketing, dashboard data, robot monitoring, reports, or other feature pages.
- Do not add another UI library.
- Do not redesign the app shell beyond the minimum needed for user identity and logout.

## Acceptance Criteria

- Visiting `/dashboard` without a valid session redirects to `/login`.
- Visiting `/login` with a valid session redirects to `/dashboard`.
- A stored token is validated with `GET /api/v1/auth/me`.
- Expired or invalid tokens are cleared and treated as unauthenticated.
- Authenticated user identity is available to the app shell.
- `Topbar` displays the authenticated user's name or email and role.
- Logout clears frontend session state and returns the user to `/login`.
- Logout calls `POST /api/v1/auth/logout` when possible.
- Protected route loading state is clear and calm.
- No role-based permissions are implemented.
- No backend API contract changes are required.

## Affected Files

- `frontend/app/login/page.tsx`
- `frontend/app/dashboard/page.tsx`
- `frontend/app/providers.tsx` if an auth provider is needed
- `frontend/components/layout/AppShell.tsx`
- `frontend/components/layout/Topbar.tsx`
- `frontend/features/auth/`
- `frontend/lib/api.ts`
- `frontend/lib/auth.ts`
- `frontend/hooks/` if auth hooks are introduced
- `frontend/types/auth.ts`
- `README.md` only if setup commands or environment instructions change
- `tasks/TASK-008-frontend-auth-session-guard.md`
- `backlog/STATUS.md`
- `backlog/BACKLOG.md`

## Validation Steps

- Run `npm run lint` from `frontend/`.
- Run `npm run typecheck` from `frontend/`.
- Run `npm run build` from `frontend/`.
- Start backend API with PostgreSQL and `JWT_SECRET`.
- Start frontend dev server.
- Verify unauthenticated `/dashboard` redirects to `/login`.
- Verify valid login from TASK-007 reaches `/dashboard`.
- Verify refresh on `/dashboard` keeps the user authenticated when the token is valid.
- Verify invalid or manually removed token redirects to `/login`.
- Verify logout returns to `/login` and clears session state.
- Verify no role-based permission behavior was added.

## Relevant Docs

- `AGENTS.md`
- `docs/PRD.md`
- `docs/ARCHITECTURE.md`
- `docs/API_CONTRACT.md`
- `docs/UI_SYSTEM.md`
- `docs/DECISIONS.md`
- `backlog/STATUS.md`

## Agent Instructions

- Work only on this task.
- Before coding, summarize the task scope, affected files, session strategy, and validation plan.
- Use Next.js App Router, Ant Design, and OpsPilot Glass Style.
- Reuse auth client/service code created in TASK-007.
- Keep implementation small and reviewable.
- Do not implement role-based permissions or non-auth feature work.
- After coding, run relevant checks.
- Update this task status.
- Update `backlog/STATUS.md`.
- Update `backlog/BACKLOG.md` if task status changes.
- Update `docs/UI_SYSTEM.md` only if reusable UI rules change.
- Update `docs/DECISIONS.md` only if a new technical decision is made.
