# LawLens-G — AI Contract Clause Mapper & Risk Explorer

Upload a contract and LawLens-G turns it into an interactive map of clauses, risks, and obligations — with a Go backend and a Next.js + Tailwind front-end.

## Stack

- Backend: Go 1.22, Gin, GORM, Postgres
- Frontend: Next.js 14 (App Router), React 18, Tailwind CSS, Recharts
- Infra: Docker, docker-compose, Postgres 16

## Running locally

```bash
# from repo root
docker compose up --build
