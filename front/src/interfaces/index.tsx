// サインアップ
export interface SignUpParams {
  username: string;
  email: string;
  password: string;
}

export interface SignInParams {
  email: string;
  password: string;
  otp_attempt?: number;
}


// ユーザーテーブル
export interface User {
  username: string;
  email: string;
  // token?: string
}
