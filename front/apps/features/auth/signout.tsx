'use client';

import { useRouter } from 'next/navigation';
import { getAuth, signOut } from '@firebase/auth';
import SignOutButton from '@/app/components/auth/SignOutButton';

const SignOut = () => {
  const router = useRouter();

  const handleSignOut = async () => {
    const auth = getAuth();
    try {
      await signOut(auth);
      router.push('/'); // サインアウト成功後に '/' へ遷移
      router.refresh();
    } catch (error) {
      console.error('サインアウトエラー:', error);
    }
  };

  return <SignOutButton onClick={handleSignOut} />;
};

export default SignOut;
