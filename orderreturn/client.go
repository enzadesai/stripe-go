package orderreturn

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /orders APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// For more details see https://stripe.com/docs/api#list_orders
func List(params *stripe.OrderReturnListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.OrderReturnListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.OrderReturnList{}
		err := c.B.CallRaw(http.MethodGet, "/order_returns", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of OrderReturns.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// OrderReturn returns the most recent OrderReturn
// visited by a call to Next.
func (i *Iter) OrderReturn() *stripe.OrderReturn {
	return i.Current().(*stripe.OrderReturn)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
