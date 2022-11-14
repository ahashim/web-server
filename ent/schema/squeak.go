package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/ahashim/web-server/types"
)

// Squeak holds the schema definition for the Squeak entity.
type Squeak struct {
	ent.Schema
}

// Fields of the Squeak.
func (Squeak) Fields() []ent.Field {
	return []ent.Field{
		field.Int("block_number").
			GoType(new(types.Uint256)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}),
		field.String("content").
			Immutable().
			MaxLen(256).
			NotEmpty(),
	}
}

// Edges of the Squeak.
func (Squeak) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("interactions", Interaction.Type),
		edge.To("pool", Pool.Type).
      Unique(),
		edge.From("creator", User.Type).
			Ref("created").
			Unique(),
		edge.From("owner", User.Type).
			Ref("owned").
			Unique(),
	}
}
