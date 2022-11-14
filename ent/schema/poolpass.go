package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
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
			}),
	}
}

// Edges of the PoolPass.
func (PoolPass) Edges() []ent.Edge {
	return nil
}
