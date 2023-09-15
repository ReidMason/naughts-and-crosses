import axios, { AxiosError, isAxiosError } from "axios";
import { z } from "zod";

const baseUrl = "http://0.0.0.0:5001/";

function createResponseSchema<T extends z.ZodTypeAny>(dataSchema: T) {
  return z.object({
    data: dataSchema,
  })
}

const newUserSchema = z.object({
  id: z.number(),
  username: z.string(),
  token: z.string(),
  wins: z.number(),
  losses: z.number(),
});
const newUserResponseSchema = createResponseSchema(newUserSchema);

const userSchema = z.object({
  id: z.number(),
  username: z.string(),
  wins: z.number(),
  losses: z.number(),
})
const userResponseSchema = createResponseSchema(userSchema);

export type NewUserResponse = z.infer<typeof newUserResponseSchema>;
export type UserResponseSchema = z.infer<typeof userResponseSchema>;

interface ResponseError {
  message: string;
}

type Result<Ok, Err> = { result: 'ok', data: Ok } | { result: 'err', data: Err };

export async function createUser(username: string): Promise<Result<NewUserResponse, ResponseError>> {
  try {
    const response = await axios.post<NewUserResponse>(`${baseUrl}user`, { username });
    const new_user = await newUserResponseSchema.parseAsync(response.data);
    return { result: 'ok', data: new_user };
  } catch (err) {
    if (isAxiosError(err)) {
      return {
        result: 'err', data: {
          message: err.response?.data.message
        }
      }
    }

    return {
      result: 'err', data: { message: "Unknown error" }
    }
  }
}

export async function getUserFromToken(): Promise<UserResponseSchema> {
  const response = await axios.get(`${baseUrl}user`, { withCredentials: true });
  return response.data;
}
