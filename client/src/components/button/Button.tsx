import type { ReactNode } from 'react'
import { cva, VariantProps } from "class-variance-authority";

const appendClasses = (
  currentClasses: string,
  newClasses?: string
): string | undefined => {
  if (newClasses) currentClasses = (currentClasses + ` ${newClasses}`).trim();

  return currentClasses ? currentClasses : undefined;
};

const buttonStyles = cva(
  ["inline-flex items-center justify-center transition rounded border border-transparent px-2.5 py-1.5 text-xs font-medium shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2"],
  {
    variants: {
      intent: {
        primary: "hover:dark:bg-indigo-700 hover:bg-indigo-700 text-white bg-indigo-600 focus:ring-indigo-500",
        secondary: "dark:bg-slate-500 hover:dark:bg-slate-400",
      },
      pilled: {
        true: "rounded-full",
        false: "rounded-lg",
      },
      fullwidth: {
        true: "w-full",
      },
    },
    defaultVariants: {
      intent: "primary",
      pilled: false,
      fullwidth: false,
    },
  }
);

interface ButtonProps
  extends VariantProps<typeof buttonStyles>,
  React.ButtonHTMLAttributes<HTMLButtonElement> { }

interface ButtonProps {
  children: ReactNode,
}

export const Button = ({ children, className, ...props }: ButtonProps) => {
  const classes = appendClasses(buttonStyles(props), className);

  return (
    <button
      type="button"
      className={classes} {...props}
    >
      {children}
    </button>
  )
}

