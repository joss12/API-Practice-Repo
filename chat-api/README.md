# Chat API -- Real-Time Multi-Room Chat Backend (Go + Fiber + Redis + PostgreSQL)

A **production-grade real-time chat backend** built with **Go**,
**Fiber**, **WebSockets**, **Redis Pub/Sub**, **PostgreSQL**,
**Prometheus**, and **Nginx**.\
Fully tested, scalable, multi-room, and portfolio-ready.

------------------------------------------------------------------------

## ✨ Features

### 🔥 Real-Time Chat Engine

-   Multi-room WebSocket chat
-   JWT-secured WebSocket handshake
-   Distributed messaging via Redis Pub/Sub
-   Presence system (online users)
-   Typing indicators
-   Delivered receipts (ACK)
-   Seen receipts (read)
-   Room creation API
-   Admin control panel

### 🗄 Persistence + Caching

-   PostgreSQL for messages
-   Redis for:
    -   Presence
    -   Online users
    -   Active rooms
    -   Ban/mute lists

### 📡 Observability

-   Built-in Prometheus metrics:
    -   Active WS connections
    -   Messages per room
    -   Redis latency
    -   Database latency
    -   HTTP request count

------------------------------------------------------------------------

## 🧱 Architecture Diagram

              ┌───────────────┐
              │    Client      │
              └───────┬───────┘
                      │ WebSocket + REST
                      ▼
               ┌───────────────┐
               │     Nginx      │
               └───────┬───────┘
                      ▼
            ┌──────────────────────┐
            │       Fiber API       │
            │  (Go WebSocket Hub)   │
            └─────────┬────────────┘
       ┌──────────────┴──────────────┐
       ▼                             ▼
    ┌─────────┐               ┌─────────────┐
    │ Redis   │               │ PostgreSQL  │
    │ Pub/Sub │               │ Message DB  │
    └─────────┘               └─────────────┘

------------------------------------------------------------------------

## 📁 File Tree

    chat-api/
    │
    ├── cmd/server
    │   └── main.go
    │
    ├── internal/
    │   ├── app/
    │   ├── config/
    │   ├── db/
    │   ├── handlers/
    │   ├── middleware/
    │   ├── metrics/
    │   ├── models/
    │   ├── redis/
    │   └── ws/
    │
    ├── nginx/
    │   └── nginx.conf
    │
    ├── docker-compose.yml
    ├── Dockerfile
    └── README.md

------------------------------------------------------------------------

## 🔑 REST API Overview

### **Auth**

    POST /auth/register
    POST /auth/login

### **Rooms**

    POST /rooms
    GET  /chat/history?room=name&limit=50

### **Admin**

    GET /admin/rooms
    GET /admin/users
    POST /admin/ban/:id
    POST /admin/mute/:id
    POST /admin/unmute/:id

------------------------------------------------------------------------

## 🔌 WebSocket Event Specification

### Send Message

``` json
{ "type": "message", "body": "hello" }
```

### Delivered ACK

``` json
{ "type": "ack", "messageID": 12 }
```

### Seen

``` json
{ "type": "seen", "messageID": 12 }
```

### Typing

``` json
{ "type": "typing", "isTyping": true }
```

------------------------------------------------------------------------

## 🐳 Docker Usage

### Run the stack

    make up

### View logs

    make logs

### Stop everything

    make down

------------------------------------------------------------------------

## 📊 Metrics

Prometheus endpoint:

    GET /metrics

Includes: - `chat_ws_active_connections` - `chat_messages_total` -
`chat_redis_latency_seconds` - `chat_db_latency_seconds` -
`chat_http_requests_total`

------------------------------------------------------------------------

## 🛠 Tech Stack

-   **Go 1.22+**
-   **Fiber**
-   **WebSockets**
-   **Redis**
-   **PostgreSQL**
-   **Docker**
-   **Nginx**
-   **Prometheus**

------------------------------------------------------------------------

## 🧪 Testing

Run all tests:

    go test ./... -v

Everything is already tested: - Handlers - WebSocket events - Presence
system - Redis integration - Message repo

------------------------------------------------------------------------

## 🎯 Production-Ready

This backend can support: - Scalable chat apps - Customer support
dashboards - Real-time collaboration tools - Gaming chat systems -
Multi-instance deployments

------------------------------------------------------------------------

## 🙌 Author

Built by **Eddy** with the goal of mastering full-stack real-time
backend engineering.

