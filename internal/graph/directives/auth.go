package directives

import (
	"context"
	"ecobake/internal/graph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	gc, _ := graph.GinContextFromContext(ctx)
	log.Println(obj)
	tokenData, _ := gc.Get("authorization_payload")
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
