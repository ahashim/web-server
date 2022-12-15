package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").
			Immutable().
			Match(regexp.MustCompile("^0x[0-9a-fA-F]{40}$")). // eth address
			NotEmpty(),
		field.String("username").
			MaxLen(32).
			NotEmpty(),
		field.Enum("status").
			Values("UNKNOWN", "ACTIVE", "SUSPENDED", "BANNED"),
		field.Int8("level").
			Default(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("interactions", Interaction.Type),
		edge.To("pool_passes", PoolPass.Type),
		edge.To("roles", Role.Type),
		edge.To("created", Squeak.Type),
		edge.To("owned", Squeak.Type),
		edge.To("following", User.Type).
			From("followers"),
	}
}

// Mixins of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
