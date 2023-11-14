import { MouseEventHandler } from "react";

export const Button = ({
  children,
  disabled = false,
  type = "button",
  form,
  onClick,
}: {
  onMouseEnter?: React.MouseEventHandler<HTMLButtonElement>;
  onMouseLeave?: React.MouseEventHandler<HTMLButtonElement>;
  children: React.ReactNode;
  color?: "purple" | "white" | "navColor" | "transparent" | "sub";
  sub?: boolean;
  fullWidth?: boolean;
  disabled?: boolean;
  type?: React.ComponentProps<"button">["type"];
  form?: React.ComponentProps<"button">["form"];
  onClick?: MouseEventHandler<HTMLButtonElement>;
}): JSX.Element => {

  const handleSubmit: MouseEventHandler<HTMLButtonElement> = (event) => {
    if (onClick && !disabled) {
      onClick(event);
    }
  };

  return (
    <>
      <button
        onClick={handleSubmit}
        type={type}
        form={form}
      >
        {children}
      </button>
    </>
  );
};

export default Button;
