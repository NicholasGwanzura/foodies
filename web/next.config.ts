import path from "node:path";
import type { NextConfig } from "next";

// .env at the repository root is the one place configuration lives, for the API
// and the web app alike. CI has no .env, so it uses the example values.
for (const file of [".env", ".env.example"]) {
  try {
    process.loadEnvFile(path.resolve(import.meta.dirname, "..", file));
  } catch (error) {
    if ((error as NodeJS.ErrnoException).code !== "ENOENT") throw error;
  }
}

const nextConfig: NextConfig = {
  cacheComponents: true,
  partialPrefetching: true,
};

export default nextConfig;
