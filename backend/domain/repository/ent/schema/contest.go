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
		field.String("name"),
		field.String("slug").Unique(),
		field.String("description"),
		field.Int("penalty_seconds"),
		field.String("contest_type"),
		field.Bool("is_public"),
		field.Time("start_at"),
		field.Time("end_at"),
	}
}

// Edges of the Contest.
func (Contest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submits", Submit.Type),
		edge.To("users", User.Type).Through("contest_user", ContestUser.Type),
		edge.To("tasks", Task.Type).Through("contest_task", ContestTask.Type),
	}
}
