/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// checks if the VolumeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeProperties{}

// VolumeProperties struct for VolumeProperties
type VolumeProperties struct {
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
	// Hardware type of the volume. DAS (Direct Attached Storage) could be used only in a composite call with a Cube server.
	Type *string `json:"type,omitempty"`
	// The size of the volume in GB.
	Size float32 `json:"size"`
	// The availability zone in which the volume should be provisioned. The storage volume will be provisioned on as few physical storage devices as possible, but this cannot be guaranteed upfront. This is uavailable for DAS (Direct Attached Storage), and subject to availability for SSD.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// Image or snapshot ID to be used as template for this volume.
	Image *string `json:"image,omitempty"`
	// Initial password to be set for installed OS. Works with public images only. Not modifiable, forbidden in update requests. Password rules allows all characters from a-z, A-Z, 0-9.
	ImagePassword *string `json:"imagePassword,omitempty"`
	ImageAlias    *string `json:"imageAlias,omitempty"`
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	SshKeys []string `json:"sshKeys,omitempty"`
	// The bus type for this volume; default is VIRTIO.
	Bus *string `json:"bus,omitempty"`
	// OS type for this volume.
	LicenceType *string `json:"licenceType,omitempty"`
	// Hot-plug capable CPU (no reboot required).
	CpuHotPlug *bool `json:"cpuHotPlug,omitempty"`
	// Hot-plug capable RAM (no reboot required).
	RamHotPlug *bool `json:"ramHotPlug,omitempty"`
	// Hot-plug capable NIC (no reboot required).
	NicHotPlug *bool `json:"nicHotPlug,omitempty"`
	// Hot-unplug capable NIC (no reboot required).
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty"`
	// Hot-plug capable Virt-IO drive (no reboot required).
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty"`
	// Hot-unplug capable Virt-IO drive (no reboot required). Not supported with Windows VMs.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty"`
	// If set to `true` will expose the serial id of the disk attached to the server. If set to `false` will not expose the serial id. Some operating systems or software solutions require the serial id to be exposed to work properly. Exposing the serial  can influence licensed software (e.g. Windows) behavior
	ExposeSerial *bool `json:"exposeSerial,omitempty"`
	// The Logical Unit Number of the storage volume. Null for volumes, not mounted to a VM.
	DeviceNumber *int64 `json:"deviceNumber,omitempty"`
	// The PCI slot number of the storage volume. Null for volumes, not mounted to a VM.
	PciSlot *int32 `json:"pciSlot,omitempty"`
	// The ID of the backup unit that the user has access to. The property is immutable and is only allowed to be set on creation of a new a volume. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	BackupunitId *string `json:"backupunitId,omitempty"`
	// The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on creation of a new a volume. It is mandatory to provide either 'public image' or 'imageAlias' that has cloud-init compatibility in conjunction with this property.
	UserData *string `json:"userData,omitempty"`
	// The UUID of the attached server.
	BootServer *string `json:"bootServer,omitempty"`
	// Determines whether the volume will be used as a boot volume. Set to `NONE`, the volume will not be used as boot volume. Set to `PRIMARY`, the volume will be used as boot volume and all other volumes must be set to `NONE`. Set to `AUTO` or `null` requires all volumes to be set to `AUTO` or `null`; this will use the legacy behavior, which is to use the volume as a boot volume only if there are no other volumes or cdrom devices.
	BootOrder NullableString `json:"bootOrder,omitempty"`
}

// NewVolumeProperties instantiates a new VolumeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeProperties(size float32) *VolumeProperties {
	this := VolumeProperties{}

	this.Size = size
	var exposeSerial bool = false
	this.ExposeSerial = &exposeSerial
	var bootOrder = "AUTO"
	this.BootOrder = *NewNullableString(&bootOrder)

	return &this
}

// NewVolumePropertiesWithDefaults instantiates a new VolumeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePropertiesWithDefaults() *VolumeProperties {
	this := VolumeProperties{}
	var exposeSerial bool = false
	this.ExposeSerial = &exposeSerial
	var bootOrder = "AUTO"
	this.BootOrder = *NewNullableString(&bootOrder)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumeProperties) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumeProperties) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumeProperties) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VolumeProperties) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VolumeProperties) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VolumeProperties) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value
func (o *VolumeProperties) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *VolumeProperties) SetSize(v float32) {
	o.Size = v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *VolumeProperties) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *VolumeProperties) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *VolumeProperties) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *VolumeProperties) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *VolumeProperties) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *VolumeProperties) SetImage(v string) {
	o.Image = &v
}

// GetImagePassword returns the ImagePassword field value if set, zero value otherwise.
func (o *VolumeProperties) GetImagePassword() string {
	if o == nil || IsNil(o.ImagePassword) {
		var ret string
		return ret
	}
	return *o.ImagePassword
}

// GetImagePasswordOk returns a tuple with the ImagePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetImagePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ImagePassword) {
		return nil, false
	}
	return o.ImagePassword, true
}

// HasImagePassword returns a boolean if a field has been set.
func (o *VolumeProperties) HasImagePassword() bool {
	if o != nil && !IsNil(o.ImagePassword) {
		return true
	}

	return false
}

// SetImagePassword gets a reference to the given string and assigns it to the ImagePassword field.
func (o *VolumeProperties) SetImagePassword(v string) {
	o.ImagePassword = &v
}

// GetImageAlias returns the ImageAlias field value if set, zero value otherwise.
func (o *VolumeProperties) GetImageAlias() string {
	if o == nil || IsNil(o.ImageAlias) {
		var ret string
		return ret
	}
	return *o.ImageAlias
}

// GetImageAliasOk returns a tuple with the ImageAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetImageAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ImageAlias) {
		return nil, false
	}
	return o.ImageAlias, true
}

// HasImageAlias returns a boolean if a field has been set.
func (o *VolumeProperties) HasImageAlias() bool {
	if o != nil && !IsNil(o.ImageAlias) {
		return true
	}

	return false
}

// SetImageAlias gets a reference to the given string and assigns it to the ImageAlias field.
func (o *VolumeProperties) SetImageAlias(v string) {
	o.ImageAlias = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *VolumeProperties) GetSshKeys() []string {
	if o == nil || IsNil(o.SshKeys) {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetSshKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *VolumeProperties) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *VolumeProperties) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *VolumeProperties) GetBus() string {
	if o == nil || IsNil(o.Bus) {
		var ret string
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetBusOk() (*string, bool) {
	if o == nil || IsNil(o.Bus) {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *VolumeProperties) HasBus() bool {
	if o != nil && !IsNil(o.Bus) {
		return true
	}

	return false
}

// SetBus gets a reference to the given string and assigns it to the Bus field.
func (o *VolumeProperties) SetBus(v string) {
	o.Bus = &v
}

// GetLicenceType returns the LicenceType field value if set, zero value otherwise.
func (o *VolumeProperties) GetLicenceType() string {
	if o == nil || IsNil(o.LicenceType) {
		var ret string
		return ret
	}
	return *o.LicenceType
}

// GetLicenceTypeOk returns a tuple with the LicenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetLicenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LicenceType) {
		return nil, false
	}
	return o.LicenceType, true
}

// HasLicenceType returns a boolean if a field has been set.
func (o *VolumeProperties) HasLicenceType() bool {
	if o != nil && !IsNil(o.LicenceType) {
		return true
	}

	return false
}

// SetLicenceType gets a reference to the given string and assigns it to the LicenceType field.
func (o *VolumeProperties) SetLicenceType(v string) {
	o.LicenceType = &v
}

// GetCpuHotPlug returns the CpuHotPlug field value if set, zero value otherwise.
func (o *VolumeProperties) GetCpuHotPlug() bool {
	if o == nil || IsNil(o.CpuHotPlug) {
		var ret bool
		return ret
	}
	return *o.CpuHotPlug
}

// GetCpuHotPlugOk returns a tuple with the CpuHotPlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetCpuHotPlugOk() (*bool, bool) {
	if o == nil || IsNil(o.CpuHotPlug) {
		return nil, false
	}
	return o.CpuHotPlug, true
}

// HasCpuHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasCpuHotPlug() bool {
	if o != nil && !IsNil(o.CpuHotPlug) {
		return true
	}

	return false
}

// SetCpuHotPlug gets a reference to the given bool and assigns it to the CpuHotPlug field.
func (o *VolumeProperties) SetCpuHotPlug(v bool) {
	o.CpuHotPlug = &v
}

// GetRamHotPlug returns the RamHotPlug field value if set, zero value otherwise.
func (o *VolumeProperties) GetRamHotPlug() bool {
	if o == nil || IsNil(o.RamHotPlug) {
		var ret bool
		return ret
	}
	return *o.RamHotPlug
}

// GetRamHotPlugOk returns a tuple with the RamHotPlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetRamHotPlugOk() (*bool, bool) {
	if o == nil || IsNil(o.RamHotPlug) {
		return nil, false
	}
	return o.RamHotPlug, true
}

// HasRamHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasRamHotPlug() bool {
	if o != nil && !IsNil(o.RamHotPlug) {
		return true
	}

	return false
}

// SetRamHotPlug gets a reference to the given bool and assigns it to the RamHotPlug field.
func (o *VolumeProperties) SetRamHotPlug(v bool) {
	o.RamHotPlug = &v
}

// GetNicHotPlug returns the NicHotPlug field value if set, zero value otherwise.
func (o *VolumeProperties) GetNicHotPlug() bool {
	if o == nil || IsNil(o.NicHotPlug) {
		var ret bool
		return ret
	}
	return *o.NicHotPlug
}

// GetNicHotPlugOk returns a tuple with the NicHotPlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetNicHotPlugOk() (*bool, bool) {
	if o == nil || IsNil(o.NicHotPlug) {
		return nil, false
	}
	return o.NicHotPlug, true
}

// HasNicHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasNicHotPlug() bool {
	if o != nil && !IsNil(o.NicHotPlug) {
		return true
	}

	return false
}

// SetNicHotPlug gets a reference to the given bool and assigns it to the NicHotPlug field.
func (o *VolumeProperties) SetNicHotPlug(v bool) {
	o.NicHotPlug = &v
}

// GetNicHotUnplug returns the NicHotUnplug field value if set, zero value otherwise.
func (o *VolumeProperties) GetNicHotUnplug() bool {
	if o == nil || IsNil(o.NicHotUnplug) {
		var ret bool
		return ret
	}
	return *o.NicHotUnplug
}

// GetNicHotUnplugOk returns a tuple with the NicHotUnplug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetNicHotUnplugOk() (*bool, bool) {
	if o == nil || IsNil(o.NicHotUnplug) {
		return nil, false
	}
	return o.NicHotUnplug, true
}

// HasNicHotUnplug returns a boolean if a field has been set.
func (o *VolumeProperties) HasNicHotUnplug() bool {
	if o != nil && !IsNil(o.NicHotUnplug) {
		return true
	}

	return false
}

// SetNicHotUnplug gets a reference to the given bool and assigns it to the NicHotUnplug field.
func (o *VolumeProperties) SetNicHotUnplug(v bool) {
	o.NicHotUnplug = &v
}

// GetDiscVirtioHotPlug returns the DiscVirtioHotPlug field value if set, zero value otherwise.
func (o *VolumeProperties) GetDiscVirtioHotPlug() bool {
	if o == nil || IsNil(o.DiscVirtioHotPlug) {
		var ret bool
		return ret
	}
	return *o.DiscVirtioHotPlug
}

// GetDiscVirtioHotPlugOk returns a tuple with the DiscVirtioHotPlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetDiscVirtioHotPlugOk() (*bool, bool) {
	if o == nil || IsNil(o.DiscVirtioHotPlug) {
		return nil, false
	}
	return o.DiscVirtioHotPlug, true
}

// HasDiscVirtioHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasDiscVirtioHotPlug() bool {
	if o != nil && !IsNil(o.DiscVirtioHotPlug) {
		return true
	}

	return false
}

// SetDiscVirtioHotPlug gets a reference to the given bool and assigns it to the DiscVirtioHotPlug field.
func (o *VolumeProperties) SetDiscVirtioHotPlug(v bool) {
	o.DiscVirtioHotPlug = &v
}

// GetDiscVirtioHotUnplug returns the DiscVirtioHotUnplug field value if set, zero value otherwise.
func (o *VolumeProperties) GetDiscVirtioHotUnplug() bool {
	if o == nil || IsNil(o.DiscVirtioHotUnplug) {
		var ret bool
		return ret
	}
	return *o.DiscVirtioHotUnplug
}

// GetDiscVirtioHotUnplugOk returns a tuple with the DiscVirtioHotUnplug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetDiscVirtioHotUnplugOk() (*bool, bool) {
	if o == nil || IsNil(o.DiscVirtioHotUnplug) {
		return nil, false
	}
	return o.DiscVirtioHotUnplug, true
}

// HasDiscVirtioHotUnplug returns a boolean if a field has been set.
func (o *VolumeProperties) HasDiscVirtioHotUnplug() bool {
	if o != nil && !IsNil(o.DiscVirtioHotUnplug) {
		return true
	}

	return false
}

// SetDiscVirtioHotUnplug gets a reference to the given bool and assigns it to the DiscVirtioHotUnplug field.
func (o *VolumeProperties) SetDiscVirtioHotUnplug(v bool) {
	o.DiscVirtioHotUnplug = &v
}

// GetExposeSerial returns the ExposeSerial field value if set, zero value otherwise.
func (o *VolumeProperties) GetExposeSerial() bool {
	if o == nil || IsNil(o.ExposeSerial) {
		var ret bool
		return ret
	}
	return *o.ExposeSerial
}

// GetExposeSerialOk returns a tuple with the ExposeSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetExposeSerialOk() (*bool, bool) {
	if o == nil || IsNil(o.ExposeSerial) {
		return nil, false
	}
	return o.ExposeSerial, true
}

// HasExposeSerial returns a boolean if a field has been set.
func (o *VolumeProperties) HasExposeSerial() bool {
	if o != nil && !IsNil(o.ExposeSerial) {
		return true
	}

	return false
}

// SetExposeSerial gets a reference to the given bool and assigns it to the ExposeSerial field.
func (o *VolumeProperties) SetExposeSerial(v bool) {
	o.ExposeSerial = &v
}

// GetDeviceNumber returns the DeviceNumber field value if set, zero value otherwise.
func (o *VolumeProperties) GetDeviceNumber() int64 {
	if o == nil || IsNil(o.DeviceNumber) {
		var ret int64
		return ret
	}
	return *o.DeviceNumber
}

// GetDeviceNumberOk returns a tuple with the DeviceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetDeviceNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceNumber) {
		return nil, false
	}
	return o.DeviceNumber, true
}

// HasDeviceNumber returns a boolean if a field has been set.
func (o *VolumeProperties) HasDeviceNumber() bool {
	if o != nil && !IsNil(o.DeviceNumber) {
		return true
	}

	return false
}

// SetDeviceNumber gets a reference to the given int64 and assigns it to the DeviceNumber field.
func (o *VolumeProperties) SetDeviceNumber(v int64) {
	o.DeviceNumber = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *VolumeProperties) GetPciSlot() int32 {
	if o == nil || IsNil(o.PciSlot) {
		var ret int32
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetPciSlotOk() (*int32, bool) {
	if o == nil || IsNil(o.PciSlot) {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *VolumeProperties) HasPciSlot() bool {
	if o != nil && !IsNil(o.PciSlot) {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given int32 and assigns it to the PciSlot field.
func (o *VolumeProperties) SetPciSlot(v int32) {
	o.PciSlot = &v
}

// GetBackupunitId returns the BackupunitId field value if set, zero value otherwise.
func (o *VolumeProperties) GetBackupunitId() string {
	if o == nil || IsNil(o.BackupunitId) {
		var ret string
		return ret
	}
	return *o.BackupunitId
}

// GetBackupunitIdOk returns a tuple with the BackupunitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetBackupunitIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackupunitId) {
		return nil, false
	}
	return o.BackupunitId, true
}

// HasBackupunitId returns a boolean if a field has been set.
func (o *VolumeProperties) HasBackupunitId() bool {
	if o != nil && !IsNil(o.BackupunitId) {
		return true
	}

	return false
}

// SetBackupunitId gets a reference to the given string and assigns it to the BackupunitId field.
func (o *VolumeProperties) SetBackupunitId(v string) {
	o.BackupunitId = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *VolumeProperties) GetUserData() string {
	if o == nil || IsNil(o.UserData) {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetUserDataOk() (*string, bool) {
	if o == nil || IsNil(o.UserData) {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *VolumeProperties) HasUserData() bool {
	if o != nil && !IsNil(o.UserData) {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *VolumeProperties) SetUserData(v string) {
	o.UserData = &v
}

// GetBootServer returns the BootServer field value if set, zero value otherwise.
func (o *VolumeProperties) GetBootServer() string {
	if o == nil || IsNil(o.BootServer) {
		var ret string
		return ret
	}
	return *o.BootServer
}

// GetBootServerOk returns a tuple with the BootServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeProperties) GetBootServerOk() (*string, bool) {
	if o == nil || IsNil(o.BootServer) {
		return nil, false
	}
	return o.BootServer, true
}

// HasBootServer returns a boolean if a field has been set.
func (o *VolumeProperties) HasBootServer() bool {
	if o != nil && !IsNil(o.BootServer) {
		return true
	}

	return false
}

// SetBootServer gets a reference to the given string and assigns it to the BootServer field.
func (o *VolumeProperties) SetBootServer(v string) {
	o.BootServer = &v
}

// GetBootOrder returns the BootOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VolumeProperties) GetBootOrder() string {
	if o == nil || IsNil(o.BootOrder.Get()) {
		var ret string
		return ret
	}
	return *o.BootOrder.Get()
}

// GetBootOrderOk returns a tuple with the BootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetBootOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BootOrder.Get(), o.BootOrder.IsSet()
}

// HasBootOrder returns a boolean if a field has been set.
func (o *VolumeProperties) HasBootOrder() bool {
	if o != nil && o.BootOrder.IsSet() {
		return true
	}

	return false
}

// SetBootOrder gets a reference to the given NullableString and assigns it to the BootOrder field.
func (o *VolumeProperties) SetBootOrder(v string) {
	o.BootOrder.Set(&v)
}

// SetBootOrderNil sets the value for BootOrder to be an explicit nil
func (o *VolumeProperties) SetBootOrderNil() {
	o.BootOrder.Set(nil)
}

// UnsetBootOrder ensures that no value is present for BootOrder, not even an explicit nil
func (o *VolumeProperties) UnsetBootOrder() {
	o.BootOrder.Unset()
}

func (o VolumeProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsZero(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ImagePassword) {
		toSerialize["imagePassword"] = o.ImagePassword
	}
	if !IsNil(o.ImageAlias) {
		toSerialize["imageAlias"] = o.ImageAlias
	}
	if !IsNil(o.SshKeys) {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if !IsNil(o.Bus) {
		toSerialize["bus"] = o.Bus
	}
	if !IsNil(o.LicenceType) {
		toSerialize["licenceType"] = o.LicenceType
	}
	if !IsNil(o.CpuHotPlug) {
		toSerialize["cpuHotPlug"] = o.CpuHotPlug
	}
	if !IsNil(o.RamHotPlug) {
		toSerialize["ramHotPlug"] = o.RamHotPlug
	}
	if !IsNil(o.NicHotPlug) {
		toSerialize["nicHotPlug"] = o.NicHotPlug
	}
	if !IsNil(o.NicHotUnplug) {
		toSerialize["nicHotUnplug"] = o.NicHotUnplug
	}
	if !IsNil(o.DiscVirtioHotPlug) {
		toSerialize["discVirtioHotPlug"] = o.DiscVirtioHotPlug
	}
	if !IsNil(o.DiscVirtioHotUnplug) {
		toSerialize["discVirtioHotUnplug"] = o.DiscVirtioHotUnplug
	}
	if !IsNil(o.ExposeSerial) {
		toSerialize["exposeSerial"] = o.ExposeSerial
	}
	if !IsNil(o.DeviceNumber) {
		toSerialize["deviceNumber"] = o.DeviceNumber
	}
	if !IsNil(o.PciSlot) {
		toSerialize["pciSlot"] = o.PciSlot
	}
	if !IsNil(o.BackupunitId) {
		toSerialize["backupunitId"] = o.BackupunitId
	}
	if !IsNil(o.UserData) {
		toSerialize["userData"] = o.UserData
	}
	if !IsNil(o.BootServer) {
		toSerialize["bootServer"] = o.BootServer
	}
	if o.BootOrder.IsSet() {
		toSerialize["bootOrder"] = o.BootOrder.Get()
	}
	return toSerialize, nil
}

type NullableVolumeProperties struct {
	value *VolumeProperties
	isSet bool
}

func (v NullableVolumeProperties) Get() *VolumeProperties {
	return v.value
}

func (v *NullableVolumeProperties) Set(val *VolumeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeProperties(val *VolumeProperties) *NullableVolumeProperties {
	return &NullableVolumeProperties{value: val, isSet: true}
}

func (v NullableVolumeProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
