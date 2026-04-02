# 🔐 Secure CI/CD Pipeline with Integrated DevSecOps Practices

This project implements an end-to-end DevSecOps pipeline for a Go REST API, shifting security left by embedding automated security validation at every stage of the CI/CD workflow. The application is containerized with Docker, scanned for vulnerabilities, pushed to Amazon ECR, and deployed to AWS ECS Fargate — all through GitHub Actions.

---

## ⚙️ Tech Stack

| Layer | Tool / Technology |
|---|---|
| Application | Go (REST API) |
| CI/CD | GitHub Actions |
| Secret Scanning | Gitleaks, TruffleHog |
| SAST | Semgrep |
| SCA | Snyk |
| Container Scanning | Trivy |
| Containerization | Docker |
| Image Registry | Amazon ECR |
| Deployment | Amazon ECS Fargate |
| Networking | AWS VPC, ALB, NAT Gateway |
| Authentication | AWS OIDC (no static keys) |
| Monitoring | Amazon CloudWatch |
| Version Control | GitHub |

---

## 📁 Project Structure

```
go-devsecops-app/
├── .github/
│   └── workflows/
│       └── devsecops-pipeline.yml   # GitHub Actions pipeline
├── main.go                          # Go REST API source code
├── Dockerfile                       # Multi-stage Docker build
├── .dockerignore                    # Docker ignore rules
├── go.mod                           # Go module definition
└── README.md
```

---

## 🏗️ Architecture

> Add your architecture diagram here (draw.io export)

```
GitHub Push
    │
    ▼
GitHub Actions
    │
    ├── Gitleaks         (secrets scan)
    ├── TruffleHog       (git history secrets scan)
    ├── Semgrep          (SAST)
    ├── Snyk             (SCA)
    ├── Docker Build
    ├── Trivy            (container image scan)
    │
    ▼
Amazon ECR
    │
    ▼
ECS Fargate (private subnet)
    │
    ▼
Application Load Balancer (public)
    │
    ▼
Internet
```

---

## 🔒 DevSecOps Pipeline Stages

### 1. Gitleaks — Secret Scanning
Scans the entire repository for hardcoded secrets, API keys, and credentials on every push and pull request.

### 2. TruffleHog — Git History Secret Scanning
Performs a deep scan of the full git commit history to detect any secrets that may have been committed and later removed.

### 3. Semgrep — SAST
Runs static application security testing on the Go source code, detecting common vulnerabilities and insecure coding patterns.

### 4. Snyk — SCA
Scans `go.mod` for known vulnerabilities in third-party dependencies. Fails the pipeline on any HIGH or CRITICAL severity findings.

### 5. Docker Build
Builds the Go application using a multi-stage Dockerfile — the final image is based on `alpine:3.19` and runs as a non-root user to minimise the attack surface.

### 6. Trivy — Container Image Scan
Scans the built Docker image for OS-level and application-level vulnerabilities before it is pushed to ECR. Fails on HIGH or CRITICAL findings.

### 7. Push to ECR & Deploy to ECS
On successful completion of all security stages, the image is tagged with the git SHA, pushed to Amazon ECR, and deployed to ECS Fargate via a rolling update.

---

## 🌐 API Endpoints

| Method | Endpoint | Response |
|---|---|---|
| GET | `/` | `{"message": "Hello from DevSecOps pipeline"}` |
| GET | `/health` | `{"status": "healthy"}` |
| GET | `/version` | `{"version": "1.0.0"}` |

---

## ✅ Prerequisites

Before running this project, ensure the following are in place:

- AWS account with the following resources provisioned:
  - VPC with public and private subnets
  - NAT Gateway
  - Application Load Balancer
  - ECS Cluster and Fargate Service
  - ECR Repository
  - IAM OIDC Role for GitHub Actions
- GitHub repository secrets configured:
  - `AWS_ROLE_ARN` — IAM role ARN for OIDC authentication
  - `SNYK_TOKEN` — Snyk API token (free tier)
- Go 1.21+ installed locally for development
- Docker installed locally for testing

---

## 🚀 How the Pipeline Works

### On Pull Request
All security stages run (Gitleaks → TruffleHog → Semgrep → Snyk → Docker Build → Trivy). Deployment is skipped — this ensures every PR is security-validated before merging.

### On Push to Main
All security stages run sequentially. If every stage passes, the image is pushed to ECR and deployed to ECS Fargate automatically. Any stage failure stops the pipeline immediately — fail fast approach.

---

## 🔐 Security Design Decisions

- **OIDC authentication** — GitHub Actions authenticates with AWS using short-lived OIDC tokens. No static AWS access keys are stored anywhere.
- **Non-root container** — The Docker image runs as a non-root user (`appuser`) to follow the principle of least privilege.
- **Multi-stage Docker build** — The final image contains only the compiled binary and minimal runtime dependencies, reducing the attack surface.
- **Tag immutability** — ECR tag immutability is enabled, preventing image tags from being overwritten.
- **Private subnet deployment** — ECS Fargate tasks run in a private subnet with no direct internet exposure. All inbound traffic goes through the ALB.
- **Fail fast pipeline** — Security scans run before the Docker build. A vulnerability detected early stops the pipeline before any image is built or pushed.

---

## 📸 Screenshots

<img width="1862" height="562" alt="Screenshot 2026-04-02 213645" src="https://github.com/user-attachments/assets/40d3527d-b9a3-47d6-b661-95ba916bb30c" />

<img width="1390" height="852" alt="Screenshot 2026-04-02 213800" src="https://github.com/user-attachments/assets/7227dec0-d660-4f63-9ab7-7bd781c7820b" />

<img width="1413" height="768" alt="Screenshot 2026-04-02 214121" src="https://github.com/user-attachments/assets/1c7c7fc1-d971-41a9-834b-e71b7d6eaae2" />

<img width="1397" height="775" alt="Screenshot 2026-04-02 214215" src="https://github.com/user-attachments/assets/6a109a15-2132-4c5f-b726-69ab2dd12a3c" />

<img width="1428" height="607" alt="Screenshot 2026-04-02 214259" src="https://github.com/user-attachments/assets/9c500fe1-bc3b-429e-90be-d642e8ad9eab" />

<img width="1411" height="770" alt="Screenshot 2026-04-02 214423" src="https://github.com/user-attachments/assets/0e9d8763-3707-41d0-9c36-9547e9657659" />

<img width="1865" height="828" alt="Screenshot 2026-04-02 214529" src="https://github.com/user-attachments/assets/5aee026b-9762-4058-9db8-7c2fdd727929" />

<img width="1893" height="713" alt="ECR" src="https://github.com/user-attachments/assets/ad41a016-a341-4e12-adcf-b811be28fb8a" />


<img width="1893" height="657" alt="ECS" src="https://github.com/user-attachments/assets/1fa5966b-10e0-43f9-b014-6964f2b2c00a" />

<img width="891" height="145" alt="Screenshot 2026-04-02 215115" src="https://github.com/user-attachments/assets/a02f73b0-6775-4930-8e43-0cffdb0c3f40" />

<img width="865" height="147" alt="Screenshot 2026-04-02 215158" src="https://github.com/user-attachments/assets/fa0ff421-2644-4335-bc29-7ff586ad1f08" />

<img width="856" height="152" alt="Screenshot 2026-04-02 215228" src="https://github.com/user-attachments/assets/85227f76-4c6d-434e-9427-cbfb36059fbd" />














