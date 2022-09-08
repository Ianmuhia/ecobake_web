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

//func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
//	ginContext := ctx.Value("GinContextKey")
//	if ginContext == nil {
//		err := fmt.Errorf("could not retrieve gin.Context")
//		return nil, err
//	}
//
//	gc, ok := ginContext.(*gin.Context)
//	if !ok {
//		err := fmt.Errorf("gin.Context has wrong type")
//		return nil, err
//	}
//	return gc, nil
//}
//func (r *mutationResolver) CheckPasswordHash(password, passwordHash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
//	return err == nil
//}
