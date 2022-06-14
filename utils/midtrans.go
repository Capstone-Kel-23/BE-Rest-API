package utils

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"os"
)

type SnapMidtrans struct {
	snap snap.Client
}

func NewPayment() SnapMidtrans {
	return SnapMidtrans{}
}

func (s SnapMidtrans) GeneratePayment(inv domain.Invoice, orderID string) (string, *midtrans.Error) {
	s.snap.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	var items []midtrans.ItemDetails
	if len(inv.Items) != 0 {
		for _, v := range inv.Items {
			items = append(items, midtrans.ItemDetails{
				ID:    v.ID.String(),
				Name:  v.Name,
				Price: int64(v.Price),
				Qty:   int32(v.Quantity),
			})
		}
	}
	if len(inv.AdditionalCosts) != 0 {
		for _, v := range inv.AdditionalCosts {
			if v.Type == "tax" {
				items = append(items, midtrans.ItemDetails{
					ID:       v.ID.String(),
					Name:     v.Name,
					Price:    int64(v.Total),
					Qty:      1,
					Category: v.Type,
				})
			} else {
				items = append(items, midtrans.ItemDetails{
					ID:       v.ID.String(),
					Name:     v.Name,
					Price:    -int64(v.Total),
					Qty:      1,
					Category: v.Type,
				})
			}
		}
	}

	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(inv.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: inv.Client.FirstName,
			LName: inv.Client.LastName,
			Email: inv.Client.Email,
			Phone: inv.Client.PhoneNumber,
		},
		Items: &items,
	}

	snapResp, err := s.snap.CreateTransactionToken(request)
	if err != nil {
		return "", err
	}
	return snapResp, nil
}
