'use client';

import { useAuth } from '@/app/components/layout/providers/FirebaseAuth';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

function Home() {
  const { user, isLoading, isError } = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (!isLoading && user) {
      // ログイン済みの場合はページトップにリダイレクトさせる
      router.push('/');
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
        <main className="text-2xl font-bold">
          {/* 検証のため一時的にfirebase側で発行したユニークなIDを表示する。これをAPI Requestのヘッダーに詰めてバックエンドに渡す */}
          {user ? `provider_id: ${user.uid}` : 'Firebase Authentication sample by Next.js'}
        </main>
      </div>
    );
  }

  return (
    <div className="flex items-center justify-center h-screen">
      <main className="text-4xl font-bold">
        Firebase Authentication sample by Next.js
      </main>
    </div>
  );
}

export default Home;
