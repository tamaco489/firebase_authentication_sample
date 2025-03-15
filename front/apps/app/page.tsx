'use client';

import { useAuth } from '@/app/components/layout/providers/FirebaseAuth';
import { useRouter } from 'next/navigation';
import { FirebaseAuthProvider } from '@/app/components/layout/providers/FirebaseAuth';
import { useState, useEffect } from 'react';
import SignIn from '@/app/components/auth/SignIn';

function Home() {
  const { user, isLoading, isError } = useAuth();
  const router = useRouter();
  const [showSignIn, setShowSignIn] = useState(false);

  useEffect(() => {
    if (!isLoading && !user) {
      // 未ログインの場合はSignInコンポーネントを表示
      setShowSignIn(true);
    } else {
      setShowSignIn(false);
    }
  }, [user, isLoading, router]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error...</div>;
  }

  if (showSignIn) {
    return <SignIn />;
  }

  return (
    <div className="flex items-center justify-center h-screen">
      <main className="text-4xl font-bold">
      {user ? `Welcome, ${user.displayName}!` : 'Firebase Authentication sample by Next.js'}
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
