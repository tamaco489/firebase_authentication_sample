'use client';

import { useState } from 'react';
import { useAuth } from '@/features/auth/hooks/useAuth';
import { useRouter } from 'next/navigation';
import { apiClient } from '@/utils/apiClient';

const useSignUp = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { signUp, error } = useAuth();
  const router = useRouter();

  const handleSignUp = async (e: React.FormEvent) => {
    e.preventDefault();

    const authResult = await signUp(email, password);
    if (!authResult) return;

    const body = {
      provider_type: 'firebase',
    };

    try {
      // APIリクエスト: `POST /users`
      const userData = await apiClient.post('/users', authResult.idToken, body);
      // todo: 検証後削除
      console.log('User data:', userData);

      // サインアップ成功後に '/' へ遷移
      router.push('/');
    } catch (err) {
      console.error('API request error:', err);
    };
  };

  return (
    <div
      className="flex justify-center items-center h-screen bg-cover bg-center font-pixelify text-white"
      style={{ backgroundImage: "url(/images/signup-bg.jpg)" }}
    >
      <form onSubmit={handleSignUp} className="bg-black bg-opacity-80 p-5 rounded-lg w-80">
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
          className="w-full p-2 bg-blue-500 hover:bg-blue-600 rounded text-white transition"
        >
          Sign Up
        </button>
      </form>
    </div>
  );
};

export default useSignUp;
