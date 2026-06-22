# OpsPilot

OpsPilot is an internal RPA operations dashboard for ticketing, robot monitoring, asset monitoring, and operational reporting.

This repository is prepared as a monorepo with separate backend and frontend applications. Feature implementation is driven by one task file at a time from `tasks/`.

## Project Structure

```txt
opspilot/
  backend/            # Go API application
  frontend/           # Next.js / React application
  docs/               # Product, architecture, API, and decision docs
  backlog/            # Backlog and project status
  tasks/              # One executable task file per work item
  prompts/            # Agent prompt references
  AGENTS.md           # Agent working rules
  README.md
  docker-compose.yml
```

## Initial Setup

The current foundation prepares the repository layout, environment placeholders, local PostgreSQL dependency, backend health API, and frontend app shell.

1. Copy backend environment defaults when backend implementation starts:

   ```sh
   cp backend/.env.example backend/.env
   ```

2. Copy frontend environment defaults when frontend implementation starts:

   ```sh
   cp frontend/.env.example frontend/.env.local
   ```

## Local Dependencies

Start PostgreSQL:

```sh
docker compose up -d postgres
```

Check service status:

```sh
docker compose ps
```

Stop local dependencies:

```sh
docker compose down
```

The PostgreSQL credentials match `backend/.env.example`:

```txt
database: opspilot
user: opspilot
password: opspilot
port: 5432
```

## Backend

Run the backend API:

```sh
cd backend
APP_ENV=local APP_PORT=8080 go run ./cmd/api
```

Health check:

```sh
curl http://localhost:8080/api/v1/health
```

Run database migrations:

```sh
cd backend
DATABASE_URL="postgres://opspilot:opspilot@localhost:5432/opspilot?sslmode=disable" go run ./cmd/migrate -direction=up
```

Rollback the latest migration:

```sh
cd backend
DATABASE_URL="postgres://opspilot:opspilot@localhost:5432/opspilot?sslmode=disable" go run ./cmd/migrate -direction=down
```

## Frontend

Install frontend dependencies:

```sh
cd frontend
npm install
```

Run the frontend app:

```sh
cd frontend
npm run dev
```

Frontend checks:

```sh
cd frontend
npm run lint
npm run typecheck
npm run build
```

## Working Rules

- Read `AGENTS.md` before making changes.
- Execute exactly one task from `tasks/` per session.
- Keep changes inside the active task scope.
- Update the active task status and `backlog/STATUS.md` after completing work.
- Record technical decisions in `docs/DECISIONS.md` when a new decision is made.
