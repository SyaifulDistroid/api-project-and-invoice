package models

type Project struct {
	ProjectID   int
	ProjectName string
	CompanyName string
	StartDate   string
	EndDate     string
	ArrayOfID   []int
}

type InvoiceProject struct {
	InvoiceID          int
	ProjectID          int
	ProjectName        string
	CompanyName        string
	StartDate          string
	EndDate            string
	ArrayOfID          []int
	InvoiceDueDate     string
	InvoiceCreateDate  string
	InvoiceTitle       string
	ArrayOfInvoiceItem []InvoiceItem
	Notes              string
	Currency           string
	Discount           int
	TotalAmount        int
}
