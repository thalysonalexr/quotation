package quotation

import (
	"encoding/json"
	"strconv"
)

// OptionsTableNonce data
type OptionsTableNonce struct {
	CotacoesOpcoesTableNonce string `json:"cotacoes_opcoes_table_nonce"`
	AjaxURL                  string `json:"ajax_url"`
	HomeURL                  string `json:"home_url"`
	TemplateURL              string `json:"template_url"`
}

// Quotation data
type Quotation struct {
	Volume            int     `json:"Volume"`
	Price             float32 `json:"Price"`
	Trades            int     `json:"Trades"`
	VarDay            float64 `json:"VarDay"`
	StockCode         string  `json:"StockCode"`
	Symbol            string  `json:"Symbol"`
	StockCurrentPrice float32 `json:"StockCurrentPrice"`
	StrikePrice       float32 `json:"StrikePrice"`
	ExerciseStyle     int     `json:"ExerciseStyle"`
	MaturityDate      string  `json:"MaturityDate"`
	PutOrCall         int     `json:"PutOrCall"`
	MaturityMonthYear string  `json:"MaturityMonthYear"`
	LastTradeTime     string  `json:"LastTradeTime"`
}

// Response data
type Response struct {
	HasMore bool        `json:"HasMore"`
	Options []Quotation `json:"Options"`
}

// NewQuotationResponse create new quotation json data
func NewQuotationResponse(data []byte) (*Response, error) {
	q := Response{}
	err := json.Unmarshal([]byte(string(data)), &q)
	return &q, err
}

// GetDescriptionKeys get array string description fields
func GetDescriptionKeys() *[]string {
	descriptions := []string{
		"Cod empresa",      // Symbol - ok
		"Pre de exercicio", // PutOrCall 1 == call 0 else put - ok
		"Cotacao",          // StrikePrice - ok
		"Variacao",         // Price - ok
		"Vol negociado",    // VarDay - ok
		"N de negocios",    // Volume - ok
		"Data/hora",        // Trades - ok
		"Vencimento",       // LastTradeTime - ok
		"Modelo",           // MaturityDate - ok
		"Nome da acao",     // ExerciseStyle 1 == EU else US - ok
		"Cotacao da acao",  // StockCode - ok
	}
	return &descriptions
}

// ToArrayString transform values quotation in array string
func ToArrayString(data *Quotation) *[]string {
	var putOrCall, exerciseStyle string
	if data.PutOrCall == 1 {
		putOrCall = "CALL"
	} else {
		putOrCall = "PUT"
	}

	if data.ExerciseStyle == 1 {
		exerciseStyle = "EU"
	} else {
		exerciseStyle = "US"
	}

	var keyValue []string
	keyValue = append(keyValue, data.Symbol)
	keyValue = append(keyValue, putOrCall)
	keyValue = append(keyValue, strconv.FormatFloat(float64(data.StrikePrice), 'f', 2, 64))
	keyValue = append(keyValue, strconv.FormatFloat(float64(data.Price), 'f', 2, 64))
	keyValue = append(keyValue, strconv.FormatFloat(float64(data.VarDay), 'f', 2, 64))
	keyValue = append(keyValue, strconv.Itoa(data.Volume))
	keyValue = append(keyValue, strconv.Itoa(data.Trades))
	keyValue = append(keyValue, data.LastTradeTime)
	keyValue = append(keyValue, data.MaturityDate)
	keyValue = append(keyValue, exerciseStyle)
	keyValue = append(keyValue, data.StockCode)

	return &keyValue
}
