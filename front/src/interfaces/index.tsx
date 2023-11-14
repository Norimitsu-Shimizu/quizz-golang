// サインアップ
export interface SignUpParams {
  username: string;
  email: string;
  password: string;
}

// ユーザーテーブル
export interface User {
  email: string;
}
