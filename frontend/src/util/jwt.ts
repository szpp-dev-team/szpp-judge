/**
 * JWT のペイロード部分をデコードする。
 * ハッシュ値や有効期限などの検証はしない。
 * @throws SyntaxError JWT のフォーマットにそぐわない場合、あるいは JSON.parse に失敗した場合。
 */
export const decodeJwtPayload = (jwt: string): unknown => {
  const ldot = jwt.indexOf(".");
  const rdot = jwt.lastIndexOf(".");
  // ldot と rdot の間にさらにドットがあるケースを弾けないが、そんなコーナーケースは滅多にないので無視する
  if (ldot < 0 || rdot < 0 || ldot === rdot) {
    throw new SyntaxError("Invalid JWT: expected 3 parts splited by dot");
  }
  const payload = jwt.substring(ldot + 1, rdot);
  return JSON.parse(Buffer.from(payload, "base64").toString());
};
