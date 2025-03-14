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

All URIs are relative to *https://conversations.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigurationApi* | [**FetchConfiguration**](docs/ConfigurationApi.md#fetchconfiguration) | **Get** /v1/Configuration | 
*ConfigurationApi* | [**UpdateConfiguration**](docs/ConfigurationApi.md#updateconfiguration) | **Post** /v1/Configuration | 
*ConfigurationAddressesApi* | [**CreateConfigurationAddress**](docs/ConfigurationAddressesApi.md#createconfigurationaddress) | **Post** /v1/Configuration/Addresses | 
*ConfigurationAddressesApi* | [**DeleteConfigurationAddress**](docs/ConfigurationAddressesApi.md#deleteconfigurationaddress) | **Delete** /v1/Configuration/Addresses/{Sid} | 
*ConfigurationAddressesApi* | [**FetchConfigurationAddress**](docs/ConfigurationAddressesApi.md#fetchconfigurationaddress) | **Get** /v1/Configuration/Addresses/{Sid} | 
*ConfigurationAddressesApi* | [**ListConfigurationAddress**](docs/ConfigurationAddressesApi.md#listconfigurationaddress) | **Get** /v1/Configuration/Addresses | 
*ConfigurationAddressesApi* | [**UpdateConfigurationAddress**](docs/ConfigurationAddressesApi.md#updateconfigurationaddress) | **Post** /v1/Configuration/Addresses/{Sid} | 
*ConfigurationWebhooksApi* | [**FetchConfigurationWebhook**](docs/ConfigurationWebhooksApi.md#fetchconfigurationwebhook) | **Get** /v1/Configuration/Webhooks | 
*ConfigurationWebhooksApi* | [**UpdateConfigurationWebhook**](docs/ConfigurationWebhooksApi.md#updateconfigurationwebhook) | **Post** /v1/Configuration/Webhooks | 
*ConversationsApi* | [**CreateConversation**](docs/ConversationsApi.md#createconversation) | **Post** /v1/Conversations | 
*ConversationsApi* | [**DeleteConversation**](docs/ConversationsApi.md#deleteconversation) | **Delete** /v1/Conversations/{Sid} | 
*ConversationsApi* | [**FetchConversation**](docs/ConversationsApi.md#fetchconversation) | **Get** /v1/Conversations/{Sid} | 
*ConversationsApi* | [**ListConversation**](docs/ConversationsApi.md#listconversation) | **Get** /v1/Conversations | 
*ConversationsApi* | [**UpdateConversation**](docs/ConversationsApi.md#updateconversation) | **Post** /v1/Conversations/{Sid} | 
*ConversationsMessagesApi* | [**CreateConversationMessage**](docs/ConversationsMessagesApi.md#createconversationmessage) | **Post** /v1/Conversations/{ConversationSid}/Messages | 
*ConversationsMessagesApi* | [**DeleteConversationMessage**](docs/ConversationsMessagesApi.md#deleteconversationmessage) | **Delete** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
*ConversationsMessagesApi* | [**FetchConversationMessage**](docs/ConversationsMessagesApi.md#fetchconversationmessage) | **Get** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
*ConversationsMessagesApi* | [**ListConversationMessage**](docs/ConversationsMessagesApi.md#listconversationmessage) | **Get** /v1/Conversations/{ConversationSid}/Messages | 
*ConversationsMessagesApi* | [**UpdateConversationMessage**](docs/ConversationsMessagesApi.md#updateconversationmessage) | **Post** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
*ConversationsMessagesReceiptsApi* | [**FetchConversationMessageReceipt**](docs/ConversationsMessagesReceiptsApi.md#fetchconversationmessagereceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
*ConversationsMessagesReceiptsApi* | [**ListConversationMessageReceipt**](docs/ConversationsMessagesReceiptsApi.md#listconversationmessagereceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 
*ConversationsParticipantsApi* | [**CreateConversationParticipant**](docs/ConversationsParticipantsApi.md#createconversationparticipant) | **Post** /v1/Conversations/{ConversationSid}/Participants | 
*ConversationsParticipantsApi* | [**DeleteConversationParticipant**](docs/ConversationsParticipantsApi.md#deleteconversationparticipant) | **Delete** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
*ConversationsParticipantsApi* | [**FetchConversationParticipant**](docs/ConversationsParticipantsApi.md#fetchconversationparticipant) | **Get** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
*ConversationsParticipantsApi* | [**ListConversationParticipant**](docs/ConversationsParticipantsApi.md#listconversationparticipant) | **Get** /v1/Conversations/{ConversationSid}/Participants | 
*ConversationsParticipantsApi* | [**UpdateConversationParticipant**](docs/ConversationsParticipantsApi.md#updateconversationparticipant) | **Post** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
*ConversationsWebhooksApi* | [**CreateConversationScopedWebhook**](docs/ConversationsWebhooksApi.md#createconversationscopedwebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks | 
*ConversationsWebhooksApi* | [**DeleteConversationScopedWebhook**](docs/ConversationsWebhooksApi.md#deleteconversationscopedwebhook) | **Delete** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
*ConversationsWebhooksApi* | [**FetchConversationScopedWebhook**](docs/ConversationsWebhooksApi.md#fetchconversationscopedwebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
*ConversationsWebhooksApi* | [**ListConversationScopedWebhook**](docs/ConversationsWebhooksApi.md#listconversationscopedwebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks | 
*ConversationsWebhooksApi* | [**UpdateConversationScopedWebhook**](docs/ConversationsWebhooksApi.md#updateconversationscopedwebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
*CredentialsApi* | [**CreateCredential**](docs/CredentialsApi.md#createcredential) | **Post** /v1/Credentials | 
*CredentialsApi* | [**DeleteCredential**](docs/CredentialsApi.md#deletecredential) | **Delete** /v1/Credentials/{Sid} | 
*CredentialsApi* | [**FetchCredential**](docs/CredentialsApi.md#fetchcredential) | **Get** /v1/Credentials/{Sid} | 
*CredentialsApi* | [**ListCredential**](docs/CredentialsApi.md#listcredential) | **Get** /v1/Credentials | 
*CredentialsApi* | [**UpdateCredential**](docs/CredentialsApi.md#updatecredential) | **Post** /v1/Credentials/{Sid} | 
*ParticipantConversationsApi* | [**ListParticipantConversation**](docs/ParticipantConversationsApi.md#listparticipantconversation) | **Get** /v1/ParticipantConversations | 
*RolesApi* | [**CreateRole**](docs/RolesApi.md#createrole) | **Post** /v1/Roles | 
*RolesApi* | [**DeleteRole**](docs/RolesApi.md#deleterole) | **Delete** /v1/Roles/{Sid} | 
*RolesApi* | [**FetchRole**](docs/RolesApi.md#fetchrole) | **Get** /v1/Roles/{Sid} | 
*RolesApi* | [**ListRole**](docs/RolesApi.md#listrole) | **Get** /v1/Roles | 
*RolesApi* | [**UpdateRole**](docs/RolesApi.md#updaterole) | **Post** /v1/Roles/{Sid} | 
*ServicesApi* | [**CreateService**](docs/ServicesApi.md#createservice) | **Post** /v1/Services | 
*ServicesApi* | [**DeleteService**](docs/ServicesApi.md#deleteservice) | **Delete** /v1/Services/{Sid} | 
*ServicesApi* | [**FetchService**](docs/ServicesApi.md#fetchservice) | **Get** /v1/Services/{Sid} | 
*ServicesApi* | [**ListService**](docs/ServicesApi.md#listservice) | **Get** /v1/Services | 
*ServicesBindingsApi* | [**DeleteServiceBinding**](docs/ServicesBindingsApi.md#deleteservicebinding) | **Delete** /v1/Services/{ChatServiceSid}/Bindings/{Sid} | 
*ServicesBindingsApi* | [**FetchServiceBinding**](docs/ServicesBindingsApi.md#fetchservicebinding) | **Get** /v1/Services/{ChatServiceSid}/Bindings/{Sid} | 
*ServicesBindingsApi* | [**ListServiceBinding**](docs/ServicesBindingsApi.md#listservicebinding) | **Get** /v1/Services/{ChatServiceSid}/Bindings | 
*ServicesConfigurationApi* | [**FetchServiceConfiguration**](docs/ServicesConfigurationApi.md#fetchserviceconfiguration) | **Get** /v1/Services/{ChatServiceSid}/Configuration | 
*ServicesConfigurationApi* | [**UpdateServiceConfiguration**](docs/ServicesConfigurationApi.md#updateserviceconfiguration) | **Post** /v1/Services/{ChatServiceSid}/Configuration | 
*ServicesConfigurationNotificationsApi* | [**FetchServiceNotification**](docs/ServicesConfigurationNotificationsApi.md#fetchservicenotification) | **Get** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 
*ServicesConfigurationNotificationsApi* | [**UpdateServiceNotification**](docs/ServicesConfigurationNotificationsApi.md#updateservicenotification) | **Post** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 
*ServicesConfigurationWebhooksApi* | [**FetchServiceWebhookConfiguration**](docs/ServicesConfigurationWebhooksApi.md#fetchservicewebhookconfiguration) | **Get** /v1/Services/{ChatServiceSid}/Configuration/Webhooks | 
*ServicesConfigurationWebhooksApi* | [**UpdateServiceWebhookConfiguration**](docs/ServicesConfigurationWebhooksApi.md#updateservicewebhookconfiguration) | **Post** /v1/Services/{ChatServiceSid}/Configuration/Webhooks | 
*ServicesConversationsApi* | [**CreateServiceConversation**](docs/ServicesConversationsApi.md#createserviceconversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations | 
*ServicesConversationsApi* | [**DeleteServiceConversation**](docs/ServicesConversationsApi.md#deleteserviceconversation) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
*ServicesConversationsApi* | [**FetchServiceConversation**](docs/ServicesConversationsApi.md#fetchserviceconversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
*ServicesConversationsApi* | [**ListServiceConversation**](docs/ServicesConversationsApi.md#listserviceconversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations | 
*ServicesConversationsApi* | [**UpdateServiceConversation**](docs/ServicesConversationsApi.md#updateserviceconversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
*ServicesConversationsMessagesApi* | [**CreateServiceConversationMessage**](docs/ServicesConversationsMessagesApi.md#createserviceconversationmessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
*ServicesConversationsMessagesApi* | [**DeleteServiceConversationMessage**](docs/ServicesConversationsMessagesApi.md#deleteserviceconversationmessage) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
*ServicesConversationsMessagesApi* | [**FetchServiceConversationMessage**](docs/ServicesConversationsMessagesApi.md#fetchserviceconversationmessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
*ServicesConversationsMessagesApi* | [**ListServiceConversationMessage**](docs/ServicesConversationsMessagesApi.md#listserviceconversationmessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
*ServicesConversationsMessagesApi* | [**UpdateServiceConversationMessage**](docs/ServicesConversationsMessagesApi.md#updateserviceconversationmessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
*ServicesConversationsMessagesReceiptsApi* | [**FetchServiceConversationMessageReceipt**](docs/ServicesConversationsMessagesReceiptsApi.md#fetchserviceconversationmessagereceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
*ServicesConversationsMessagesReceiptsApi* | [**ListServiceConversationMessageReceipt**](docs/ServicesConversationsMessagesReceiptsApi.md#listserviceconversationmessagereceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 
*ServicesConversationsParticipantsApi* | [**CreateServiceConversationParticipant**](docs/ServicesConversationsParticipantsApi.md#createserviceconversationparticipant) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants | 
*ServicesConversationsParticipantsApi* | [**DeleteServiceConversationParticipant**](docs/ServicesConversationsParticipantsApi.md#deleteserviceconversationparticipant) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
*ServicesConversationsParticipantsApi* | [**FetchServiceConversationParticipant**](docs/ServicesConversationsParticipantsApi.md#fetchserviceconversationparticipant) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
*ServicesConversationsParticipantsApi* | [**ListServiceConversationParticipant**](docs/ServicesConversationsParticipantsApi.md#listserviceconversationparticipant) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants | 
*ServicesConversationsParticipantsApi* | [**UpdateServiceConversationParticipant**](docs/ServicesConversationsParticipantsApi.md#updateserviceconversationparticipant) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
*ServicesConversationsWebhooksApi* | [**CreateServiceConversationScopedWebhook**](docs/ServicesConversationsWebhooksApi.md#createserviceconversationscopedwebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
*ServicesConversationsWebhooksApi* | [**DeleteServiceConversationScopedWebhook**](docs/ServicesConversationsWebhooksApi.md#deleteserviceconversationscopedwebhook) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
*ServicesConversationsWebhooksApi* | [**FetchServiceConversationScopedWebhook**](docs/ServicesConversationsWebhooksApi.md#fetchserviceconversationscopedwebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
*ServicesConversationsWebhooksApi* | [**ListServiceConversationScopedWebhook**](docs/ServicesConversationsWebhooksApi.md#listserviceconversationscopedwebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
*ServicesConversationsWebhooksApi* | [**UpdateServiceConversationScopedWebhook**](docs/ServicesConversationsWebhooksApi.md#updateserviceconversationscopedwebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
*ServicesParticipantConversationsApi* | [**ListServiceParticipantConversation**](docs/ServicesParticipantConversationsApi.md#listserviceparticipantconversation) | **Get** /v1/Services/{ChatServiceSid}/ParticipantConversations | 
*ServicesRolesApi* | [**CreateServiceRole**](docs/ServicesRolesApi.md#createservicerole) | **Post** /v1/Services/{ChatServiceSid}/Roles | 
*ServicesRolesApi* | [**DeleteServiceRole**](docs/ServicesRolesApi.md#deleteservicerole) | **Delete** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
*ServicesRolesApi* | [**FetchServiceRole**](docs/ServicesRolesApi.md#fetchservicerole) | **Get** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
*ServicesRolesApi* | [**ListServiceRole**](docs/ServicesRolesApi.md#listservicerole) | **Get** /v1/Services/{ChatServiceSid}/Roles | 
*ServicesRolesApi* | [**UpdateServiceRole**](docs/ServicesRolesApi.md#updateservicerole) | **Post** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
*ServicesUsersApi* | [**CreateServiceUser**](docs/ServicesUsersApi.md#createserviceuser) | **Post** /v1/Services/{ChatServiceSid}/Users | 
*ServicesUsersApi* | [**DeleteServiceUser**](docs/ServicesUsersApi.md#deleteserviceuser) | **Delete** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
*ServicesUsersApi* | [**FetchServiceUser**](docs/ServicesUsersApi.md#fetchserviceuser) | **Get** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
*ServicesUsersApi* | [**ListServiceUser**](docs/ServicesUsersApi.md#listserviceuser) | **Get** /v1/Services/{ChatServiceSid}/Users | 
*ServicesUsersApi* | [**UpdateServiceUser**](docs/ServicesUsersApi.md#updateserviceuser) | **Post** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
*ServicesUsersConversationsApi* | [**DeleteServiceUserConversation**](docs/ServicesUsersConversationsApi.md#deleteserviceuserconversation) | **Delete** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | 
*ServicesUsersConversationsApi* | [**FetchServiceUserConversation**](docs/ServicesUsersConversationsApi.md#fetchserviceuserconversation) | **Get** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | 
*ServicesUsersConversationsApi* | [**ListServiceUserConversation**](docs/ServicesUsersConversationsApi.md#listserviceuserconversation) | **Get** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations | 
*ServicesUsersConversationsApi* | [**UpdateServiceUserConversation**](docs/ServicesUsersConversationsApi.md#updateserviceuserconversation) | **Post** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | 
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /v1/Users | 
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /v1/Users/{Sid} | 
*UsersApi* | [**FetchUser**](docs/UsersApi.md#fetchuser) | **Get** /v1/Users/{Sid} | 
*UsersApi* | [**ListUser**](docs/UsersApi.md#listuser) | **Get** /v1/Users | 
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Post** /v1/Users/{Sid} | 
*UsersConversationsApi* | [**DeleteUserConversation**](docs/UsersConversationsApi.md#deleteuserconversation) | **Delete** /v1/Users/{UserSid}/Conversations/{ConversationSid} | 
*UsersConversationsApi* | [**FetchUserConversation**](docs/UsersConversationsApi.md#fetchuserconversation) | **Get** /v1/Users/{UserSid}/Conversations/{ConversationSid} | 
*UsersConversationsApi* | [**ListUserConversation**](docs/UsersConversationsApi.md#listuserconversation) | **Get** /v1/Users/{UserSid}/Conversations | 
*UsersConversationsApi* | [**UpdateUserConversation**](docs/UsersConversationsApi.md#updateuserconversation) | **Post** /v1/Users/{UserSid}/Conversations/{ConversationSid} | 


## Documentation For Models

 - [ListUserResponse](docs/ListUserResponse.md)
 - [ConversationsV1ServiceWebhookConfiguration](docs/ConversationsV1ServiceWebhookConfiguration.md)
 - [ListServiceUserConversationResponse](docs/ListServiceUserConversationResponse.md)
 - [ConversationsV1ConversationMessageReceipt](docs/ConversationsV1ConversationMessageReceipt.md)
 - [ListServiceResponse](docs/ListServiceResponse.md)
 - [ListServiceConversationMessageReceiptResponse](docs/ListServiceConversationMessageReceiptResponse.md)
 - [ConversationsV1ServiceParticipantConversation](docs/ConversationsV1ServiceParticipantConversation.md)
 - [ListServiceConversationResponse](docs/ListServiceConversationResponse.md)
 - [ListConversationScopedWebhookResponse](docs/ListConversationScopedWebhookResponse.md)
 - [ConversationsV1Role](docs/ConversationsV1Role.md)
 - [ListServiceUserResponse](docs/ListServiceUserResponse.md)
 - [ConversationsV1ServiceNotification](docs/ConversationsV1ServiceNotification.md)
 - [ListConversationMessageReceiptResponse](docs/ListConversationMessageReceiptResponse.md)
 - [ListConfigurationAddressResponseMeta](docs/ListConfigurationAddressResponseMeta.md)
 - [ListConversationMessageResponse](docs/ListConversationMessageResponse.md)
 - [ConversationsV1ConfigurationWebhook](docs/ConversationsV1ConfigurationWebhook.md)
 - [ConversationsV1Credential](docs/ConversationsV1Credential.md)
 - [ConversationsV1UserConversation](docs/ConversationsV1UserConversation.md)
 - [ConversationsV1Conversation](docs/ConversationsV1Conversation.md)
 - [ConversationsV1ServiceConversationMessageReceipt](docs/ConversationsV1ServiceConversationMessageReceipt.md)
 - [ConversationsV1ServiceRole](docs/ConversationsV1ServiceRole.md)
 - [ConversationsV1User](docs/ConversationsV1User.md)
 - [ConversationsV1ConversationMessage](docs/ConversationsV1ConversationMessage.md)
 - [ConversationsV1ServiceBinding](docs/ConversationsV1ServiceBinding.md)
 - [ConversationsV1Configuration](docs/ConversationsV1Configuration.md)
 - [ConversationsV1ServiceConversationParticipant](docs/ConversationsV1ServiceConversationParticipant.md)
 - [ListParticipantConversationResponse](docs/ListParticipantConversationResponse.md)
 - [ListConversationResponse](docs/ListConversationResponse.md)
 - [ListServiceRoleResponse](docs/ListServiceRoleResponse.md)
 - [ConversationsV1ServiceUser](docs/ConversationsV1ServiceUser.md)
 - [ListConversationParticipantResponse](docs/ListConversationParticipantResponse.md)
 - [ListServiceConversationParticipantResponse](docs/ListServiceConversationParticipantResponse.md)
 - [ListCredentialResponse](docs/ListCredentialResponse.md)
 - [ListServiceBindingResponse](docs/ListServiceBindingResponse.md)
 - [ListRoleResponse](docs/ListRoleResponse.md)
 - [ConversationsV1ConversationScopedWebhook](docs/ConversationsV1ConversationScopedWebhook.md)
 - [ListConfigurationAddressResponse](docs/ListConfigurationAddressResponse.md)
 - [ConversationsV1ConfigurationAddress](docs/ConversationsV1ConfigurationAddress.md)
 - [ConversationsV1ServiceUserConversation](docs/ConversationsV1ServiceUserConversation.md)
 - [ConversationsV1ServiceConversationScopedWebhook](docs/ConversationsV1ServiceConversationScopedWebhook.md)
 - [ConversationsV1ServiceConversationMessage](docs/ConversationsV1ServiceConversationMessage.md)
 - [ListServiceConversationScopedWebhookResponse](docs/ListServiceConversationScopedWebhookResponse.md)
 - [ConversationsV1Service](docs/ConversationsV1Service.md)
 - [ConversationsV1ConversationParticipant](docs/ConversationsV1ConversationParticipant.md)
 - [ConversationsV1ServiceConversation](docs/ConversationsV1ServiceConversation.md)
 - [ListServiceConversationMessageResponse](docs/ListServiceConversationMessageResponse.md)
 - [ListUserConversationResponse](docs/ListUserConversationResponse.md)
 - [ConversationsV1ParticipantConversation](docs/ConversationsV1ParticipantConversation.md)
 - [ConversationsV1ServiceConfiguration](docs/ConversationsV1ServiceConfiguration.md)
 - [ListServiceParticipantConversationResponse](docs/ListServiceParticipantConversationResponse.md)


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

