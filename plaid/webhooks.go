package plaid

import (
	"encoding/json"
	"errors"
)

// Plaid Webhook Types
// PlaidWebhook is basis of webhook return type
type Webhook struct {

	// Available on all Plaid Webhooks
	WebhookType string        `json:"webhook_type"`
	WebhookCode string        `json:"webhook_code"`
	Error       *WebhookError `json:"error"`

	// Assets
	AssetReportID *string `json:"asset_report_id"`

	// Available on Auth
	AccountID *string `json:"account_id"`

	// Available on Item, Transaction Webhooks
	ItemID *string `json:"item_id"`

	// Available on Item Webhooks
	NewWebhookUrl         *string `json:"new_webhook_url"`
	ConsentExpirationTime *string `json:"consent_expiration_time"`

	// Available on Transaction Webhooks
	NewTransactions     *int      `json:"new_transactions"`
	RemovedTransactions *[]string `json:"removed_transactions"`
}

// PlaidWebhookError exists in all Plaid Webhooks
type WebhookError struct {
	DisplayMessage string `json:"display_message"`
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	ErrorType      string `json:"error_type"`
}

// JWT Payload
type JWTPayload struct {
	IAT     int64  `json:"iat"`
	BodySHA string `json:"request_body_sha256"`
}

type WebhookVerificationKey struct {
	Alg       string `json:"alg"`
	CreatedAt int64  `json:"created_at"`
	Crv       string `json:"crv"`
	ExpiredAt int64  `json:"expired_at"`
	Kid       string `json:"kid"`
	Kty       string `json:"kty"`
	Use       string `json:"use"`
	X         string `json:"x"`
	Y         string `json:"y"`
}

type GetWebhookVerificationKeyResponse struct {
	APIResponse
	Key WebhookVerificationKey `json:"key"`
}

type getWebhookVerificationKeyRequest struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	KeyID    string `json:"key_id"`
}

// GetWebhookVerificationKey retrieves the verification key for a given webhook verification key ID
// See https://plaid.com/docs/api/webhook-verification/.
func (c *Client) GetWebhookVerificationKey(
	keyID string,
) (resp GetWebhookVerificationKeyResponse, err error) {
	if keyID == "" {
		return resp, errors.New("/webhook_verification_key/get - key ID must be specified")
	}

	req := getWebhookVerificationKeyRequest{
		ClientID: c.clientID,
		Secret:   c.secret,
		KeyID:    keyID,
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/webhook_verification_key/get", jsonBody, &resp)
	return resp, err
}
