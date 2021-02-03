package plaid

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fatih/structs"
)

type UnofficialCurrencyCode string

const (
	unofficialCurrencyCodeADA  UnofficialCurrencyCode = "ADA"
	unofficialCurrencyCodeBAT  UnofficialCurrencyCode = "BAT"
	unofficialCurrencyCodeBCH  UnofficialCurrencyCode = "BCH"
	unofficialCurrencyCodeBNB  UnofficialCurrencyCode = "BNB"
	unofficialCurrencyCodeBTC  UnofficialCurrencyCode = "BTC"
	unofficialCurrencyCodeBTG  UnofficialCurrencyCode = "BTG"
	unofficialCurrencyCodeCNH  UnofficialCurrencyCode = "CNH"
	unofficialCurrencyCodeDASH UnofficialCurrencyCode = "DASH"
	unofficialCurrencyCodeDOGE UnofficialCurrencyCode = "DOGE"
	unofficialCurrencyCodeETC  UnofficialCurrencyCode = "ETC"
	unofficialCurrencyCodeETH  UnofficialCurrencyCode = "ETH"
	unofficialCurrencyCodeGBX  UnofficialCurrencyCode = "GBX"
	unofficialCurrencyCodeLSK  UnofficialCurrencyCode = "LSK"
	unofficialCurrencyCodeNEO  UnofficialCurrencyCode = "NEO"
	unofficialCurrencyCodeOMG  UnofficialCurrencyCode = "OMG"
	unofficialCurrencyCodeQTUM UnofficialCurrencyCode = "QTUM"
	unofficialCurrencyCodeUSDT UnofficialCurrencyCode = "USDT"
	unofficialCurrencyCodeXLM  UnofficialCurrencyCode = "XLM"
	unofficialCurrencyCodeXMR  UnofficialCurrencyCode = "XMR"
	unofficialCurrencyCodeXRP  UnofficialCurrencyCode = "XRP"
	unofficialCurrencyCodeZEC  UnofficialCurrencyCode = "ZEC"
	unofficialCurrencyCodeZRX  UnofficialCurrencyCode = "ZRX"
)

// MarshalUUID converts a uuid.UUID to a graphql string
func MarshalUnofficialCurrencyCode(s UnofficialCurrencyCode) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(string(s)))
	})
}

// UnmarshalUUID converts a graphql uuid string type into a uuid.UUID
func UnmarshalUnofficialCurrencyCode(v interface{}) (UnofficialCurrencyCode, error) {
	switch v := v.(type) {
	case string:
		if IsValidCode(v) {
			return UnofficialCurrencyCode(v), nil
		} else {
			return "", fmt.Errorf("%T is not a valid unofficial currency code", v)
		}
	default:
		return "", fmt.Errorf("%T is not a unofficial currency code", v)
	}
}

type unofficialCurrencyCodes struct {
	ADA  UnofficialCurrencyCode
	BAT  UnofficialCurrencyCode
	BCH  UnofficialCurrencyCode
	BNB  UnofficialCurrencyCode
	BTC  UnofficialCurrencyCode
	BTG  UnofficialCurrencyCode
	CNH  UnofficialCurrencyCode
	DASH UnofficialCurrencyCode
	DOGE UnofficialCurrencyCode
	ETC  UnofficialCurrencyCode
	ETH  UnofficialCurrencyCode
	GBX  UnofficialCurrencyCode
	LSK  UnofficialCurrencyCode
	NEO  UnofficialCurrencyCode
	OMG  UnofficialCurrencyCode
	QTUM UnofficialCurrencyCode
	USDT UnofficialCurrencyCode
	XLM  UnofficialCurrencyCode
	XMR  UnofficialCurrencyCode
	XRP  UnofficialCurrencyCode
	ZEC  UnofficialCurrencyCode
	ZRX  UnofficialCurrencyCode
}

var UnofficialCurrencyCodes unofficialCurrencyCodes = unofficialCurrencyCodes{
	ADA:  unofficialCurrencyCodeADA,
	BAT:  unofficialCurrencyCodeBAT,
	BCH:  unofficialCurrencyCodeBCH,
	BNB:  unofficialCurrencyCodeBNB,
	BTC:  unofficialCurrencyCodeBTC,
	BTG:  unofficialCurrencyCodeBTG,
	CNH:  unofficialCurrencyCodeCNH,
	DASH: unofficialCurrencyCodeDASH,
	DOGE: unofficialCurrencyCodeDOGE,
	ETC:  unofficialCurrencyCodeETC,
	ETH:  unofficialCurrencyCodeETH,
	GBX:  unofficialCurrencyCodeGBX,
	LSK:  unofficialCurrencyCodeLSK,
	NEO:  unofficialCurrencyCodeNEO,
	OMG:  unofficialCurrencyCodeOMG,
	QTUM: unofficialCurrencyCodeQTUM,
	USDT: unofficialCurrencyCodeUSDT,
	XLM:  unofficialCurrencyCodeXLM,
	XMR:  unofficialCurrencyCodeXMR,
	XRP:  unofficialCurrencyCodeXRP,
	ZEC:  unofficialCurrencyCodeZEC,
	ZRX:  unofficialCurrencyCodeZRX,
}

func IsValidCode(s string) bool {
	for _, value := range structs.Map(UnofficialCurrencyCodes) {
		if string(value.(string)) == s {
			return true
		}
	}
	return false
}
