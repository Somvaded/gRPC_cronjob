# Go gRPC Cron Job Service

A simple Go microservice that demonstrates:

- gRPC API
- Scheduled cron job
- In-memory report storage
- Graceful shutdown
- Logging
- gRPC reflection (for easier testing)

---

## How to run this service

### 1️⃣ Clone the repository

```bash
git clone https://github.com/Somvaded/cronjob-task.git
cd cronjob-task
```

### 2️⃣ Install dependencies

```bash
go mod tidy
```

- proto files are already generated

---

### 4️⃣ Run the service

```bash
go run main.go
```

You should see logs indicating:

- gRPC server started
- Cron jobs running every 10 seconds

---

## 📡 Testing

### Using grpcurl

Please install grpcurl

Since reflection is enabled:

#### Health Check

```bash
grpcurl -plaintext localhost:1234 report.ReportService/HealthCheck
```

#### Generate Report

```bash
grpcurl -plaintext -d '{"user_id":"test1"}' localhost:1234 report.ReportService/GenerateReport
```

**On Windows for generate Report (PowerShell):**

```powershell
grpcurl -plaintext -d '{\"user_id\":\"test_user\"}' -- localhost:1234 report.ReportService/GenerateReport
```

---

## 🗄 Project Structure

```
.
├── proto/          # Protobuf definitions
│   └── report.proto
├── server/         # gRPC server logic
├── cron/           # Cron scheduler logic
├── main.go         # Entry point
└── README.md
```

## 📖 Tech Stack

- Go (Golang)
- gRPC
- Protocol Buffers
- `robfig/cron/v3`
- Go's standard `log` package
