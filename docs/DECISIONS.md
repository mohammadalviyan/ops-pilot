# Decisions — OpsPilot

This file records technical and product decisions so AI agents do not re-litigate the same choices repeatedly.

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

**Status:** Proposed  
**Date:** 2026-06-19

### Decision

Use Tailwind CSS and shadcn/ui for UI components.

### Reason

This speeds up dashboard implementation while keeping a modern, clean UI.

### Guardrail

Do not over-customize components before core MVP flow works.

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
