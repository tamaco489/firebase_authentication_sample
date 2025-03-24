'use client';

import { useState } from 'react';
import { getAuth, createUserWithEmailAndPassword } from '@firebase/auth';
import { initializeApp } from '@firebase/app';
import { FIREBASE_CONFIG } from '@/constants/auth';
import { useRouter } from 'next/navigation';
import { API_HOST } from '@/constants/api';

const useSignUp = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleSignUp = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      const app = initializeApp(FIREBASE_CONFIG);
      const auth = getAuth(app);
      const userCredential = await createUserWithEmailAndPassword(auth, email, password);

      // Firebase から ID トークンを取得
      const user = userCredential.user;
      const idToken = await user.getIdToken();

      // todo: 検証後削除。
      console.log('[debug:1] user info:' , user);
      console.log('[debug:2] id_token:' , idToken);

      // APIリクエスト設定

      // リクエストヘッダーの設定
      // Bearer トークン: Firebase Authentication から取得したIDトークンを設定
      // todo: これもconstants/api.tsで定義してそれを使いまわすかたちにする。※GetMeのリクエスト完成させてからの方が早いかも
      const headers = {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${idToken}`,
      };

      // リクエストボディの設定
      // todo: このエンドポイント専用のリクエストボディの型と、provider_typeも型定義する
      const body = {
        provider_type: 'firebase',
      };

      // todo: この辺りもクラスメソッドに変更
      const response = await fetch(`${API_HOST}/users`, {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(body),
      });

      // レスポンスの処理
      // todo: これは不要になる想定
      const data = await response.json();
      console.log('user created:', data);

      // レスポンスのエラーチェック
      if (!response.ok) {
        throw new Error(`API request failed: ${response.status} ${response.statusText}`);
      };

      // サインアップ成功後に '/' へ遷移
      router.push('/');
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError('An unknown error occurred.');
      };
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
