'use client';

interface SignInButtonProps {
  onClick: () => void;
}

const SignInButton: React.FC<SignInButtonProps> = ({ onClick }) => {
  return (
    <button
      className="
        bg-[#4CAF50]
        text-white
        py-2 px-5
        text-lg
        font-medium
        rounded-md
        transition-colors
        duration-300
        hover:bg-[#45a049]
        focus:outline-none
      "
      onClick={onClick}
    >
      ログイン
    </button>
  );
};

export default SignInButton;
