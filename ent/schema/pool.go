package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/ahashim/web-server/types"
)

// Pool holds the schema definition for the Pool entity.
type Pool struct {
	ent.Schema
}

// Fields of the Pool.
func (Pool) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").
			GoType(new(types.Uint256)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}),
		field.Int("shares").
			GoType(new(types.Uint256)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}),
		field.Int("block_number").
			GoType(new(types.Uint256)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}).
			Immutable(),
		field.Int("score").
			Immutable(),
	}
}

// Edges of the Pool.
func (Pool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pool_passes", PoolPass.Type),
		edge.From("squeak", Squeak.Type).
			Ref("pool").
			Unique(),
	}
}

// Mixins of the Pool.
func (Pool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
