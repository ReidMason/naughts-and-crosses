
import { TRPCError } from "@trpc/server";
import { z } from "zod";

import { createTRPCRouter, publicProcedure } from "~/server/api/trpc";
import { createUser, NewUserResponse } from "~/server/services/userService";

export const userRouter = createTRPCRouter({
  createUser: publicProcedure
    .input(z.object({ username: z.string() }))
    .mutation(async ({ input }): Promise<NewUserResponse> => {
      try {
        const { result, data } = await createUser(input.username);
        if (result == 'err') {
          throw new TRPCError({
            code: "INTERNAL_SERVER_ERROR",
            message: data.message,
          })
        }

        return data;
      } catch (err) {
        throw new TRPCError({
          code: "INTERNAL_SERVER_ERROR",
          message: "An error occurred when creating user",
          cause: err
        })
      }
    }),
});
