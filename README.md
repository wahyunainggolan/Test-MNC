# Wallet API with Background Worker and Dashboard

A RESTful API for wallet transactions using Golang, Gin, PostgreSQL, and Redis-based background workers.

## Features
- User Registration, Login (with JWT)
- Top-up, Payment, and Transfer
- Transaction Reports
- Background Job Processing for Transfers using Asynq
- Monitoring Dashboard via Asynqmon
- Auto-migration on server start

---

## Prerequisites

- Go 1.20+
- PostgreSQL (e.g., via Docker)
- Redis (e.g., via Docker)

---

## 1. Clone the Project

```bash
unzip wallet-api.zip && cd wallet-api
```

---

## 2. Start PostgreSQL (Docker Example)

```bash
docker run --name wallet-db -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=walletuser -e POSTGRES_DB=mnc_wallet_db -p 5432:5432 -d postgres
```

---

## 3. Start Redis (Docker Example)

```bash
docker run --name redis -p 6379:6379 -d redis
```

---

## 4. Set Environment Variables

Create a `.env` file or export manually:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=mnc_wallet_db
```

Optional: Load with [godotenv](https://github.com/joho/godotenv).

---

## 5. Run the API Server (Auto-Migrates DB)

```bash
go run cmd/main.go
```

Server runs at: `http://localhost:8080`

---

## 6. Run the Background Worker

```bash
go run worker/main.go
```

---

## 7. Start the Asynq Dashboard

```bash
docker run --rm -p 8081:8080 hibiken/asynqmon   --redis-addr=host.docker.internal:6379
```

Open [http://localhost:8081](http://localhost:8081) to view job status.

---

## 8. Run SQL Migration Manually (Optional)

```bash
psql -h localhost -U walletuser -d wallet -f migrations/init.sql
```

---

## License

MIT