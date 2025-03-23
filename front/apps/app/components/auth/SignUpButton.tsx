'use client';

interface SignUpButtonProps {
  onClick: () => void;
}

const SignUpButton: React.FC<SignUpButtonProps> = ({ onClick }) => {
  return (
    <button
      className="
        bg-[#008CBA]
        text-white
        py-2 px-5
        text-lg
        font-medium
        rounded-md
        transition-colors
        duration-300
        hover:bg-[#006f8f]
        focus:outline-none
      "
      onClick={onClick}
    >
      新規登録
    </button>
  );
};

export default SignUpButton;
