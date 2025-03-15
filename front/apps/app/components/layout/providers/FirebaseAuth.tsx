'use client';

import { createContext, useContext, ReactNode } from "react";
import { useQuery } from '@tanstack/react-query';
import { initializeApp } from '@firebase/app';
import { getAuth, onAuthStateChanged, User } from '@firebase/auth';
import { FIREBASE_CONFIG } from "@/constants/auth";


// Firebase Authenticationのコンテキスト型
interface AuthContextType {
  user: User | null | undefined;
  isLoading: boolean;
  isError: boolean;
}

// Firebase Authenticationのコンテキストを作成
const AuthContext = createContext<AuthContextType | undefined>(undefined);

// Firebase Authentication Providerコンポーネント
export const FirebaseAuthProvider = ({ children }: { children: ReactNode }) => {
  const { data: user, isLoading, isError } = useQuery({
    queryKey: ['firebaseAuthUser'],
    queryFn: () => {
      return new Promise<User | null>((resolve) => {
        // Firebase初期化
        const app = initializeApp(FIREBASE_CONFIG);
        const auth = getAuth(app);

        // 認証状態の監視
        const unsubscribe = onAuthStateChanged(auth, (user) => {
          resolve(user);
          unsubscribe(); // 1回のみ監視
        });
      });
    },
    staleTime: Infinity, // キャッシュを常に最新に保つ
  });

  const value: AuthContextType = { user, isLoading, isError };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
};

// Firebase Authenticationコンテキストのカスタムフック
export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within a FirebaseAuthProvider');
  }
  return context;
};
