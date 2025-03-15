'use client';

import { FirebaseAuthProvider } from '@/app/components/layout/providers/FirebaseAuth';

function Home() {
  return (
    <div className="flex items-center justify-center h-screen">
      <main className="text-4xl font-bold">
        Firebase Authentication sample by Next.js
      </main>
      <footer className="">
      </footer>
    </div>
  );
}

export default function HomePage() {
  return (
    <FirebaseAuthProvider>
      <Home />
    </FirebaseAuthProvider>
  );
};
