package http

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/request"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
	"io"
	"net/http"
	"os"
)

type M map[string]interface{}

type InvoiceController interface {
	CreateNewInvoice(c echo.Context) error
	GetListAllInvoices(c echo.Context) error
	UpdateStatusInvoice(c echo.Context) error
	UpdateStatusPayment(c echo.Context) error
	GetListInvoicesByUserIDAndStatus(c echo.Context) error
	GetListInvoicesByUserID(c echo.Context) error
	GetDetailInvoice(c echo.Context) error
	GetDetailInvoiceByID(c echo.Context) error
	PaymentHandling(c echo.Context) error
	GeneratePayment(c echo.Context) error
	SendPaymentAndInvoice(c echo.Context) error
	DeleteInvoiceByID(c echo.Context) error
	CreateInvoiceWithExcel(c echo.Context) error
}

type invoiceController struct {
	invoiceUsecase domain.InvoiceUsecase
	userUsecase    domain.UserUsecase
}

func NewInvoiceController(iu domain.InvoiceUsecase, uu domain.UserUsecase) InvoiceController {
	return &invoiceController{
		invoiceUsecase: iu,
		userUsecase:    uu,
	}
}

// CreateNewInvoice godoc
// @Summary Create new invoice
// @Description Create invoice
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoice [post]
// @param data body request.InvoiceCreateRequest true "required"
// @Success 201 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) CreateNewInvoice(c echo.Context) error {
	var req request.InvoiceCreateRequest
	if err := c.Bind(&req); err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	jwtBearer := c.Get("user").(*jwt.Token)
	claims := jwtBearer.Claims.(jwt.MapClaims)
	userid := claims["UserID"].(string)

	invoice, err := i.invoiceUsecase.CreateNewInvoice(req, userid)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	var invStruct struct {
		InvoiceNumber string `json:"invoice_number"`
	}
	invoiceRes, _ := json.Marshal(invoice)
	json.Unmarshal(invoiceRes, &invStruct)
	i.invoiceUsecase.SendPaymentToClient(invStruct.InvoiceNumber)
	return response.SuccessResponse(c, http.StatusCreated, true, "success create invoice", invoice)
}

// GetListAllInvoices godoc
// @Summary Get list all invoice
// @Description Get list all invoice
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoices [get]
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) GetListAllInvoices(c echo.Context) error {
	invoices, err := i.invoiceUsecase.GetListAllInvoices()
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, true, "success get list invoices", invoices)
}

// UpdateStatusInvoice godoc
// @Summary Update status invoice
// @Description Update status invoice
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoice/status/{id} [put]
// @Param id path string true "id"
// @Param status query string true "status" Enums(paid, unpaid)
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) UpdateStatusInvoice(c echo.Context) error {
	id := c.Param("id")
	status := c.QueryParam("status")

	_, err := i.invoiceUsecase.GetDetailInvoiceByID(id)
	if err != nil {
		return response.FailResponse(c, http.StatusNotFound, false, err.Error())
	}

	invoice, err := i.invoiceUsecase.UpdateStatusInvoiceByID(status, id)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "success update status invoice", invoice)
}

// GetListInvoicesByUserIDAndStatus godoc
// @Summary Get list invoice by status and user
// @Description Get list invoice by status and user
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoices/status [get]
// @Param status query string true "status" Enums(paid, unpaid)
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) GetListInvoicesByUserIDAndStatus(c echo.Context) error {
	status := c.QueryParam("status")
	jwtBearer := c.Get("user").(*jwt.Token)
	claims := jwtBearer.Claims.(jwt.MapClaims)
	userid := claims["UserID"].(string)

	invoices, err := i.invoiceUsecase.GetListInvoicesByUserIDAndStatus(status, userid)
	if err != nil {
		return response.FailResponse(c, http.StatusNotFound, false, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, true, "success get list invoices", invoices)
}

// GetListInvoicesByUserID godoc
// @Summary Get list invoice by  user
// @Description Get list invoice by user
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoices/user [get]
// @Param userid query string true "userid"
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) GetListInvoicesByUserID(c echo.Context) error {
	userid := c.QueryParam("userid")

	invoices, err := i.invoiceUsecase.GetListInvoicesByUserID(userid)
	if err != nil {
		return response.FailResponse(c, http.StatusNotFound, false, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, true, "success get list invoices", invoices)
}

// GetDetailInvoice godoc
// @Summary Get detail invoice by invoice number
// @Description Get detail invoice by invoice number
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoice/detail [get]
// @Param invoice_number query string true "invoice_number"
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) GetDetailInvoice(c echo.Context) error {
	invoiceNumber := c.QueryParam("invoice_number")
	invoice, err := i.invoiceUsecase.GetDetailInvoice(invoiceNumber)
	if err != nil {
		return response.FailResponse(c, http.StatusNotFound, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "success get detail invoice", invoice)
}

// GetDetailInvoiceByID godoc
// @Summary Get detail invoice by id
// @Description Get detail invoice by id
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoice/detail/{id} [get]
// @Param id path string true "id"
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) GetDetailInvoiceByID(c echo.Context) error {
	id := c.Param("id")
	invoice, err := i.invoiceUsecase.GetDetailInvoiceByID(id)
	if err != nil {
		return response.FailResponse(c, http.StatusNotFound, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "success get detail invoice", invoice)
}

func (i invoiceController) PaymentHandling(c echo.Context) error {
	var req snap.RequestParamWithMap
	if err := c.Bind(&req); err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	encode, _ := json.Marshal(req)
	resArray := make(map[string]string)
	_ = json.Unmarshal(encode, &resArray)
	enc := sha512.New()
	enc.Write([]byte(resArray["order_id"] + resArray["status_code"] + resArray["gross_amount"] + os.Getenv("SERVER_KEY")))
	validateKey := fmt.Sprintf("%x", enc.Sum(nil))
	if validateKey != resArray["signature_key"] {
		return response.FailResponse(c, http.StatusForbidden, false, "invalid signature")
	}
	transaction := resArray["transaction_status"]
	paymentType := resArray["payment_type"]
	orderID := resArray["order_id"]
	fraud := resArray["fraud_status"]

	i.invoiceUsecase.UpdateStatusInvoiceAndPayment(transaction, orderID, fraud, paymentType)

	return response.SuccessResponse(c, http.StatusOK, true, "success handling invoice", resArray)
}

func (i invoiceController) GeneratePayment(c echo.Context) error {
	numInvoice := c.Param("inv")
	inv, err := i.invoiceUsecase.GetDetailInvoice(numInvoice)
	if err != nil {
		return response.FailResponse(c, http.StatusNotFound, false, err.Error())
	}
	var dataInv = inv.(map[string]interface{})
	var invoice struct {
		Status    string `json:"status"`
		SnapToken string `json:"snap_token"`
		Client    domain.Client
	}
	var user domain.User
	invoiceRes, _ := json.Marshal(dataInv["invoice"])
	userRes, _ := json.Marshal(dataInv["user"])
	json.Unmarshal(invoiceRes, &invoice)
	json.Unmarshal(userRes, &user)

	if invoice.Status == "paid" {
		return response.FailResponse(c, http.StatusNotFound, false, "Invoice already paid")
	}

	return c.Render(http.StatusOK, "Payment.html", M{
		"token":          invoice.SnapToken,
		"number_invoice": numInvoice,
		"to":             invoice.Client.FirstName + " " + invoice.Client.LastName,
		"from":           user.Profile.FirstName + " " + user.Profile.LastName,
	})
}

// SendPaymentAndInvoice godoc
// @Summary Send payment and invoice
// @Description Send Payment And Invoice
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoice/send/{id} [post]
// @Param id path string true "id"
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 404 {object} response.JSONBadRequestResult{}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) SendPaymentAndInvoice(c echo.Context) error {
	id := c.Param("id")
	inv, err := i.invoiceUsecase.GetDetailInvoiceByID(id)
	if err != nil {
		return response.FailResponse(c, http.StatusNotFound, false, err.Error())
	}
	var dataInv = inv.(map[string]interface{})
	var invoice struct {
		InvoiceNumber string `json:"invoice_number"`
	}
	invoiceRes, _ := json.Marshal(dataInv["invoice"])
	json.Unmarshal(invoiceRes, &invoice)

	_, err = i.invoiceUsecase.SendPaymentToClient(invoice.InvoiceNumber)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, "payment already paid")
	}
	return response.SuccessResponse(c, http.StatusOK, true, "success send payment and invoice to client", inv)
}

// UpdateStatusPayment godoc
// @Summary Update status payment invoice
// @Description Update status payment invoice
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoice/status-payment/{id} [put]
// @Param id path string true "id"
// @Param status query string true "status" Enums(pending, success, expired, failed)
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) UpdateStatusPayment(c echo.Context) error {
	id := c.Param("id")
	status := c.QueryParam("status")
	byID, err := i.invoiceUsecase.UpdateStatusInvoicePaymentByID(status, id)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, true, "success update payment status", byID)
}

// DeleteInvoiceByID godoc
// @Summary Delete invoice
// @Description Delete invoice
// @Tags Invoice
// @accept json
// @Produce json
// @Router /invoice/{id} [delete]
// @Param id path string true "id"
// @Success 200 {object} response.JSONSuccessDeleteResult{}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) DeleteInvoiceByID(c echo.Context) error {
	id := c.Param("id")
	err := i.invoiceUsecase.DeletePaymentByID(id)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	return response.SuccessDeleteResponse(c, http.StatusOK, true, "success delete invoice")
}

// CreateInvoiceWithExcel godoc
// @Summary Create new invoice with Excel
// @Description Create new invoice with Excele
// @Tags Invoice
// @accept multipart/form-data
// @Produce json
// @Router /invoice/file [post]
// @Param files formData file true  "upload file excel"
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (i invoiceController) CreateInvoiceWithExcel(c echo.Context) error {
	file, err := c.FormFile("files")
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./public/" + file.Filename)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	jwtBearer := c.Get("user").(*jwt.Token)
	claims := jwtBearer.Claims.(jwt.MapClaims)
	userid := claims["UserID"].(string)

	excel, err := i.invoiceUsecase.CreateInvoiceWithExcel(file.Filename, userid)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "success upload file and generate invoice", excel)
}
