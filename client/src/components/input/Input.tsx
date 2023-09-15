import React from 'react'
import { ExclamationCircleIcon } from '@heroicons/react/20/solid'
import { UseFormRegisterReturn } from 'react-hook-form'

interface InputProps {
  label?: string,
  errorMessage?: string,
  register?: UseFormRegisterReturn,
  className?: string,
  transparent?: boolean,
}

export const Input = ({ label, errorMessage, register, className, transparent }: InputProps) => {
  const errorStyles = "border-red-300 text-red900 placeholder-red-300 focus:border-red-500 focus:ring-red-500"

  return (
    <input
      type="text"
      name="username"
      id="username"
      className={`block rounded-md pr-10 sm:text-sm ${errorMessage && errorStyles} ${className} ${transparent && "bg-transparent border-none focus:ring-transparent"}`}
      placeholder="Username"
      aria-invalid="true"
      aria-describedby="username-error"
      {...register}
    />
    // <div>
    //   {label &&
    //     <label htmlFor="username" className="block text-sm font-medium text-slate-100">
    //       {label}
    //     </label>
    //   }
    //   <div className={`relative rounded-md shadow-sm ${label && "mt-1"}`}>
    //     <input
    //       type="text"
    //       name="username"
    //       id="username"
    //       className={`block w-full rounded-md pr-10 sm:text-sm ${errorMessage && errorStyles} ${className} ${transparent && "bg-transparent border-none focus:ring-transparent"}`}
    //       placeholder="Username"
    //       aria-invalid="true"
    //       aria-describedby="username-error"
    //       {...register}
    //     />
    //     {errorMessage &&
    //       <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
    //         <ExclamationCircleIcon className="h-5 w-5 text-red-500" aria-hidden="true" />
    //       </div>
    //     }
    //   </div>
    //   {errorMessage &&
    //     <p className="mt-2 text-sm text-red-600" id="username-error">
    //       {errorMessage}
    //     </p>
    //   }
    // </div>
  )
}
