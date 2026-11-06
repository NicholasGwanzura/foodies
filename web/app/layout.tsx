import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Foodies",
  description: "Operations and compliance for the Foodies group.",
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
