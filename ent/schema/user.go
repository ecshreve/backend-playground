package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			Annotations(
				entproto.Field(4),
			),
		field.String("email").NotEmpty().
			Annotations(
				entproto.Field(5),
			),
		field.String("avatar_image_url").
			Optional().
			Nillable().
			Annotations(
				entproto.Field(6),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("todos", Todo.Type).
			Ref("user").
			Annotations(
				entproto.Field(7),
			),
	}
}
