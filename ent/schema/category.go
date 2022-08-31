package schema

import (
	"entgo.io/ent"
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
		field.Time("updated_at"),
		field.Time("deleted_at"),
		field.String("icon").Unique().NotEmpty(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
