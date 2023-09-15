import { TRPCError } from '@trpc/server';
import React, { useState } from 'react'
import { useForm } from 'react-hook-form';
import { Input } from "~/components/input/Input";
import { api } from '~/utils/api';
import { setUserToken } from '~/utils/cookieHandler';
import { Button } from '../button/Button';

interface UsernameInput {
  username: string;
}

export const UsernameInput = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<UsernameInput>();
  const [errorMessage, setErrorMessage] = useState<string>();
  const userCreationMutation = api.user.createUser.useMutation();

  const onSubmit = async (data: UsernameInput) => {
    try {
      const response = await userCreationMutation.mutateAsync(data);
      setUserToken(response.data.token);
    } catch (err: any) {
      console.log(err.message);
      setErrorMessage(err.message);
    }
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col items-center gap-4 w-full">
      <div className="flex bg-white p-2 rounded-xl w-full">
        <Input className="flex-auto" transparent errorMessage={errorMessage} register={register("username")} />
        <Button className="inline-flex">Get started</Button>
      </div>
    </form>
  )
}
