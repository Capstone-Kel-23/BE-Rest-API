package usecase

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/invoice/utils"
	utils2 "github.com/Capstone-Kel-23/BE-Rest-API/utils"
	"github.com/Capstone-Kel-23/BE-Rest-API/utils/mail/body_email"
	"github.com/Capstone-Kel-23/BE-Rest-API/utils/mail/send_mail"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/request"
	"github.com/midtrans/midtrans-go"
	uuid "github.com/satori/go.uuid"
	"net/mail"
	"os"
	"strconv"
	"time"
)

type invoiceUsecase struct {
	invoiceRepository domain.InvoiceRepository
	userRepository    domain.UserRepository
}

func NewInvoiceUsecase(ir domain.InvoiceRepository, ur domain.UserRepository) domain.InvoiceUsecase {
	return &invoiceUsecase{
		invoiceRepository: ir,
		userRepository:    ur,
	}
}

func (i *invoiceUsecase) CreateNewInvoice(req request.InvoiceCreateRequest, userid string) (interface{}, error) {
	var (
		date      time.Time
		invoiceID uuid.UUID = uuid.NewV4()
	)
	date, _ = time.Parse("2006-01-02", req.Date)
	dateDue := utils.TimeDueCount(date, req.DateDue)

	var items []domain.Item
	if len(req.Items) != 0 {
		for _, item := range req.Items {
			items = append(items, domain.Item{
				ID:          uuid.NewV4(),
				InvoiceID:   invoiceID,
				Name:        item.Name,
				Quantity:    item.Quantity,
				Description: item.Description,
				Price:       item.Price,
			})
		}
	}
	var costs []domain.AdditionalCost
	if len(req.Items) != 0 {
		for _, cost := range req.AdditionalCosts {
			costs = append(costs, domain.AdditionalCost{
				ID:        uuid.NewV4(),
				InvoiceID: invoiceID,
				Type:      cost.Type,
				Name:      cost.Name,
				Total:     cost.Total,
			})
		}
	}
	var client domain.Client
	client = domain.Client{
		ID:          uuid.NewV4(),
		InvoiceID:   invoiceID,
		FirstName:   req.Client.FirstName,
		LastName:    req.Client.LastName,
		Email:       req.Client.Email,
		Address:     req.Client.Address,
		PostalCode:  req.Client.PostalCode,
		City:        req.Client.City,
		State:       req.Client.State,
		Country:     req.Client.Country,
		PhoneNumber: req.Client.PhoneNumber,
	}

	invoice := &domain.Invoice{
		ID:            invoiceID,
		UserID:        uuid.FromStringOrNil(userid),
		InvoiceNumber: req.InvoiceNumber,
		TypePayment:   req.TypePayment,
		Status:        "unpaid",
		StatusPayment: "pending",
		Description:   req.Description,
		Date:          date,
		DateDue:       dateDue,
		Total:         req.Total,
		SubTotal:      req.SubTotal,
		LogoURL:       req.LogoURL,
	}

	_, emailInvalid := mail.ParseAddress(client.Email)
	if emailInvalid != nil {
		return nil, errors.New("email client invalid")
	}

	invByInvoiceNumber, _ := i.invoiceRepository.FindByInvoiceNumber(invoice.InvoiceNumber)
	if invByInvoiceNumber.ID != uuid.FromStringOrNil("") {
		return nil, errors.New("invoice number already exist")
	}

	resInvoice, err := i.invoiceRepository.Save(invoice, &items, &client, &costs)
	if err != nil {
		return nil, err
	}
	finalInvoice, err := i.invoiceRepository.FindByID(resInvoice.ID.String())
	if err != nil {
		return nil, err
	}

	return finalInvoice, nil
}

func (i *invoiceUsecase) UpdateStatusInvoice(status, invoiceNumber string) (interface{}, error) {
	_, err := i.invoiceRepository.UpdateStatus(status, invoiceNumber)
	if err != nil {
		return nil, err
	}
	invoiceExist, _ := i.invoiceRepository.FindByInvoiceNumber(invoiceNumber)
	if invoiceExist.ID == uuid.FromStringOrNil("") {
		return nil, err
	}
	return invoiceExist, nil
}

func (i *invoiceUsecase) GetListAllInvoices() (*domain.Invoices, error) {
	all, err := i.invoiceRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (i *invoiceUsecase) GetListInvoicesByUserIDAndStatus(status string, userid string) (*domain.Invoices, error) {
	invoices, err := i.invoiceRepository.FindByUserIDAndStatus(userid, status)
	if err != nil {
		return nil, err
	}
	return invoices, err
}

func (i *invoiceUsecase) GetListInvoicesByUserID(userid string) (*domain.Invoices, error) {
	invoices, err := i.invoiceRepository.FindByUserID(userid)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (i *invoiceUsecase) GetDetailInvoice(invoiceNumber string) (interface{}, error) {
	invoice, err := i.invoiceRepository.FindByInvoiceNumber(invoiceNumber)
	if err != nil {
		return nil, err
	}
	if invoice.ID == uuid.FromStringOrNil("") {
		return nil, errors.New("invoice not found")
	}
	user, err := i.userRepository.FindWithProfile(invoice.UserID.String())
	if err != nil {
		return nil, err
	}

	finalData := map[string]interface{}{
		"user": map[string]interface{}{
			"username": user.Username,
			"id":       user.ID,
			"profile":  user.Profile,
		},
		"invoice": invoice,
	}
	return finalData, err
}

func (i *invoiceUsecase) GetDetailInvoiceByID(id string) (interface{}, error) {
	invoice, err := i.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if invoice.ID == uuid.FromStringOrNil("") {
		return nil, errors.New("invoice not found")
	}
	user, err := i.userRepository.FindWithProfile(invoice.UserID.String())
	if err != nil {
		return nil, err
	}

	finalData := map[string]interface{}{
		"user": map[string]interface{}{
			"username": user.Username,
			"id":       user.ID,
			"profile":  user.Profile,
		},
		"invoice": invoice,
	}
	return finalData, err
}

func (i *invoiceUsecase) UpdateStatusInvoiceByID(status, id string) (interface{}, error) {
	_, err := i.invoiceRepository.UpdateStatusByID(status, id)
	if err != nil {
		return nil, err
	}
	invoiceExist, _ := i.invoiceRepository.FindByID(id)
	if invoiceExist.ID == uuid.FromStringOrNil("") {
		return nil, errors.New("error get detail invoice")
	}
	return invoiceExist, nil
}

func (i *invoiceUsecase) UpdateStatusInvoicePayment(status, invoiceNumber string) (interface{}, error) {
	_, err := i.invoiceRepository.UpdateStatusPayment(status, invoiceNumber)
	if err != nil {
		return nil, err
	}
	invoiceExist, _ := i.invoiceRepository.FindByInvoiceNumber(invoiceNumber)
	if invoiceExist.ID == uuid.FromStringOrNil("") {
		return nil, errors.New("error get detail invoice")
	}
	return invoiceExist, nil
}

func (i *invoiceUsecase) SendPaymentToClient(inv string) (interface{}, error) {
	invExist, err := i.invoiceRepository.FindByInvoiceNumber(inv)
	if err != nil {
		return nil, err
	}
	profileUser, err := i.userRepository.FindWithProfile(invExist.UserID.String())
	if err != nil {
		return nil, err
	}
	if invExist.Status == "paid" {
		return nil, errors.New("invoice already paid")
	}
	var snapToken string
	randomStringNumber := utils.GenerateNumberInvoice()
	orderID := invExist.InvoiceNumber + "-" + randomStringNumber
	if invExist.TypePayment != "cash" {
		errGenerate := &midtrans.Error{}
		snapToken, errGenerate = utils2.NewPayment().GeneratePayment(*invExist, orderID)
		if errGenerate != nil {
			return nil, errors.New(errGenerate.Error())
		}
	} else {
		templateData := body_email.SendPaymentCashEmailBody{
			From:          profileUser.Profile.FirstName + " " + profileUser.Profile.LastName,
			To:            invExist.Client.FirstName + " " + invExist.Client.LastName,
			InvoiceNumber: inv,
			Subject:       "Invoice and Payment Email",
		}
		err = send_mail.SendPaymentCashEmail(invExist.Client.Email, templateData)
		return "success send invoice and payment to email client", nil
	}
	err = i.invoiceRepository.UpdateOrderID(orderID, invExist.ID.String())
	if err != nil {
		return nil, err
	}
	err = i.invoiceRepository.UpdateSnapToken(snapToken, invExist.ID.String())
	if err != nil {
		return nil, err
	}
	templateData := body_email.SendPaymentEmailBody{
		URL:           "http://prodapi.tagihin.my.id/api/v1/invoice/payment-generate/" + invExist.InvoiceNumber,
		From:          profileUser.Profile.FirstName + " " + profileUser.Profile.LastName,
		To:            invExist.Client.FirstName + " " + invExist.Client.LastName,
		InvoiceNumber: inv,
		Subject:       "Invoice and Payment Email",
	}
	err = send_mail.SendPaymentEmail(invExist.Client.Email, templateData)

	return "success send invoice and payment to email client", nil
}

func (i *invoiceUsecase) UpdateStatusInvoiceAndPayment(transaction, orderID, fraud, paymentType string) {
	inv, _ := i.invoiceRepository.FindByOrderID(orderID)
	if transaction == "capture" {
		if paymentType == "credit_card" {
			if fraud == "challenge" {
				i.invoiceRepository.UpdateStatusPayment("pending", inv.InvoiceNumber)
			} else {
				i.invoiceRepository.UpdateStatusPayment("success", inv.InvoiceNumber)
				i.invoiceRepository.UpdateStatus("paid", inv.InvoiceNumber)
			}
		}
	} else if transaction == "settlement" {
		i.invoiceRepository.UpdateStatusPayment("success", inv.InvoiceNumber)
		i.invoiceRepository.UpdateStatus("paid", inv.InvoiceNumber)
	} else if transaction == "pending" {
		i.invoiceRepository.UpdateStatusPayment("pending", inv.InvoiceNumber)
	} else if transaction == "deny" {
		i.invoiceRepository.UpdateStatusPayment("failed", inv.InvoiceNumber)
	} else if transaction == "expired" {
		i.invoiceRepository.UpdateStatusPayment("expired", inv.InvoiceNumber)
	} else if transaction == "cancel" {
		i.invoiceRepository.UpdateStatusPayment("failed", inv.InvoiceNumber)
	}
}

func (i *invoiceUsecase) UpdateStatusInvoicePaymentByID(status, id string) (interface{}, error) {
	_, err := i.invoiceRepository.UpdateStatusPaymentByID(status, id)
	if err != nil {
		return nil, err
	}
	invoiceExist, _ := i.invoiceRepository.FindByID(id)
	if invoiceExist.ID == uuid.FromStringOrNil("") {
		return nil, errors.New("error get detail invoice")
	}
	return invoiceExist, nil
}

func (i *invoiceUsecase) DeletePaymentByID(id string) error {
	invExist, _ := i.invoiceRepository.FindByID(id)
	if invExist.ID == uuid.FromStringOrNil("") {
		return errors.New("not found invoice")
	}
	err := i.invoiceRepository.Delete(invExist)
	if err != nil {
		return err
	}
	return nil
}

func (i *invoiceUsecase) CreateInvoiceWithExcel(filename string, userid string) (interface{}, error) {
	var (
		date      time.Time
		invoiceID uuid.UUID = uuid.NewV4()
	)
	xlsx, err := excelize.OpenFile("./public/" + filename)
	if err != nil {
		return nil, err
	}
	sheetClientProfile := "client_profile"
	sheetItems := "items"
	sheetCosts := "costs"
	sheetInvoice := "invoice"

	var items []domain.Item
	if len(xlsx.GetRows(sheetItems)) > 1 {
		for i := 2; i < len(xlsx.GetRows(sheetItems))+1; i++ {
			qty, _ := strconv.Atoi(xlsx.GetCellValue(sheetItems, fmt.Sprintf("B%d", i)))
			price, _ := strconv.Atoi(xlsx.GetCellValue(sheetItems, fmt.Sprintf("D%d", i)))
			items = append(items, domain.Item{
				ID:          uuid.NewV4(),
				InvoiceID:   invoiceID,
				Name:        xlsx.GetCellValue(sheetItems, fmt.Sprintf("A%d", i)),
				Quantity:    qty,
				Description: xlsx.GetCellValue(sheetItems, fmt.Sprintf("C%d", i)),
				Price:       float64(price),
			})
		}
	}

	var costs []domain.AdditionalCost
	if len(xlsx.GetRows(sheetCosts)) > 1 {
		for i := 2; i < len(xlsx.GetRows(sheetCosts))+1; i++ {
			total, _ := strconv.Atoi(xlsx.GetCellValue(sheetCosts, fmt.Sprintf("C%d", i)))
			costs = append(costs, domain.AdditionalCost{
				ID:        uuid.NewV4(),
				InvoiceID: invoiceID,
				Type:      xlsx.GetCellValue(sheetCosts, fmt.Sprintf("A%d", i)),
				Name:      xlsx.GetCellValue(sheetCosts, fmt.Sprintf("B%d", i)),
				Total:     float64(total),
			})
		}
	}

	var client domain.Client
	client = domain.Client{
		ID:          uuid.NewV4(),
		InvoiceID:   invoiceID,
		FirstName:   xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("A%d", 2)),
		LastName:    xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("B%d", 2)),
		Email:       xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("C%d", 2)),
		Address:     xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("D%d", 2)),
		PostalCode:  xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("E%d", 2)),
		City:        xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("F%d", 2)),
		State:       xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("G%d", 2)),
		Country:     xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("H%d", 2)),
		PhoneNumber: xlsx.GetCellValue(sheetClientProfile, fmt.Sprintf("I%d", 2)),
	}

	date, _ = time.Parse("2006-01-02", utils.ExcelSerialDateToTime(xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("D%d", 2))))
	dateDue := utils.TimeDueCount(date, utils.ExcelSerialDateToTime(xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("E%d", 2))))
	total, _ := strconv.Atoi(xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("F%d", 2)))
	subTotal, _ := strconv.Atoi(xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("G%d", 2)))
	invoice := &domain.Invoice{
		ID:            invoiceID,
		UserID:        uuid.FromStringOrNil(userid),
		InvoiceNumber: xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("A%d", 2)),
		TypePayment:   xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("B%d", 2)),
		Status:        "unpaid",
		StatusPayment: "pending",
		Description:   xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("C%d", 2)),
		Date:          date,
		DateDue:       dateDue,
		Total:         float64(total),
		SubTotal:      float64(subTotal),
		LogoURL:       xlsx.GetCellValue(sheetInvoice, fmt.Sprintf("H%d", 2)),
	}

	_, emailInvalid := mail.ParseAddress(client.Email)
	if emailInvalid != nil {
		return nil, errors.New("email client invalid")
	}

	invByInvoiceNumber, _ := i.invoiceRepository.FindByInvoiceNumber(invoice.InvoiceNumber)
	if invByInvoiceNumber.ID != uuid.FromStringOrNil("") {
		return nil, errors.New("invoice number already exist")
	}

	resInvoice, err := i.invoiceRepository.Save(invoice, &items, &client, &costs)
	if err != nil {
		return nil, err
	}
	finalInvoice, err := i.invoiceRepository.FindByID(resInvoice.ID.String())
	if err != nil {
		return nil, err
	}

	_ = os.Remove("./public/" + filename)

	_, _ = i.SendPaymentToClient(finalInvoice.InvoiceNumber)

	return finalInvoice, nil
}
