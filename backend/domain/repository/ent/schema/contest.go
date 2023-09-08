package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Contest holds the schema definition for the Contest entity.
type Contest struct {
	ent.Schema
}

// Fields of the Contest.
func (Contest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("slug"),
		field.String("description"),
		field.Int("task_id"),
		field.Int("clarification_id"),
		field.Time("start_at"),
		field.Time("end_at"),
	}
}

// Edges of the Contest.
func (Contest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
		edge.From("contest_users", User.Type).Ref("contests"),
	}
}
