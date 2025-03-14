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

All URIs are relative to *https://events.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SchemasApi* | [**FetchSchema**](docs/SchemasApi.md#fetchschema) | **Get** /v1/Schemas/{Id} | 
*SchemasVersionsApi* | [**FetchSchemaVersion**](docs/SchemasVersionsApi.md#fetchschemaversion) | **Get** /v1/Schemas/{Id}/Versions/{SchemaVersion} | 
*SchemasVersionsApi* | [**ListSchemaVersion**](docs/SchemasVersionsApi.md#listschemaversion) | **Get** /v1/Schemas/{Id}/Versions | 
*SinksApi* | [**CreateSink**](docs/SinksApi.md#createsink) | **Post** /v1/Sinks | 
*SinksApi* | [**DeleteSink**](docs/SinksApi.md#deletesink) | **Delete** /v1/Sinks/{Sid} | 
*SinksApi* | [**FetchSink**](docs/SinksApi.md#fetchsink) | **Get** /v1/Sinks/{Sid} | 
*SinksApi* | [**ListSink**](docs/SinksApi.md#listsink) | **Get** /v1/Sinks | 
*SinksApi* | [**UpdateSink**](docs/SinksApi.md#updatesink) | **Post** /v1/Sinks/{Sid} | 
*SinksTestApi* | [**CreateSinkTest**](docs/SinksTestApi.md#createsinktest) | **Post** /v1/Sinks/{Sid}/Test | 
*SinksValidateApi* | [**CreateSinkValidate**](docs/SinksValidateApi.md#createsinkvalidate) | **Post** /v1/Sinks/{Sid}/Validate | 
*SubscriptionsApi* | [**CreateSubscription**](docs/SubscriptionsApi.md#createsubscription) | **Post** /v1/Subscriptions | 
*SubscriptionsApi* | [**DeleteSubscription**](docs/SubscriptionsApi.md#deletesubscription) | **Delete** /v1/Subscriptions/{Sid} | 
*SubscriptionsApi* | [**FetchSubscription**](docs/SubscriptionsApi.md#fetchsubscription) | **Get** /v1/Subscriptions/{Sid} | 
*SubscriptionsApi* | [**ListSubscription**](docs/SubscriptionsApi.md#listsubscription) | **Get** /v1/Subscriptions | 
*SubscriptionsApi* | [**UpdateSubscription**](docs/SubscriptionsApi.md#updatesubscription) | **Post** /v1/Subscriptions/{Sid} | 
*SubscriptionsSubscribedEventsApi* | [**CreateSubscribedEvent**](docs/SubscriptionsSubscribedEventsApi.md#createsubscribedevent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
*SubscriptionsSubscribedEventsApi* | [**DeleteSubscribedEvent**](docs/SubscriptionsSubscribedEventsApi.md#deletesubscribedevent) | **Delete** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
*SubscriptionsSubscribedEventsApi* | [**FetchSubscribedEvent**](docs/SubscriptionsSubscribedEventsApi.md#fetchsubscribedevent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
*SubscriptionsSubscribedEventsApi* | [**ListSubscribedEvent**](docs/SubscriptionsSubscribedEventsApi.md#listsubscribedevent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
*SubscriptionsSubscribedEventsApi* | [**UpdateSubscribedEvent**](docs/SubscriptionsSubscribedEventsApi.md#updatesubscribedevent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
*TypesApi* | [**FetchEventType**](docs/TypesApi.md#fetcheventtype) | **Get** /v1/Types/{Type} | 
*TypesApi* | [**ListEventType**](docs/TypesApi.md#listeventtype) | **Get** /v1/Types | 


## Documentation For Models

 - [EventsV1SinkValidate](docs/EventsV1SinkValidate.md)
 - [EventsV1EventType](docs/EventsV1EventType.md)
 - [EventsV1SchemaVersion](docs/EventsV1SchemaVersion.md)
 - [EventsV1Sink](docs/EventsV1Sink.md)
 - [ListSubscriptionResponse](docs/ListSubscriptionResponse.md)
 - [EventsV1Schema](docs/EventsV1Schema.md)
 - [EventsV1Subscription](docs/EventsV1Subscription.md)
 - [ListEventTypeResponse](docs/ListEventTypeResponse.md)
 - [ListSubscribedEventResponse](docs/ListSubscribedEventResponse.md)
 - [EventsV1SubscribedEvent](docs/EventsV1SubscribedEvent.md)
 - [ListSchemaVersionResponse](docs/ListSchemaVersionResponse.md)
 - [EventsV1SinkTest](docs/EventsV1SinkTest.md)
 - [ListEventTypeResponseMeta](docs/ListEventTypeResponseMeta.md)
 - [ListSinkResponse](docs/ListSinkResponse.md)


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

