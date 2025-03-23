'use client';

import { useState } from 'react';
import { getAuth, signInWithEmailAndPassword } from '@firebase/auth';
import { initializeApp } from '@firebase/app';
import { FIREBASE_CONFIG } from '@/constants/auth';
import { useRouter } from 'next/navigation';

const useSignIn = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleSignIn = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      const app = initializeApp(FIREBASE_CONFIG);
      const auth = getAuth(app);
      await signInWithEmailAndPassword(auth, email, password);
      router.push('/'); // サインイン成功後に '/' へ遷移
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
      className="flex justify-center items-center h-screen bg-cover bg-center font-pixelify text-white"
      style={{ backgroundImage: "url(/images/signin-bg.jpg)" }}
    >
      <form onSubmit={handleSignIn} className="bg-black bg-opacity-80 p-5 rounded-lg w-80">
        {error && <p className="text-red-500">{error}</p>}
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="w-full p-2 my-2 bg-gray-800 border border-gray-600 rounded text-white placeholder-gray-400"
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="w-full p-2 my-2 bg-gray-800 border border-gray-600 rounded text-white placeholder-gray-400"
        />
        <button
          type="submit"
          className="w-full p-2 bg-green-500 hover:bg-green-600 rounded text-white transition"
        >
          Sign In
        </button>
      </form>
    </div>
  );
};

export default useSignIn;
