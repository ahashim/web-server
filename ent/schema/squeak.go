package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
				dialect.Postgres: "numeric(78, 0)", // uint256
			}),
	}
}

// Edges of the Squeak.
func (Squeak) Edges() []ent.Edge {
	return nil
}
