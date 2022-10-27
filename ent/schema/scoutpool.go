package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"github.com/ahashim/web-server/types"
)

// ScoutPool holds the schema definition for the ScoutPool entity.
type ScoutPool struct {
	ent.Schema
}

// Fields of the ScoutPool.
func (ScoutPool) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").
			GoType(new(types.Uint256)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}),
	}
}

// Edges of the ScoutPool.
func (ScoutPool) Edges() []ent.Edge {
	return nil
}
