# MetadataWithSupportedRegions

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
|**Status** | **string** | The status of the object. The status can be: * &#x60;AVAILABLE&#x60; - resource exists and is healthy. * &#x60;PROVISIONING&#x60; - resource is being created or updated. * &#x60;DESTROYING&#x60; - delete command was issued, the resource is being deleted. * &#x60;FAILED&#x60; - resource failed, details in &#x60;failureMessage&#x60;.  | [readonly] |
|**StatusMessage** | Pointer to **string** | The message of the failure if the status is &#x60;FAILED&#x60;.  | [optional] [readonly] |
|**Administrative** | Pointer to **bool** | Indicates if the key is an administrative key. Administrative keys can create buckets and set bucket policies.  | [optional] [readonly] |
|**SupportedRegions** | **[]string** | The list of supported regions.  | |

## Methods

### NewMetadataWithSupportedRegions

`func NewMetadataWithSupportedRegions(status string, supportedRegions []string, ) *MetadataWithSupportedRegions`

NewMetadataWithSupportedRegions instantiates a new MetadataWithSupportedRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithSupportedRegionsWithDefaults

`func NewMetadataWithSupportedRegionsWithDefaults() *MetadataWithSupportedRegions`

NewMetadataWithSupportedRegionsWithDefaults instantiates a new MetadataWithSupportedRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *MetadataWithSupportedRegions) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *MetadataWithSupportedRegions) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *MetadataWithSupportedRegions) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *MetadataWithSupportedRegions) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MetadataWithSupportedRegions) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MetadataWithSupportedRegions) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MetadataWithSupportedRegions) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MetadataWithSupportedRegions) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *MetadataWithSupportedRegions) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *MetadataWithSupportedRegions) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *MetadataWithSupportedRegions) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *MetadataWithSupportedRegions) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *MetadataWithSupportedRegions) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *MetadataWithSupportedRegions) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *MetadataWithSupportedRegions) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *MetadataWithSupportedRegions) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *MetadataWithSupportedRegions) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MetadataWithSupportedRegions) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MetadataWithSupportedRegions) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MetadataWithSupportedRegions) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *MetadataWithSupportedRegions) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *MetadataWithSupportedRegions) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *MetadataWithSupportedRegions) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *MetadataWithSupportedRegions) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *MetadataWithSupportedRegions) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *MetadataWithSupportedRegions) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *MetadataWithSupportedRegions) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *MetadataWithSupportedRegions) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetStatus

`func (o *MetadataWithSupportedRegions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetadataWithSupportedRegions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetadataWithSupportedRegions) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *MetadataWithSupportedRegions) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *MetadataWithSupportedRegions) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *MetadataWithSupportedRegions) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *MetadataWithSupportedRegions) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAdministrative

`func (o *MetadataWithSupportedRegions) GetAdministrative() bool`

GetAdministrative returns the Administrative field if non-nil, zero value otherwise.

### GetAdministrativeOk

`func (o *MetadataWithSupportedRegions) GetAdministrativeOk() (*bool, bool)`

GetAdministrativeOk returns a tuple with the Administrative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrative

`func (o *MetadataWithSupportedRegions) SetAdministrative(v bool)`

SetAdministrative sets Administrative field to given value.

### HasAdministrative

`func (o *MetadataWithSupportedRegions) HasAdministrative() bool`

HasAdministrative returns a boolean if a field has been set.

### GetSupportedRegions

`func (o *MetadataWithSupportedRegions) GetSupportedRegions() []string`

GetSupportedRegions returns the SupportedRegions field if non-nil, zero value otherwise.

### GetSupportedRegionsOk

`func (o *MetadataWithSupportedRegions) GetSupportedRegionsOk() (*[]string, bool)`

GetSupportedRegionsOk returns a tuple with the SupportedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRegions

`func (o *MetadataWithSupportedRegions) SetSupportedRegions(v []string)`

SetSupportedRegions sets SupportedRegions field to given value.



