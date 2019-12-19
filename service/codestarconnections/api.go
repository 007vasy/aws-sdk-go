// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarconnections

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

const opCreateConnection = "CreateConnection"

// CreateConnectionRequest generates a "aws/request.Request" representing the
// client's request for the CreateConnection operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateConnection for more information on using the CreateConnection
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateConnectionRequest method.
//    req, resp := client.CreateConnectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/CreateConnection
func (c *CodeStarConnections) CreateConnectionRequest(input *CreateConnectionInput) (req *request.Request, output *CreateConnectionOutput) {
	op := &request.Operation{
		Name:       opCreateConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateConnectionInput{}
	}

	output = &CreateConnectionOutput{}
	req = c.newRequest(op, input, output)
	return
}

// CreateConnection API operation for AWS CodeStar connections.
//
// Creates a connection that can then be given to other AWS services like CodePipeline
// so that it can access third-party code repositories. The connection is in
// pending status until the third-party connection handshake is completed from
// the console.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS CodeStar connections's
// API operation CreateConnection for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeLimitExceededException "LimitExceededException"
//   Exceeded the maximum limit for connections.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/CreateConnection
func (c *CodeStarConnections) CreateConnection(input *CreateConnectionInput) (*CreateConnectionOutput, error) {
	req, out := c.CreateConnectionRequest(input)
	return out, req.Send()
}

// CreateConnectionWithContext is the same as CreateConnection with the addition of
// the ability to pass a context and additional request options.
//
// See CreateConnection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeStarConnections) CreateConnectionWithContext(ctx aws.Context, input *CreateConnectionInput, opts ...request.Option) (*CreateConnectionOutput, error) {
	req, out := c.CreateConnectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteConnection = "DeleteConnection"

// DeleteConnectionRequest generates a "aws/request.Request" representing the
// client's request for the DeleteConnection operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteConnection for more information on using the DeleteConnection
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteConnectionRequest method.
//    req, resp := client.DeleteConnectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/DeleteConnection
func (c *CodeStarConnections) DeleteConnectionRequest(input *DeleteConnectionInput) (req *request.Request, output *DeleteConnectionOutput) {
	op := &request.Operation{
		Name:       opDeleteConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConnectionInput{}
	}

	output = &DeleteConnectionOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(jsonrpc.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// DeleteConnection API operation for AWS CodeStar connections.
//
// The connection to be deleted.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS CodeStar connections's
// API operation DeleteConnection for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   Resource not found. Verify the connection resource ARN and try again.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/DeleteConnection
func (c *CodeStarConnections) DeleteConnection(input *DeleteConnectionInput) (*DeleteConnectionOutput, error) {
	req, out := c.DeleteConnectionRequest(input)
	return out, req.Send()
}

// DeleteConnectionWithContext is the same as DeleteConnection with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteConnection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeStarConnections) DeleteConnectionWithContext(ctx aws.Context, input *DeleteConnectionInput, opts ...request.Option) (*DeleteConnectionOutput, error) {
	req, out := c.DeleteConnectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetConnection = "GetConnection"

// GetConnectionRequest generates a "aws/request.Request" representing the
// client's request for the GetConnection operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetConnection for more information on using the GetConnection
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetConnectionRequest method.
//    req, resp := client.GetConnectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/GetConnection
func (c *CodeStarConnections) GetConnectionRequest(input *GetConnectionInput) (req *request.Request, output *GetConnectionOutput) {
	op := &request.Operation{
		Name:       opGetConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConnectionInput{}
	}

	output = &GetConnectionOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetConnection API operation for AWS CodeStar connections.
//
// Returns the connection ARN and details such as status, owner, and provider
// type.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS CodeStar connections's
// API operation GetConnection for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   Resource not found. Verify the connection resource ARN and try again.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/GetConnection
func (c *CodeStarConnections) GetConnection(input *GetConnectionInput) (*GetConnectionOutput, error) {
	req, out := c.GetConnectionRequest(input)
	return out, req.Send()
}

// GetConnectionWithContext is the same as GetConnection with the addition of
// the ability to pass a context and additional request options.
//
// See GetConnection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeStarConnections) GetConnectionWithContext(ctx aws.Context, input *GetConnectionInput, opts ...request.Option) (*GetConnectionOutput, error) {
	req, out := c.GetConnectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListConnections = "ListConnections"

// ListConnectionsRequest generates a "aws/request.Request" representing the
// client's request for the ListConnections operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListConnections for more information on using the ListConnections
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListConnectionsRequest method.
//    req, resp := client.ListConnectionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/ListConnections
func (c *CodeStarConnections) ListConnectionsRequest(input *ListConnectionsInput) (req *request.Request, output *ListConnectionsOutput) {
	op := &request.Operation{
		Name:       opListConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListConnectionsInput{}
	}

	output = &ListConnectionsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListConnections API operation for AWS CodeStar connections.
//
// Lists the connections associated with your account.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS CodeStar connections's
// API operation ListConnections for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/ListConnections
func (c *CodeStarConnections) ListConnections(input *ListConnectionsInput) (*ListConnectionsOutput, error) {
	req, out := c.ListConnectionsRequest(input)
	return out, req.Send()
}

// ListConnectionsWithContext is the same as ListConnections with the addition of
// the ability to pass a context and additional request options.
//
// See ListConnections for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeStarConnections) ListConnectionsWithContext(ctx aws.Context, input *ListConnectionsInput, opts ...request.Option) (*ListConnectionsOutput, error) {
	req, out := c.ListConnectionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListConnectionsPages iterates over the pages of a ListConnections operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListConnections method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListConnections operation.
//    pageNum := 0
//    err := client.ListConnectionsPages(params,
//        func(page *codestarconnections.ListConnectionsOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *CodeStarConnections) ListConnectionsPages(input *ListConnectionsInput, fn func(*ListConnectionsOutput, bool) bool) error {
	return c.ListConnectionsPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListConnectionsPagesWithContext same as ListConnectionsPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CodeStarConnections) ListConnectionsPagesWithContext(ctx aws.Context, input *ListConnectionsInput, fn func(*ListConnectionsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListConnectionsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListConnectionsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*ListConnectionsOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

// The configuration that allows a service such as CodePipeline to connect to
// a third-party code repository.
type Connection struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the connection. The ARN is used as the
	// connection reference when the connection is shared between AWS services.
	//
	// The ARN is never reused if the connection is deleted.
	ConnectionArn *string `type:"string"`

	// The name of the connection. Connection names must be unique in an AWS user
	// account.
	ConnectionName *string `min:"1" type:"string"`

	// The current status of the connection.
	ConnectionStatus *string `type:"string" enum:"ConnectionStatus"`

	// The name of the external provider where your third-party code repository
	// is configured. For Bitbucket, this is the account ID of the owner of the
	// Bitbucket repository.
	OwnerAccountId *string `min:"12" type:"string"`

	// The name of the external provider where your third-party code repository
	// is configured. Currently, the valid provider type is Bitbucket.
	ProviderType *string `type:"string" enum:"ProviderType"`
}

// String returns the string representation
func (s Connection) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Connection) GoString() string {
	return s.String()
}

// SetConnectionArn sets the ConnectionArn field's value.
func (s *Connection) SetConnectionArn(v string) *Connection {
	s.ConnectionArn = &v
	return s
}

// SetConnectionName sets the ConnectionName field's value.
func (s *Connection) SetConnectionName(v string) *Connection {
	s.ConnectionName = &v
	return s
}

// SetConnectionStatus sets the ConnectionStatus field's value.
func (s *Connection) SetConnectionStatus(v string) *Connection {
	s.ConnectionStatus = &v
	return s
}

// SetOwnerAccountId sets the OwnerAccountId field's value.
func (s *Connection) SetOwnerAccountId(v string) *Connection {
	s.OwnerAccountId = &v
	return s
}

// SetProviderType sets the ProviderType field's value.
func (s *Connection) SetProviderType(v string) *Connection {
	s.ProviderType = &v
	return s
}

type CreateConnectionInput struct {
	_ struct{} `type:"structure"`

	// The name of the connection to be created. The name must be unique in the
	// calling AWS account.
	//
	// ConnectionName is a required field
	ConnectionName *string `min:"1" type:"string" required:"true"`

	// The name of the external provider where your third-party code repository
	// is configured. Currently, the valid provider type is Bitbucket.
	//
	// ProviderType is a required field
	ProviderType *string `type:"string" required:"true" enum:"ProviderType"`
}

// String returns the string representation
func (s CreateConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateConnectionInput"}
	if s.ConnectionName == nil {
		invalidParams.Add(request.NewErrParamRequired("ConnectionName"))
	}
	if s.ConnectionName != nil && len(*s.ConnectionName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ConnectionName", 1))
	}
	if s.ProviderType == nil {
		invalidParams.Add(request.NewErrParamRequired("ProviderType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetConnectionName sets the ConnectionName field's value.
func (s *CreateConnectionInput) SetConnectionName(v string) *CreateConnectionInput {
	s.ConnectionName = &v
	return s
}

// SetProviderType sets the ProviderType field's value.
func (s *CreateConnectionInput) SetProviderType(v string) *CreateConnectionInput {
	s.ProviderType = &v
	return s
}

type CreateConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the connection to be created. The ARN is
	// used as the connection reference when the connection is shared between AWS
	// services.
	//
	// The ARN is never reused if the connection is deleted.
	//
	// ConnectionArn is a required field
	ConnectionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateConnectionOutput) GoString() string {
	return s.String()
}

// SetConnectionArn sets the ConnectionArn field's value.
func (s *CreateConnectionOutput) SetConnectionArn(v string) *CreateConnectionOutput {
	s.ConnectionArn = &v
	return s
}

type DeleteConnectionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the connection to be deleted.
	//
	// The ARN is never reused if the connection is deleted.
	//
	// ConnectionArn is a required field
	ConnectionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteConnectionInput"}
	if s.ConnectionArn == nil {
		invalidParams.Add(request.NewErrParamRequired("ConnectionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetConnectionArn sets the ConnectionArn field's value.
func (s *DeleteConnectionInput) SetConnectionArn(v string) *DeleteConnectionInput {
	s.ConnectionArn = &v
	return s
}

type DeleteConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteConnectionOutput) GoString() string {
	return s.String()
}

type GetConnectionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of a connection.
	//
	// ConnectionArn is a required field
	ConnectionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetConnectionInput"}
	if s.ConnectionArn == nil {
		invalidParams.Add(request.NewErrParamRequired("ConnectionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetConnectionArn sets the ConnectionArn field's value.
func (s *GetConnectionInput) SetConnectionArn(v string) *GetConnectionInput {
	s.ConnectionArn = &v
	return s
}

type GetConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The connection details, such as status, owner, and provider type.
	Connection *Connection `type:"structure"`
}

// String returns the string representation
func (s GetConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConnectionOutput) GoString() string {
	return s.String()
}

// SetConnection sets the Connection field's value.
func (s *GetConnectionOutput) SetConnection(v *Connection) *GetConnectionOutput {
	s.Connection = v
	return s
}

type ListConnectionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token that was returned from the previous ListConnections call, which
	// can be used to return the next set of connections in the list.
	NextToken *string `min:"1" type:"string"`

	// Filters the list of connections to those associated with a specified provider,
	// such as Bitbucket.
	ProviderTypeFilter *string `type:"string" enum:"ProviderType"`
}

// String returns the string representation
func (s ListConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListConnectionsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListConnectionsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListConnectionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListConnectionsInput) SetMaxResults(v int64) *ListConnectionsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListConnectionsInput) SetNextToken(v string) *ListConnectionsInput {
	s.NextToken = &v
	return s
}

// SetProviderTypeFilter sets the ProviderTypeFilter field's value.
func (s *ListConnectionsInput) SetProviderTypeFilter(v string) *ListConnectionsInput {
	s.ProviderTypeFilter = &v
	return s
}

type ListConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of connections and the details for each connection, such as status,
	// owner, and provider type.
	Connections []*Connection `type:"list"`

	// A token that can be used in the next ListConnections call. To view all items
	// in the list, continue to call this operation with each subsequent token until
	// no more nextToken values are returned.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListConnectionsOutput) GoString() string {
	return s.String()
}

// SetConnections sets the Connections field's value.
func (s *ListConnectionsOutput) SetConnections(v []*Connection) *ListConnectionsOutput {
	s.Connections = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListConnectionsOutput) SetNextToken(v string) *ListConnectionsOutput {
	s.NextToken = &v
	return s
}

const (
	// ConnectionStatusPending is a ConnectionStatus enum value
	ConnectionStatusPending = "PENDING"

	// ConnectionStatusAvailable is a ConnectionStatus enum value
	ConnectionStatusAvailable = "AVAILABLE"

	// ConnectionStatusError is a ConnectionStatus enum value
	ConnectionStatusError = "ERROR"
)

const (
	// ProviderTypeBitbucket is a ProviderType enum value
	ProviderTypeBitbucket = "Bitbucket"
)
