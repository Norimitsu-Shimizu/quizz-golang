import { client } from "./URL";
import { SignInParams } from "@/interfaces";

// サインアップ（新規アカウント作成）
export const signUp = (params: any) => {
  return client.post("api/v1/signup", params);
};

// サインイン（ログイン）
export const signIn = (params: SignInParams) => {
  return client.post("api/v1/signin", params);
};

