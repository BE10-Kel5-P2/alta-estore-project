package midtrans

import (
	"fmt"
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func InitConnection(serverkey string) snap.Client {
	var c = snap.Client{}
	c.New(serverkey, midtrans.Sandbox)

	return c
}

func CreateConnection(s snap.Client, total, id int) *snap.Response {
	req := &snap.Request{

		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("be10-%d", id),
			GrossAmt: int64(total),
		},
	}

	apiRes, err := s.CreateTransaction(req)

	if err != nil {
		log.Println("payment err: ", err)
	}

	return apiRes
}
