package quotation

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"quotation/config"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func unmarshalQuotationOptionsTableNonce(node *html.Node) (*OptionsTableNonce, error) {
	tag := RenderNode(node)
	start := strings.Index(tag, "{")
	last := strings.Index(tag, "}")

	var quotationOptions OptionsTableNonce
	err := json.Unmarshal([]byte(tag[start:last+1]), &quotationOptions)
	return &quotationOptions, err
}

func getQuotationOptionsTableNonce() (*OptionsTableNonce, error) {
	res, err := http.Get(config.ScrapEndpoint)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body := res.Body
	htmlParsed, err := html.Parse(body)

	if err != nil {
		log.Fatalln(err)
	}
	condition := func(node *html.Attribute) bool {
		return node.Key == "id" && node.Val == "tool-cotacoes-opcoes-js-extra"
	}
	b, err := GetContentHTML(htmlParsed, "script", condition)
	if err != nil {
		log.Fatalln(err)
	}
	q, err := unmarshalQuotationOptionsTableNonce(b)
	return q, err
}

func makeFormData(page int, optionsTableOnce string) *url.Values {
	form := url.Values{}
	form.Add("action", "tool_cotacoes_opcoes")
	form.Add("pagination", strconv.Itoa(page))
	form.Add("cotacoes_opcoes_table_nonce", optionsTableOnce)
	return &form
}

// GetQuotation get data quotations
func GetQuotation(URI string, form *url.Values) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", URI, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, err
}
