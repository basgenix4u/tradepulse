# TradePulse

Real-time trading platform (Go backend + React frontend).

## Local development

1. Set the required secrets in your environment (see `docker-compose.yml`):
   - `POSTGRES_PASSWORD`
   - `JWT_SECRET`
2. Start the stack:

```bash
POSTGRES_PASSWORD=your-strong-password JWT_SECRET=your-strong-secret docker compose up
```

Generate secrets with `openssl rand -base64 32`.
