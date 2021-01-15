package quotation

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/net/html"
)

type fnCondition = func(node *html.Attribute) bool

// RenderNode render a node html to bytes
func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

// GetContentHTML search in content html by tag and attribute
func GetContentHTML(doc *html.Node, tag string, fn fnCondition) (*html.Node, error) {
	var crawler func(*html.Node)
	var content *html.Node

	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			for i := range node.Attr {
				if fn(&node.Attr[i]) {
					content = node
					return
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if content != nil {
		return content, nil
	}
	return nil, errors.New(ErrMissingTagOrAttibute.Error())
}
