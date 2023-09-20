package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ContestTask struct {
	ent.Schema
}

func (ContestTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("score"),
		field.Int("contest_id"),
		field.Int("task_id"),
	}
}

func (ContestTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contest", Contest.Type).
			Field("contest_id").
			Unique().
			Required(),
		edge.To("task", Task.Type).
			Field("task_id").
			Unique().
			Required(),
	}
}
