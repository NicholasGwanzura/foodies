# Deploy Foodies with Dokploy

This deployment uses a Dokploy-managed PostgreSQL database and Docker Compose for the API and web app. Uploaded evidence is kept in a persistent named volume.

## Configure Dokploy

1. Create a **PostgreSQL** service in this project and environment. Set database and user names to `foodies`, and generate a database password in Dokploy.
2. Create a **Docker Compose** service and connect the `vuskhoza/foodies-platform` repository on branch `deploy/dokploy-foodies`.
3. Set the Compose Path to `./docker-compose.dokploy.yml`.
4. Add these variables in the Compose service's Environment tab:

   | Variable | Value |
   |---|---|
   | `DATABASE_URL` | The internal connection URL shown by the Dokploy PostgreSQL service |
   | `AUTH_JWT_SECRET` | Output of `openssl rand -base64 48` |
   | `PUBLIC_API_URL` | `https://foodies.spiritusglobal.tech` |
   | `WHATSAPP_VERIFY_TOKEN` | Leave empty until WhatsApp is configured |
   | `WHATSAPP_APP_SECRET` | Leave empty until WhatsApp is configured |

   Dokploy Compose variables are written to its `.env`; the compose file explicitly passes only the needed values into each service.

5. Add `foodies.spiritusglobal.tech` in the Compose service's Domains tab and route it to the web container on port `3000`.
6. Create an `A` record for `foodies.spiritusglobal.tech` pointing to `169.58.37.240`, then enable HTTPS for the domain in Dokploy.
7. Deploy the PostgreSQL service first. Then deploy Compose. Goose applies pending migrations, and the seed job loads reference roles and rules before the API and web app start.

The web app proxies `/api/v1` to the API over the private Compose network. The database should use Dokploy's internal connection URL; do not enable external database access. Uploaded evidence persists in a named volume. Configure Dokploy database backups before production use.

## First login

The seed job loads reference data by default. Without demo mode, the repository has no first-admin bootstrap flow, so the services can start but users cannot sign in until an administrator account is provisioned separately. To also load the made-up organisation and demo users, set `LOAD_DEMO_DATA=true`, `SEED_PASSWORD`, and `SEED_PIN` in Dokploy. That adds fictional employees E1001–E1028 across all 26 roles, using the same demo password and PIN for each account. Demo data is for testing only; do not enable it on a live database.
