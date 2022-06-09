package body_email

type VerificationEmailBody struct {
	URL      string
	Username string
	Subject  string
}

type SendPaymentEmailBody struct {
	URL           string
	From          string
	To            string
	Subject       string
	InvoiceNumber string
}

type SendPaymentCashEmailBody struct {
	From          string
	To            string
	Subject       string
	InvoiceNumber string
}
