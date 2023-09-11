package paidy

import "time"

type Token struct {
	ID         string `json:"id,omitempty"`
	MerchantID string `json:"merchant_id,omitempty"`
	WalletID   string `json:"wallet_id,omitempty"`
	Status     string `json:"status,omitempty"`
	Origin     struct {
		Name1   string `json:"name1,omitempty"`
		Name2   string `json:"name2,omitempty"`
		Email   string `json:"email,omitempty"`
		Phone   string `json:"phone,omitempty"`
		Address struct {
			Line1 string `json:"line1,omitempty"`
			Line2 string `json:"line2,omitempty"`
			State string `json:"state,omitempty"`
			City  string `json:"city,omitempty"`
			Zip   string `json:"zip,omitempty"`
		} `json:"address,omitempty"`
	} `json:"origin,omitempty"`
	Description string `json:"description,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Metadata    struct {
	} `json:"metadata,omitempty"`
	WebhookURL  string `json:"webhook_url,omitempty"`
	ConsumerID  string `json:"consumer_id,omitempty"`
	Suspensions []struct {
		Timestamp time.Time `json:"timestamp,omitempty"`
		Authority string    `json:"authority,omitempty"`
	} `json:"suspensions,omitempty"`
	Test        bool      `json:"test,omitempty"`
	VersionNr   int       `json:"version_nr,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	ActivatedAt time.Time `json:"activated_at,omitempty"`
	DeletedAt   string    `json:"deleted_at,omitempty"`
}
