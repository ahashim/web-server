package schema

import "entgo.io/ent"

// Squeak holds the schema definition for the Squeak entity.
type Squeak struct {
	ent.Schema
}

// Fields of the Squeak.
func (Squeak) Fields() []ent.Field {
	return nil
}

// Edges of the Squeak.
func (Squeak) Edges() []ent.Edge {
	return nil
}
