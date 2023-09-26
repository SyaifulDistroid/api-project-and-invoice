package service

import (
	"EMI/controller"
	"EMI/models"
)

type EmiServiceImpl struct {
	emiRepo controller.EmiRepoController
}

func CreateEmiServiceImpl(emiRepo controller.EmiRepoController) controller.EmiServiceController {
	return &EmiServiceImpl{emiRepo}
}
func (e *EmiServiceImpl) GetAllProject() (*[]models.Project, error) {
	return e.emiRepo.GetAllProject()
}
func (e *EmiServiceImpl) UpdateProject(id int, data models.Project) (*models.Project, error) {
	return e.emiRepo.UpdateProject(id, data)
}
func (e *EmiServiceImpl) GetAllProjectByUserID(id int) (*[]models.Project, error) {
	return e.emiRepo.GetAllProjectByUserID(id)
}
func (e *EmiServiceImpl) InsertInvoice(data models.Invoice) (*models.Invoice, error) {
	return e.emiRepo.InsertInvoice(data)
}
func (e *EmiServiceImpl) DeleteInvoice(id int) (*models.Invoice, error) {
	return e.emiRepo.DeleteInvoice(id)
}
func (e *EmiServiceImpl) UpdateInvoice(id int, data models.Invoice) (*models.Invoice, error) {
	return e.emiRepo.UpdateInvoice(id, data)
}
func (e *EmiServiceImpl) GetInvoiceByUserID(id int) (*[]models.InvoiceProject, error) {
	return e.emiRepo.GetInvoiceByUserID(id)
}
func (e *EmiServiceImpl) GetInvoiceByDate(id int) (*[]models.InvoiceProject, error) {
	return e.emiRepo.GetInvoiceByDate(id)
}
func (e *EmiServiceImpl) GetInvoiceByDateFuture(id int) (*[]models.InvoiceProject, error) {
	return e.emiRepo.GetInvoiceByDateFuture(id)
}
func (e *EmiServiceImpl) GetAllUser() (*[]models.User, error) {
	return e.emiRepo.GetAllUser()
}
