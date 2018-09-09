package handler

import (
	"context"
	"net/http"
	"strconv"

	gpay "../../payment-service/proto"

	"../db"
	"../domain"

	"google.golang.org/grpc"
)

var addr = "localhost:50051"

func Charge(c Context) {
	t := domain.Payment{}
	c.Bind(&t)
	identifier, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	res, err := db.SelectItem(int64(identifier))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	greq := &gpay.PayRequest{
		Id:          int64(identifier),
		Token:       t.Token,
		Amount:      res.Amount,
		Name:        res.Name,
		Description: res.Description,
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	defer conn.Close()
	client := gpay.NewPayManagerClient(conn)
	gres, err := client.Charge(context.Background(), greq)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, gres)
}
