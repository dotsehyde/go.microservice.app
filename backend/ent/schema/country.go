package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("code").Unique(),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}
