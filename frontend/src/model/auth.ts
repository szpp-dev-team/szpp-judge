export type Credential = {
  accessToken: string;
  refreshToken: string;
};

/** accessToken から導出される、自身のユーザ情報 */
export type AuthUser = {
  id: number;
  username: string;
  isAdmin: boolean;
};
