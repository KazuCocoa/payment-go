package main

import (
	"context"
	"os"

	gpay "payment/payment-service/proto"

	payjp "github.com/payjp/payjp-go/v1" // go get github.com/payjp/payjp-go/v1
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Charge(ctx context.Context, req *gpay.PayRequest) (*gpay.PayResponse, error) {
	pay := payjp.New(os.Getenv("PAYJP_TEST_SECRET_KEY"), nil)

	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency:    "jpy",
		CardToken:   req.Token,
		Capture:     true,
		Description: req.Name + ":" + req.Description,
	})
	if err != nil {
		return nil, err
	}

	res := &gpay.PayResponse{
		Paid:     charge.Paid,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}
	return res, nil
}
