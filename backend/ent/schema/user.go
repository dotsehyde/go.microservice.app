package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("phone").Unique().NotEmpty(),
		field.Enum("role").Values("farmer", "worker", "supplier", "officer").Default("farmer"),
		field.String("id_type").Optional(),
		field.String("id_number").Optional(),
		field.String("id_photo").Optional(),
		field.String("profile_photo").Optional(),
		field.String("address").Optional(),
		field.String("city").Optional(),
		field.String("nationality").Optional(),
		field.Enum("language").Values("en", "fr").Default("en"),
		field.Int("country_id").Default(1),
		field.Bool("is_worker").Default(false),
		field.Bool("is_verified").Default(false),
		field.Bool("is_blocked").Default(false),
		field.String("reason").Optional(),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country", Country.Type).
			Ref("user").
			Field("country_id").
			Unique().
			Required(),
	}
}
