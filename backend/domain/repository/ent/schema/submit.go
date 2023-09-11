package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Submit struct {
	ent.Schema
}

func (Submit) Field() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("status").Optional(),
		field.Int("exec_time").Optional(),   // ms
		field.Int("exec_memory").Optional(), // kib
		field.Time("submitted_at"),
		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("submits"),
		edge.From("task", Task.Type).Ref("submits").Unique(),
		edge.From("language", Language.Type).Ref("submits").Unique(),
		edge.To("testcase_results", TestcaseResult.Type),
	}
}
