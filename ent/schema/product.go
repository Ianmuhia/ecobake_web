package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.String("price").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("ingredients").NotEmpty(),
		field.Float("totalRating").Default(0),
		field.JSON("images", []string{}).
			Default([]string{""}),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}).Ref("product").
			Unique(),
		edge.To("favourites", Favourites.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}).Unique(),
	}
}
