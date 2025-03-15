'use client';

interface SignOutButtonProps {
  onClick: () => void;
}

const SignOutButton: React.FC<SignOutButtonProps> = ({ onClick }) => {
  return (
    <button
      style={{
        backgroundColor: '#f44336',
        border: 'none',
        color: 'white',
        padding: '10px 20px',
        textAlign: 'center',
        textDecoration: 'none',
        display: 'inline-block',
        fontSize: '1em',
        margin: '4px 2px',
        cursor: 'pointer',
        borderRadius: '5px',
        transition: 'background-color 0.3s ease',
      }}
      onClick={onClick}
    >
      サインアウト
    </button>
  );
};

export default SignOutButton;
