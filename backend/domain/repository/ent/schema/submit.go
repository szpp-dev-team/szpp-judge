package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Submit struct {
	ent.Schema
}

func (Submit) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("status").Optional().Nillable(),
		field.Int("exec_time").Optional(),   // ms
		field.Int("exec_memory").Optional(), // kib
		field.Int("score").Optional(),
		field.String("compile_message").Optional(),
		field.Time("submitted_at"),
		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("submits").Unique().Required(),
		edge.From("task", Task.Type).Ref("submits").Unique().Required(),
		edge.From("language", Language.Type).Ref("submits").Unique().Required(),
		edge.From("contest", Contest.Type).Ref("submits").Unique(),
		edge.To("testcase_results", TestcaseResult.Type),
	}
}
