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

package aiplatform

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newMigrationClientHook clientHook

// MigrationCallOptions contains the retry settings for each method of MigrationClient.
type MigrationCallOptions struct {
	SearchMigratableResources []gax.CallOption
	BatchMigrateResources     []gax.CallOption
}

func defaultMigrationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMigrationCallOptions() *MigrationCallOptions {
	return &MigrationCallOptions{
		SearchMigratableResources: []gax.CallOption{},
		BatchMigrateResources:     []gax.CallOption{},
	}
}

// internalMigrationClient is an interface that defines the methods availaible from Vertex AI API.
type internalMigrationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SearchMigratableResources(context.Context, *aiplatformpb.SearchMigratableResourcesRequest, ...gax.CallOption) *MigratableResourceIterator
	BatchMigrateResources(context.Context, *aiplatformpb.BatchMigrateResourcesRequest, ...gax.CallOption) (*BatchMigrateResourcesOperation, error)
	BatchMigrateResourcesOperation(name string) *BatchMigrateResourcesOperation
}

// MigrationClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service that migrates resources from automl.googleapis.com (at http://automl.googleapis.com),
// datalabeling.googleapis.com (at http://datalabeling.googleapis.com) and ml.googleapis.com (at http://ml.googleapis.com) to Vertex AI.
type MigrationClient struct {
	// The internal transport-dependent client.
	internalClient internalMigrationClient

	// The call options for this service.
	CallOptions *MigrationCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MigrationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MigrationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *MigrationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SearchMigratableResources searches all of the resources in automl.googleapis.com (at http://automl.googleapis.com),
// datalabeling.googleapis.com (at http://datalabeling.googleapis.com) and ml.googleapis.com (at http://ml.googleapis.com) that can be migrated to
// Vertex AI’s given location.
func (c *MigrationClient) SearchMigratableResources(ctx context.Context, req *aiplatformpb.SearchMigratableResourcesRequest, opts ...gax.CallOption) *MigratableResourceIterator {
	return c.internalClient.SearchMigratableResources(ctx, req, opts...)
}

// BatchMigrateResources batch migrates resources from ml.googleapis.com (at http://ml.googleapis.com), automl.googleapis.com (at http://automl.googleapis.com),
// and datalabeling.googleapis.com (at http://datalabeling.googleapis.com) to Vertex AI.
func (c *MigrationClient) BatchMigrateResources(ctx context.Context, req *aiplatformpb.BatchMigrateResourcesRequest, opts ...gax.CallOption) (*BatchMigrateResourcesOperation, error) {
	return c.internalClient.BatchMigrateResources(ctx, req, opts...)
}

// BatchMigrateResourcesOperation returns a new BatchMigrateResourcesOperation from a given name.
// The name must be that of a previously created BatchMigrateResourcesOperation, possibly from a different process.
func (c *MigrationClient) BatchMigrateResourcesOperation(name string) *BatchMigrateResourcesOperation {
	return c.internalClient.BatchMigrateResourcesOperation(name)
}

// migrationGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type migrationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing MigrationClient
	CallOptions **MigrationCallOptions

	// The gRPC API client.
	migrationClient aiplatformpb.MigrationServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMigrationClient creates a new migration service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service that migrates resources from automl.googleapis.com (at http://automl.googleapis.com),
// datalabeling.googleapis.com (at http://datalabeling.googleapis.com) and ml.googleapis.com (at http://ml.googleapis.com) to Vertex AI.
func NewMigrationClient(ctx context.Context, opts ...option.ClientOption) (*MigrationClient, error) {
	clientOpts := defaultMigrationGRPCClientOptions()
	if newMigrationClientHook != nil {
		hookOpts, err := newMigrationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MigrationClient{CallOptions: defaultMigrationCallOptions()}

	c := &migrationGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		migrationClient:  aiplatformpb.NewMigrationServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *migrationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *migrationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *migrationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *migrationGRPCClient) SearchMigratableResources(ctx context.Context, req *aiplatformpb.SearchMigratableResourcesRequest, opts ...gax.CallOption) *MigratableResourceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SearchMigratableResources[0:len((*c.CallOptions).SearchMigratableResources):len((*c.CallOptions).SearchMigratableResources)], opts...)
	it := &MigratableResourceIterator{}
	req = proto.Clone(req).(*aiplatformpb.SearchMigratableResourcesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.MigratableResource, string, error) {
		resp := &aiplatformpb.SearchMigratableResourcesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.migrationClient.SearchMigratableResources(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMigratableResources(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *migrationGRPCClient) BatchMigrateResources(ctx context.Context, req *aiplatformpb.BatchMigrateResourcesRequest, opts ...gax.CallOption) (*BatchMigrateResourcesOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchMigrateResources[0:len((*c.CallOptions).BatchMigrateResources):len((*c.CallOptions).BatchMigrateResources)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.migrationClient.BatchMigrateResources(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchMigrateResourcesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// BatchMigrateResourcesOperation manages a long-running operation from BatchMigrateResources.
type BatchMigrateResourcesOperation struct {
	lro *longrunning.Operation
}

// BatchMigrateResourcesOperation returns a new BatchMigrateResourcesOperation from a given name.
// The name must be that of a previously created BatchMigrateResourcesOperation, possibly from a different process.
func (c *migrationGRPCClient) BatchMigrateResourcesOperation(name string) *BatchMigrateResourcesOperation {
	return &BatchMigrateResourcesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchMigrateResourcesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.BatchMigrateResourcesResponse, error) {
	var resp aiplatformpb.BatchMigrateResourcesResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *BatchMigrateResourcesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.BatchMigrateResourcesResponse, error) {
	var resp aiplatformpb.BatchMigrateResourcesResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *BatchMigrateResourcesOperation) Metadata() (*aiplatformpb.BatchMigrateResourcesOperationMetadata, error) {
	var meta aiplatformpb.BatchMigrateResourcesOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchMigrateResourcesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchMigrateResourcesOperation) Name() string {
	return op.lro.Name()
}

// MigratableResourceIterator manages a stream of *aiplatformpb.MigratableResource.
type MigratableResourceIterator struct {
	items    []*aiplatformpb.MigratableResource
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.MigratableResource, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *MigratableResourceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *MigratableResourceIterator) Next() (*aiplatformpb.MigratableResource, error) {
	var item *aiplatformpb.MigratableResource
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *MigratableResourceIterator) bufLen() int {
	return len(it.items)
}

func (it *MigratableResourceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}