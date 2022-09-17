package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Favourites holds the schema definition for the Favourites entity.
type Favourites struct {
	ent.Schema
}

// Fields of the Favourites.
func (Favourites) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Favourites.
func (Favourites) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}).Ref("favourites").
			Unique(),
		edge.From("user", User.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}).Ref("favourites").
			Unique(),
	}
}
