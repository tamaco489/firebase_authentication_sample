import type { Metadata } from "next";
import "./styles/globals.css";
import QueryProvider from "@/app/components/layout/providers/QueryProvider"
import { FirebaseAuthProvider } from '@/app/components/layout/providers/FirebaseAuth';
import Header from '@/app/components/layout/header/Header';
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
          <FirebaseAuthProvider>
            <Header />
            {/* フッターの高さに合わせて調整 */}
            <div style={{ paddingBottom: '60px' }}>
              {children}
            </div>
          </FirebaseAuthProvider>
        </QueryProvider>
        <Footer />
      </body>
    </html>
  );
};
