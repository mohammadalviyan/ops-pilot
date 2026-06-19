# TASK-004 — Database Migration Foundation and Seed Data

## Status

done

## Priority

P0

## Goal

Set up database connection, migration tooling, and initial seed strategy.

## Scope

- Add PostgreSQL connection.
- Add migration folder.
- Add initial migrations for users, roles, departments, applications, ticket categories.
- Add seed data for MVP demo users and master data.

## Out of Scope

- Do not build ticket module yet.
- Do not build auth endpoint yet, except password hash seed if needed.

## Acceptance Criteria

- Migrations can run locally.
- Seed users and roles exist.
- Passwords are hashed.
- README documents migration command.

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
