# UserAccessMetadata

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
|**State** | **string** | State of the resource. Resource states: &#x60;AVAILABLE&#x60;: There are no pending modification requests for this item. &#x60;BUSY&#x60;: There is at least one modification request pending and all following requests will be queued. &#x60;DEPLOYING&#x60;: The resource is being created. &#x60;FAILED&#x60;: The creation of the resource failed. &#x60;UPDATING&#x60;: The resource is being updated. &#x60;FAILED_UPDATING&#x60;: An update to the resource was not successful. &#x60;DESTROYING&#x60;: A delete command was issued, and the resource is being deleted.  | |
|**Message** | Pointer to **string** | A human readable message describing the current state. In case of an error, the message will contain a detailed error message.  | [optional] |
|**CertificateAuthority** | Pointer to **string** | PEM for the certificate authority. | [optional] |
|**PrivateKey** | Pointer to **string** | PEM for the private key. | [optional] |
|**Certificate** | Pointer to **string** | PEM for the certificate. | [optional] |

## Methods

### NewUserAccessMetadata

`func NewUserAccessMetadata(state string, ) *UserAccessMetadata`

NewUserAccessMetadata instantiates a new UserAccessMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessMetadataWithDefaults

`func NewUserAccessMetadataWithDefaults() *UserAccessMetadata`

NewUserAccessMetadataWithDefaults instantiates a new UserAccessMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *UserAccessMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *UserAccessMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *UserAccessMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *UserAccessMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UserAccessMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UserAccessMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UserAccessMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UserAccessMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *UserAccessMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *UserAccessMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *UserAccessMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *UserAccessMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *UserAccessMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *UserAccessMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *UserAccessMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *UserAccessMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *UserAccessMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *UserAccessMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *UserAccessMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *UserAccessMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *UserAccessMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *UserAccessMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *UserAccessMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *UserAccessMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *UserAccessMetadata) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *UserAccessMetadata) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *UserAccessMetadata) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *UserAccessMetadata) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetState

`func (o *UserAccessMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserAccessMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserAccessMetadata) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *UserAccessMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserAccessMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserAccessMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserAccessMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCertificateAuthority

`func (o *UserAccessMetadata) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *UserAccessMetadata) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *UserAccessMetadata) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *UserAccessMetadata) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UserAccessMetadata) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UserAccessMetadata) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UserAccessMetadata) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UserAccessMetadata) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetCertificate

`func (o *UserAccessMetadata) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UserAccessMetadata) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UserAccessMetadata) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *UserAccessMetadata) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


