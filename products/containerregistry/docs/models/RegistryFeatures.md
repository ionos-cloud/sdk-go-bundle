# RegistryFeatures

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**VulnerabilityScanning** | Pointer to [**FeatureVulnerabilityScanning**](FeatureVulnerabilityScanning.md) |  | [optional] |

## Methods

### NewRegistryFeatures

`func NewRegistryFeatures() *RegistryFeatures`

NewRegistryFeatures instantiates a new RegistryFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryFeaturesWithDefaults

`func NewRegistryFeaturesWithDefaults() *RegistryFeatures`

NewRegistryFeaturesWithDefaults instantiates a new RegistryFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVulnerabilityScanning

`func (o *RegistryFeatures) GetVulnerabilityScanning() FeatureVulnerabilityScanning`

GetVulnerabilityScanning returns the VulnerabilityScanning field if non-nil, zero value otherwise.

### GetVulnerabilityScanningOk

`func (o *RegistryFeatures) GetVulnerabilityScanningOk() (*FeatureVulnerabilityScanning, bool)`

GetVulnerabilityScanningOk returns a tuple with the VulnerabilityScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanning

`func (o *RegistryFeatures) SetVulnerabilityScanning(v FeatureVulnerabilityScanning)`

SetVulnerabilityScanning sets VulnerabilityScanning field to given value.

### HasVulnerabilityScanning

`func (o *RegistryFeatures) HasVulnerabilityScanning() bool`

HasVulnerabilityScanning returns a boolean if a field has been set.


