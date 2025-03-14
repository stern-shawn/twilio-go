# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.40.0
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://content.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContentApi* | [**DeleteContent**](docs/ContentApi.md#deletecontent) | **Delete** /v1/Content/{Sid} | 
*ContentApi* | [**FetchContent**](docs/ContentApi.md#fetchcontent) | **Get** /v1/Content/{Sid} | 
*ContentApi* | [**ListContent**](docs/ContentApi.md#listcontent) | **Get** /v1/Content | 
*ContentAndApprovalsApi* | [**ListContentAndApprovals**](docs/ContentAndApprovalsApi.md#listcontentandapprovals) | **Get** /v1/ContentAndApprovals | 
*ContentApprovalRequestsApi* | [**FetchApprovalFetch**](docs/ContentApprovalRequestsApi.md#fetchapprovalfetch) | **Get** /v1/Content/{Sid}/ApprovalRequests | 
*LegacyContentApi* | [**ListLegacyContent**](docs/LegacyContentApi.md#listlegacycontent) | **Get** /v1/LegacyContent | 


## Documentation For Models

 - [ListLegacyContentResponse](docs/ListLegacyContentResponse.md)
 - [ContentV1ContentAndApprovals](docs/ContentV1ContentAndApprovals.md)
 - [ContentV1ApprovalFetch](docs/ContentV1ApprovalFetch.md)
 - [ContentV1Content](docs/ContentV1Content.md)
 - [ListContentResponseMeta](docs/ListContentResponseMeta.md)
 - [ListContentResponse](docs/ListContentResponse.md)
 - [ContentV1LegacyContent](docs/ContentV1LegacyContent.md)
 - [ListContentAndApprovalsResponse](docs/ListContentAndApprovalsResponse.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

