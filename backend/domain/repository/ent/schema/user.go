package schema

import (
	"entgo.io/ent"
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
		field.String("username"),
		field.String("email").Unique(),
		field.String("role"), // this field's type must be managed with proto
		field.String("encrypted_password"),
		field.Time("created_at"),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
