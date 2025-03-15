'use client';

interface SignUpButtonProps {
  onClick: () => void;
}

const SignUpButton: React.FC<SignUpButtonProps> = ({ onClick }) => {
  return (
    <button
      style={{
        backgroundColor: '#008CBA',
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
      新規登録
    </button>
  );
};

export default SignUpButton;
