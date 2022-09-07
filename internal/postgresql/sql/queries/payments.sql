-- name: ListPaymentTransactions :many
select payment_transactions.id,
       created_at,
       updated_at,
       status,
       transaction_ref,
       transaction_complete,
       data,
       code,
       payment_integration_type,
       payment_purpose,
       amount,
       payment_mode,
       currency,
       phone_number
from payment_transactions;

-- name: GetPaymentTransactionByID :one
select payment_transactions.id,
       created_at,
       updated_at,
       transaction_ref,
       status,
       transaction_complete,
       data,
       code,
       payment_integration_type,
       payment_purpose,
       amount,
       payment_mode,
       currency,
       phone_number
from payment_transactions
where id = @id;

-- name: ListPaymentTransactionsByUser :many
select payment_transactions.id,
       created_at,
       updated_at,
       transaction_ref,
       status,
       transaction_complete,
       data,
       code,
       payment_integration_type,
       payment_purpose,
       amount,
       payment_mode,
       currency,
       phone_number
from payment_transactions
where phone_number = @phone_number;


--TODO:Add indexing for paymentransactions

-- name: CreatePaymentTransaction :one
insert into payment_transactions (created_at,
                                  updated_at,
                                  transaction_ref,
                                  status,
                                  transaction_complete,
                                  data,
                                  code,
                                  payment_integration_type,
                                  payment_purpose,
                                  amount,
                                  payment_mode,
                                  currency, phone_number)
VALUES (current_timestamp,
        current_timestamp,
        @transaction_ref,
        @status,
        @transaction_complete,
        @data,
        @code,
        @payment_integration_type,
        @payment_purpose,
        @amount,
        @payment_mode,
        @currency,
        @phone_number)
returning id;

-- name: GetTotalPayments :one
SELECT COALESCE(SUM(amount), 0) AS total
FROM payment_transactions;



