export type Credential = {
  accessToken: string;
  refreshToken: string;
};

/** accessToken のペイロードから導出される情報 */
export type AccessTokenClaim = Readonly<{
  /** 期限切れになる時刻の UNIX time */
  exp: number;
  /** トークンを発行した時刻の UNIX time */
  iat: number;

  username: string;
  isAdmin: boolean;
}>;
