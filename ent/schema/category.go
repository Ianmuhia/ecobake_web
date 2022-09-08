package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("icon").Unique().NotEmpty(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}
