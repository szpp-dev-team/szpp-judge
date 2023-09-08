package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TestcaseResult struct {
	ent.Schema
}

func (TestcaseResult) Field() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("submission_id"),
		field.Int("testcase_id"),
		field.String("status"),
		field.Int("exec_time"),   // ms
		field.Int("exec_memory"), // kib
	}
}

func (TestcaseResult) Edge() []ent.Edge {
	return []ent.Edge{
		edge.From("submit", Submit.Type),
	}
}
