package controllers

//
//import (
//	"ecobake/internal/models"
//	"ecobake/pkg/resterrors"
//	"encoding/json"
//	"io"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"net/url"
//	"strings"
//	"time"
//	"unicode/utf8"
//
//	"github.com/gin-gonic/gin"
//	"github.com/nyaruka/phonenumbers"
//
//
//)
//
//type SuccessWebHook struct {
//	Body struct {
//		StkCallback struct {
//			MerchantRequestID string `json:"MerchantRequestID"`
//			CheckoutRequestID string `json:"CheckoutRequestID"`
//			ResultCode        int    `json:"ResultCode"`
//			ResultDesc        string `json:"ResultDesc"`
//			CallbackMetadata  struct {
//				Item []struct {
//					Name  string      `json:"Name"`
//					Value interface{} `json:"Value"`
//				} `json:"Item"`
//			} `json:"CallbackMetadata"`
//			TinyPesaID        string `json:"TinyPesaID"`
//			ExternalReference string `json:"ExternalReference"`
//			Amount            int    `json:"Amount"`
//			Msisdn            string `json:"Msisdn"`
//		} `json:"stkCallback"`
//	} `json:"Body"`
//}
//type ErrorWebHook struct {
//	Body struct {
//		StkCallback struct {
//			MerchantRequestID string `json:"MerchantRequestID"`
//			CheckoutRequestID string `json:"CheckoutRequestID"`
//			ResultCode        int    `json:"ResultCode"`
//			ResultDesc        string `json:"ResultDesc"`
//			TinyPesaID        string `json:"TinyPesaID"`
//			ExternalReference string `json:"ExternalReference"`
//			Amount            int    `json:"Amount"`
//			Msisdn            string `json:"Msisdn"`
//		} `json:"stkCallback"`
//	} `json:"Body"`
//}
//
//const baseURL = "https://tinypesa.com/api/v1/"
//
//type transactionReq struct {
//	Amount        string `json:"amount,omitempty"`
//	PhoneNumber   string `json:"phone_number,omitempty"`
//	AccountNumber string `json:"account_number,omitempty"`
//}
//
//func (r *Repository) InitTransaction(ctx *gin.Context) {
//	var req transactionReq
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		restErr := resterrors.NewBadRequestError("invalid json body.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	v, err := r.paymentService.GetSubscriptionBySuid(ctx, req.AccountNumber)
//	if err != nil || v == 0 {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request.Check account number.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	formData := url.Values{}
//
//	formData.Add("amount", req.Amount)
//	formData.Add("msisdn", req.PhoneNumber)
//	formData.Add("account_no", req.AccountNumber)
//
//	client := &http.Client{}
//	request, err := http.NewRequest("POST", baseURL+"express/initialize", strings.NewReader(formData.Encode()))
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	request.Header.Add("Accept", "application/json")
//	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	request.Header.Add("Apikey", "QReTPVlYXeM")
//
//	response, err := client.Do(request)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//
//		}
//	}(response.Body)
//	//dat, _ := ioutil.ReadAll(response.Body)
//	if response.StatusCode != http.StatusOK {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request.Ensure you have passed the correct details")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	dc := NewStatusOkResponse("Transaction in progress", nil)
//	ctx.AbortWithStatusJSON(dc.Status, dc)
//}
//
//type PaymentModes int
//
//const (
//	Mpesa PaymentModes = iota
//	Stripe
//	Bank
//	PayPal
//)
//
//type PaymentsTransaction struct {
//	ID                     int64       `json:"id"`
//	CreatedAt              time.Time   `json:"created_at"`
//	UpdatedAt              time.Time   `json:"updated_at"`
//	TransactionRef         string      `json:"transaction_ref"`
//	Status                 bool        `json:"status"`
//	TransactionComplete    bool        `json:"transaction_complete"`
//	Data                   interface{} `json:"data"`
//	Code                   string      `json:"code"`
//	PaymentIntegrationType int32       `json:"payment_integration_type"`
//	PaymentPurpose         int32       `json:"payment_purpose"`
//	Amount                 int64       `json:"amount"`
//	PaymentMode            int32       `json:"payment_mode"`
//	Currency               string      `json:"currency"`
//	PhoneNumber            string      `json:"phone_number"`
//}
//
//func (r *Repository) WebHook(ctx *gin.Context) {
//	dc, _ := ioutil.ReadAll(ctx.Request.Body)
//
//	var data SuccessWebHook
//
//	err := json.Unmarshal(dc, &data)
//	if err != nil {
//		log.Println(err)
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	if data.Body.StkCallback.ResultCode != 0 {
//		log.Println(err)
//		restErr := resterrors.NewBadRequestError(data.Body.StkCallback.ResultDesc)
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	var pmt = Mpesa
//	d := models.PaymentsTransaction{
//		TransactionRef:         data.Body.StkCallback.CheckoutRequestID,
//		Status:                 true,
//		TransactionComplete:    true,
//		Data:                   data,
//		Code:                   data.Body.StkCallback.MerchantRequestID,
//		PaymentIntegrationType: 1,
//		PaymentPurpose:         1,
//		Amount:                 int64(data.Body.StkCallback.Amount),
//		PaymentMode:            int32(pmt),
//		Currency:               "KES",
//		PhoneNumber:            data.Body.StkCallback.Msisdn,
//	}
//
//	id, err := r.paymentService.CreatePaymentTransaction(ctx, d)
//	if err != nil {
//		log.Println(err)
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	log.Println(data.Body.StkCallback.ExternalReference)
//
//	//TODO:Finish payment controllers
//	//TODO:Finish web ui
//	//TODO:Finish Filter and search
//
//	v, err := r.paymentService.GetSubscriptionBySuid(ctx, data.Body.StkCallback.ExternalReference)
//	if err != nil || v == 0 {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request.Check account number.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	err = r.paymentService.CreatePayment(ctx, id, v)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request.")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//
//	res := NewStatusOkResponse(data.Body.StkCallback.ResultDesc, nil)
//	ctx.JSON(res.Status, res)
//}
//
//func (r *Repository) ListTransaction(ctx *gin.Context) {
//	data, _, err := r.paymentService.ListPaymentsTransactions(ctx)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	dc := NewStatusOkResponse("All Transactions", data)
//	ctx.JSON(dc.Status, dc)
//}
//
//func trimFirstRune(s string) string {
//	_, i := utf8.DecodeRuneInString(s)
//	return s[i:]
//}
//
//func (r *Repository) ListUserTransaction(ctx *gin.Context) {
//	number := ctx.Query("number")
//	num, err := phonenumbers.Parse(number, "KE")
//	if !phonenumbers.IsValidNumber(num) || err != nil {
//		restErr := resterrors.NewBadRequestError("Phone number provided is invalid")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	// format it using national format
//	formattedNum := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
//	pn := strings.Replace(trimFirstRune(formattedNum), " ", "", -1)
//
//	data, _, err := r.paymentService.ListPaymentsTransactionsByUser(ctx, pn)
//	if err != nil {
//		restErr := resterrors.NewBadRequestError("An error occurred processing the request")
//		ctx.AbortWithStatusJSON(restErr.Status, restErr)
//		return
//	}
//	dc := NewStatusOkResponse("All Transactions", data)
//	ctx.JSON(dc.Status, dc)
//}
//
//func (r *Repository) name() {
//
//}
