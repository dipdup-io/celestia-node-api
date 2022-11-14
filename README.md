# Celestia node API

[![Test Status](https://github.com/dipdup-net/celestia-node-api/workflows/tests/badge.svg)](https:/github.com/dipdup-net/celestia-node-api/actions?query=branch%3Amaster+workflow%3A%22tests%22)
[![made_with golang](https://img.shields.io/badge/made_with-golang-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/dipdup-net/celestia-node-api.svg)](https://pkg.go.dev/github.com/dipdup-net/celestia-node-api)

Golang wrapper over Celestia node API. Library implements API from the [docs](https://docs.celestia.org/developers/node-api).

## Install

```bash
go get github.com/dipdup-net/celestia-node-api
```

## Usage

First, create API structure

```go
api := NewAPI("base URL to Celestia node API")
```

API endpoints are implemented in library

```go
// Head - returns the tip (head) of the node's current chain.
func (api API) Head(ctx context.Context) (response HeaderResponse, err error)

// Header - returns the header of the given `height`.
func (api API) Header(ctx context.Context, height uint64) (response HeaderResponse, err error) 

// NamespaceData - returns original messages of the given namespace ID `namespaceID` from the given block `height`.
func (api API) NamespaceData(ctx context.Context, namespaceID string, height uint64) (response NamespaceData, err error)

// NamespaceShares - returns shares of the given namespace ID `namespaceID` from the latest block (chain head).
func (api API) NamespaceShares(ctx context.Context, namespaceID string) (response NamespaceData, err error)

// NamespaceSharesByHeight - returns shares of the given namespace ID `namespaceID` from the block of the given `height`.
func (api API) NamespaceSharesByHeight(ctx context.Context, namespaceID string, height uint64) (response NamespaceShares, err error)

// DataAvailable - returns whether data is available at a specific block `height` and the probability that it is available based on the amount of samples collected.
func (api API) DataAvailable(ctx context.Context, height uint64) (response DataAvailableResponse, err error)

// Balance - returns the balance of the default account address of the node.
func (api API) Balance(ctx context.Context) (response Balance, err error)

// BalanceOf - returns the balance of the default account address of the node.
func (api API) BalanceOf(ctx context.Context, address string) (response Balance, err error)

// SubmitTx - submits the given transaction to a running instance of celestia-app.
func (api API) SubmitTx(ctx context.Context, tx SubmitTx) (response SubmittedTx, err error)

// SubmitPfd - Constructs, signs and submits a PayForData message to a running instance of celestia-app. The body of the /submit_pfd request should contain the hex-encoded namespace_id, the hex-encoded data, and the gas_limit as a uint64.
func (api API) SubmitPfd(ctx context.Context, tx SubmitPfd) (response SubmittedPfd, err error)
```

## Example

```go
package main

import (
	"context"
	"log"
	"time"
)

func main() {
	api := NewAPI("base URL to Celestia node API")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	head, err := api.Head(ctx)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%##v", head)
}
```