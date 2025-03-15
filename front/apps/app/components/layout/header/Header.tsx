'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { getAuth, onAuthStateChanged,  User } from '@firebase/auth';
import { initializeApp } from '@firebase/app';
import { FIREBASE_CONFIG } from '@/constants/auth';
import SignOut from '@/features/auth/signout';
import SignInButton from '@/app/components/auth/SignInButton';
import SignUpButton from '@/app/components/auth/SignUpButton';

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

  return (
    <header
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        backgroundImage: 'url(/images/header-bg.jpg)',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        padding: '20px',
        color: 'white',
        fontFamily: 'Pixelify Sans, sans-serif',
        zIndex: 100,
      }}
    >
      <nav style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h1
          style={{
            margin: 0,
            fontSize: '2em',
            cursor: 'pointer',
          }}
          onClick={() => router.push('/')}
        >
          Game Title
        </h1>
        <div>
          {/* 認証されていない場合のみ表示 */}
          {!user && (
            <>
              <SignInButton onClick={() => router.push('/sign_in')} />
              <SignUpButton onClick={() => router.push('sign_up') } />
            </>
          )}

          {/* 認証済みの場合のみ表示 */}
          {user && <SignOut />}
        </div>
      </nav>
    </header>
  );
};

export default Header;
