// @generated by protoc-gen-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file judge/v1/judge_type.proto (package judge.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * 行ごとの一致比較 (最後の改行忘れは許容, 改行コードの違いも許容)
 *
 * @generated from message judge.v1.JudgeTypeNormal
 */
export class JudgeTypeNormal extends Message<JudgeTypeNormal> {
  /**
   * @generated from field: optional bool case_insensitive = 1;
   */
  caseInsensitive?: boolean;

  constructor(data?: PartialMessage<JudgeTypeNormal>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.JudgeTypeNormal";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "case_insensitive", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeTypeNormal {
    return new JudgeTypeNormal().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeTypeNormal {
    return new JudgeTypeNormal().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeTypeNormal {
    return new JudgeTypeNormal().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeTypeNormal | PlainMessage<JudgeTypeNormal> | undefined, b: JudgeTypeNormal | PlainMessage<JudgeTypeNormal> | undefined): boolean {
    return proto3.util.equals(JudgeTypeNormal, a, b);
  }
}

/**
 * 浮動小数点数向けの許容誤差付き比較
 *
 * @generated from message judge.v1.JudgeTypeNumericalErrTolerant
 */
export class JudgeTypeNumericalErrTolerant extends Message<JudgeTypeNumericalErrTolerant> {
  /**
   * 許容誤差の桁数
   * 例えば ndigits=5 なら絶対誤差または相対誤差のいずれかが 1e-5 未満なら正解
   *
   * @generated from field: uint32 ndigits = 1;
   */
  ndigits = 0;

  constructor(data?: PartialMessage<JudgeTypeNumericalErrTolerant>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.JudgeTypeNumericalErrTolerant";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ndigits", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeTypeNumericalErrTolerant {
    return new JudgeTypeNumericalErrTolerant().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeTypeNumericalErrTolerant {
    return new JudgeTypeNumericalErrTolerant().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeTypeNumericalErrTolerant {
    return new JudgeTypeNumericalErrTolerant().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeTypeNumericalErrTolerant | PlainMessage<JudgeTypeNumericalErrTolerant> | undefined, b: JudgeTypeNumericalErrTolerant | PlainMessage<JudgeTypeNumericalErrTolerant> | undefined): boolean {
    return proto3.util.equals(JudgeTypeNumericalErrTolerant, a, b);
  }
}

/**
 * インタラクティブジャッジ
 * 提出プログラムとジャッジプログラムとでリアルタイムに入出力をやりとりする
 *
 * @generated from message judge.v1.JudgeTypeInteractive
 */
export class JudgeTypeInteractive extends Message<JudgeTypeInteractive> {
  /**
   * @generated from field: string judge_code_path = 1;
   */
  judgeCodePath = "";

  constructor(data?: PartialMessage<JudgeTypeInteractive>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.JudgeTypeInteractive";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "judge_code_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeTypeInteractive {
    return new JudgeTypeInteractive().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeTypeInteractive {
    return new JudgeTypeInteractive().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeTypeInteractive {
    return new JudgeTypeInteractive().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeTypeInteractive | PlainMessage<JudgeTypeInteractive> | undefined, b: JudgeTypeInteractive | PlainMessage<JudgeTypeInteractive> | undefined): boolean {
    return proto3.util.equals(JudgeTypeInteractive, a, b);
  }
}

/**
 * スペシャルジャッジ
 * 正解が複数考えられる場合にはこのジャッジ形式を使う
 *
 * @generated from message judge.v1.JudgeTypeCustom
 */
export class JudgeTypeCustom extends Message<JudgeTypeCustom> {
  /**
   * @generated from field: string judge_code_path = 1;
   */
  judgeCodePath = "";

  constructor(data?: PartialMessage<JudgeTypeCustom>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.JudgeTypeCustom";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "judge_code_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeTypeCustom {
    return new JudgeTypeCustom().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeTypeCustom {
    return new JudgeTypeCustom().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeTypeCustom {
    return new JudgeTypeCustom().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeTypeCustom | PlainMessage<JudgeTypeCustom> | undefined, b: JudgeTypeCustom | PlainMessage<JudgeTypeCustom> | undefined): boolean {
    return proto3.util.equals(JudgeTypeCustom, a, b);
  }
}

/**
 * @generated from message judge.v1.JudgeType
 */
export class JudgeType extends Message<JudgeType> {
  /**
   * @generated from oneof judge.v1.JudgeType.judge_type
   */
  judgeType: {
    /**
     * @generated from field: judge.v1.JudgeTypeNormal normal = 1;
     */
    value: JudgeTypeNormal;
    case: "normal";
  } | {
    /**
     * @generated from field: judge.v1.JudgeTypeNumericalErrTolerant numerical_err_torelant = 2;
     */
    value: JudgeTypeNumericalErrTolerant;
    case: "numericalErrTorelant";
  } | {
    /**
     * @generated from field: judge.v1.JudgeTypeInteractive interactive = 3;
     */
    value: JudgeTypeInteractive;
    case: "interactive";
  } | {
    /**
     * @generated from field: judge.v1.JudgeTypeCustom custom = 4;
     */
    value: JudgeTypeCustom;
    case: "custom";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<JudgeType>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "judge.v1.JudgeType";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "normal", kind: "message", T: JudgeTypeNormal, oneof: "judge_type" },
    { no: 2, name: "numerical_err_torelant", kind: "message", T: JudgeTypeNumericalErrTolerant, oneof: "judge_type" },
    { no: 3, name: "interactive", kind: "message", T: JudgeTypeInteractive, oneof: "judge_type" },
    { no: 4, name: "custom", kind: "message", T: JudgeTypeCustom, oneof: "judge_type" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JudgeType {
    return new JudgeType().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JudgeType {
    return new JudgeType().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JudgeType {
    return new JudgeType().fromJsonString(jsonString, options);
  }

  static equals(a: JudgeType | PlainMessage<JudgeType> | undefined, b: JudgeType | PlainMessage<JudgeType> | undefined): boolean {
    return proto3.util.equals(JudgeType, a, b);
  }
}
