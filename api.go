package celestia

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// API - wrapper of celestia node API.
type API struct {
	client  *http.Client
	baseURL string
}

// NewAPI - constructor of API
func NewAPI(baseURL string) API {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	client := &http.Client{
		Transport: t,
	}

	return API{client, baseURL}
}

func (api API) get(ctx context.Context, url string, output any) error {
	link := fmt.Sprintf("%s/%s", api.baseURL, url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return err
	}
	response, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(output)
	return err
}

func (api API) post(ctx context.Context, url string, input, output any) error {
	link := fmt.Sprintf("%s/%s", api.baseURL, url)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, buf)
	if err != nil {
		return err
	}
	response, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(output)
	return err
}

// Head - returns the tip (head) of the node's current chain.
func (api API) Head(ctx context.Context) (response HeaderResponse, err error) {
	err = api.get(ctx, "head", &response)
	return
}

// Header - returns the header of the given `height`.
func (api API) Header(ctx context.Context, height uint64) (response HeaderResponse, err error) {
	err = api.get(ctx, fmt.Sprintf("header/%d", height), &response)
	return
}

// NamespaceData - returns original messages of the given namespace ID `namespaceID` from the given block `height`.
func (api API) NamespaceData(ctx context.Context, namespaceID string, height uint64) (response NamespaceData, err error) {
	err = api.get(ctx, fmt.Sprintf("namespaced_data/%s/height/%d", namespaceID, height), &response)
	return
}

// NamespaceShares - returns shares of the given namespace ID `namespaceID` from the latest block (chain head).
func (api API) NamespaceShares(ctx context.Context, namespaceID string) (response NamespaceData, err error) {
	err = api.get(ctx, fmt.Sprintf("namespaced_shares/%s", namespaceID), &response)
	return
}

// NamespaceSharesByHeight - returns shares of the given namespace ID `namespaceID` from the block of the given `height`.
func (api API) NamespaceSharesByHeight(ctx context.Context, namespaceID string, height uint64) (response NamespaceShares, err error) {
	err = api.get(ctx, fmt.Sprintf("namespaced_shares/%s/height/%d", namespaceID, height), &response)
	return
}

// DataAvailable - returns whether data is available at a specific block `height` and the probability that it is available based on the amount of samples collected.
func (api API) DataAvailable(ctx context.Context, height uint64) (response DataAvailableResponse, err error) {
	err = api.get(ctx, fmt.Sprintf("data_available/%d", height), &response)
	return
}

// Balance - returns the balance of the default account address of the node.
func (api API) Balance(ctx context.Context) (response Balance, err error) {
	err = api.get(ctx, "balance", &response)
	return
}

// BalanceOf - returns the balance of the default account address of the node.
func (api API) BalanceOf(ctx context.Context, address string) (response Balance, err error) {
	err = api.get(ctx, fmt.Sprintf("balance/%s", address), &response)
	return
}

// SubmitTx - submits the given transaction to a running instance of celestia-app.
func (api API) SubmitTx(ctx context.Context, tx SubmitTx) (response SubmittedTx, err error) {
	err = api.post(ctx, "submit_tx", tx, &response)
	return
}

// SubmitPfd - Constructs, signs and submits a PayForData message to a running instance of celestia-app. The body of the /submit_pfd request should contain the hex-encoded namespace_id, the hex-encoded data, and the gas_limit as a uint64.
func (api API) SubmitPfd(ctx context.Context, tx SubmitPfd) (response SubmittedPfd, err error) {
	err = api.post(ctx, "submit_pfd", tx, &response)
	return
}
