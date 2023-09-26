package models

type Invoice struct {
	InvoiceID          int
	ProjectID          int
	Currency           string
	TotalAmount        int
	Discount           int
	Notes              string
	InvoiceDueDate     string
	InvoiceCreateDate  string
	InvoiceTitle       string
	ArrayOfInvoiceItem []InvoiceItem
}

type InvoiceItem struct {
	ItemName string
	Qty      int
	Amount   int
}
