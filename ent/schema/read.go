package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.Int("user_id"),
		field.Int("sentence_id"),
		field.Int("result"),
	}
}

func (Read) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "sentence_id").
			Unique(),
	}
}

// Edges of the Read.
func (Read) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Field("user_id").
			Ref("reads").
			Unique().
			Required(),
		edge.From("sentence", Sentense.Type).
			Field("sentence_id").
			Ref("reads").
			Unique().
			Required(),
	}
}
