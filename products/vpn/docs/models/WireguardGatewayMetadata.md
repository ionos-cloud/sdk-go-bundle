# WireguardGatewayMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 creation timestamp. | [optional] [readonly] |
|**CreatedBy** | Pointer to **string** | Unique name of the identity that created the resource. | [optional] [readonly] |
|**CreatedByUserId** | Pointer to **string** | Unique id of the identity that created the resource. | [optional] [readonly] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 modified timestamp. | [optional] [readonly] |
|**LastModifiedBy** | Pointer to **string** | Unique name of the identity that last modified the resource. | [optional] [readonly] |
|**LastModifiedByUserId** | Pointer to **string** | Unique id of the identity that last modified the resource. | [optional] [readonly] |
|**ResourceURN** | Pointer to **string** | Unique name of the resource. | [optional] [readonly] |
|**Status** | **string** | The current status of the resource. The status can be:  * &#x60;AVAILABLE&#x60; - resource exists and is healthy. * &#x60;PROVISIONING&#x60; - resource is being created or updated. * &#x60;DESTROYING&#x60; - delete command was issued, the resource is being deleted. * &#x60;FAILED&#x60;: - resource failed, details in &#x60;statusMessage&#x60;.  | [readonly] |
|**StatusMessage** | Pointer to **string** | The message of the failure if the status is &#x60;FAILED&#x60;.  | [optional] [readonly] |
|**PublicKey** | **string** | Public key correspondng to the WireGuard Server private key.  | |

## Methods

### NewWireguardGatewayMetadata

`func NewWireguardGatewayMetadata(status string, publicKey string, ) *WireguardGatewayMetadata`

NewWireguardGatewayMetadata instantiates a new WireguardGatewayMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardGatewayMetadataWithDefaults

`func NewWireguardGatewayMetadataWithDefaults() *WireguardGatewayMetadata`

NewWireguardGatewayMetadataWithDefaults instantiates a new WireguardGatewayMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *WireguardGatewayMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WireguardGatewayMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WireguardGatewayMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WireguardGatewayMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WireguardGatewayMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WireguardGatewayMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WireguardGatewayMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WireguardGatewayMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *WireguardGatewayMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *WireguardGatewayMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *WireguardGatewayMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *WireguardGatewayMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *WireguardGatewayMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *WireguardGatewayMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *WireguardGatewayMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *WireguardGatewayMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *WireguardGatewayMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *WireguardGatewayMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *WireguardGatewayMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *WireguardGatewayMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *WireguardGatewayMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *WireguardGatewayMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *WireguardGatewayMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *WireguardGatewayMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *WireguardGatewayMetadata) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *WireguardGatewayMetadata) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *WireguardGatewayMetadata) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *WireguardGatewayMetadata) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetStatus

`func (o *WireguardGatewayMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WireguardGatewayMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WireguardGatewayMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *WireguardGatewayMetadata) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *WireguardGatewayMetadata) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *WireguardGatewayMetadata) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *WireguardGatewayMetadata) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardGatewayMetadata) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardGatewayMetadata) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardGatewayMetadata) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



