package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/ahashim/web-server/enums"
)

// Interaction holds the schema definition for the Interaction entity.
type Interaction struct {
	ent.Schema
}

// Fields of the Interaction.
func (Interaction) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			GoType(enums.Interaction(0)),
	}
}

// Edges of the Interaction.
func (Interaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("interactions").
			Unique(),
		edge.From("squeak", Squeak.Type).
			Ref("interactions").
			Unique(),
	}
}

// Mixins of the Interaction.
func (Interaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
