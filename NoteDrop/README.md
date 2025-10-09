# 📝 NoteDrop – Public/Private Notes API

NoteDrop is a full-stack-ready **backend API** for securely creating, viewing, and sharing notes with optional **public slugs**, **JWT Auth**, **RBAC**, and **Dockerized PostgreSQL**.

Built with:
- ✅ Express.js + JavaScript
- ✅ Prisma ORM + PostgreSQL
- ✅ JWT Auth + Role Middleware
- ✅ Docker & Docker Compose
- ✅ Curl-ready endpoints

---

## 🚀 Features

- Register/Login users with JWT auth
- Create personal or public notes with slugs
- Auth-protected `GET /me` routes
- Publicly accessible notes by slug (no login)
- Dockerized PostgreSQL + Prisma Migrate
- Clean environment-based config

---

## 🛠 Project Setup

### 1. 🔧 Environment File

Create a `.env` file in the root:

```env
DATABASE_URL=postgresql://postgres:postgres@localhost:5433/notedrop?schema=public
JWT_SECRET=supersecretkey
PORT=3600
```

---

### 2. 🐳 Start with Docker

```bash
# Build and run containers
docker-compose up --build

# (Optional) Stop and remove containers
docker-compose down -v
```

This spins up:

| Service   | Port  | Description           |
|-----------|-------|------------------------|
| API       | 3600  | Express + JWT API      |
| Postgres  | 5433  | Dockerized Postgres DB |

---

### 3. 🧪 Test Auth + Notes

#### ✅ Register

```bash
curl -X POST http://localhost:3600/api/auth/register   -H "Content-Type: application/json"   -d '{"email": "user@mail.com", "password": "123456"}'
```

#### ✅ Login

```bash
curl -X POST http://localhost:3600/api/auth/login   -H "Content-Type: application/json"   -d '{"email": "user@mail.com", "password": "123456"}'
```

Copy the returned `token`.

#### 🛡️ Authenticated: Create Note

```bash
curl -X POST http://localhost:3600/api/notes   -H "Content-Type: application/json"   -H "Authorization: Bearer YOUR_TOKEN_HERE"   -d '{"title": "Launch Plan", "content": "This is secret", "isPublic": true}'
```

#### 🌍 Public Access (no auth)

```bash
curl http://localhost:3600/api/notes/launch-plan-bq9s8c
```

(replace with real slug)

---

## 📂 Folder Structure

```
notedrop/
├── src/                  # App source (routes, controllers, middleware)
├── prisma/               # Prisma schema & migrations
├── Dockerfile            # App container
├── docker-compose.yml    # Service orchestrator
├── .env                  # Env variables
└── package.json          # Node scripts
```

---

## 📜 Scripts

```bash
# Start dev server (no Docker)
npm run dev

# Format code
npm run format

# Prisma CLI
npx prisma studio       # GUI explorer
npx prisma migrate dev  # Re-run migration

# Run inside Docker (auto)
docker-compose up --build
```

---

## 📅 Generated on
October 09, 2025

---

> Built by Eddy for backend drills 🧠⚔️
