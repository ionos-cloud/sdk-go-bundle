# Provider

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the certificate provider. | |
|**Email** | **string** | The email address of the certificate requester. | |
|**Server** | **string** | The URL of the certificate provider. | |
|**ExternalAccountBinding** | Pointer to [**ProviderExternalAccountBinding**](ProviderExternalAccountBinding.md) |  | [optional] |

## Methods

### NewProvider

`func NewProvider(name string, email string, server string, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *Provider) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Provider) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Provider) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetServer

`func (o *Provider) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Provider) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Provider) SetServer(v string)`

SetServer sets Server field to given value.


### GetExternalAccountBinding

`func (o *Provider) GetExternalAccountBinding() ProviderExternalAccountBinding`

GetExternalAccountBinding returns the ExternalAccountBinding field if non-nil, zero value otherwise.

### GetExternalAccountBindingOk

`func (o *Provider) GetExternalAccountBindingOk() (*ProviderExternalAccountBinding, bool)`

GetExternalAccountBindingOk returns a tuple with the ExternalAccountBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountBinding

`func (o *Provider) SetExternalAccountBinding(v ProviderExternalAccountBinding)`

SetExternalAccountBinding sets ExternalAccountBinding field to given value.

### HasExternalAccountBinding

`func (o *Provider) HasExternalAccountBinding() bool`

HasExternalAccountBinding returns a boolean if a field has been set.


