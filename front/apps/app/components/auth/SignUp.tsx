'use client';

import { useState } from 'react';
import { getAuth, createUserWithEmailAndPassword } from '@firebase/auth';
import { initializeApp } from '@firebase/app';
import { FIREBASE_CONFIG } from '@/constants/auth';
import { useRouter } from 'next/navigation';

const SignUp = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      const app = initializeApp(FIREBASE_CONFIG);
      const auth = getAuth(app);
      await createUserWithEmailAndPassword(auth, email, password);
      router.push('/'); // サインアップ成功後に '/' へ遷移
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError('An unknown error occurred.');
      }
    }
  };

  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        backgroundImage: 'url(/images/signup-bg.jpg)',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        fontFamily: 'Pixelify Sans, sans-serif',
        color: 'white',
      }}
    >
      <form
        onSubmit={handleSubmit}
        style={{
          backgroundColor: 'rgba(0, 0, 0, 0.8)',
          padding: '20px',
          borderRadius: '10px',
          width: '300px',
        }}
      >
        {error && <p style={{ color: 'red' }}>{error}</p>}
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          style={{
            width: '100%',
            padding: '10px',
            margin: '10px 0',
            backgroundColor: '#333',
            border: '1px solid #555',
            borderRadius: '5px',
            color: 'white',
          }}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          style={{
            width: '100%',
            padding: '10px',
            margin: '10px 0',
            backgroundColor: '#333',
            border: '1px solid #555',
            borderRadius: '5px',
            color: 'white',
          }}
        />
        <button
          type="submit"
          style={{
            width: '100%',
            padding: '10px',
            backgroundColor: '#008CBA',
            border: 'none',
            borderRadius: '5px',
            color: 'white',
            cursor: 'pointer',
            transition: 'background-color 0.3s ease',
          }}
          onMouseOver={(e) => (e.currentTarget.style.backgroundColor = '#0077b5')}
          onMouseOut={(e) => (e.currentTarget.style.backgroundColor = '#008CBA')}
        >
          Sign Up
        </button>
      </form>
    </div>
  );
};

export default SignUp;
