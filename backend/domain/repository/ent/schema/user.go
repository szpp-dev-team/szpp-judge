package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("role"), // this field's type must be managed with proto
		field.Bytes("hashed_password"),
		field.Time("created_at"),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
		edge.To("submits", Submit.Type),
		edge.From("contests", Contest.Type).Ref("users").Through("contest_user", ContestUser.Type),
		edge.To("refresh_tokens", RefreshToken.Type),
	}
}
