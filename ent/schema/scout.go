package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/types"
)

// Scout holds the schema definition for the Scout entity.
type Scout struct {
	ent.Schema
}

// Fields of the Scout.
func (Scout) Fields() []ent.Field {
	return []ent.Field{
		field.Int("shares").
			GoType(new(types.Uint256)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}),
	}
}

// Edges of the Scout.
func (Scout) Edges() []ent.Edge {
	return nil
}
