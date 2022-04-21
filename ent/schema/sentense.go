package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sentense holds the schema definition for the Sentense entity.
type Sentense struct {
	ent.Schema
}

// Fields of the Sentense.
func (Sentense) Fields() []ent.Field {
	return []ent.Field{
		// "我是中国人"
		field.String("chinese"),
		// "wo3 shi4 zhong1 guo2 ren2"
		field.String("pinyin"),
		// "I'm chinese"
		field.String("english"),
	}
}

// Edges of the Sentense.
func (Sentense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reads", Read.Type),
	}

}
