'use client';

import { useState } from 'react';
import { useAuth } from '@/features/auth/hooks/useAuth';
import { useRouter } from 'next/navigation';
import { apiClient } from '@/utils/apiClient';

const useSignIn = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { signIn, error } = useAuth();
  const router = useRouter();

  const handleSignIn = async (e: React.FormEvent) => {
    e.preventDefault();

    const authResult = await signIn(email, password);
    if (!authResult) return;

    try {
      // APIリクエスト: `GET /core/v1/users/me`
      const userData = await apiClient.get('/users/me', authResult.idToken);
      console.log('User data:', userData);

      // 成功したらルートページへ遷移
      router.push('/');
    } catch (err) {
      console.error('API request error:', err);
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
