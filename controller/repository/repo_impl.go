package repository

import (
	"EMI/controller"
	"EMI/models"
	Utils "EMI/utils"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type EmiRepoMysqlImpl struct {
	db *sql.DB
}

func CreateEmiRepoMysqlImpl(db *sql.DB) controller.EmiRepoController {
	return &EmiRepoMysqlImpl{db}
}

func (e EmiRepoMysqlImpl) GetAllProject() (*[]models.Project, error) {
	query := "SELECT PROJECT_ID, PROJECT_NAME, COMPANY_NAME, START_DATE, END_DATE, ARRAY_OF_USERID FROM tb_projects"

	var Project models.Project
	var Projects []models.Project

	var ArrayOfID string

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Query Get All Project : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&Project.ProjectID, &Project.ProjectName, &Project.CompanyName, &Project.StartDate, &Project.EndDate, &ArrayOfID); err != nil {
			return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Rows Scan Project : %v", err)
		}
		split := strings.Split(ArrayOfID, ",")
		var intID int
		var _intID []int
		for i := 0; i < len(split); i++ {
			intID, _ = strconv.Atoi(split[i])
			_intID = append(_intID, intID)
		}
		Project.ArrayOfID = _intID
		Projects = append(Projects, Project)
	}
	return &Projects, nil
}

func (e EmiRepoMysqlImpl) UpdateProject(id int, data models.Project) (*models.Project, error) {

	var strID string
	var _strID []string
	for i := 0; i < len(data.ArrayOfID); i++ {
		strID = strconv.Itoa(data.ArrayOfID[i])
		_strID = append(_strID, strID)
	}
	strArrayOfID := strings.Join(_strID, ",")

	query := "UPDATE tb_projects SET PROJECT_NAME = ?, COMPANY_NAME = ?, START_DATE = ?, END_DATE = ?, ARRAY_OF_USERID = ? WHERE PROJECT_ID = ?"

	_, err := e.db.Exec(query, data.ProjectName, data.CompanyName, data.StartDate, data.EndDate, strArrayOfID, id)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.UpdateProject] Error When Query Get All Project : %v", err)
	}
	data.ProjectID = id
	return &data, nil
}

func (e EmiRepoMysqlImpl) GetAllProjectByUserID(id int) (*[]models.Project, error) {
	query := fmt.Sprintf("SELECT PROJECT_ID, PROJECT_NAME, COMPANY_NAME, START_DATE, END_DATE, ARRAY_OF_USERID FROM tb_projects WHERE ARRAY_OF_USERID LIKE %v%v%v", "'%", id, "%'")

	var Project models.Project
	var Projects []models.Project

	var ArrayOfID string

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Query Get All Project By User ID : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&Project.ProjectID, &Project.ProjectName, &Project.CompanyName, &Project.StartDate, &Project.EndDate, &ArrayOfID); err != nil {
			return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Rows Scan Project By User ID : %v", err)
		}
		split := strings.Split(ArrayOfID, ",")
		var intID int
		var _intID []int
		for i := 0; i < len(split); i++ {
			intID, _ = strconv.Atoi(split[i])
			_intID = append(_intID, intID)
		}
		Project.ArrayOfID = _intID
		Projects = append(Projects, Project)
	}
	return &Projects, nil
}

func (e EmiRepoMysqlImpl) InsertInvoice(data models.Invoice) (*models.Invoice, error) {
	var _strID []string

	for i := 0; i < len(data.ArrayOfInvoiceItem); i++ {
		_strID = append(_strID, Utils.JoinVarAndValStruct(reflect.ValueOf(data.ArrayOfInvoiceItem[i])))
	}

	strArrayOfID := strings.Join(_strID, ",")

	query := "INSERT INTO tb_invoices (INVOICE_ID, PROJECT_ID, CURRENCY, TOTAL_AMOUNT, DISCOUNT, NOTES, INVOICE_DUE_DATE, INVOICE_CREATE_DATE, INVOICE_TITLE, ARRAY_OF_INVOICE_ITEM) VALUES (?,?,?,?,?,?,?,?,?,?)"

	_, err := e.db.Exec(query, 0, data.ProjectID, data.Currency, data.TotalAmount, data.Discount, data.Notes, data.InvoiceDueDate, data.InvoiceCreateDate, data.InvoiceTitle, strArrayOfID)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.UpdateInvoice] Error When Query Get All Invoice : %v", err)
	}
	return &data, nil
}

func (e EmiRepoMysqlImpl) DeleteInvoice(id int) (*models.Invoice, error) {
	query := "DELETE FROM tb_invoices WHERE INVOICE_ID = ?"

	var data models.Invoice
	_, err := e.db.Exec(query, id)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.DeleteInvoice] Error When Query Delete Invoice : %v", err)
	}
	return &data, nil
}

func (e EmiRepoMysqlImpl) UpdateInvoice(id int, data models.Invoice) (*models.Invoice, error) {
	var _strID []string
	for i := 0; i < len(data.ArrayOfInvoiceItem); i++ {
		_strID = append(_strID, Utils.JoinVarAndValStruct(reflect.ValueOf(data.ArrayOfInvoiceItem[i])))
	}
	strArrayOfID := strings.Join(_strID, ",")

	query := "UPDATE tb_invoices SET PROJECT_ID = ?, CURRENCY = ?, TOTAL_AMOUNT = ?, DISCOUNT = ?, NOTES = ?, INVOICE_DUE_DATE = ?, INVOICE_CREATE_DATE = ?, INVOICE_TITLE = ?, ARRAY_OF_INVOICE_ITEM = ? WHERE INVOICE_ID = ?"

	_, err := e.db.Exec(query, data.ProjectID, data.Currency, data.TotalAmount, data.Discount, data.Notes, data.InvoiceDueDate, data.InvoiceCreateDate, data.InvoiceTitle, strArrayOfID, id)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.UpdateInvoice] Error When Query Update Invoice : %v", err)
	}
	data.InvoiceID = id
	return &data, nil
}

func (e EmiRepoMysqlImpl) GetInvoiceByUserID(id int) (*[]models.InvoiceProject, error) {
	query := fmt.Sprintf("SELECT tb_invoices.INVOICE_ID, tb_projects.PROJECT_ID, CURRENCY, TOTAL_AMOUNT, DISCOUNT, NOTES, INVOICE_DUE_DATE,INVOICE_CREATE_DATE, INVOICE_TITLE, ARRAY_OF_INVOICE_ITEM, PROJECT_NAME, COMPANY_NAME, START_DATE, END_DATE, ARRAY_OF_USERID FROM tb_invoices INNER JOIN tb_projects	ON tb_invoices.PROJECT_ID = tb_projects.PROJECT_ID WHERE tb_projects.ARRAY_OF_USERID LIKE %v%v%v", "'%", id, "%'")

	var data models.InvoiceProject
	var _data []models.InvoiceProject
	var item models.InvoiceItem
	var items []models.InvoiceItem
	var _items []models.InvoiceItem

	var ArrayOfID string
	var ArrayOfItem string
	var _ArrayOfItem []string
	var _BlankArrayOfItem []string

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Query Get All Invoice Project By User ID : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&data.InvoiceID, &data.ProjectID, &data.Currency, &data.TotalAmount, &data.Discount, &data.Notes, &data.InvoiceDueDate, &data.InvoiceCreateDate, &data.InvoiceTitle, &ArrayOfItem, &data.ProjectName, &data.CompanyName, &data.StartDate, &data.EndDate, &ArrayOfID); err != nil {
			return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Rows Scan Invoice Project By User ID : %v", err)
		}
		//For Split Array Of ID
		split := strings.Split(ArrayOfID, ",")
		var intID int
		var _intID []int
		for i := 0; i < len(split); i++ {
			intID, _ = strconv.Atoi(split[i])
			_intID = append(_intID, intID)
		}
		data.ArrayOfID = _intID

		//For Split Array Of Item
		splitItem := strings.Split(ArrayOfItem, ",")
		for i := 0; i < len(splitItem); i++ {
			_ArrayOfItem = append(_ArrayOfItem, splitItem[i])
			if len(_ArrayOfItem) == 3 {

				splitOne := strings.Split(_ArrayOfItem[0], ":")
				splitTwo := strings.Split(_ArrayOfItem[1], ":")
				splitThree := strings.Split(_ArrayOfItem[2], ":")

				item.ItemName = splitOne[1]
				item.Qty = Utils.StringToInt(splitTwo[1])
				item.Amount = Utils.StringToInt(splitThree[1])

				items = append(items, item)
				_ArrayOfItem = _BlankArrayOfItem
			}
			data.ArrayOfInvoiceItem = items
		}
		items = _items
		_data = append(_data, data)
	}
	return &_data, nil
}

func (e EmiRepoMysqlImpl) GetInvoiceByDate(id int) (*[]models.InvoiceProject, error) {
	query := fmt.Sprintf("SELECT tb_invoices.INVOICE_ID, tb_projects.PROJECT_ID, CURRENCY, TOTAL_AMOUNT, DISCOUNT, NOTES, INVOICE_DUE_DATE,INVOICE_CREATE_DATE, INVOICE_TITLE, ARRAY_OF_INVOICE_ITEM, PROJECT_NAME, COMPANY_NAME, START_DATE, END_DATE, ARRAY_OF_USERID FROM tb_invoices INNER JOIN tb_projects	ON tb_invoices.PROJECT_ID = tb_projects.PROJECT_ID WHERE tb_projects.ARRAY_OF_USERID LIKE %v%v%v AND (INVOICE_DUE_DATE BETWEEN '2010-01-01' AND DATE(NOW()))", "'%", id, "%'")

	var data models.InvoiceProject
	var _data []models.InvoiceProject
	var item models.InvoiceItem
	var items []models.InvoiceItem
	var _items []models.InvoiceItem

	var ArrayOfID string
	var ArrayOfItem string
	var _ArrayOfItem []string
	var _BlankArrayOfItem []string

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Query Get All Invoice Project By Date : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&data.InvoiceID, &data.ProjectID, &data.Currency, &data.TotalAmount, &data.Discount, &data.Notes, &data.InvoiceDueDate, &data.InvoiceCreateDate, &data.InvoiceTitle, &ArrayOfItem, &data.ProjectName, &data.CompanyName, &data.StartDate, &data.EndDate, &ArrayOfID); err != nil {
			return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Rows Scan Invoice Project By Date : %v", err)
		}
		//For Split Array Of ID
		split := strings.Split(ArrayOfID, ",")
		var intID int
		var _intID []int
		for i := 0; i < len(split); i++ {
			intID, _ = strconv.Atoi(split[i])
			_intID = append(_intID, intID)
		}
		data.ArrayOfID = _intID

		//For Split Array Of Item
		splitItem := strings.Split(ArrayOfItem, ",")
		for i := 0; i < len(splitItem); i++ {
			_ArrayOfItem = append(_ArrayOfItem, splitItem[i])
			if len(_ArrayOfItem) == 3 {

				splitOne := strings.Split(_ArrayOfItem[0], ":")
				splitTwo := strings.Split(_ArrayOfItem[1], ":")
				splitThree := strings.Split(_ArrayOfItem[2], ":")

				item.ItemName = splitOne[1]
				item.Qty = Utils.StringToInt(splitTwo[1])
				item.Amount = Utils.StringToInt(splitThree[1])

				items = append(items, item)
				_ArrayOfItem = _BlankArrayOfItem
			}
			data.ArrayOfInvoiceItem = items
		}
		items = _items
		_data = append(_data, data)
	}
	return &_data, nil
}

func (e EmiRepoMysqlImpl) GetInvoiceByDateFuture(id int) (*[]models.InvoiceProject, error) {
	query := fmt.Sprintf("SELECT tb_invoices.INVOICE_ID, tb_projects.PROJECT_ID, CURRENCY, TOTAL_AMOUNT, DISCOUNT, NOTES, INVOICE_DUE_DATE,INVOICE_CREATE_DATE, INVOICE_TITLE, ARRAY_OF_INVOICE_ITEM, PROJECT_NAME, COMPANY_NAME, START_DATE, END_DATE, ARRAY_OF_USERID FROM tb_invoices INNER JOIN tb_projects	ON tb_invoices.PROJECT_ID = tb_projects.PROJECT_ID WHERE tb_projects.ARRAY_OF_USERID LIKE %v%v%v AND (INVOICE_DUE_DATE BETWEEN DATE(NOW()) AND '2050-01-01')", "'%", id, "%'")

	var data models.InvoiceProject
	var _data []models.InvoiceProject
	var item models.InvoiceItem
	var items []models.InvoiceItem
	var _items []models.InvoiceItem

	var ArrayOfID string
	var ArrayOfItem string
	var _ArrayOfItem []string
	var _BlankArrayOfItem []string

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Query Get All Invoice Project By Date Future : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&data.InvoiceID, &data.ProjectID, &data.Currency, &data.TotalAmount, &data.Discount, &data.Notes, &data.InvoiceDueDate, &data.InvoiceCreateDate, &data.InvoiceTitle, &ArrayOfItem, &data.ProjectName, &data.CompanyName, &data.StartDate, &data.EndDate, &ArrayOfID); err != nil {
			return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllProject] Error When Rows Scan Invoice Project By Date Future : %v", err)
		}
		//For Split Array Of ID
		split := strings.Split(ArrayOfID, ",")
		var intID int
		var _intID []int
		for i := 0; i < len(split); i++ {
			intID, _ = strconv.Atoi(split[i])
			_intID = append(_intID, intID)
		}
		data.ArrayOfID = _intID

		//For Split Array Of Item
		splitItem := strings.Split(ArrayOfItem, ",")
		for i := 0; i < len(splitItem); i++ {
			_ArrayOfItem = append(_ArrayOfItem, splitItem[i])
			if len(_ArrayOfItem) == 3 {

				splitOne := strings.Split(_ArrayOfItem[0], ":")
				splitTwo := strings.Split(_ArrayOfItem[1], ":")
				splitThree := strings.Split(_ArrayOfItem[2], ":")

				item.ItemName = splitOne[1]
				item.Qty = Utils.StringToInt(splitTwo[1])
				item.Amount = Utils.StringToInt(splitThree[1])

				items = append(items, item)
				_ArrayOfItem = _BlankArrayOfItem
			}
			data.ArrayOfInvoiceItem = items
		}
		items = _items
		_data = append(_data, data)
	}
	return &_data, nil
}

func (e EmiRepoMysqlImpl) GetAllUser() (*[]models.User, error) {
	query := "SELECT USER_ID, USERNAME, EMAIL, PASSWORD FROM tb_users"

	var User models.User
	var Users []models.User

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllUser] Error When Query Get All User : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&User.ID, &User.Username, &User.Email, &User.Password); err != nil {
			return nil, fmt.Errorf("[EmiRepoMysqlImpl.GetAllUser] Error When Row Scan User : %w", err)
		}
		Users = append(Users, User)
	}
	return &Users, nil
}
