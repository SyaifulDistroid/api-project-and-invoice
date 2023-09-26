package controller

import "EMI/models"

type EmiServiceController interface {
	GetAllProject() (*[]models.Project, error)
	UpdateProject(id int, data models.Project) (*models.Project, error)
	GetAllProjectByUserID(id int) (*[]models.Project, error)
	InsertInvoice(data models.Invoice) (*models.Invoice, error)
	DeleteInvoice(id int) (*models.Invoice, error)
	UpdateInvoice(id int, data models.Invoice) (*models.Invoice, error)
	GetInvoiceByUserID(id int) (*[]models.InvoiceProject, error)
	GetInvoiceByDate(id int) (*[]models.InvoiceProject, error)
	GetInvoiceByDateFuture(id int) (*[]models.InvoiceProject, error)
	GetAllUser() (*[]models.User, error)
}
