package invoice

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/invoice/delivery/http"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/invoice/repository"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/invoice/usecase"
	mid "github.com/Capstone-Kel-23/BE-Rest-API/internal/user/delivery/http/middleware"
	repository2 "github.com/Capstone-Kel-23/BE-Rest-API/internal/user/repository"
	usecase2 "github.com/Capstone-Kel-23/BE-Rest-API/internal/user/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InvoiceRouter(c *echo.Echo, db *gorm.DB) {
	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	invoiceRepository := repository.NewInvoiceRepository(db)
	userRepository := repository2.NewUserRepository(db)

	userUsecase := usecase2.NewUserUsecase(userRepository)
	invoiceUsecase := usecase.NewInvoiceUsecase(invoiceRepository, userRepository)
	invoiceController := http.NewInvoiceController(invoiceUsecase, userUsecase)

	c.POST("/api/v1/invoice", invoiceController.CreateNewInvoice, authMiddleware)
	c.GET("/api/v1/invoices", invoiceController.GetListAllInvoices, authMiddleware)
	c.GET("/api/v1/invoice/detail", invoiceController.GetDetailInvoice, authMiddleware)
	c.GET("/api/v1/invoice/detail/:id", invoiceController.GetDetailInvoiceByID, authMiddleware)
	c.GET("/api/v1/invoices/status", invoiceController.GetListInvoicesByUserIDAndStatus, authMiddleware)
	c.PUT("/api/v1/invoice/status/:id", invoiceController.UpdateStatusInvoice, authMiddleware)
	c.GET("/api/v1/invoices/user", invoiceController.GetListInvoicesByUserID, authMiddleware)
	c.POST("/api/v1/invoice/send/:id", invoiceController.SendPaymentAndInvoice, authMiddleware)
	c.PUT("/api/v1/invoice/status-payment/:id", invoiceController.UpdateStatusPayment, authMiddleware)
	c.DELETE("/api/v1/invoice/:id", invoiceController.DeleteInvoiceByID, authMiddleware)
	c.POST("/api/v1/invoice/file", invoiceController.CreateInvoiceWithExcel, authMiddleware)

	c.POST("/api/v1/invoice/payment/handling", invoiceController.PaymentHandling)
	c.GET("/api/v1/invoice/payment-generate/:inv", invoiceController.GeneratePayment)
}
