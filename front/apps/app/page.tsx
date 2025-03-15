'use client';

import { useAuth } from '@/app/components/layout/providers/FirebaseAuth';
import { useRouter } from 'next/navigation';
import { FirebaseAuthProvider } from '@/app/components/layout/providers/FirebaseAuth';
import { useEffect } from 'react';
import Header from '@/app/components/layout/header/Header';

function Home() {
  const { user, isLoading, isError } = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (!isLoading && user) {
      // ログイン済みの場合はリダイレクト
      router.push('/'); // ログイン後のページにリダイレクト
    }
  }, [user, isLoading, router]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error...</div>;
  }

  if (user) {
    return (
      <div className="flex items-center justify-center h-screen">
        <main className="text-4xl font-bold">
          {user ? `Welcome, ${user.displayName}!` : 'Firebase Authentication sample by Next.js'}
        </main>
      </div>
    );
  }

  return (
    <div className="flex items-center justify-center h-screen">
      <Header />
      <main className="text-4xl font-bold">
        Firebase Authentication sample by Next.js
      </main>
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
