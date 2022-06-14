package repository

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) domain.InvoiceRepository {
	return &invoiceRepository{
		db: db,
	}
}

func (i *invoiceRepository) Save(invoice *domain.Invoice, items *[]domain.Item, client *domain.Client, costs *[]domain.AdditionalCost) (*domain.Invoice, error) {
	err := i.db.Create(&domain.Invoice{
		ID:              invoice.ID,
		UserID:          invoice.UserID,
		InvoiceNumber:   invoice.InvoiceNumber,
		Status:          invoice.Status,
		SnapToken:       invoice.SnapToken,
		TypePayment:     invoice.TypePayment,
		StatusPayment:   invoice.StatusPayment,
		Description:     invoice.Description,
		Date:            invoice.Date,
		DateDue:         invoice.DateDue,
		Total:           invoice.Total,
		SubTotal:        invoice.SubTotal,
		LogoURL:         invoice.LogoURL,
		Items:           *items,
		AdditionalCosts: *costs,
		Client:          *client,
	}).Error
	return invoice, err
}

func (i *invoiceRepository) FindAll() (invoices *domain.Invoices, err error) {
	err = i.db.Preload("Client").Preload("AdditionalCosts").Preload("Items").Find(&invoices).Error
	return invoices, err
}

func (i *invoiceRepository) FindByID(id string) (invoice *domain.Invoice, err error) {
	err = i.db.Preload("Client").Preload("AdditionalCosts").Preload("Items").Where("id = ?", id).Find(&invoice).Error
	return invoice, err
}

func (i *invoiceRepository) FindByUserID(userid string) (invoices *domain.Invoices, err error) {
	err = i.db.Preload("Client").Preload("AdditionalCosts").Preload("Items").Where("user_id = ?", userid).Find(&invoices).Error
	return invoices, err
}

func (i *invoiceRepository) FindByStatus(status string) (invoices *domain.Invoices, err error) {
	err = i.db.Preload("Client").Preload("AdditionalCosts").Preload("Items").Where("status = ?", status).Find(&invoices).Error
	return invoices, err
}

func (i *invoiceRepository) FindByInvoiceNumber(invoiceNumber string) (invoice *domain.Invoice, err error) {
	err = i.db.Preload("Client").Preload("AdditionalCosts").Preload("Items").Where("invoice_number = ?", invoiceNumber).Find(&invoice).Error
	return invoice, err
}

func (i *invoiceRepository) UpdateStatus(status string, invoiceNumber string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := i.db.Model(&invoice).Where("invoice_number = ?", invoiceNumber).Update("status", status).Error
	return &invoice, err
}

func (i *invoiceRepository) FindByUserIDAndStatus(userid string, status string) (invoices *domain.Invoices, err error) {
	err = i.db.Preload("Client").Preload("AdditionalCosts").Preload("Items").Where("user_id = ? AND status = ?", userid, status).Find(&invoices).Error
	return invoices, err
}

func (i *invoiceRepository) UpdateStatusByID(status string, id string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := i.db.Model(&invoice).Where("id = ?", id).Update("status", status).Error
	return &invoice, err
}

func (i *invoiceRepository) UpdateStatusPayment(status string, invoiceNumber string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := i.db.Model(&invoice).Where("invoice_number = ?", invoiceNumber).Update("status_payment", status).Error
	return &invoice, err
}

func (i *invoiceRepository) UpdateTokenSnap(invoiceNumber string, token string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := i.db.Model(&invoice).Where("invoice_number = ?", invoiceNumber).Update("snap_token", token).Error
	return &invoice, err
}

func (i *invoiceRepository) UpdateSnapToken(snapToken string, id string) error {
	err := i.db.Model(&domain.Invoice{}).Where("id = ?", id).Update("snap_token", snapToken).Error
	return err
}

func (i *invoiceRepository) UpdateOrderID(orderID string, id string) error {
	err := i.db.Model(&domain.Invoice{}).Where("id = ?", id).Update("order_id", orderID).Error
	return err
}

func (i *invoiceRepository) FindByOrderID(orderID string) (invoice *domain.Invoice, err error) {
	err = i.db.Preload("Client").Preload("AdditionalCosts").Preload("Items").Where("order_id = ?", orderID).Find(&invoice).Error
	return invoice, err
}

func (i *invoiceRepository) UpdateStatusPaymentByID(status string, id string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := i.db.Model(&invoice).Where("id = ?", id).Update("status_payment", status).Error
	return &invoice, err
}

func (i *invoiceRepository) Delete(invoice *domain.Invoice) error {
	err := i.db.Select(clause.Associations).Where("id = ?", invoice.ID).Delete(&invoice).Error
	return err
}
