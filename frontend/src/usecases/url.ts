/** ログイン成功後に指定されたURIへリダイレクトしてくれる URI を生成する */
export const getLoginUriWithRedirect = (redirectUri: string): string => {
  return "/login?redirecturi=" + encodeURIComponent(redirectUri);
};
