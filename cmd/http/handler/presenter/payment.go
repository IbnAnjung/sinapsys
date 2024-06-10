package presenter

type ManualTransferConfirmationRequest struct {
	OrderID           uint64  `json:"order_id"`
	PaymentType       uint8   `json:"payment_type"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankAccountName   string  `json:"bank_account_name"`
	Date              string  `json:"date"`
	Value             float64 `json:"value"`
}
