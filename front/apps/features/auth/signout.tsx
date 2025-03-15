'use client';

import { useRouter } from 'next/navigation';
import { getAuth, signOut } from '@firebase/auth';
import { initializeApp } from '@firebase/app';
import { FIREBASE_CONFIG } from '@/constants/auth';

const useSignOut = () => {
  const router = useRouter();

  const handleSignOut = async () => {
    const app = initializeApp(FIREBASE_CONFIG);
    const auth = getAuth(app);
    try {
      await signOut(auth);
      router.push('/'); // サインアウト成功後に '/' へ遷移
      router.refresh();
    } catch (error) {
      console.error('サインアウトエラー:', error);
    }
  };

  return { handleSignOut };
};

export default useSignOut;
