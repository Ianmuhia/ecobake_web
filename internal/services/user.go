package services

import (
	"bytes"
	"context"
	"ecobake/cmd/config"
	"ecobake/ent"
	user2 "ecobake/ent/user"
	"ecobake/internal/models"
	"ecobake/internal/postgresql/db"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type usersService struct {
	q      *db.Queries
	cfg    *config.AppConfig
	client *ent.Client
}

func NewUsersService(q db.DBTX, cfg *config.AppConfig, client *ent.Client) *usersService {
	return &usersService{q: db.New(q), cfg: cfg, client: client}
}

type UsersService interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	UpdateUserStatus(context.Context, string) (*models.User, error)
	UpdateUserImage(ctx context.Context, email, imageName string) (*models.User, error)
	GetAllUsers(ctx context.Context) (int, []*models.User, error)
	VerifyPasswordResetCode(key string) VerificationData
	UpdateUserDetails(context.Context, int64, *models.User) (*models.User, error)
	UpdateUserPassword(context.Context, int64, string) error
	DeleteUser(context.Context, int64) error
	CleanDB()
}

func (us *usersService) CreateUser(ctx context.Context, user models.User) (*models.User, error) {

	data, err := us.client.User.
		Create().
		SetEmail(user.Email).
		SetIsVerified(false).
		SetUserName(user.Email).
		SetPhoneNumber(user.PhoneNumber).
		SetPasswordHash(user.Password).Save(ctx)

	if err != nil {
		return new(models.User), err
	}
	return &models.User{
		ID:           int(data.ID),
		CreatedAt:    data.CreatedAt.String(),
		PhoneNumber:  data.PhoneNumber,
		Email:        data.Email,
		ProfileImage: fmt.Sprintf("%s/%s/%s", us.cfg.StorageURL.String(), us.cfg.StorageBucket, data.ProfileImage),
	}, nil

}
func (us *usersService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := us.client.User.Query().Where(user2.Email(email)).Where(user2.IsVerified(false)).Only(ctx)
	if err != nil {
		return new(models.User), err
	}
	return &models.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.String(),
		//PasswordHash: user.PasswordHash,
		//UserName:     user.UserName,
		Password:     user.PasswordHash,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		ProfileImage: fmt.Sprintf("%s/%s/%s", us.cfg.StorageURL.String(), us.cfg.StorageBucket, user.ProfileImage),
		//IsVerified:   user.IsVerified,
	}, nil
}

func (us *usersService) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	//user, err := us.q.GetUserById(ctx, id)
	//
	//if err != nil {
	//	return new(models.User), err
	//}
	//
	//return &models.User{
	//	ID:          user.ID,
	//	CreatedAt:   user.CreatedAt.Time,
	//	UpdatedAt:   user.UpdatedAt.Time,
	//	UserName:    user.UserName,
	//	PhoneNumber: user.PhoneNumber,
	//
	//	Email:        user.Email,
	//	ProfileImage: fmt.Sprintf("%s/%s/%s", us.cfg.StorageURL.String(), us.cfg.StorageBucket, user.ProfileImage.String),
	//	IsVerified:   user.IsVerified.Bool,
	//}, nil
	panic('e')

}
func (us *usersService) DeleteUser(ctx context.Context, id int64) error {
	err := us.q.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (us *usersService) GetAllUsers(ctx context.Context) (int, []*models.User, error) {
	//users, err := us.client.User.Query().All(ctx)
	//d := make([]*models.User, len(users))
	//if err != nil {
	//	return 0, d, err
	//}
	//
	//for i, user := range users {
	//	d[i] = &models.User{
	//		ID:           int64(user.ID),
	//		CreatedAt:    user.CreatedAt,
	//		UpdatedAt:    user.UpdatedAt,
	//		UserName:     user.UserName,
	//		Email:        user.Email,
	//		PhoneNumber:  user.PhoneNumber,
	//		ProfileImage: fmt.Sprintf("%s/%s/%s", us.cfg.StorageURL.String(), us.cfg.StorageBucket, user.ProfileImage),
	//		IsVerified:   user.IsVerified,
	//	}

	//}

	//total := len(d)
	//return total, d, nil
	panic('e')

}
func (us *usersService) UpdateUserImage(ctx context.Context, email, imageName string) (*models.User, error) {
	//v := db.UpdateUserProfileImageParams{
	//	ProfileImage: sql.NullString{
	//		String: imageName,
	//		Valid:  true,
	//	},
	//	Email: email,
	//}
	//user, err := us.q.UpdateUserProfileImage(ctx, v)
	//if err != nil {
	//	log.Println(err)
	//	return new(models.User), err
	//}
	//return &models.User{
	//	//ID:           string(int64(user.ID)),
	//	CreatedAt: user.CreatedAt.Time,
	//	UpdatedAt: user.UpdatedAt.Time,
	//	//UserName:     user.UserName,
	//	PhoneNumber:  user.PhoneNumber,
	//	Email:        user.Email,
	//	ProfileImage: fmt.Sprintf("%s/%s/%s", us.cfg.StorageURL.String(), us.cfg.StorageBucket, user.ProfileImage.String),
	//	IsVerified:   user.IsVerified.Bool,
	//}, nil
	panic('e')

}
func (us *usersService) UpdateUserStatus(ctx context.Context, email string) (*models.User, error) {
	//_, err := us.client.User.Update().Where(user2.Email(email)).SetIsVerified(true).Save(ctx)
	//if err != nil {
	//	return new(models.User), err
	//}
	//return new(models.User), nil
	panic('e')

}
func (us *usersService) UpdateUserDetails(ctx context.Context, id int64, userModel *models.User) (*models.User, error) {
	//
	//d := db.UpdateUserParams{
	//	UserName:    userModel.UserName,
	//	Email:       userModel.Email,
	//	PhoneNumber: userModel.PhoneNumber,
	//	ID:          id,
	//}
	//user, err := us.q.UpdateUser(ctx, d)
	//if err != nil {
	//	log.Println(err)
	//	return new(models.User), err
	//}
	//return &models.User{
	//	ID:           user.ID,
	//	CreatedAt:    user.CreatedAt.Time,
	//	UpdatedAt:    user.UpdatedAt.Time,
	//	UserName:     user.UserName,
	//	PhoneNumber:  user.PhoneNumber,
	//	Email:        user.Email,
	//	ProfileImage: fmt.Sprintf("%s/%s/%s", us.cfg.StorageURL.String(), us.cfg.StorageBucket, user.ProfileImage.String),
	//	IsVerified:   user.IsVerified.Bool,
	//}, nil
	panic('e')
}
func (us *usersService) UpdateUserPassword(ctx context.Context, id int64, newPasswd string) error {
	_, err := us.client.User.Update().Where(user2.ID(int(id))).SetPasswordHash(newPasswd).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (us *usersService) VerifyPasswordResetCode(key string) VerificationData {
	var a VerificationData

	data := us.cfg.RedisConn.Get(context.TODO(), key)
	cmdb, err := data.Bytes()
	if err != nil {
		log.Println(err)
		return a
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&a); err != nil {
		log.Println(err)
		return a
	}
	return a
}
func (us *usersService) GetRefreshToken(key string) string {
	data, err := us.cfg.RedisConn.Get(context.TODO(), key).Result()
	if err != nil {
		log.Println(err)
		return ""
	}

	return data
}
func (us *usersService) CleanDB() {
	go func() {
		for now := range time.Tick(time.Second) {
			fmt.Println(now)
		}
	}()
}
