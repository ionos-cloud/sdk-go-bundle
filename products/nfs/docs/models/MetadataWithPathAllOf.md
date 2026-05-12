# MetadataWithPathAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**NfsPath** | **string** | The path of the NFS export (currently equal to the UUID of the share). On a machine with access to the share, mount it using the following command: &#x60;mount -t nfs &lt;cluster-ip&gt;:&lt;nfs-path&gt; &lt;target-dir&gt;&#x60;  | [readonly] |

## Methods

### NewMetadataWithPathAllOf

`func NewMetadataWithPathAllOf(nfsPath string, ) *MetadataWithPathAllOf`

NewMetadataWithPathAllOf instantiates a new MetadataWithPathAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithPathAllOfWithDefaults

`func NewMetadataWithPathAllOfWithDefaults() *MetadataWithPathAllOf`

NewMetadataWithPathAllOfWithDefaults instantiates a new MetadataWithPathAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfsPath

`func (o *MetadataWithPathAllOf) GetNfsPath() string`

GetNfsPath returns the NfsPath field if non-nil, zero value otherwise.

### GetNfsPathOk

`func (o *MetadataWithPathAllOf) GetNfsPathOk() (*string, bool)`

GetNfsPathOk returns a tuple with the NfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsPath

`func (o *MetadataWithPathAllOf) SetNfsPath(v string)`

SetNfsPath sets NfsPath field to given value.



