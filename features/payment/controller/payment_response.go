package controller

import "be-api/features"

type ResponMidtrans struct {
	TransactionTime   string
	TransactionStatus string
	TransactionId     string
	StatusMessage     string
	StatusCode        string
	SignatureKey      string
	PaymentType       string
	SettlementTime    string
	OrderId           string
	MerchatId         string
	GrossAmount       string
	FraudStatus       string
	Currency          string
	ApprovalCode      string
	VANumber          string
	bank              string
}

func PaymentMidstransToModel(payment ResponMidtrans) features.Payment {
	return features.Payment{
		Name:   payment.bank,
		Status: payment.StatusMessage,
	}
}