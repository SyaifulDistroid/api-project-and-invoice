package handler

import (
	"EMI/controller"
	"EMI/middleware"
	"EMI/models"
	Utils "EMI/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EmiHandler struct {
	emiService controller.EmiServiceController
}

func CreateEmiHandler(r *gin.Engine, emiService controller.EmiServiceController) {
	EmiHandler := &EmiHandler{emiService}

	r.Use(middleware.RequestLoggerActivity())
	r.GET("/Projects", EmiHandler.getAllProject)
	r.PUT("/Project/:id", EmiHandler.updateProject)
	r.GET("/Project/User/:id", EmiHandler.getAllProjectByUserID)
	r.POST("/Invoice", EmiHandler.insertInvoice)
	r.DELETE("/Invoice/:id", EmiHandler.deleteInvoice)
	r.PUT("/Invoice/:id", EmiHandler.updateInvoice)
	r.GET("/Invoice/User/:id", EmiHandler.getInvoiceByUserID)
	r.GET("/Invoice/Past/:id", EmiHandler.getInvoiceByDate)
	r.GET("/Invoice/Next/:id", EmiHandler.getInvoiceByDateFuture)
	r.GET("/User", EmiHandler.getAllUser)
}

func (e *EmiHandler) getAllProject(c *gin.Context) {

	Project, err := e.emiService.GetAllProject()

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Data Project", Project)
}

func (e *EmiHandler) updateProject(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Println(err)
		return
	}

	var data models.Project
	c.BindJSON(&data)
	Project, err := e.emiService.UpdateProject(id, data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Update Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Update Data Project", Project)
}

func (e *EmiHandler) getAllProjectByUserID(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Println(err)
		return
	}

	Project, err := e.emiService.GetAllProjectByUserID(id)

	if err != nil {
		MessageError := fmt.Sprintf("Failed Show Data Project From User ID = %v", id)
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, MessageError, nil)
		logrus.Println(err)
		return
	}
	MessageSucces := fmt.Sprintf("Succes Show Data Project From User ID = %v", id)
	Utils.Response(c, http.StatusOK, http.StatusOK, MessageSucces, Project)
}

func (e *EmiHandler) insertInvoice(c *gin.Context) {
	var data models.Invoice
	c.BindJSON(&data)

	Invoice, err := e.emiService.InsertInvoice(data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Save Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Save Data Invoice", Invoice)
}

func (e *EmiHandler) deleteInvoice(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Println(err)
		return
	}

	_, err = e.emiService.DeleteInvoice(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Delete Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Delete Data Invoice", nil)
}

func (e *EmiHandler) updateInvoice(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Println(err)
		return
	}

	var data models.Invoice
	c.BindJSON(&data)

	Invoice, err := e.emiService.UpdateInvoice(id, data)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Update Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Update Data Invoice", Invoice)
}

func (e *EmiHandler) getInvoiceByUserID(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Println(err)
		return
	}

	InvoiceProject, err := e.emiService.GetInvoiceByUserID(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Invoice Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Invoice Data", InvoiceProject)
}

func (e *EmiHandler) getInvoiceByDate(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Println(err)
		return
	}

	InvoiceProject, err := e.emiService.GetInvoiceByDate(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Invoice Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Invoice Data", InvoiceProject)
}

func (e *EmiHandler) getInvoiceByDateFuture(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusBadRequest, "Invalid Request", nil)
		logrus.Println(err)
		return
	}

	InvoiceProject, err := e.emiService.GetInvoiceByDateFuture(id)

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Invoice Data", nil)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Invoice Data", InvoiceProject)
}

func (e *EmiHandler) getAllUser(c *gin.Context) {
	User, err := e.emiService.GetAllUser()

	if err != nil {
		Utils.Response(c, http.StatusNotFound, http.StatusNotFound, "Failed Show Data", err)
		logrus.Println(err)
		return
	}
	Utils.Response(c, http.StatusOK, http.StatusOK, "Succes Show Data User", User)
}
