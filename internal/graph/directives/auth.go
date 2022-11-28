package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func Auth(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
	//TODO: rework this.
	// gc, _ := randomcode.GinContextFromContext(ctx)
	// log.Println(obj)
	// tokenData, _ := gc.Get("authorization_payload")
	// if tokenData == nil {
	// 	return nil, &gqlerror.Error{
	// 		Message: "Access Denied",
	// 	}
	// }

	return next(ctx)
}

// func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
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
