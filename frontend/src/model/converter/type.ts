/**
 * proto から生成されたものをそのまま使うとどうしても都合が悪いので変換したいときがある.
 * これは決まったインターフェースでその手続きを書いて治安を保ちたいという意図で作られた型.
 */
export type Converter<FrontendModel = unknown, ProtoDef = unknown> = {
  from(obj: ProtoDef): FrontendModel;
  to(obj: FrontendModel): ProtoDef;
};
