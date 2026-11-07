import type { Metadata, Viewport } from "next";
import { Nunito } from "next/font/google";
import "./globals.css";

const nunito = Nunito({
  subsets: ["latin"],
  variable: "--font-nunito",
  display: "swap",
});

export const metadata: Metadata = {
  title: { default: "Foodies", template: "%s | Foodies" },
  description: "Operations and compliance for the Foodies group.",
  applicationName: "Foodies",
};

export const viewport: Viewport = {
  themeColor: "#f97316",
  width: "device-width",
  initialScale: 1,
  viewportFit: "cover",
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html lang="en-ZW" className={nunito.variable}>
      <body>{children}</body>
    </html>
  );
}
