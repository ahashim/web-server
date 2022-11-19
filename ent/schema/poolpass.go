package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/ahashim/web-server/types"
)

// PoolPass holds the schema definition for the PoolPass entity.
type PoolPass struct {
	ent.Schema
}

// Fields of the PoolPass.
func (PoolPass) Fields() []ent.Field {
	return []ent.Field{
		field.Int("shares").
			GoType(new(types.Uint256)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}).
			Immutable(),
	}
}

// Edges of the PoolPass.
func (PoolPass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("pool_passes").
			Unique(),
		edge.From("pool", Pool.Type).
			Ref("pool_passes").
			Unique(),
	}
}

// Mixins of the PoolPass
func (PoolPass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
