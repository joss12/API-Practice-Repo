
# PortLog 🔌
Simple port-scanning API built with Go + Fiber + Redis

## 🔧 Features
- Scan common TCP ports (22, 80, 443, etc.)
- Caches results with Redis for 10 minutes
- Protects access via API Key
- Runs in Docker with Redis container

## 🚀 Run Locally

```bash
make run

docker-compose down 
docker-compose up --build


docker-compose down -v --remove-orphans
docker-compose build --no-cache
docker-compose up


