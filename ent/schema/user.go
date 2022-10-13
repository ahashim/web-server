package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/enums"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").
			Annotations(entsql.Annotation{
				Size: 42, // eth address with `0x` prefix
			}).
			Immutable().
			Match(regexp.MustCompile("^0x[0-9a-fA-F]{40}$")).
			NotEmpty(),
		field.String("username").
			Annotations(entsql.Annotation{
				Size: 32,
			}).
			MaxLen(32).
			NotEmpty(),
		field.Enum("status").
			GoType(enums.Status(0)),
		field.Int8("scout_level").
			Default(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			From("followers"),
		edge.To("roles", Role.Type),
	}
}
