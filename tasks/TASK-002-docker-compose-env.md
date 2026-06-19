# TASK-002 — Docker Compose and Environment Setup

## Status

done

## Priority

P0

## Goal

Create local development runtime using Docker Compose.

## Scope

- Add `docker-compose.yml` with PostgreSQL service.
- Add optional pgAdmin only if simple.
- Configure backend and frontend environment examples.
- Document how to start local dependencies.

## Out of Scope

- Do not implement application features.
- Do not add Redis unless needed by active task.

## Acceptance Criteria

- `docker compose up` starts PostgreSQL.
- Database credentials match backend `.env.example`.
- README contains local setup command.

## Relevant Docs

- `AGENTS.md`
- `docs/PRD.md`
- `docs/ARCHITECTURE.md`
- `docs/API_CONTRACT.md`
- `docs/DECISIONS.md`
- `backlog/STATUS.md`

## Agent Instructions

- Work only on this task.
- Before coding, write a short implementation plan.
- Do not implement out-of-scope items.
- After coding, run relevant checks.
- Update this task status.
- Update `backlog/STATUS.md`.
