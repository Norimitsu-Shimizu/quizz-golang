// サインアップ
export interface SignUpParams {
  username: string;
  mail_address: string;
  password: string;
}

// ユーザーテーブル
export interface User {
  mail_address: string;
}
