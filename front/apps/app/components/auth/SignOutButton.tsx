'use client';

interface SignOutButtonProps {
  onClick: () => void;
}

const SignOutButton: React.FC<SignOutButtonProps> = ({ onClick }) => {
  return (
    <button
      className="
        bg-[#f44336]
        text-white
        py-2 px-5
        text-lg
        font-medium
        rounded-md
        transition-colors
        duration-300
        hover:bg-[#e53935]
        focus:outline-none
      "
      onClick={onClick}
    >
      サインアウト
    </button>
  );
};

export default SignOutButton;
