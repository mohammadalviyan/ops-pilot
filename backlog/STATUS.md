# Status — OpsPilot

## Current Phase

Foundation

## Active Task

None

## Last Completed Task

`tasks/TASK-004-db-migrations-seed.md`

## Next Recommended Task

`tasks/TASK-005-frontend-shell-layout.md`

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

## Open Questions

- Confirm whether auth token will be stored via httpOnly cookie or temporary local storage for MVP.
- Confirm whether Next.js should use App Router or Pages Router. Current proposal: App Router.

## Blockers

None.

## Recent Decisions

- Use one task per agent session.
- Keep future integrations out of MVP.
- Use monorepo with backend and frontend folders.
- Use Gin for the backend HTTP router.
