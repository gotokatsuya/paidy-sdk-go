package paidy

import "time"

type Payment struct {
	ID          string    `json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
	Amount      int       `json:"amount,omitempty"`
	Currency    string    `json:"currency,omitempty"`
	Description string    `json:"description,omitempty"`
	StoreName   string    `json:"store_name,omitempty"`
	Test        bool      `json:"test,omitempty"`
	Status      string    `json:"status,omitempty"`
	Tier        string    `json:"tier,omitempty"`
	Buyer       struct {
		Name1 string `json:"name1,omitempty"`
		Name2 string `json:"name2,omitempty"`
		Email string `json:"email,omitempty"`
		Phone string `json:"phone,omitempty"`
	} `json:"buyer,omitempty"`
	Order struct {
		Items []struct {
			ID          string `json:"id,omitempty"`
			Title       string `json:"title,omitempty"`
			Description string `json:"description,omitempty"`
			UnitPrice   int    `json:"unit_price,omitempty"`
			Quantity    int    `json:"quantity,omitempty"`
		} `json:"items,omitempty"`
		Tax       int    `json:"tax,omitempty"`
		Shipping  int    `json:"shipping,omitempty"`
		OrderRef  string `json:"order_ref,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	} `json:"order,omitempty"`
	ShippingAddress struct {
		Line1 string `json:"line1,omitempty"`
		Line2 string `json:"line2,omitempty"`
		City  string `json:"city,omitempty"`
		State string `json:"state,omitempty"`
		Zip   string `json:"zip,omitempty"`
	} `json:"shipping_address,omitempty"`
	Captures []struct {
		ID        string    `json:"id,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		Amount    int       `json:"amount,omitempty"`
		Tax       int       `json:"tax,omitempty"`
		Shipping  int       `json:"shipping,omitempty"`
		Items     []struct {
			ID          string `json:"id,omitempty"`
			Title       string `json:"title,omitempty"`
			Description string `json:"description,omitempty"`
			UnitPrice   int    `json:"unit_price,omitempty"`
			Quantity    int    `json:"quantity,omitempty"`
		} `json:"items,omitempty"`
		Metadata struct {
			Key1 string `json:"key1,omitempty"`
			Key2 string `json:"key2,omitempty"`
		} `json:"metadata,omitempty"`
	} `json:"captures,omitempty"`
	Refunds  []any `json:"refunds,omitempty"`
	Metadata struct {
	} `json:"metadata,omitempty"`
}
