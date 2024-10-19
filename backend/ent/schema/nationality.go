package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Nationality holds the schema definition for the Nationality entity.
type Nationality struct {
	ent.Schema
}

func (Nationality) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Nationality.
func (Nationality) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
	}
}

// Edges of the Nationality.
func (Nationality) Edges() []ent.Edge {
	return nil
}
