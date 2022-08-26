package controllers

//
//import (
//	"bytes"
//	"encoding/gob"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"log"
//	"net/http"
//	"net/url"
//	"strconv"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"github.com/jackc/pgx"
//
//	"fainda/internal/models"
//	"fainda/internal/services"
//	"fainda/pkg/randomcode"
//	"fainda/pkg/resterrors"
//)
//
//type createUserRequest struct {
//	Username        string  `json:"username" binding:"required"`
//	Password        string  `json:"password" binding:"required,min=6"`
//	PasswordConfirm string  `json:"password_confirm" binding:"required,min=6"`
//	Email           string  `json:"email" binding:"required,email"`
//	PhoneNumber     string  `json:"phone_number" binding:"required"`
//	Latitude        float64 `json:"latitude" binding:"required"`
//	Longitude       float64 `json:"longitude" binding:"required"`
//}
//
//type userResponse struct {
//	ID           int       `json:"ID"`
//	CreatedAt    time.Time `json:"created_at"`
//	Username     string    `json:"username"`
//	Email        string    `json:"email"`
//	ProfileImage string    `json:"profile_image"`
//	IsVerified   bool      `json:"is_verified"`
//	PhoneNumber  string    `json:"phone_number"`
//}
//
//func (u *userResponse) FormatDate() {
//	layout := "2006-01-02"
//	t, _ := time.Parse(layout, u.CreatedAt.String())
//	u.CreatedAt = t
//}
//
//type registerUserResponse struct {
//	Message  string `json:"data"`
//	EmailURL string `json:"email-url"`
//}
//
//func newUserResponse(user *models.User) userResponse {
//	return userResponse{
//		ID:           int(user.ID),
//		CreatedAt:    user.CreatedAt,
//		Username:     user.UserName,
//		Email:        user.Email,
//		ProfileImage: user.ProfileImage,
//		PhoneNumber:  user.PhoneNumber,
//		IsVerified:   user.IsVerified,
//	}
//}
//
//// RegisterUser new user.
//func (r *Repository) RegisterUser(ctx *gin.Context) {
//	var req createUserRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	user := models.User{
//		UserName:    req.Username,
//		Email:       req.Email,
//		Password:    req.Password,
//		PhoneNumber: req.PhoneNumber,
//		Geo: models.Geo{
//			Lat: req.Latitude,
//			Lng: req.Longitude,
//		},
//	}
//	if !user.ComparePassword(req.PasswordConfirm) {
//		r.app.Logger.Println("password and password re-enter did not match")
//		e := resterrors.NewBadRequestError("password and password re-enter did not match")
//		ctx.AbortWithStatusJSON(e.Status, e)
//		return
//	}
//	err := user.Hash()
//	if err != nil {
//		data := resterrors.NewBadRequestError("An error occurred registering the user.")
//		ctx.AbortWithStatusJSON(data.Status, data)
//		return
//	}
//	data, saveErr := r.userServices.CreateUser(ctx, user)
//	if saveErr != nil {
//		err := resterrors.NewBadRequestError("Could not save user.")
//		ctx.AbortWithStatusJSON(err.Status, err)
//		log.Println(saveErr)
//		return
//	}
//	userChan := make(chan *models.User)
//	go r.searchService.IndexUser(userChan)
//	userChan <- data
//
//	dg := <-r.app.ZincRcvChan
//	log.Println("finished indexing user: %w", dg)
//
//	code := randomcode.Code(6)
//	err = r.app.RedisConn.Set(ctx, user.Email, code, time.Minute*time.Duration(r.app.PasswordResetCodeExpiry)).Err()
//	if err != nil {
//		log.Println(err)
//		restErr := resterrors.NewBadRequestError("Unable to set user verification code. Please try again later")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	u := &url.URL{
//		Scheme: "http",
//		Host:   "2423-41-80-96-152.ngrok.io",
//		Path:   "/api/v1/auth/users/email/verify/",
//	}
//	rq := u.Query()
//	rq.Set("code", code)
//	rq.Set("email", user.Email)
//
//	u.RawQuery = rq.Encode()
//	from := "ianmuhia3@gmail.com"
//	to := user.Email
//	subject := "Email Verification"
//	mailType := services.MailConfirmation
//	mailData := &services.MailData{
//		Username: user.UserName,
//		Code:     code,
//		URL:      u,
//	}
//
//	mailReq := r.mailService.NewMail(from, to, subject, mailType, mailData)
//
//	v, err := json.Marshal(mailReq)
//	if err != nil {
//		r.app.Logger.Println(err)
//	}
//	err = r.natService.Publish("Mail.Verification", v)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred publishing the mail")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	message := fmt.Sprintf("Thank %s you for creating and account.Please verify your email %s code is %s", user.UserName, user.Email, code)
//	response := registerUserResponse{
//		Message:  message,
//		EmailURL: "localhost:8090/api/users/",
//	}
//
//	ctx.JSON(http.StatusOK, response)
//}
//
//type loginUserRequest struct {
//	Email    string `json:"email" binding:"required"`
//	Password string `json:"password" binding:"required,min=6"`
//}
//
//type loginUserResponse struct {
//	AccessToken  string       `json:"access_token"`
//	RefreshToken string       `json:"refresh_token"`
//	User         userResponse `json:"user"`
//}
//
//func (r *Repository) Login(ctx *gin.Context) {
//	var req loginUserRequest
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	user, err := r.userServices.GetUserByEmail(ctx, req.Email)
//	if errors.Is(err, pgx.ErrNoRows) {
//		data := resterrors.NewBadRequestError("The user does not exist.Please create an account to continue")
//		ctx.AbortWithStatusJSON(data.Status, data)
//		return
//	}
//
//	log.Println(randomcode.Code(6))
//
//	//if !user.IsVerified {
//	//	data := resterrors.NewBadRequestError("Please verify your email address to login")
//	//	ctx.AbortWithStatusJSON(data.Status, data)
//	//	return
//	//}
//	ok := user.CheckPasswordHash(req.Password)
//
//	user.FormatDate(user.CreatedAt.String())
//
//	if !ok {
//		data := resterrors.NewBadRequestError("invalid email or password ")
//		ctx.AbortWithStatusJSON(data.Status, data)
//		return
//	}
//
//	// t := r.app.TokenLifeTime
//	duration := 30 * time.Hour
//	rtduration := time.Duration(time.Now().Add(time.Hour * 100).Unix())
//
//	accessToken, err := r.tokenService.CreateToken(user.Email, user.UserName, duration, rtduration, user.ID)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred please login again")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	refreshToken, err := r.tokenService.CreateRefreshToken(user.Email, user.UserName, duration, rtduration, user.ID)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred please login again")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	val := 100
//	go func() {
//		dc := r.app.RedisConn.Set(ctx, strconv.Itoa(int(user.ID)), refreshToken, time.Minute*time.Duration(val)).Err()
//		if dc != nil {
//			restErr := resterrors.NewBadRequestError("Unable to send password reset code. Please try again later")
//			ctx.AbortWithStatusJSON(restErr.Status, restErr)
//			return
//		}
//	}()
//
//	response := loginUserResponse{
//		AccessToken:  accessToken,
//		RefreshToken: refreshToken,
//		User:         newUserResponse(user),
//	}
//
//	res := NewStatusOkResponse("Login Successful.", response)
//	ctx.JSON(res.Status, res)
//}
//
//type loginCodeRequest struct {
//	Email string `json:"email"`
//}
//
//func (r *Repository) LoginCode(ctx *gin.Context) {
//	var req loginCodeRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	user, err := r.userServices.GetUnVerifiedUserByEmail(ctx, req.Email)
//	if err != nil {
//		resp := resterrors.NewBadRequestError("No account associated with that email account.")
//		ctx.AbortWithStatusJSON(resp.Status, resp)
//		return
//	}
//	code := randomcode.Code(6)
//
//	err = r.app.RedisConn.Set(ctx, user.Email, code, time.Minute*time.Duration(r.app.PasswordResetCodeExpiry)).Err()
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("Unable to set user verification code. Please try again later")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	from := "me@here.com"
//	to := user.Email
//	subject := "Email Verification for Pass Change"
//	mailType := services.MailConfirmation
//	mailData := &services.MailData{
//		Username: user.UserName,
//		Code:     code,
//	}
//
//	mailReq := r.mailService.NewMail(from, to, subject, mailType, mailData)
//
//	v, _ := json.Marshal(mailReq)
//	// err = r.natService.CreateStream("stream for testing stuff and mail", "test_str")
//	// if err != nil {
//	//	r.app.Logger.Println(err)
//	// }
//	err = r.natService.Publish("Mail.Verification", v)
//	if err != nil {
//		r.app.Logger.Println(err)
//	}
//
//	message := fmt.Sprintf("Thank %s you for creating and account.Please verify your email %s code is %s", user.UserName, user.Email, code)
//
//	response := registerUserResponse{
//		Message:  message,
//		EmailURL: "localhost:8090/api/users/",
//	}
//
//	res := NewStatusOkResponse("Password reset code sent successfully.", response)
//	ctx.JSON(res.Status, res)
//}
//
//func (r *Repository) VerifyEmailCode(ctx *gin.Context) {
//	email := ctx.Query("email")
//
//	data := r.mailService.VerifyMailCode(email)
//	if ctx.Query("code") == data {
//		r.mailService.RemoveMailCode(email)
//		user, err := r.userServices.UpdateUserStatus(ctx, email)
//		if err != nil {
//			resp := resterrors.NewBadRequestError("Error Updating user status")
//			log.Println(resp)
//			ctx.HTML(resp.Status, "404", resp)
//			return
//		}
//		userChan := make(chan *models.User)
//		go r.searchService.IndexUser(userChan)
//		userChan <- user
//
//		r.app.ZincChan <- user
//
//		dg := <-r.app.ZincRcvChan
//		log.Println("finished indexing user: %w", dg)
//
//		resp := SuccessResponse{
//			TimeStamp: time.Now(),
//			Message:   "Email has been verified you can now login to your account",
//			Status:    http.StatusOK,
//			Data:      nil,
//		}
//
//		ctx.HTML(resp.Status, "welcome", resp)
//		return
//	}
//
//	resp := resterrors.NewBadRequestError("Email verification failed invalid code provided")
//	ctx.HTML(resp.Status, "404", resp)
//}
//
//type getAllUsersResponse struct {
//	Total int         `json:"total"`
//	Users interface{} `json:"users"`
//}
//
//func (r *Repository) GetAllUsers(ctx *gin.Context) {
//	_ = r.GetPayloadFromContext(ctx)
//	total, users, err := r.userServices.GetAllUsers(ctx)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("Error getting all users.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	data := getAllUsersResponse{
//		Total: total,
//		Users: users,
//	}
//	resp := NewStatusOkResponse("All users.", data)
//	ctx.JSON(resp.Status, resp)
//}
//
//func (r *Repository) GetPayloadFromContext(ctx *gin.Context) *models.Payload {
//	payload, exists := ctx.Get("authorization_payload")
//	if !exists {
//		restErr := resterrors.NewBadRequestError("could not get auth_payload from context")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return nil
//	}
//	data := payload.(*models.Payload)
//	_, err := r.userServices.GetUserByID(ctx, data.ID)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("Error Processing request")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return nil
//	}
//	return data
//}
//
//func (r *Repository) GetUser(ctx *gin.Context) {
//	id := r.GetInt(ctx, "id")
//	data, err := r.userServices.GetUserByID(ctx, id)
//	if err != nil {
//		restErr := resterrors.NewNotFoundError("user does not exits")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	res := NewStatusOkResponse("Successfully got user", data)
//	ctx.JSON(res.Status, res)
//}
//
//func (r *Repository) GetUserProfile(ctx *gin.Context) {
//	//user := r.GetPayloadFromContext(ctx)
//	id := r.GetUserPayload(ctx)
//	data, err := r.userServices.GetUserByID(ctx, id)
//	if err != nil {
//		log.Println(err)
//		restErr := resterrors.NewNotFoundError("user does not exits")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	res := NewStatusOkResponse("Successfully got user", data)
//	ctx.JSON(res.Status, res)
//}
//
//func (r *Repository) DeleteUser(ctx *gin.Context) {
//	r.GetPayloadFromContext(ctx)
//
//	id, _ := strconv.Atoi(ctx.Param("id"))
//
//	err := r.userServices.DeleteUser(ctx, int64(id))
//	if err != nil {
//		restErr := resterrors.NewNotFoundError("user does not exits")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	res := NewStatusOkResponse("Successfully deleted user", nil)
//	ctx.JSON(res.Status, res)
//}
//
//func (r *Repository) UpdateUserProfileImage(ctx *gin.Context) {
//	data := r.GetPayloadFromContext(ctx)
//	file, m, err := ctx.Request.FormFile("profile_image")
//
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("Please attach image to the request")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	fileContentType := m.Header["Content-Type"][0]
//
//	uploadFile, err := r.storageService.UploadFile(m.Filename, file, m.Size, fileContentType)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("could not upload image to server")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	userDoc, err := r.userServices.UpdateUserImage(ctx, data.Username, uploadFile.Key)
//
//	userchan := make(chan *models.User)
//	go r.searchService.UpdateUserDoc(userchan)
//	userchan <- userDoc
//
//	dg := <-r.app.ZincRcvChan
//	log.Println("finished indexing user: %w", dg)
//	if err != nil {
//		data := resterrors.NewBadRequestError("Error Processing upload profile image request")
//		ctx.AbortWithStatusJSON(data.Status, data)
//		return
//	}
//	resp := NewStatusOkResponse("Profile image upload successful", nil)
//	ctx.JSON(resp.Status, resp)
//}
//
//type GetPasswordResetCode struct {
//	Email string `json:"email"`
//}
//
//func (r *Repository) ForgotPassword(ctx *gin.Context) {
//	var req GetPasswordResetCode
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	user, err := r.userServices.GetUserByEmail(ctx, req.Email)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("User with that email does not exits")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	// Send verification mail.
//	from := "me@gmail.com"
//	to := user.Email
//	subject := "Password Reset for User"
//
//	mailType := services.PassReset
//	mailData := &services.MailData{
//		Username: user.UserName,
//		Code:     randomcode.Code(6),
//	}
//	mailReq := &services.Mail{
//		From:     from,
//		To:       to,
//		Subject:  subject,
//		Body:     mailData,
//		MailType: mailType,
//	}
//
//	log.Println(mailReq)
//	v, _ := json.Marshal(mailReq)
//	// err = r.natService.CreateStream("stream for testing stuff and mail", "test_str")
//	// if err != nil {
//	//	r.app.Logger.Println(err)
//	// }
//	err = r.natService.Publish("Mail.Verification", v)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("Unable to send mail.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	// store the password reset code to db
//	verificationData := &services.VerificationData{
//		Email:     user.Email,
//		Code:      mailData.Code,
//		Type:      string(rune(services.PassReset)),
//		ExpiresAt: time.Now().Add(time.Minute * time.Duration(r.app.PasswordResetCodeExpiry)),
//	}
//
//	var b bytes.Buffer
//	if err := gob.NewEncoder(&b).Encode(verificationData); err != nil {
//		restErr := resterrors.NewBadRequestError("Unable to send password reset code. Please try again later")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	err = r.app.RedisConn.Set(ctx, verificationData.Email, b.Bytes(), time.Minute*time.Duration(r.app.PasswordResetCodeExpiry)).Err()
//	if err != nil {
//		log.Println(err)
//		restErr := resterrors.NewBadRequestError("Unable to send password reset code. Please try again later")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	resp := NewStatusOkResponse("Password reset code sent successfully", mailData.Code)
//	ctx.JSON(resp.Status, resp)
//}
//
//type PasswordResetCode struct {
//	Code  string `json:"code"`
//	Email string `json:"email"`
//}
//
//// VerifyPassWordResetCode   handles the password reset request.
//func (r *Repository) VerifyPassWordResetCode(ctx *gin.Context) {
//	var req PasswordResetCode
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	data := r.userServices.VerifyPasswordResetCode(req.Email)
//	if req.Code != data.Code {
//		e := resterrors.NewBadRequestError("Invalid code provided.")
//		ctx.AbortWithStatusJSON(e.Status, e)
//		return
//	}
//
//	resp := NewStatusOkResponse("Code verification successful", nil)
//	ctx.JSON(resp.Status, resp)
//}
//
//type PasswordResetReq struct {
//	Password        string `json:"password"`
//	PasswordConfirm string `json:"password_confirm"`
//	Email           string `json:"email"`
//	Code            string `json:"code"`
//}
//
//// ResetPassword  handles the password reset request.
//func (r *Repository) ResetPassword(ctx *gin.Context) {
//	var req PasswordResetReq
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	user, err := r.userServices.GetUserByEmail(ctx, req.Email)
//	if user == nil {
//		log.Println(err)
//		data := resterrors.NewBadRequestError("Unable to reset password. Please try again later")
//		ctx.AbortWithStatusJSON(data.Status, data)
//		return
//	}
//
//	data := r.userServices.VerifyPasswordResetCode(req.Email)
//	if req.Code != data.Code {
//		log.Println(data)
//		e := resterrors.NewBadRequestError("Unable to reset password. Please try again later")
//		ctx.AbortWithStatusJSON(e.Status, e)
//		return
//	}
//
//	if req.Code != data.Code {
//		r.app.Logger.Println("verification code did not match even after verifying PassReset")
//		e := resterrors.NewBadRequestError("Unable to reset password. Please try again later")
//		ctx.AbortWithStatusJSON(e.Status, e)
//		return
//	}
//
//	if req.Password != req.PasswordConfirm {
//		r.app.Logger.Println("password and password re-enter did not match")
//		e := resterrors.NewBadRequestError("Unable to reset password. Please try again later")
//		ctx.AbortWithStatusJSON(e.Status, e)
//		return
//	}
//	var pwd = &models.User{
//		Password: req.Password,
//	}
//	err = pwd.Hash()
//	if err != nil {
//		data := resterrors.NewBadRequestError("An error occurred ")
//		ctx.AbortWithStatusJSON(data.Status, data)
//		return
//	}
//
//	user, err = r.userServices.UpdateUserDetails(ctx, user.ID, pwd)
//	if err != nil {
//		r.app.Logger.Println("update user failed")
//		e := resterrors.NewBadRequestError("Unable to reset password. Please try again later")
//		ctx.AbortWithStatusJSON(e.Status, e)
//		return
//	}
//
//	userchan := make(chan *models.User)
//	go r.searchService.UpdateUserDoc(userchan)
//	userchan <- user
//
//	r.mailService.RemoveMailCode(req.Email)
//	resp := &SuccessResponse{
//		TimeStamp: time.Now(),
//		Message:   "Password reset successful",
//		Status:    http.StatusOK,
//		Data:      nil,
//	}
//
//	ctx.JSON(resp.Status, resp)
//}
//
//type UpdateUserPasswordRequest struct {
//	NewPassword string `json:"new_password"`
//}
//
//type updateUserDetails struct {
//	Username    string  `json:"username"`
//	Email       string  `json:"email"`
//	PhoneNumber string  `json:"phone_number"`
//	Latitude    float64 `json:"latitude" `
//	Longitude   float64 `json:"longitude" `
//}
//
//func (r *Repository) UpdateUserDetails(ctx *gin.Context) {
//	id := r.GetUserPayload(ctx)
//
//	var req updateUserDetails
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	um := &models.User{
//
//		UserName:    req.Username,
//		Email:       req.Email,
//		PhoneNumber: req.PhoneNumber,
//		Geo:         models.Geo{Lat: req.Latitude, Lng: req.Longitude},
//	}
//	userDoc, err := r.userServices.UpdateUserDetails(ctx, id, um)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("Unable to update user details.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	userchan := make(chan *models.User)
//	go r.searchService.UpdateUserDoc(userchan)
//	userchan <- userDoc
//	data := NewStatusOkResponse("Successfully updated profile", nil)
//	ctx.JSON(data.Status, data)
//}
//
//func (r *Repository) UpdateUserPassword(ctx *gin.Context) {
//	id := r.GetUserPayload(ctx)
//
//	var req UpdateUserPasswordRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	newPass := &models.User{
//		Password: req.NewPassword,
//	}
//
//	err := newPass.Hash()
//	if err != nil {
//		data := resterrors.NewBadRequestError("Unable to update password. Please try again later")
//		ctx.AbortWithStatusJSON(data.Status, data)
//		return
//	}
//	err = r.userServices.UpdateUserPassword(ctx, id, newPass.PasswordHash)
//	if err != nil {
//		r.app.Logger.Println("update user failed")
//		e := resterrors.NewBadRequestError("Unable to update password. Please try again later")
//		ctx.AbortWithStatusJSON(e.Status, e)
//		return
//	}
//
//	data := NewStatusOkResponse("Successfully updated password", nil)
//	ctx.JSON(data.Status, data)
//}
