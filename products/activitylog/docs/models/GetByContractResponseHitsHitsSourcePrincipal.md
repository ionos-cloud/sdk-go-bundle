# GetByContractResponseHitsHitsSourcePrincipal

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SourceIP** | Pointer to **string** | The remote IP address which triggered the activity. | [optional] |
|**Identity** | Pointer to [**GetByContractResponseHitsHitsSourcePrincipalIdentity**](GetByContractResponseHitsHitsSourcePrincipalIdentity.md) |  | [optional] |
|**SourceService** | Pointer to **string** | The principal triggered the activity using this IONOS service.  Example values are  * &#x60;PUBLIC_REST_V6&#x60; for the CloudAPI V6, * &#x60;Reseller_V1&#x60; for the Reseller API V1 or * &#x60;DCD&#x60; for the Data Center Designer.  | [optional] |
|**ServiceHost** | Pointer to **string** | The FQDN of the hostname on which the service runs.  | [optional] |

## Methods

### NewGetByContractResponseHitsHitsSourcePrincipal

`func NewGetByContractResponseHitsHitsSourcePrincipal() *GetByContractResponseHitsHitsSourcePrincipal`

NewGetByContractResponseHitsHitsSourcePrincipal instantiates a new GetByContractResponseHitsHitsSourcePrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByContractResponseHitsHitsSourcePrincipalWithDefaults

`func NewGetByContractResponseHitsHitsSourcePrincipalWithDefaults() *GetByContractResponseHitsHitsSourcePrincipal`

NewGetByContractResponseHitsHitsSourcePrincipalWithDefaults instantiates a new GetByContractResponseHitsHitsSourcePrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceIP

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetSourceIP() string`

GetSourceIP returns the SourceIP field if non-nil, zero value otherwise.

### GetSourceIPOk

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetSourceIPOk() (*string, bool)`

GetSourceIPOk returns a tuple with the SourceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIP

`func (o *GetByContractResponseHitsHitsSourcePrincipal) SetSourceIP(v string)`

SetSourceIP sets SourceIP field to given value.

### HasSourceIP

`func (o *GetByContractResponseHitsHitsSourcePrincipal) HasSourceIP() bool`

HasSourceIP returns a boolean if a field has been set.

### GetIdentity

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetIdentity() GetByContractResponseHitsHitsSourcePrincipalIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetIdentityOk() (*GetByContractResponseHitsHitsSourcePrincipalIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *GetByContractResponseHitsHitsSourcePrincipal) SetIdentity(v GetByContractResponseHitsHitsSourcePrincipalIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *GetByContractResponseHitsHitsSourcePrincipal) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetSourceService

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetSourceService() string`

GetSourceService returns the SourceService field if non-nil, zero value otherwise.

### GetSourceServiceOk

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetSourceServiceOk() (*string, bool)`

GetSourceServiceOk returns a tuple with the SourceService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceService

`func (o *GetByContractResponseHitsHitsSourcePrincipal) SetSourceService(v string)`

SetSourceService sets SourceService field to given value.

### HasSourceService

`func (o *GetByContractResponseHitsHitsSourcePrincipal) HasSourceService() bool`

HasSourceService returns a boolean if a field has been set.

### GetServiceHost

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *GetByContractResponseHitsHitsSourcePrincipal) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *GetByContractResponseHitsHitsSourcePrincipal) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *GetByContractResponseHitsHitsSourcePrincipal) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.


