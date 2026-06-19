# TASK-003 — Backend Base API and Health Check

## Status

done

## Priority

P0

## Goal

Initialize Go backend with config, router, logger, standard response, and health endpoint.

## Scope

- Initialize Go module.
- Add backend folder structure.
- Implement config loader.
- Implement HTTP server.
- Implement `GET /api/v1/health`.
- Implement standard response package.

## Out of Scope

- Do not implement auth.
- Do not implement ticket API.
- Do not connect frontend yet.

## Acceptance Criteria

- Backend starts locally.
- `GET /api/v1/health` returns standard success response.
- Config comes from environment variable.
- README contains backend run command.

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
