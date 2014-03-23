package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type QuoteRequestReject struct {
	quickfixgo.Message
}

func (m *QuoteRequestReject) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequestReject) RFQReqID() (*field.RFQReqID, error) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequestReject) QuoteRequestRejectReason() (*field.QuoteRequestRejectReason, error) {
	f := new(field.QuoteRequestRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequestReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequestReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequestReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}