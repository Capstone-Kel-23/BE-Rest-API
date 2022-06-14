package domain

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/web/request"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Invoice struct {
	ID              uuid.UUID        `json:"id" gorm:"PrimaryKey"`
	UserID          uuid.UUID        `json:"user_id" gorm:"type:varchar;size:191"`
	InvoiceNumber   string           `json:"invoice_number"`
	Status          string           `json:"status"`
	StatusPayment   string           `json:"status_payment"`
	TypePayment     string           `json:"type_payment"`
	SnapToken       string           `json:"snap_token"`
	OrderID         string           `json:"order_id"`
	Description     string           `json:"description"`
	Date            time.Time        `json:"date"`
	DateDue         time.Time        `json:"date_due"`
	Total           float64          `json:"total"`
	SubTotal        float64          `json:"sub_total"`
	LogoURL         string           `json:"logo_url"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Items           []Item           `json:"items,omitempty" gorm:"foreignKey:InvoiceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AdditionalCosts []AdditionalCost `json:"additional_costs,omitempty" gorm:"foreignKey:InvoiceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Client          Client           `json:"client,omitempty" gorm:"foreignKey:InvoiceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Invoices []Invoice

type InvoiceRepository interface {
	FindAll() (*Invoices, error)
	FindByID(id string) (*Invoice, error)
	FindByUserID(userid string) (*Invoices, error)
	FindByUserIDAndStatus(userid string, status string) (*Invoices, error)
	FindByStatus(status string) (*Invoices, error)
	FindByInvoiceNumber(invoiceNumber string) (*Invoice, error)
	FindByOrderID(orderID string) (*Invoice, error)
	UpdateStatus(status string, invoiceNumber string) (*Invoice, error)
	UpdateStatusPayment(status string, invoiceNumber string) (*Invoice, error)
	UpdateStatusPaymentByID(status string, id string) (*Invoice, error)
	UpdateStatusByID(status string, id string) (*Invoice, error)
	UpdateTokenSnap(invoiceNumber string, token string) (*Invoice, error)
	UpdateSnapToken(snapToken string, id string) error
	UpdateOrderID(orderID string, id string) error
	Save(invoice *Invoice, items *[]Item, client *Client, costs *[]AdditionalCost) (*Invoice, error)
	Delete(invoice *Invoice) error
}

type InvoiceUsecase interface {
	CreateNewInvoice(req request.InvoiceCreateRequest, userid string) (interface{}, error)
	UpdateStatusInvoice(status, invoiceNumber string) (interface{}, error)
	UpdateStatusInvoicePayment(status, invoiceNumber string) (interface{}, error)
	UpdateStatusInvoicePaymentByID(status, id string) (interface{}, error)
	UpdateStatusInvoiceByID(status, id string) (interface{}, error)
	UpdateStatusInvoiceAndPayment(transaction, orderID, fraud, paymentType string)
	GetListAllInvoices() (*Invoices, error)
	GetListInvoicesByUserIDAndStatus(status string, userid string) (*Invoices, error)
	GetListInvoicesByUserID(userid string) (*Invoices, error)
	GetDetailInvoice(invoiceNumber string) (interface{}, error)
	GetDetailInvoiceByID(id string) (interface{}, error)
	SendPaymentToClient(inv string) (interface{}, error)
	DeletePaymentByID(id string) error
	CreateInvoiceWithExcel(filename string, userid string) (interface{}, error)
}
