# 🚀 Full Cycle Challenge: Stress Test CLI in Go

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![Status](https://img.shields.io/badge/Status-Completed-success?style=flat)]()

This project is a CLI (Command Line Interface) tool implemented in Go to perform load testing on web services. It allows users to specify a target URL, the total number of requests, and the desired level of concurrency.

## 🧠 Architecture

The system is designed to efficiently distribute HTTP requests using Go's powerful concurrency primitives:

1.  **Worker Pool**: The application spawns a configurable number of goroutines (concurrency) that act as workers.
2.  **Job Distribution**: A channel is used to distribute the total number of requests among the available workers.
3.  **Result Collection**: Each worker sends the result of its HTTP request (status code or error) to a buffered results channel.
4.  **Reporting**: Once all requests are completed, the system aggregates the results and generates a detailed report.

---

## 📁 Project Structure

```text
.
├── cmd/
│   └── stress-test/     # Application entry point
├── internal/
│   └── tester/          # Core logic for load testing and reporting
├── Dockerfile           # Multi-stage build for the application
├── go.mod               # Go module definition
└── README.md            # Documentation
```

---

## 🚀 How to Run

### Using Docker (Recommended)

You can build and run the application using Docker:

1.  **Build the Image**:
    ```bash
    docker build -t go-stress-test .
    ```

2.  **Execute the Test**:
    ```bash
    docker run go-stress-test --url=http://google.com --requests=100 --concurrency=10
    ```

### Running Locally

If you have Go installed, you can run it directly:

```bash
go run cmd/stress-test/main.go --url=http://google.com --requests=100 --concurrency=10
```

---

## 📊 Parameters

| Parameter | Description | Required |
|-----------|-------------|----------|
| `--url` | URL of the service to be tested | Yes |
| `--requests` | Total number of requests to be performed | Yes |
| `--concurrency` | Number of simultaneous calls | Yes |

---

## 📋 Report Example

Upon completion, the tool will output a report similar to this:

```text
--------------------------------------------------
             Stress Test Report                  
--------------------------------------------------
Total time spent:      1.45s
Total requests:        100
HTTP 200 responses:    100
Status code distribution:
  HTTP 200:              100
--------------------------------------------------
```
