package gopaypal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type paymentCreateResponse struct {
	ID            string        `json:"id"`
	CreateTime    time.Time     `json:"create_time"`
	UpdateTime    time.Time     `json:"update_time"`
	State         string        `json:"state"`
	Intent        string        `json:"intent,omitempty"`
	Payer         Payer         `json:"payer,omitempty"`
	Transactions  []Transaction `json:"transactions,omitempty"`
	FailureReason string        `json:"failure_reason,omitempty"`
	RedirectURL   RedirectURL   `json:"redirect_urls,omitempty"`
	Links         []Link        `json:"links"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type Payment struct {
	Intent       string        `json:"intent,omitempty"`
	Payer        Payer         `json:"payer,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
	RedirectURL  RedirectURL   `json:"redirect_urls,omitempty"`
}

type Payer struct {
	PaymentMethod string    `json:"payment_method,omitempty"`
	Status        string    `json:"status,omitempty"`
	Info          PayerInfo `json:"payer_info,omitempty"`
}

type PayerInfo struct {
	ID string `json:"payer_id,omitempty"`
}

type Transaction struct {
	Amount           Amount              `json:"amount,omitempty"`
	Description      string              `json:"description,omitempty"`
	NoteToPayee      string              `json:"note_to_payee,omitempty"`
	Custom           string              `json:"custom,omitempty"`
	InvoiceNumber    string              `json:"invoice_number,omitempty"`
	SoftDescriptor   string              `json:"soft_descriptor,omitempty"`
	PaymentOptions   PaymentOptions      `json:"payment_options,omitempty"`
	ItemList         ItemList            `json:"item_list,omitempty"`
	NotifyURL        string              `json:"notify_url,omitempty"`
	OrderURL         string              `json:"order_url,omitempty"`
	RelatedResources []*RelatedResources `json:"related_resources,omitempty"`
}

type RelatedResources struct {
	Sale Sale `json:"sale,omitempty"`
}

type Sale struct {
	ID                      string `json:"id,omitempty"`
	PurchaseUnitReferenceID string `json:"purchase_unit_reference_id,omitempty"`
	Amount                  Amount `json:"amount,omitempty"`
	PaymentMode             string `json:"payment_mode,omitempty"`
	State                   string `json:"state,omitempty"`
	ReasonCode              string `json:"reason_code,omitempty"`
	ClearingTime            string `json:"clearing_time,omitempty"`
	ReceiptID               string `json:"receipt_id,omitempty"`
}

type Amount struct {
	Currency string  `json:"currency,omitempty"`
	Total    string  `json:"total,omitempty"`
	Details  Details `json:"details,omitempty"`
}

type Details struct {
	SubTotal         string `json:"subtotal,omitempty"`
	Shipping         string `json:"shipping,omitempty"`
	Tax              string `json:"tax,omitempty"`
	HandlingFee      string `json:"handling_fee,omitempty"`
	ShippingDiscount string `json:"shipping_discount,omitempty"`
	Insurance        string `json:"insurance,omitempty"`
	GiftWrap         string `json:"gift_wrap,omitempty"`
}

type PaymentOptions struct {
	AllowedPaymentMethod string `json:"allowed_payment_method,omitempty"`
}

type ItemList struct {
	Items               []Item `json:"items,omitempty"`
	ShippingMethod      string `json:"shipping_method,omitempty"`
	ShippingPhoneNumber string `json:"shipping_phone_number,omitempty"`
}

type Item struct {
	Sku         string `json:"sku,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Price       string `json:"price,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Tax         string `json:"tax,omitempty"`
	URL         string `json:"url,omitempty"`
}

type RedirectURL struct {
	ReturnURL string `json:"return_url,omitempty"`
	CancelURL string `json:"cancel_url,omitempty"`
}

func (c Client) PaymentInformation(paymentID string) (*paymentCreateResponse, error) {
	// Create auth request
	req, err := c.AuthRequest(fmt.Sprintf(
		PaymentInfoURL,
		paymentID,
	), nil, http.MethodGet)

	if err != nil {
		return nil, err
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	res, err := c.Execute(req)

	if err != nil {
		return nil, err
	}

	// Hold payment creation response
	d := paymentCreateResponse{}

	// Marshal response
	if err := json.Unmarshal(res, &d); err != nil {
		return nil, err
	}

	return &d, err
}

func (c Client) ExecutePayment(paymentID, payerID string) (*paymentCreateResponse, error) {
	// Create auth request
	req, err := c.AuthRequest(fmt.Sprintf(
		PaymentExecuteURL,
		paymentID,
	), []byte("{\"payer_id\": \""+payerID+"\"}"), http.MethodPost)

	if err != nil {
		return nil, err
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	res, err := c.Execute(req)

	if err != nil {
		return nil, err
	}

	// Hold payment creation response
	d := paymentCreateResponse{}

	// Marshal response
	if err := json.Unmarshal(res, &d); err != nil {
		return nil, err
	}

	return &d, err
}

// CreatePayment creates a PayPal payment with the given payment object
func (c Client) CreatePayment(payment Payment) (*paymentCreateResponse, error) {
	// Marshal payment object to byte array
	buff, err := json.Marshal(&payment)

	if err != nil {
		return nil, err
	}

	// Create auth request
	req, err := c.AuthRequest(PaymentCreateURL, buff, http.MethodPost)

	if err != nil {
		return nil, err
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	res, err := c.Execute(req)

	if err != nil {
		return nil, err
	}

	// Hold payment creation response
	d := paymentCreateResponse{}

	// Marshal response
	if err := json.Unmarshal(res, &d); err != nil {
		return nil, err
	}

	return &d, err
}
