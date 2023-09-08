// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTitle, v))
}

// Statement applies equality check predicate on the "statement" field. It's identical to StatementEQ.
func Statement(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatement, v))
}

// Difficulty applies equality check predicate on the "difficulty" field. It's identical to DifficultyEQ.
func Difficulty(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDifficulty, v))
}

// ExecTimeLimit applies equality check predicate on the "exec_time_limit" field. It's identical to ExecTimeLimitEQ.
func ExecTimeLimit(v uint) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldExecTimeLimit, v))
}

// ExecMemoryLimit applies equality check predicate on the "exec_memory_limit" field. It's identical to ExecMemoryLimitEQ.
func ExecMemoryLimit(v uint) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldExecMemoryLimit, v))
}

// CaseInsensitive applies equality check predicate on the "case_insensitive" field. It's identical to CaseInsensitiveEQ.
func CaseInsensitive(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCaseInsensitive, v))
}

// Ndigits applies equality check predicate on the "ndigits" field. It's identical to NdigitsEQ.
func Ndigits(v uint) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldNdigits, v))
}

// JudgeCodePath applies equality check predicate on the "judge_code_path" field. It's identical to JudgeCodePathEQ.
func JudgeCodePath(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldJudgeCodePath, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTitle, v))
}

// StatementEQ applies the EQ predicate on the "statement" field.
func StatementEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatement, v))
}

// StatementNEQ applies the NEQ predicate on the "statement" field.
func StatementNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldStatement, v))
}

// StatementIn applies the In predicate on the "statement" field.
func StatementIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldStatement, vs...))
}

// StatementNotIn applies the NotIn predicate on the "statement" field.
func StatementNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldStatement, vs...))
}

// StatementGT applies the GT predicate on the "statement" field.
func StatementGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldStatement, v))
}

// StatementGTE applies the GTE predicate on the "statement" field.
func StatementGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldStatement, v))
}

// StatementLT applies the LT predicate on the "statement" field.
func StatementLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldStatement, v))
}

// StatementLTE applies the LTE predicate on the "statement" field.
func StatementLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldStatement, v))
}

// StatementContains applies the Contains predicate on the "statement" field.
func StatementContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldStatement, v))
}

// StatementHasPrefix applies the HasPrefix predicate on the "statement" field.
func StatementHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldStatement, v))
}

// StatementHasSuffix applies the HasSuffix predicate on the "statement" field.
func StatementHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldStatement, v))
}

// StatementEqualFold applies the EqualFold predicate on the "statement" field.
func StatementEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldStatement, v))
}

// StatementContainsFold applies the ContainsFold predicate on the "statement" field.
func StatementContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldStatement, v))
}

// DifficultyEQ applies the EQ predicate on the "difficulty" field.
func DifficultyEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDifficulty, v))
}

// DifficultyNEQ applies the NEQ predicate on the "difficulty" field.
func DifficultyNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDifficulty, v))
}

// DifficultyIn applies the In predicate on the "difficulty" field.
func DifficultyIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDifficulty, vs...))
}

// DifficultyNotIn applies the NotIn predicate on the "difficulty" field.
func DifficultyNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDifficulty, vs...))
}

// DifficultyGT applies the GT predicate on the "difficulty" field.
func DifficultyGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDifficulty, v))
}

// DifficultyGTE applies the GTE predicate on the "difficulty" field.
func DifficultyGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDifficulty, v))
}

// DifficultyLT applies the LT predicate on the "difficulty" field.
func DifficultyLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDifficulty, v))
}

// DifficultyLTE applies the LTE predicate on the "difficulty" field.
func DifficultyLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDifficulty, v))
}

// DifficultyContains applies the Contains predicate on the "difficulty" field.
func DifficultyContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldDifficulty, v))
}

// DifficultyHasPrefix applies the HasPrefix predicate on the "difficulty" field.
func DifficultyHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldDifficulty, v))
}

// DifficultyHasSuffix applies the HasSuffix predicate on the "difficulty" field.
func DifficultyHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldDifficulty, v))
}

// DifficultyEqualFold applies the EqualFold predicate on the "difficulty" field.
func DifficultyEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldDifficulty, v))
}

// DifficultyContainsFold applies the ContainsFold predicate on the "difficulty" field.
func DifficultyContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldDifficulty, v))
}

// ExecTimeLimitEQ applies the EQ predicate on the "exec_time_limit" field.
func ExecTimeLimitEQ(v uint) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldExecTimeLimit, v))
}

// ExecTimeLimitNEQ applies the NEQ predicate on the "exec_time_limit" field.
func ExecTimeLimitNEQ(v uint) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldExecTimeLimit, v))
}

// ExecTimeLimitIn applies the In predicate on the "exec_time_limit" field.
func ExecTimeLimitIn(vs ...uint) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldExecTimeLimit, vs...))
}

// ExecTimeLimitNotIn applies the NotIn predicate on the "exec_time_limit" field.
func ExecTimeLimitNotIn(vs ...uint) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldExecTimeLimit, vs...))
}

// ExecTimeLimitGT applies the GT predicate on the "exec_time_limit" field.
func ExecTimeLimitGT(v uint) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldExecTimeLimit, v))
}

// ExecTimeLimitGTE applies the GTE predicate on the "exec_time_limit" field.
func ExecTimeLimitGTE(v uint) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldExecTimeLimit, v))
}

// ExecTimeLimitLT applies the LT predicate on the "exec_time_limit" field.
func ExecTimeLimitLT(v uint) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldExecTimeLimit, v))
}

// ExecTimeLimitLTE applies the LTE predicate on the "exec_time_limit" field.
func ExecTimeLimitLTE(v uint) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldExecTimeLimit, v))
}

// ExecMemoryLimitEQ applies the EQ predicate on the "exec_memory_limit" field.
func ExecMemoryLimitEQ(v uint) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldExecMemoryLimit, v))
}

// ExecMemoryLimitNEQ applies the NEQ predicate on the "exec_memory_limit" field.
func ExecMemoryLimitNEQ(v uint) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldExecMemoryLimit, v))
}

// ExecMemoryLimitIn applies the In predicate on the "exec_memory_limit" field.
func ExecMemoryLimitIn(vs ...uint) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldExecMemoryLimit, vs...))
}

// ExecMemoryLimitNotIn applies the NotIn predicate on the "exec_memory_limit" field.
func ExecMemoryLimitNotIn(vs ...uint) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldExecMemoryLimit, vs...))
}

// ExecMemoryLimitGT applies the GT predicate on the "exec_memory_limit" field.
func ExecMemoryLimitGT(v uint) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldExecMemoryLimit, v))
}

// ExecMemoryLimitGTE applies the GTE predicate on the "exec_memory_limit" field.
func ExecMemoryLimitGTE(v uint) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldExecMemoryLimit, v))
}

// ExecMemoryLimitLT applies the LT predicate on the "exec_memory_limit" field.
func ExecMemoryLimitLT(v uint) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldExecMemoryLimit, v))
}

// ExecMemoryLimitLTE applies the LTE predicate on the "exec_memory_limit" field.
func ExecMemoryLimitLTE(v uint) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldExecMemoryLimit, v))
}

// JudgeTypeEQ applies the EQ predicate on the "judge_type" field.
func JudgeTypeEQ(v JudgeType) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldJudgeType, v))
}

// JudgeTypeNEQ applies the NEQ predicate on the "judge_type" field.
func JudgeTypeNEQ(v JudgeType) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldJudgeType, v))
}

// JudgeTypeIn applies the In predicate on the "judge_type" field.
func JudgeTypeIn(vs ...JudgeType) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldJudgeType, vs...))
}

// JudgeTypeNotIn applies the NotIn predicate on the "judge_type" field.
func JudgeTypeNotIn(vs ...JudgeType) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldJudgeType, vs...))
}

// CaseInsensitiveEQ applies the EQ predicate on the "case_insensitive" field.
func CaseInsensitiveEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCaseInsensitive, v))
}

// CaseInsensitiveNEQ applies the NEQ predicate on the "case_insensitive" field.
func CaseInsensitiveNEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCaseInsensitive, v))
}

// CaseInsensitiveIsNil applies the IsNil predicate on the "case_insensitive" field.
func CaseInsensitiveIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldCaseInsensitive))
}

// CaseInsensitiveNotNil applies the NotNil predicate on the "case_insensitive" field.
func CaseInsensitiveNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldCaseInsensitive))
}

// NdigitsEQ applies the EQ predicate on the "ndigits" field.
func NdigitsEQ(v uint) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldNdigits, v))
}

// NdigitsNEQ applies the NEQ predicate on the "ndigits" field.
func NdigitsNEQ(v uint) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldNdigits, v))
}

// NdigitsIn applies the In predicate on the "ndigits" field.
func NdigitsIn(vs ...uint) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldNdigits, vs...))
}

// NdigitsNotIn applies the NotIn predicate on the "ndigits" field.
func NdigitsNotIn(vs ...uint) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldNdigits, vs...))
}

// NdigitsGT applies the GT predicate on the "ndigits" field.
func NdigitsGT(v uint) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldNdigits, v))
}

// NdigitsGTE applies the GTE predicate on the "ndigits" field.
func NdigitsGTE(v uint) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldNdigits, v))
}

// NdigitsLT applies the LT predicate on the "ndigits" field.
func NdigitsLT(v uint) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldNdigits, v))
}

// NdigitsLTE applies the LTE predicate on the "ndigits" field.
func NdigitsLTE(v uint) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldNdigits, v))
}

// NdigitsIsNil applies the IsNil predicate on the "ndigits" field.
func NdigitsIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldNdigits))
}

// NdigitsNotNil applies the NotNil predicate on the "ndigits" field.
func NdigitsNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldNdigits))
}

// JudgeCodePathEQ applies the EQ predicate on the "judge_code_path" field.
func JudgeCodePathEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldJudgeCodePath, v))
}

// JudgeCodePathNEQ applies the NEQ predicate on the "judge_code_path" field.
func JudgeCodePathNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldJudgeCodePath, v))
}

// JudgeCodePathIn applies the In predicate on the "judge_code_path" field.
func JudgeCodePathIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldJudgeCodePath, vs...))
}

// JudgeCodePathNotIn applies the NotIn predicate on the "judge_code_path" field.
func JudgeCodePathNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldJudgeCodePath, vs...))
}

// JudgeCodePathGT applies the GT predicate on the "judge_code_path" field.
func JudgeCodePathGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldJudgeCodePath, v))
}

// JudgeCodePathGTE applies the GTE predicate on the "judge_code_path" field.
func JudgeCodePathGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldJudgeCodePath, v))
}

// JudgeCodePathLT applies the LT predicate on the "judge_code_path" field.
func JudgeCodePathLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldJudgeCodePath, v))
}

// JudgeCodePathLTE applies the LTE predicate on the "judge_code_path" field.
func JudgeCodePathLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldJudgeCodePath, v))
}

// JudgeCodePathContains applies the Contains predicate on the "judge_code_path" field.
func JudgeCodePathContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldJudgeCodePath, v))
}

// JudgeCodePathHasPrefix applies the HasPrefix predicate on the "judge_code_path" field.
func JudgeCodePathHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldJudgeCodePath, v))
}

// JudgeCodePathHasSuffix applies the HasSuffix predicate on the "judge_code_path" field.
func JudgeCodePathHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldJudgeCodePath, v))
}

// JudgeCodePathIsNil applies the IsNil predicate on the "judge_code_path" field.
func JudgeCodePathIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldJudgeCodePath))
}

// JudgeCodePathNotNil applies the NotNil predicate on the "judge_code_path" field.
func JudgeCodePathNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldJudgeCodePath))
}

// JudgeCodePathEqualFold applies the EqualFold predicate on the "judge_code_path" field.
func JudgeCodePathEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldJudgeCodePath, v))
}

// JudgeCodePathContainsFold applies the ContainsFold predicate on the "judge_code_path" field.
func JudgeCodePathContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldJudgeCodePath, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldUpdatedAt))
}

// HasTestcaseSets applies the HasEdge predicate on the "testcase_sets" edge.
func HasTestcaseSets() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcaseSetsTable, TestcaseSetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestcaseSetsWith applies the HasEdge predicate on the "testcase_sets" edge with a given conditions (other predicates).
func HasTestcaseSetsWith(preds ...predicate.TestcaseSet) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newTestcaseSetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTestcases applies the HasEdge predicate on the "testcases" edge.
func HasTestcases() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcasesTable, TestcasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestcasesWith applies the HasEdge predicate on the "testcases" edge with a given conditions (other predicates).
func HasTestcasesWith(preds ...predicate.Testcase) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newTestcasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaskContests applies the HasEdge predicate on the "task_contests" edge.
func HasTaskContests() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TaskContestsTable, TaskContestsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskContestsWith applies the HasEdge predicate on the "task_contests" edge with a given conditions (other predicates).
func HasTaskContestsWith(preds ...predicate.Contest) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newTaskContestsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
