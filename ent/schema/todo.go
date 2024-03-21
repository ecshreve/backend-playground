package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
		field.Int("priority").
			Default(0),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}
