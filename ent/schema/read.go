package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Read holds the schema definition for the Read entity.
type Read struct {
	ent.Schema
}

// Fields of the Read.
// result field being
// 0 : flash;
// 1 : done;
// 2 : hard
func (Read) Fields() []ent.Field {
	return []ent.Field{
		field.Int("result"),
	}
}

// Edges of the Read.
func (Read) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("reads").
			Unique().
			Required(),
		edge.From("sentence", Sentense.Type).
			Ref("reads").
			Unique().
			Required(),
	}
}
