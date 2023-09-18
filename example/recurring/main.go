package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gotokatsuya/paidy-sdk-go/paidy"
)

func main() {

	paidyCli, err := paidy.New(os.Getenv("PAIDY_SECRET_KEY"))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			PaidyPublicKey string
		}{
			PaidyPublicKey: os.Getenv("PAIDY_PUBLIC_KEY"),
		}
		if err := template.Must(template.ParseFiles("index.html")).Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/recurring", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		status := strings.ToLower(r.URL.Query().Get("status"))
		if status != "active" {
			http.Error(w, "invalid token status", http.StatusInternalServerError)
			return
		}
		// NOTE: save token to storage
		tokenID := r.URL.Query().Get("id")
		payment, err := paidyCli.PaymentCreate(ctx, &paidy.PaymentCreateRequest{
			TokenID:  tokenID,
			Amount:   100,
			Currency: "JPY",
			BuyerData: paidy.BuyerData{
				UserID:          "aqwsedrftgyhujiklp",
				Age:             20,
				OrderCount:      0,
				Ltv:             0,
				LastOrderAmount: 0,
				LastOrderAt:     0,
			},
			Order: paidy.Order{
				Items: []paidy.Item{
					{
						Quantity:  1,
						UnitPrice: 100,
					},
				},
			},
			ShippingAddress: paidy.ShippingAddress{
				Zip:  "106-2004",
				City: "東京都",
			},
		})
		if err != nil {
			switch err := err.(type) {
			case *paidy.APIError:
				data, _ := json.MarshalIndent(err, "", " ")
				fmt.Printf("%v\n", string(data))
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(payment, "", " ")
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Printf("%v\n", string(data))
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})

	fmt.Println("http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
