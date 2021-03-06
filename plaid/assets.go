package plaid

import (
	"encoding/json"
	"errors"
)

type AssetReport struct {
	AssetReportID  string            `json:"asset_report_id"`
	ClientReportID string            `json:"client_report_id"`
	DateGenerated  string            `json:"date_generated"`
	DaysRequested  int               `json:"days_requested"`
	Items          []AssetReportItem `json:"items"`
	User           AssetReportUser   `json:"user"`
}

type AssetReportItem struct {
	Accounts        []AssetReportAccount `json:"accounts"`
	DateLastUpdated string               `json:"date_last_updated"`
	InstitutionID   string               `json:"institution_id"`
	InstitutionName string               `json:"institution_name"`
	ItemID          string               `json:"item_id"`
}

type AssetReportAccountOwner struct {
	Addresses    []AssetReportAccountOwnerAddress     `json:"addresses"`
	Emails       []AssetReportAccountOwnerEmail       `json:"emails"`
	Names        []string                             `json:"names"`
	PhoneNumbers []AssetReportAccountOwnerPhoneNumber `json:"phone_numbers"`
}

type AssetReportAccountOwnerAddress struct {
	Data    AssetReportAccountOwnerAddressData `json:"data"`
	Primary bool                               `json:"primary"`
}

type AssetReportAccountOwnerAddressData struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	Region     string `json:"region"`
	Street     string `json:"street"`
}

type AssetReportAccountOwnerEmail struct {
	Data    string `json:"data"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}

type AssetReportAccountOwnerPhoneNumber struct {
	Data    string `json:"data"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}

type AssetReportAccount struct {
	Balances           AccountBalances           `json:"balances"`
	HistoricalBalances []AccountBalances         `json:"historical_balances"`
	AccountID          string                    `json:"account_id"`
	DaysAvailable      int                       `json:"days_available"`
	Mask               string                    `json:"mask"`
	Name               string                    `json:"name"`
	OfficialName       string                    `json:"official_name"`
	Owners             []AssetReportAccountOwner `json:"owners"`
	Subtype            string                    `json:"subtype"`
	Type               string                    `json:"type"`
}

type AssetReportUser struct {
	ClientID    string `json:"client_user_id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	PhoneNumber string `json:"phone_number"`
	SSN         string `json:"ssn"`
}

type getAssetReportRequest struct {
	ClientID         string `json:"client_id"`
	Secret           string `json:"secret"`
	AssetReportToken string `json:"asset_report_token"`
}

type GetAssetReportResponse struct {
	APIResponse
	Report   AssetReport `json:"report"`
	Warnings []string    `json:"warnings"`
}

type createAssetReportRequest struct {
	ClientID      string                   `json:"client_id"`
	Secret        string                   `json:"secret"`
	AccessTokens  []string                 `json:"access_tokens"`
	DaysRequested int                      `json:"days_requested"`
	Options       CreateAssetReportOptions `json:"options"`
}

type CreateAssetReportOptions struct {
	ClientReportID string `json:"client_report_id,omitempty"`
	Webhook        string `json:"webhook,omitempty"`
	User           struct {
		ClientUserID string `json:"client_user_id,omitempty"`
		FirstName    string `json:"first_name,omitempty"`
		LastName     string `json:"last_name,omitempty"`
		MiddleName   string `json:"middle_name,omitempty"`
		Ssn          string `json:"ssn,omitempty"`
		PhoneNumber  string `json:"phone_number,omitempty"`
		Email        string `json:"email,omitempty"`
	} `json:"user,omitempty"`
}

type CreateAssetReportResponse struct {
	AssetReportToken string `json:"asset_report_token"`
	AssetReportID    string `json:"asset_report_id"`
	RequestID        string `json:"request_id"`
}

type removeAssetReportRequest struct {
	ClientID         string `json:"client_id"`
	Secret           string `json:"secret"`
	AssetReportToken string `json:"asset_report_token"`
}

type RemoveAssetReportResponse struct {
	APIResponse
	Removed bool `json:"removed"`
}

type createAuditCopyRequest struct {
	ClientID         string `json:"client_id"`
	Secret           string `json:"secret"`
	AssetReportToken string `json:"asset_report_token"`
	AuditorID        string `json:"auditor_id"`
}

type CreateAuditCopyTokenResponse struct {
	APIResponse
	AuditCopyToken string `json:"audit_copy_token"`
}

func (c *Client) GetAssetReport(assetReportToken string) (resp GetAssetReportResponse, err error) {
	if assetReportToken == "" {
		return resp, errors.New("/asset_report/get - asset report token must be specified")
	}

	jsonBody, err := json.Marshal(getAssetReportRequest{
		ClientID:         c.clientID,
		Secret:           c.secret,
		AssetReportToken: assetReportToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/get", jsonBody, &resp)
	return resp, err
}

func (c *Client) CreateAssetReportWithOptions(itemAccessTokens []string, daysRequested int, options CreateAssetReportOptions) (resp CreateAssetReportResponse, err error) {
	if itemAccessTokens == nil || len(itemAccessTokens) == 0 {
		return resp, errors.New("/asset_report/create - asset report token must be specified")
	}

	jsonBody, err := json.Marshal(createAssetReportRequest{
		ClientID:      c.clientID,
		Secret:        c.secret,
		AccessTokens:  itemAccessTokens,
		DaysRequested: daysRequested,
		Options:       options,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/create", jsonBody, &resp)
	return resp, err
}

func (c *Client) CreateAssetReport(itemAccessTokens []string, daysRequested int) (resp CreateAssetReportResponse, err error) {
	if itemAccessTokens == nil || len(itemAccessTokens) == 0 {
		return resp, errors.New("/asset_report/create - asset report token must be specified")
	}

	jsonBody, err := json.Marshal(createAssetReportRequest{
		ClientID:      c.clientID,
		Secret:        c.secret,
		AccessTokens:  itemAccessTokens,
		DaysRequested: daysRequested,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/create", jsonBody, &resp)
	return resp, err
}

func (c *Client) CreateAuditCopy(assetReportToken, auditorID string) (resp CreateAuditCopyTokenResponse, err error) {
	if assetReportToken == "" || auditorID == "" {
		return resp, errors.New("/asset_report/audit_copy/create - asset report token and auditor id must be specified")
	}

	jsonBody, err := json.Marshal(createAuditCopyRequest{
		ClientID:         c.clientID,
		Secret:           c.secret,
		AssetReportToken: assetReportToken,
		AuditorID:        auditorID,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/audit_copy/create", jsonBody, &resp)
	return resp, err
}

func (c *Client) RemoveAssetReport(assetReportToken string) (resp RemoveAssetReportResponse, err error) {
	if assetReportToken == "" {
		return resp, errors.New("/asset_report/remove - asset report token must be specified")
	}

	jsonBody, err := json.Marshal(removeAssetReportRequest{
		ClientID:         c.clientID,
		Secret:           c.secret,
		AssetReportToken: assetReportToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/remove", jsonBody, &resp)
	return resp, err
}
