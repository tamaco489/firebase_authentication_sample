'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { getAuth, onAuthStateChanged, User } from '@firebase/auth';
import { initializeApp } from '@firebase/app';
import { FIREBASE_CONFIG } from '@/constants/auth';
import useSignOut from '@/features/auth/signout';
import SignInButton from '@/app/components/auth/SignInButton';
import SignUpButton from '@/app/components/auth/SignUpButton';
import SignOutButton from '@/app/components/auth/SignOutButton';

const Header = () => {
  const router = useRouter();
  const [user, setUser] = useState<null | User>(null);

  useEffect(() => {
    const app = initializeApp(FIREBASE_CONFIG);
    const auth = getAuth(app);

    const unsubscribe = onAuthStateChanged(auth, (authUser) => {
      if (authUser) {
        setUser(authUser);
      } else {
        setUser(null);
      }
    });

    return () => unsubscribe();
  }, []);

  const { handleSignOut } = useSignOut();

  return (
    <header className="fixed top-0 left-0 w-full bg-cover bg-center p-5 text-white font-[Pixelify Sans, sans-serif] z-50"
            style={{ backgroundImage: 'url(/images/header-bg.jpg)' }}>
      <nav className="flex justify-between items-center">
        <h1 className="text-2xl cursor-pointer" onClick={() => router.push('/')}>
          Game Title
        </h1>
        <div className="flex gap-4">
          {/* 認証されていない場合のみ表示 */}
          {!user && (
            <>
              <SignInButton onClick={() => router.push('/sign_in')} />
              <SignUpButton onClick={() => router.push('/sign_up')} />
            </>
          )}

          {/* 認証済みの場合のみ表示 */}
          {user && <SignOutButton onClick={handleSignOut} />}
        </div>
      </nav>
    </header>
  );
};

export default Header;
