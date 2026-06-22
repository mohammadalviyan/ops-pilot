# Codex Prompts — OpsPilot

Use these prompts to run OpsPilot development with Codex using a task-based workflow.

Core rule:

```txt
1 task = 1 agent session = 1 small reviewable diff
```

Do not ask Codex to implement multiple unrelated tasks in one session.

---

## 1. Initial Repo Setup

Use this only for the first project foundation task.

```txt
Read:
- AGENTS.md
- docs/PRD.md
- docs/ARCHITECTURE.md
- docs/API_CONTRACT.md
- docs/DECISIONS.md
- backlog/BACKLOG.md
- backlog/STATUS.md

Execute only:
tasks/TASK-001-project-foundation.md

Before coding:
- Summarize the task scope.
- Provide a short implementation plan.
- List files that will be created or changed.

Rules:
- Do not implement business features yet.
- Do not implement auth, ticketing, dashboard data, robot monitoring, or reporting yet.
- Only prepare the initial repo foundation.
- Follow AGENTS.md.
- Keep changes small and reviewable.

After coding:
- Run relevant checks if available.
- Update tasks/TASK-001-project-foundation.md status.
- Update backlog/STATUS.md.
- Summarize completed work, changed files, validation result, and next recommended task.
```

---

## 2. Execute One Task

Use this as the default prompt for any normal task.

Replace `tasks/TASK-XXX-name.md` with the active task file.

```txt
Read:
- AGENTS.md
- docs/PRD.md
- docs/ARCHITECTURE.md
- docs/API_CONTRACT.md
- docs/DB_SCHEMA.md
- docs/DECISIONS.md
- backlog/BACKLOG.md
- backlog/STATUS.md
- active task file below

Execute only:
tasks/TASK-XXX-name.md

Before coding:
- Summarize the active task scope.
- Identify whether this is backend, frontend, database, documentation, or QA work.
- Provide a short implementation plan.
- List files that will be created or changed.
- Mention any risks, assumptions, or unclear requirements before editing files.

Rules:
- Work on one task only.
- Do not implement out-of-scope items.
- Do not implement future features unless the active task explicitly requires it.
- Do not change architecture decisions unless necessary.
- If a decision changes, update docs/DECISIONS.md.
- If API request/response changes, update docs/API_CONTRACT.md.
- If database schema changes, update docs/DB_SCHEMA.md and migrations.
- If setup commands or environment variables change, update README.md and .env.example.
- If task status changes, update the active task file and backlog/STATUS.md.
- Keep changes small and reviewable.
- Reuse existing project patterns before creating new ones.

Validation:
- Run relevant checks for the changed area.
- If checks cannot be run, explain why and provide the exact command the developer should run.

After coding:
- Update the active task status.
- Update backlog/STATUS.md.
- Summarize completed work.
- List changed files.
- Summarize validation result.
- Recommend the next task.
```

---

## 3. Execute Backend Task

Use this for backend-focused tasks such as auth, ticket API, repository, service, or migration integration.

Replace `tasks/TASK-XXX-name.md` with the active backend task file.

```txt
Read:
- AGENTS.md
- docs/PRD.md
- docs/ARCHITECTURE.md
- docs/API_CONTRACT.md
- docs/DB_SCHEMA.md
- docs/DECISIONS.md
- .agents/skills/backend-go/SKILL.md if it exists
- backlog/STATUS.md
- active task file below

Execute only:
tasks/TASK-XXX-name.md

Backend rules:
- Use Go + Gin + PostgreSQL.
- Follow the existing backend structure.
- Keep handler, service, repository, DTO, and model responsibilities separated.
- Use the standard API response format.
- Use environment variables for config.
- Use migrations for schema changes.
- Use bcrypt for password hashing when handling passwords.
- Use JWT only where auth task requires it.
- Do not implement frontend changes unless explicitly required by the task.
- Do not implement ticketing, dashboard, robot, asset, or report features unless the active task requires it.

Before coding:
- Summarize task scope.
- Provide implementation plan.
- List affected backend files.
- Mention API and database docs that may need updates.

After coding:
- Run relevant backend checks.
- Update docs/API_CONTRACT.md if endpoint contract changes.
- Update docs/DB_SCHEMA.md if schema changes.
- Update docs/DECISIONS.md if a technical decision changes.
- Update the active task status.
- Update backlog/STATUS.md.
- Summarize changed files and validation result.
```

---

## 4. Execute Frontend Task

Use this for frontend-focused tasks such as Ant Design setup, layout, login page, dashboard shell, ticket list, forms, tables, and UI states.

Replace `tasks/TASK-XXX-name.md` with the active frontend task file.

```txt
Read:
- AGENTS.md
- docs/PRD.md
- docs/ARCHITECTURE.md
- docs/API_CONTRACT.md
- docs/DECISIONS.md
- docs/UI_SYSTEM.md
- .agents/skills/frontend-ant-design-glass/SKILL.md
- backlog/STATUS.md
- active task file below

Execute only:
tasks/TASK-XXX-name.md

Frontend rules:
- Use Next.js + React + TypeScript.
- Use Ant Design as the primary UI component library.
- Use OpsPilot Glass Style.
- Do not use shadcn/ui, Material UI, Bootstrap, Chakra UI, or another UI library.
- Reuse existing components before creating new ones.
- Use Ant Design components as the base.
- Use wrapper components for consistent OpsPilot UI.
- Keep shared components under frontend/components.
- Keep feature-specific logic under frontend/features.
- Keep theme tokens in frontend/config/antd-theme.ts.
- Keep shared glass styles in frontend/styles/glass.css.
- Do not implement backend logic.
- Do not implement features outside the active task.

UX rules:
- Avoid generic AI-looking UI.
- Avoid plain default Ant Design look.
- Every visible page must have meaningful loading, empty, and error states when relevant.
- Use human operational copy.
- Use StatusTag and PriorityTag for status and priority display.
- Avoid placeholder copy like Lorem ipsum or Manage your data.

Before coding:
- Summarize task scope.
- Summarize UX direction.
- List reusable components to create or reuse.
- Confirm whether client components are needed.
- List files that will be created or changed.

After coding:
- Run relevant frontend checks.
- Update docs/UI_SYSTEM.md only if UI rules or reusable component strategy changed.
- Update docs/DECISIONS.md only if a frontend decision changed.
- Update the active task status.
- Update backlog/STATUS.md.
- Summarize changed files, UX decisions, and validation result.
```

---

## 5. Review Latest Diff

Use this after Codex finishes a task and before committing or merging.

```txt
Review the latest diff against the active task acceptance criteria.

Read:
- AGENTS.md
- active task file
- backlog/STATUS.md
- docs/ARCHITECTURE.md
- docs/API_CONTRACT.md
- docs/DB_SCHEMA.md if database files changed
- docs/UI_SYSTEM.md if frontend files changed
- docs/DECISIONS.md if architecture or library decisions changed

Check:
- Correctness.
- Scope creep.
- Missing validation.
- API contract mismatch.
- Database schema or migration mismatch.
- UI loading, error, and empty states.
- Ant Design Glass consistency for frontend changes.
- Missing README updates.
- Missing task/status updates.
- Unnecessary dependencies.
- Repeated code that should be reusable.
- Security issues such as plain text password, exposed secrets, or unsafe JWT handling.

Do not implement changes yet.

Return findings grouped by:
- Critical
- Major
- Minor
- Nice to have

For each finding, include:
- File or area.
- Problem.
- Why it matters.
- Suggested fix.
```

---

## 6. Fix Review Findings

Use this only after reviewing latest diff and deciding which findings should be fixed.

```txt
Read:
- AGENTS.md
- active task file
- latest review findings
- relevant docs for the changed area

Fix only the approved findings from the review.

Rules:
- Do not add new features.
- Do not refactor unrelated code.
- Do not change task scope.
- Keep changes minimal.
- Update documentation only if the fix changes API, database, UI rules, setup, or decisions.

After fixing:
- Run relevant checks.
- Update the active task if needed.
- Update backlog/STATUS.md if needed.
- Summarize fixes and validation result.
```

---

## 7. Generate More Task Files

Use this when the next epic needs task files, or when there is a missing task number.

```txt
Read:
- AGENTS.md
- docs/PRD.md
- docs/ARCHITECTURE.md
- docs/API_CONTRACT.md
- docs/DB_SCHEMA.md
- docs/UI_SYSTEM.md
- docs/DECISIONS.md
- backlog/BACKLOG.md
- backlog/STATUS.md
- existing tasks in /tasks

Generate missing task files for the next epic only.

Rules:
- Do not implement code.
- Do not modify source files.
- Generate only task markdown files.
- Keep each task small and reviewable.
- Each task must represent one clear unit of work.
- Follow the current architecture decisions.
- For frontend tasks, use Next.js + Ant Design + OpsPilot Glass Style.
- For backend tasks, use Go + Gin + PostgreSQL.
- Do not generate tasks for out-of-scope MVP features unless explicitly requested.

Each task must include:
- Title.
- Status.
- Priority.
- Goal.
- Scope.
- Out of scope.
- Acceptance criteria.
- Relevant docs.
- Expected files or areas.
- Validation steps.
- Agent instructions.
- Notes if any.

After generating:
- Update backlog/BACKLOG.md if needed.
- Update backlog/STATUS.md if needed.
- Summarize the generated tasks and recommended execution order.
```

---

## 8. Update One Task Before Execution

Use this when a task exists but is outdated, for example when it still mentions Tailwind/shadcn while the project already uses Ant Design.

```txt
Read:
- AGENTS.md
- docs/ARCHITECTURE.md
- docs/DECISIONS.md
- docs/UI_SYSTEM.md
- backlog/STATUS.md
- active task file below

Update only:
tasks/TASK-XXX-name.md

Goal:
Make the task align with the latest architecture and decisions.

Rules:
- Do not code.
- Do not modify source files.
- Do not change unrelated task files.
- Keep the task small and executable.
- Make scope, out-of-scope, acceptance criteria, expected files, and validation steps clear.
- For frontend tasks, align with Next.js + Ant Design + OpsPilot Glass Style.
- For backend tasks, align with Go + Gin + PostgreSQL.
- Update backlog/STATUS.md only if task status or ordering changes.

After updating:
- Summarize what changed in the task file.
```

---

## 9. Antigravity UI Validation

Use this after a frontend task is implemented.

```txt
Run the app locally if possible and validate the current frontend task.

Read:
- docs/UI_SYSTEM.md
- docs/ARCHITECTURE.md
- active task file
- backlog/STATUS.md

Check:
- Page loads.
- Layout consistency.
- Ant Design Glass Style consistency.
- Sidebar navigation.
- Topbar behavior.
- Page header clarity.
- Form states.
- Table states.
- Chart visibility if relevant.
- Loading state.
- Empty state.
- Error state.
- Responsive desktop behavior.
- Responsive tablet behavior.
- Copywriting feels operational and not generic.

Rules:
- Do not implement unrelated fixes.
- Do not change backend files.
- Capture findings clearly.
- Update the active task with QA notes if appropriate.

Return:
- Passed checks.
- Issues found.
- Suggested fixes.
- Screenshots or walkthrough notes if available.
```

---

## 10. Start Next Task Checklist

Use this before starting a new task if you are unsure whether the repo is ready.

```txt
Read:
- backlog/STATUS.md
- backlog/BACKLOG.md
- tasks directory
- docs/DECISIONS.md

Tell me:
- What task should be executed next?
- Is the previous task marked done?
- Are there missing task numbers?
- Are there docs that should be updated before coding?
- Which branch name should be used?
- Which prompt from prompts/CODEX_PROMPTS.md should be used?

Do not code.
Do not modify files.
```
