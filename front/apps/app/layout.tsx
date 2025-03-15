import type { Metadata } from "next";
import "./styles/globals.css";
import QueryProvider from "@/app/components/layout/providers/QueryProvider"
import Footer from "./components/layout/footer/Footer";

export const metadata: Metadata = {
  title: "firebase authentication sample",
  description: "firebase authentication sample",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>
        <QueryProvider>
          {children}
          <Footer />
        </QueryProvider>
      </body>
    </html>
  );
};
