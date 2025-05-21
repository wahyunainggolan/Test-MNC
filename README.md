## Prerequisites
- Go 1.20+
- Redis (e.g., via Docker)
- PostgreSql

----

##  Clone the Project
```bash
https://github.com/wahyunainggolan/Test-MNC.git
```


# 2. Test Tahap 2

A RESTful API for wallet transactions using Golang, Gin, PostgreSQL, and Redis-based background workers.

## Features
- User Registration, Login (with JWT)
- Top-up, Payment, and Transfer
- Transaction Reports
- Background Job Processing for Transfers using Asynq
- Monitoring Dashboard via Asynqmon

---

## 1. Go to the Project

```bash
cd test-tahap-2
```

---

## 2. Prepare Database

Create Database in PostgreeSql with name 

```bash
mnc_wallet_db
```
---

## 3. Start Redis (Docker Example)

```bash
docker run --name redis -p 6379:6379 -d redis
```

---

## 4. Set Environment Variables

Rename file `env.example` to `.env`, and setting the variable:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=mnc_wallet_db
```

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
psql -h localhost -U postgres -d mnc_wallet_db -f migrations/init.sql
```

---

## License

Wahyu Adi P Nainggolan