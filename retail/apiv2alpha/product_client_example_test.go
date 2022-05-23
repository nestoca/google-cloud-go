// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package retail_test

import (
	"context"

	retail "cloud.google.com/go/retail/apiv2alpha"
	"google.golang.org/api/iterator"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2alpha"
)

func ExampleNewProductClient() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleProductClient_CreateProduct() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.CreateProductRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#CreateProductRequest.
	}
	resp, err := c.CreateProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_GetProduct() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetProductRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#GetProductRequest.
	}
	resp, err := c.GetProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_ListProducts() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.ListProductsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#ListProductsRequest.
	}
	it := c.ListProducts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleProductClient_UpdateProduct() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateProductRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#UpdateProductRequest.
	}
	resp, err := c.UpdateProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_DeleteProduct() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.DeleteProductRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#DeleteProductRequest.
	}
	err = c.DeleteProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProductClient_PurgeProducts() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.PurgeProductsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#PurgeProductsRequest.
	}
	op, err := c.PurgeProducts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_ImportProducts() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.ImportProductsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#ImportProductsRequest.
	}
	op, err := c.ImportProducts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_SetInventory() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.SetInventoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#SetInventoryRequest.
	}
	op, err := c.SetInventory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_AddFulfillmentPlaces() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.AddFulfillmentPlacesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#AddFulfillmentPlacesRequest.
	}
	op, err := c.AddFulfillmentPlaces(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_RemoveFulfillmentPlaces() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.RemoveFulfillmentPlacesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#RemoveFulfillmentPlacesRequest.
	}
	op, err := c.RemoveFulfillmentPlaces(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_AddLocalInventories() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.AddLocalInventoriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#AddLocalInventoriesRequest.
	}
	op, err := c.AddLocalInventories(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_RemoveLocalInventories() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.RemoveLocalInventoriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2alpha#RemoveLocalInventoriesRequest.
	}
	op, err := c.RemoveLocalInventories(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}