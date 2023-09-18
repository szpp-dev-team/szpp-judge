package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TestcaseResult struct {
	ent.Schema
}

func (TestcaseResult) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("status"),
		field.Int("exec_time"),   // ms
		field.Int("exec_memory"), // kib
	}
}

func (TestcaseResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submit", Submit.Type).Ref("testcase_results").Unique(),
		edge.To("testcase", Testcase.Type).Unique(),
	}
}
