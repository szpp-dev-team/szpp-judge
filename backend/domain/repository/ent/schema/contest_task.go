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
		field.Int("contest_id").Optional(),
		field.Int("task_id").Optional(),
	}
}

func (ContestTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contest", Contest.Type).Ref("contest_task").Field("contest_id").Unique(),
		edge.To("task", Task.Type).Field("task_id").Unique(),
	}
}
