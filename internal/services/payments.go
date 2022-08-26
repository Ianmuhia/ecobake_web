package services

import (
	"context"
	"ecobake/cmd/config"
	"ecobake/internal/models"
	"ecobake/internal/postgresql/db"
	"encoding/json"

	"time"

	"github.com/jackc/pgtype"
)

type PaymentsService interface {
	CreatePaymentTransaction(ctx context.Context, pm models.PaymentsTransaction) (int64, error)
	ListPaymentsTransactions(ctx context.Context) ([]*models.PaymentsTransaction, int, error)
	GetPaymentTransactions(ctx context.Context, id int64) (*models.PaymentsTransaction, error)
	GetTotalPayments(ctx context.Context) float64
	ListPaymentsTransactionsByUser(ctx context.Context, phoneNumber string) ([]*models.PaymentsTransaction, float64, error)
	CreateSubscription(ctx context.Context, csp createSubParams) error
	DeleteSubscription(ctx context.Context, id int64) error
	ListSubscriptionsByPlan(ctx context.Context, id int64) (data interface{}, err error)
	DeletePlan(ctx context.Context, id int64) error
	CreatePayment(ctx context.Context, pid, sid int64) error
	UpdatePayment(ctx context.Context, pid, sid int64) error
	ListUserPayment(ctx context.Context, id int64) (interface{}, error)
	DeletePayment(ctx context.Context, id int64) error
	GetSubscriptionBySuid(ctx context.Context, subUid string) (int64, error)
}

type paymentsService struct {
	cfg *config.AppConfig
	q   *db.Queries
}

func NewPaymentsService(cfg *config.AppConfig, q db.DBTX) *paymentsService {
	return &paymentsService{cfg: cfg, q: db.New(q)}
}

func value(c interface{}) ([]byte, error) {
	j, err := json.Marshal(c)
	return j, err
}

func (ps *paymentsService) CreatePaymentTransaction(ctx context.Context, pm models.PaymentsTransaction) (int64, error) {
	v, e := value(pm)
	if e != nil {
		ps.cfg.Logger.Println(e)
	}
	data, err := ps.q.CreatePaymentTransaction(ctx, db.CreatePaymentTransactionParams{
		TransactionRef:         pm.TransactionRef,
		Status:                 pm.Status,
		TransactionComplete:    true,
		Data:                   pgtype.JSONB{Bytes: v, Status: pgtype.Present},
		Code:                   pm.Code,
		PaymentIntegrationType: pm.PaymentIntegrationType,
		PaymentPurpose:         pm.PaymentPurpose,
		Amount:                 float64(pm.Amount),
		PaymentMode:            pm.PaymentMode,
		Currency:               "KES",
		PhoneNumber:            pm.PhoneNumber,
	})
	if err != nil {
		return 0, err
	}
	return data, nil
}

func (ps *paymentsService) ListPaymentsTransactions(ctx context.Context) ([]*models.PaymentsTransaction, int, error) {
	data, err := ps.q.ListPaymentTransactions(ctx)
	bks := make([]*models.PaymentsTransaction, len(data))
	f := make(map[string]interface{}, 0)
	for _, v := range data {

		err := json.Unmarshal(v.Data.Bytes, &f)
		if err != nil {
			ps.cfg.Logger.Println(err)
		}
		//TODO: Rework this loop
		g := &models.PaymentsTransaction{
			ID:                     v.ID,
			CreatedAt:              v.CreatedAt,
			UpdatedAt:              v.UpdatedAt,
			TransactionRef:         v.TransactionRef,
			Status:                 v.Status,
			TransactionComplete:    v.TransactionComplete,
			Data:                   f,
			Code:                   v.Code,
			PaymentIntegrationType: v.PaymentIntegrationType,
			PaymentPurpose:         v.PaymentPurpose,
			Amount:                 int64(v.Amount),
			PaymentMode:            v.PaymentMode,
			Currency:               v.Currency,
			PhoneNumber:            v.PhoneNumber,
		}
		bks = append(bks, g)
	}

	if err != nil {
		return bks, 0, err
	}
	return bks, len(bks), nil
}

func (ps *paymentsService) GetPaymentTransactions(ctx context.Context, id int64) (*models.PaymentsTransaction, error) {
	v, err := ps.q.GetPaymentTransactionByID(ctx, id)
	f := make(map[string]interface{}, 0)

	er := json.Unmarshal(v.Data.Bytes, &f)
	if er != nil {
		ps.cfg.Logger.Println(err)
	}
	data := &models.PaymentsTransaction{
		ID:                     v.ID,
		CreatedAt:              v.CreatedAt,
		UpdatedAt:              v.UpdatedAt,
		TransactionRef:         v.TransactionRef,
		Status:                 v.Status,
		TransactionComplete:    v.TransactionComplete,
		Data:                   f,
		Code:                   v.Code,
		PaymentIntegrationType: v.PaymentIntegrationType,
		PaymentPurpose:         v.PaymentPurpose,
		Amount:                 int64(v.Amount),
		PaymentMode:            v.PaymentMode,
		Currency:               v.Currency,
	}
	if err != nil {
		ps.cfg.Logger.Println(err)
		return data, err
	}

	return data, nil
}

func sum(array []float64) float64 {
	result := 0.0
	for _, v := range array {
		result += v
	}
	return result
}

func (ps *paymentsService) ListPaymentsTransactionsByUser(ctx context.Context, phoneNumber string) ([]*models.PaymentsTransaction, float64, error) {

	data, err := ps.q.ListPaymentTransactionsByUser(ctx, phoneNumber)
	transactions := make([]*models.PaymentsTransaction, 0)

	amtTotal := make([]float64, 0)

	f := make(map[string]interface{}, 0)
	for _, v := range data {
		amtTotal = append(amtTotal, v.Amount)
		err := json.Unmarshal(v.Data.Bytes, &f)
		if err != nil {
			ps.cfg.Logger.Println(err)
		}
		g := &models.PaymentsTransaction{
			ID:                     v.ID,
			CreatedAt:              v.CreatedAt,
			UpdatedAt:              v.UpdatedAt,
			TransactionRef:         v.TransactionRef,
			Status:                 v.Status,
			TransactionComplete:    v.TransactionComplete,
			Data:                   f,
			Code:                   v.Code,
			PaymentIntegrationType: v.PaymentIntegrationType,
			PaymentPurpose:         v.PaymentPurpose,
			Amount:                 int64(v.Amount),
			PaymentMode:            v.PaymentMode,
			Currency:               v.Currency,
			PhoneNumber:            v.PhoneNumber,
		}
		transactions = append(transactions, g)
	}

	total := sum(amtTotal)

	if err != nil {
		return transactions, 0, err
	}
	return transactions, total, nil
}

func (ps *paymentsService) GetTotalPayments(ctx context.Context) float64 {
	amount, err := ps.q.GetTotalPayments(ctx)
	if err != nil {
		return 0
	}
	return amount.(float64)

}

type createSubParams struct {
	UserID    int64
	PlanID    int64
	Status    bool
	SubUid    string
	StartDate time.Time
	EndDate   time.Time
}
