package paymentmanualtransfer

type PaymentManualTransferConfirmation struct {
	OrderID           uint64
	PaymentType       uint8
	BankAccountNumber string
	BankAccountName   string
	Date              string
	Value             float64
}
