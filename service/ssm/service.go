// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ssm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Amazon EC2 Simple Systems Manager (SSM) enables you to configure and manage
// your EC2 instances. You can create a configuration document and then associate
// it with one or more running instances.
//
// You can use a configuration document to automate the following tasks for
// your Windows instances:
//
//  Join an AWS Directory
//
// Install, repair, or uninstall software using an MSI package
//
// Run PowerShell scripts
//
// Configure CloudWatch Logs to monitor applications and systems
//
//  Note that configuration documents are not supported on Linux instances.
type SSM struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new SSM client.
func New(config *aws.Config) *SSM {
	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "ssm",
		APIVersion:   "2014-11-06",
		JSONVersion:  "1.1",
		TargetPrefix: "AmazonSSM",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &SSM{service}
}

// newRequest creates a new request for a SSM operation and runs any
// custom request initialization.
func (c *SSM) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
