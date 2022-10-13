package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
    field.String("title").
			Annotations(entsql.Annotation{
				Size: 32,
			}).
      Immutable().
      NotEmpty(),
    field.String("hash").
			Annotations(entsql.Annotation{
				Size: 256, // keccak256
			}).
      Immutable().
      NotEmpty(),
  }
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
    edge.From("users", User.Type).
      Ref("roles"),
  }
}