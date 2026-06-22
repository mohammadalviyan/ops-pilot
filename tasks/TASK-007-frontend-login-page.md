# TASK-007 — Frontend Login Page

## Status

todo

## Priority

P0

## Goal

Implement the OpsPilot frontend login page using Next.js App Router, Ant Design, and OpsPilot Glass Style, connected to the existing backend `POST /api/v1/auth/login` endpoint.

## Scope

- Replace the current `/login` placeholder with a real login page.
- Build a login form using Ant Design `Form`, `Input`, `Button`, `Alert`, and related feedback components.
- Use OpsPilot Glass Style for the login surface without creating a generic marketing page.
- Add a frontend auth feature structure for login-only concerns.
- Add or reuse a shared frontend API client for calling the backend auth endpoint.
- Submit email and password to `POST /api/v1/auth/login`.
- Handle loading, validation, and error states.
- Store the returned access token and user profile using the temporary MVP auth storage approach chosen for this task.
- Redirect to `/dashboard` after successful login.
- Keep copy operational and aligned with OpsPilot.
- Update README only if new setup or environment commands are introduced.

## Out of Scope

- Do not implement route protection or auth guards; that belongs to TASK-008.
- Do not implement role-based permission checks.
- Do not implement refresh tokens.
- Do not implement password reset, registration, SSO, MFA, or profile management.
- Do not implement backend changes.
- Do not implement dashboard data or ticketing features.
- Do not add another UI library.

## Acceptance Criteria

- `/login` renders a polished OpsPilot login form.
- The login page uses Ant Design components and existing OpsPilot Glass Style classes.
- The form validates required email and password fields before submitting.
- The form calls `POST /api/v1/auth/login` through a shared API client or auth service.
- Successful login stores the returned `access_token` and user profile for MVP frontend use.
- Successful login redirects to `/dashboard`.
- Invalid credentials show a clear error state without crashing the page.
- Submit button shows loading state while the request is in progress.
- No plaintext password is stored anywhere in frontend state or storage.
- No protected-route behavior is implemented in this task.
- No backend API contract changes are required.

## Affected Files

- `frontend/app/login/page.tsx`
- `frontend/features/auth/`
- `frontend/lib/api.ts`
- `frontend/lib/auth.ts`
- `frontend/types/api.ts`
- `frontend/types/auth.ts`
- `frontend/components/feedback/*` if existing feedback components need small reuse-friendly adjustments
- `README.md` only if setup commands or environment instructions change
- `tasks/TASK-007-frontend-login-page.md`
- `backlog/STATUS.md`
- `backlog/BACKLOG.md`

## Validation Steps

- Run `npm run lint` from `frontend/`.
- Run `npm run typecheck` from `frontend/`.
- Run `npm run build` from `frontend/`.
- Start backend API with PostgreSQL and `JWT_SECRET`.
- Start frontend dev server.
- Verify `/login` renders correctly.
- Verify valid seed login works with `support@opspilot.local` and `password`.
- Verify invalid credentials show an error message.
- Verify successful login redirects to `/dashboard`.
- Verify no route guard behavior was added yet.

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
- Before coding, summarize the task scope, UX direction, affected files, and validation plan.
- Use Next.js App Router, Ant Design, and OpsPilot Glass Style.
- Keep API calls behind shared frontend client/service code.
- Keep implementation small and reviewable.
- Do not implement auth guards or protected routes.
- After coding, run relevant checks.
- Update this task status.
- Update `backlog/STATUS.md`.
- Update `backlog/BACKLOG.md` if task status changes.
- Update `docs/UI_SYSTEM.md` only if reusable UI rules change.
- Update `docs/DECISIONS.md` only if a new technical decision is made.
