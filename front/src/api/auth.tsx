import { client } from "./URL";

// サインアップ（新規アカウント作成）
export const signUp = (params: any) => {
  return client.post("auth", params);
};
