# Status — OpsPilot

## Current Phase

Auth

## Active Task

None

## Last Completed Task

`tasks/TASK-007-frontend-login-page.md`

## Next Recommended Task

`tasks/TASK-008-frontend-auth-session-guard.md`

## Progress Summary

- PRD exists.
- AI-ready documentation pack prepared.
- Architecture, API contract, decisions, backlog, task files, and skills are prepared as starting point.
- Initial monorepo foundation has been created with `backend/` and `frontend/` folders.
- Root `.gitignore`, environment examples, and README project structure notes are in place.
- Docker Compose local dependency setup has been added for PostgreSQL.
- Backend database environment example matches the Docker Compose PostgreSQL credentials.
- Backend Go module, Gin router, config loader, standard response package, and `GET /api/v1/health` endpoint are in place.
- Database connection helper, migration runner, master table migrations, and MVP seed data are in place.
- Local migrations were verified against PostgreSQL with seed rows for roles, users, departments, applications, and ticket categories.
- Frontend Next.js App Router foundation, Ant Design registry/provider setup, theme tokens, OpsPilot glass CSS, reusable app shell components, and basic feedback states are in place.
- `/dashboard` and `/login` placeholder routes render without auth, API integration, or dashboard business data.
- Backend auth login, JWT generation/verification, auth middleware, `/auth/me`, and logout response are in place.
- Seed user passwords were migrated to a bcrypt hash that accepts the documented demo password.
- Frontend login page is connected to `POST /api/v1/auth/login`, stores the temporary MVP auth session, and redirects to `/dashboard` after successful login.

## Open Questions

- Confirm when MVP auth should move from temporary local storage to httpOnly cookie storage.

## Blockers

None.

## Recent Decisions

- Use one task per agent session.
- Keep future integrations out of MVP.
- Use monorepo with backend and frontend folders.
- Use Gin for the backend HTTP router.
- Use Next.js App Router with Ant Design and OpsPilot Glass Style for frontend foundation.
- Use temporary local storage for MVP frontend auth state until cookie auth is implemented.
