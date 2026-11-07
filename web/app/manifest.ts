import type { MetadataRoute } from "next";

export default function manifest(): MetadataRoute.Manifest {
  return {
    name: "Foodies",
    short_name: "Foodies",
    description: "Operations and compliance for the Foodies group.",
    start_url: "/",
    display: "standalone",
    background_color: "#fff7f0",
    theme_color: "#f97316",
    icons: [{ src: "/icon.svg", sizes: "any", type: "image/svg+xml", purpose: "any" }],
  };
}
