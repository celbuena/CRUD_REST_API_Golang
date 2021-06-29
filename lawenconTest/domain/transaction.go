package domain

import "errors"

var (
	ErrNotTransactionNotFound = errors.New("transaction not found")
)

type Transaction struct {
	ID            string `json:"-"`
	TrxCode       string `json:"trx_code"`
	UserId        int    `json:"user_id"`
	SchoolId      int    `json:"school_id"`
	ParentId      int    `json:"parent_id"`
	ParentEmail   string `json:"parent_email"`
	StudentEmail  string `json:"student_email"`
	ClassName     string `json:"class_name"`
	SchoolName    string `json:"school_name"`
	PaymentType   string `json:"payment_type"`
	PaymentMethod string `json:"payment_method"`
	Url           string `json:"url"`
	Price         int    `json:"price"`
	PaymentStatus string `json:"payment_status"`
}

type TransactionInfo struct {
	PaymentStatus string `json:"payment_status"`
	TrxCode       string `json:"trx_code"`
	PaymentMethod string `json:"payment_method"`
}


