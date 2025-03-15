'use client';

import { useRouter } from 'next/navigation';

const Header = () => {
  const router = useRouter();

  return (
    <header>
      <nav>
        <button onClick={() => router.push('/sign_in')}>ログイン</button>
        <button onClick={() => router.push('/sign_up')}>新規登録</button>
      </nav>
    </header>
  );
};

export default Header;
