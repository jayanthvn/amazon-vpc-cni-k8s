// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package qldbsession provides the client and types for making API
// requests to Amazon QLDB Session.
//
// The transactional data APIs for Amazon QLDB
//
// Instead of interacting directly with this API, we recommend using the QLDB
// driver or the QLDB shell to execute data transactions on a ledger.
//
//    * If you are working with an AWS SDK, use the QLDB driver. The driver
//    provides a high-level abstraction layer above this QLDB Session data plane
//    and manages SendCommand API calls for you. For information and a list
//    of supported programming languages, see Getting started with the driver
//    (https://docs.aws.amazon.com/qldb/latest/developerguide/getting-started-driver.html)
//    in the Amazon QLDB Developer Guide.
//
//    * If you are working with the AWS Command Line Interface (AWS CLI), use
//    the QLDB shell. The shell is a command line interface that uses the QLDB
//    driver to interact with a ledger. For information, see Accessing Amazon
//    QLDB using the QLDB shell (https://docs.aws.amazon.com/qldb/latest/developerguide/data-shell.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/qldb-session-2019-07-11 for more information on this service.
//
// See qldbsession package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/qldbsession/
//
// Using the Client
//
// To contact Amazon QLDB Session with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon QLDB Session client QLDBSession for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/qldbsession/#New
package qldbsession
