#!/bin/sh
set -eu

if [ "${LOAD_DEMO_DATA:-false}" = "true" ]; then
  : "${SEED_PASSWORD:?Set SEED_PASSWORD in Dokploy for demo seeding}"
  : "${SEED_PIN:?Set SEED_PIN in Dokploy for demo seeding}"
  exec /usr/local/bin/seed -demo
fi

exec /usr/local/bin/seed
