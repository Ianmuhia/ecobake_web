package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("password_hash").Unique().NotEmpty(),
		field.String("phone_number").Unique().NotEmpty(),
		field.Bool("is_verified").Default(false),
		field.String("profile_image").Default(""),
		field.String("email").
			Unique(),
	}
}

func (User) Edges() []ent.Edge {
	//return []ent.Edge{
	//	edge.To("groups", Group.Type),
	//	edge.To("friends", User.Type),
	//}
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_name", "email").
			Unique(),
	}
}
