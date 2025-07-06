package dto

type NewTransactionResponse struct {
	TransactionId   int     `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	NewBalance      float64 `json:"new_balance"`
}
