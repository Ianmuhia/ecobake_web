package randomcode

import (
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Code(length int) string {
	max := time.Now().UnixNano()
	d := strconv.FormatInt(max, 10)
	log.Println(d)
	return d[len(d)-length:]
}

// func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
// 	ginContext := ctx.Value("GinContextKey")
// 	if ginContext == nil {
// 		err := fmt.Errorf("could not retrieve gin.Context")
// 		return nil, err
// 	}

// 	gc, ok := ginContext.(*gin.Context)
// 	if !ok {
// 		err := fmt.Errorf("gin.Context has wrong type")
// 		return nil, err
// 	}
// 	return gc, nil
// }

func CheckPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
