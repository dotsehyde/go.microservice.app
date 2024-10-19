package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("code").Unique(),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return []ent.Edge{}
}
