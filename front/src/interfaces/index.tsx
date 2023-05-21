// サインアップ
export interface SignUpParams {
  last_name: string;
  first_name: string;
  mail_address: string;
  password: string;
}

// ユーザーテーブル
export interface User {
  mail_address: string;
}
