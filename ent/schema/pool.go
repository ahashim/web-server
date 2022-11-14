package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
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
	}
}

// Edges of the Pool.
func (Pool) Edges() []ent.Edge {
	return nil
}
