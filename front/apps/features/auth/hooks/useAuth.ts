import { useState } from 'react';
import { getAuth, signInWithEmailAndPassword, User } from '@firebase/auth';
import { initializeApp } from '@firebase/app';
import { FIREBASE_CONFIG } from '@/constants/auth';

const app = initializeApp(FIREBASE_CONFIG);
const auth = getAuth(app);

export const useAuth = () => {
  const [error, setError] = useState<string | null>(null);

  const signIn = async (email: string, password: string): Promise<{ user: User; idToken: string } | null> => {
    setError(null);

    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password);
      const user = userCredential.user;
      const idToken = await user.getIdToken();
      return { user, idToken };
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError('An unknown error occurred.');
      }
      return null;
    }
  };

  return { signIn, error };
};
