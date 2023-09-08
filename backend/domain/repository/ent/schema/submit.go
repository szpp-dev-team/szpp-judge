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
		field.Int("user_id"),
		field.Int("contest_id").Optional().Nillable(),
		field.Int("task_id"),
		field.Int("language_id"),
		field.String("status"),
		field.Int("exec_time"),   // ms
		field.Int("exec_memory"), // kib
		field.Time("submitted_at"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		// edge.To("task", Task.Type),
		edge.To("language", Language.Type),
		edge.To("testcase_result", TestcaseResult.Type),
	}
}
