# Metadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ETag** | Pointer to **string** | The Entity Tag of the resource as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11 | [optional] |
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The time the resource was created, ISO 8601 timestamp (UTC). | [optional] |
|**CreatedBy** | Pointer to **string** | The user that created the resource | [optional] |
|**CreatedByUserId** | Pointer to **string** | The ID of the user that created the resource | [optional] |
|**CreatedInContractNumber** | Pointer to **string** | The creators contractNumber | [optional] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The last time the resource was modified, ISO 8601 timestamp (UTC). | [optional] |
|**LastModifiedBy** | Pointer to **string** | The user that last modified the resource | [optional] |
|**LastModifiedByUserId** | Pointer to **string** | The ID of the user that last modified the resource | [optional] |
|**CurrentDataPlatformVersion** | Pointer to **string** | The version of the data platform.  | [optional] |
|**CurrentDataPlatformRevision** | Pointer to **int64** | The current data platform revision of a resource. This internal revision is used to roll out non-breaking internal changes. This attribute is read-only.  | [optional] |
|**AvailableUpgradeVersions** | Pointer to **[]string** | List of available upgrades for this cluster | [optional] |
|**State** | Pointer to **string** | State of the resource. *AVAILABLE* There are no pending modification requests for this item; *BUSY* There is at least one modification request pending and all following requests will be queued; *DEPLOYING* Resource state DEPLOYING - the resource is being created; *FAILED* Resource state FAILED - creation of the resource failed; *UPDATING* Resource state UPDATING - the resource is being updated; *FAILED_UPDATING* Resource state FAILED_UPDATING - an update to the resource was not successful; *DESTROYING* Resource state DESTROYING - the resource is being deleted; *FAILED_DESTROYING* Resource state FAILED_DESTROYING - deletion of the resource was not successful; *TERMINATED* Resource state TERMINATED - the resource was deleted.  | [optional] |

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetETag

`func (o *Metadata) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *Metadata) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *Metadata) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *Metadata) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetCreatedDate

`func (o *Metadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Metadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Metadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Metadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Metadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Metadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Metadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Metadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *Metadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *Metadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *Metadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *Metadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedInContractNumber

`func (o *Metadata) GetCreatedInContractNumber() string`

GetCreatedInContractNumber returns the CreatedInContractNumber field if non-nil, zero value otherwise.

### GetCreatedInContractNumberOk

`func (o *Metadata) GetCreatedInContractNumberOk() (*string, bool)`

GetCreatedInContractNumberOk returns a tuple with the CreatedInContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedInContractNumber

`func (o *Metadata) SetCreatedInContractNumber(v string)`

SetCreatedInContractNumber sets CreatedInContractNumber field to given value.

### HasCreatedInContractNumber

`func (o *Metadata) HasCreatedInContractNumber() bool`

HasCreatedInContractNumber returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *Metadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *Metadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *Metadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *Metadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *Metadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *Metadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *Metadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *Metadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *Metadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *Metadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *Metadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *Metadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetCurrentDataPlatformVersion

`func (o *Metadata) GetCurrentDataPlatformVersion() string`

GetCurrentDataPlatformVersion returns the CurrentDataPlatformVersion field if non-nil, zero value otherwise.

### GetCurrentDataPlatformVersionOk

`func (o *Metadata) GetCurrentDataPlatformVersionOk() (*string, bool)`

GetCurrentDataPlatformVersionOk returns a tuple with the CurrentDataPlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDataPlatformVersion

`func (o *Metadata) SetCurrentDataPlatformVersion(v string)`

SetCurrentDataPlatformVersion sets CurrentDataPlatformVersion field to given value.

### HasCurrentDataPlatformVersion

`func (o *Metadata) HasCurrentDataPlatformVersion() bool`

HasCurrentDataPlatformVersion returns a boolean if a field has been set.

### GetCurrentDataPlatformRevision

`func (o *Metadata) GetCurrentDataPlatformRevision() int64`

GetCurrentDataPlatformRevision returns the CurrentDataPlatformRevision field if non-nil, zero value otherwise.

### GetCurrentDataPlatformRevisionOk

`func (o *Metadata) GetCurrentDataPlatformRevisionOk() (*int64, bool)`

GetCurrentDataPlatformRevisionOk returns a tuple with the CurrentDataPlatformRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDataPlatformRevision

`func (o *Metadata) SetCurrentDataPlatformRevision(v int64)`

SetCurrentDataPlatformRevision sets CurrentDataPlatformRevision field to given value.

### HasCurrentDataPlatformRevision

`func (o *Metadata) HasCurrentDataPlatformRevision() bool`

HasCurrentDataPlatformRevision returns a boolean if a field has been set.

### GetAvailableUpgradeVersions

`func (o *Metadata) GetAvailableUpgradeVersions() []string`

GetAvailableUpgradeVersions returns the AvailableUpgradeVersions field if non-nil, zero value otherwise.

### GetAvailableUpgradeVersionsOk

`func (o *Metadata) GetAvailableUpgradeVersionsOk() (*[]string, bool)`

GetAvailableUpgradeVersionsOk returns a tuple with the AvailableUpgradeVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUpgradeVersions

`func (o *Metadata) SetAvailableUpgradeVersions(v []string)`

SetAvailableUpgradeVersions sets AvailableUpgradeVersions field to given value.

### HasAvailableUpgradeVersions

`func (o *Metadata) HasAvailableUpgradeVersions() bool`

HasAvailableUpgradeVersions returns a boolean if a field has been set.

### GetState

`func (o *Metadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Metadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Metadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Metadata) HasState() bool`

HasState returns a boolean if a field has been set.


