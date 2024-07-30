# WireguardGatewayMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**PublicKey** | **string** | Public key correspondng to the WireGuard Server private key.  | |

## Methods

### NewWireguardGatewayMetadataAllOf

`func NewWireguardGatewayMetadataAllOf(publicKey string, ) *WireguardGatewayMetadataAllOf`

NewWireguardGatewayMetadataAllOf instantiates a new WireguardGatewayMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardGatewayMetadataAllOfWithDefaults

`func NewWireguardGatewayMetadataAllOfWithDefaults() *WireguardGatewayMetadataAllOf`

NewWireguardGatewayMetadataAllOfWithDefaults instantiates a new WireguardGatewayMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *WireguardGatewayMetadataAllOf) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardGatewayMetadataAllOf) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardGatewayMetadataAllOf) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



