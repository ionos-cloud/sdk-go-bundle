# Quota

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**QuotaUsage** | [**QuotaDetail**](QuotaDetail.md) |  | |
|**QuotaLimits** | [**QuotaDetail**](QuotaDetail.md) |  | |

## Methods

### NewQuota

`func NewQuota(quotaUsage QuotaDetail, quotaLimits QuotaDetail, ) *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotaUsage

`func (o *Quota) GetQuotaUsage() QuotaDetail`

GetQuotaUsage returns the QuotaUsage field if non-nil, zero value otherwise.

### GetQuotaUsageOk

`func (o *Quota) GetQuotaUsageOk() (*QuotaDetail, bool)`

GetQuotaUsageOk returns a tuple with the QuotaUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaUsage

`func (o *Quota) SetQuotaUsage(v QuotaDetail)`

SetQuotaUsage sets QuotaUsage field to given value.


### GetQuotaLimits

`func (o *Quota) GetQuotaLimits() QuotaDetail`

GetQuotaLimits returns the QuotaLimits field if non-nil, zero value otherwise.

### GetQuotaLimitsOk

`func (o *Quota) GetQuotaLimitsOk() (*QuotaDetail, bool)`

GetQuotaLimitsOk returns a tuple with the QuotaLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimits

`func (o *Quota) SetQuotaLimits(v QuotaDetail)`

SetQuotaLimits sets QuotaLimits field to given value.



