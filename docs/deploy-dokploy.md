# Deploy Foodies with Dokploy

This deployment uses Dokploy Docker Compose, with PostgreSQL and persistent named volumes for the database and uploaded evidence.

## Configure Dokploy

1. Create a **Docker Compose** service and connect the `vuskhoza/foodies-platform` Git repository and `main` branch.
2. Set the Compose Path to `./docker-compose.dokploy.yml`.
3. Add these variables in the service's Environment tab:

   | Variable | Value |
   |---|---|
   | `POSTGRES_PASSWORD` | A unique, randomly generated 32-character hexadecimal password |
   | `AUTH_JWT_SECRET` | Output of `openssl rand -base64 48` |
   | `PUBLIC_API_URL` | `https://foodies.spiritusglobal.tech` |
   | `WHATSAPP_VERIFY_TOKEN` | Leave empty until WhatsApp is configured |
   | `WHATSAPP_APP_SECRET` | Leave empty until WhatsApp is configured |

   The database password is used in a PostgreSQL URL, so use hexadecimal characters to avoid URL escaping issues. Dokploy Compose variables are written to its `.env`; the compose file explicitly passes only the needed values into each service.

4. Add `foodies.spiritusglobal.tech` in the Compose service's Domains tab and route it to the web container on port `3000`.
5. Create an `A` record for `foodies.spiritusglobal.tech` pointing to `169.58.37.240`, then enable HTTPS for the domain in Dokploy.
6. Deploy. PostgreSQL starts first, Goose applies pending migrations, then the API and web app start.

The web app proxies `/api/v1` to the API over the private Compose network, so the API and PostgreSQL have no host-published ports and only one public domain is needed. PostgreSQL data and uploaded evidence persist in named volumes; configure Dokploy volume backups before production use.

## First login

This compose file does not seed demo users or sample business data. The repository currently has no first-admin bootstrap flow, so the services can start but users cannot sign in until an administrator account is provisioned separately. `make seed-demo` creates a made-up organisation and shared demo credentials; use it only on a disposable staging deployment, never on a live database.
