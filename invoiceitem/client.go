// Package invoiceitem provides the /invoiceitems APIs
package invoiceitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /invoiceitems APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new invoice items.
// For more details see https://stripe.com/docs/api#create_invoiceitem.
func New(params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	invoiceItem := &stripe.InvoiceItem{}
	err := c.B.Call(http.MethodPost, "/invoiceitems", c.Key, params, invoiceItem)
	return invoiceItem, err
}

// Get returns the details of an invoice item.
// For more details see https://stripe.com/docs/api#retrieve_invoiceitem.
func Get(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	path := stripe.FormatURLPath("/invoiceitems/%s", id)
	invoiceItem := &stripe.InvoiceItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoiceItem)
	return invoiceItem, err
}

// Update updates an invoice item's properties.
// For more details see https://stripe.com/docs/api#update_invoiceitem.
func Update(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	path := stripe.FormatURLPath("/invoiceitems/%s", id)
	invoiceItem := &stripe.InvoiceItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoiceItem)
	return invoiceItem, err
}

// Del removes an invoice item.
// For more details see https://stripe.com/docs/api#delete_invoiceitem.
func Del(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	path := stripe.FormatURLPath("/invoiceitems/%s", id)
	ii := &stripe.InvoiceItem{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, ii)
	return ii, err
}

// List returns a list of invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
func List(params *stripe.InvoiceItemListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.InvoiceItemListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.InvoiceItemList{}
		err := c.B.CallRaw(http.MethodGet, "/invoiceitems", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of InvoiceItems.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// InvoiceItem returns the most recent InvoiceItem
// visited by a call to Next.
func (i *Iter) InvoiceItem() *stripe.InvoiceItem {
	return i.Current().(*stripe.InvoiceItem)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
