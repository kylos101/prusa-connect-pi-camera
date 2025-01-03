/*
connect-api

# API Rules * Null values are not sent / received. * Object IDs are **always** sent in list methods, but are **ignored** in create / update methods. * All request and response objects are at the root of the returned structure, if they contain only one entity. * Response of multiple entities is returned as an object that contains the list of entities and a structure `pager`, if necessary. ### Additional documentation: * [Camera registration](../camera_registration/) * [Camera communication](../camera_communication/) ### HTTP Status * 200 - OK, response contains data (usually the entire entity) * 201 - OK, entry created; response contains data as required * 204 - OK, no response * 304 - Response has not been modified * 400 - Invalid request / invalid input (SDK error) * 401 - Endpoint is being accessed without credentials (SDK error) * 403 - Request can't be served, usually due to insufficient rights (user error) * 404 - Entity not found (user error or outdated data) * 409 - Conflict with the state of target resource (user error) * 50x - Server side error

API version: 0.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CameraConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CameraConfig{}

// CameraConfig Configuration of the camera
type CameraConfig struct {
	CameraId *string `json:"camera_id,omitempty"`
	// Path to the driver
	Path *string `json:"path,omitempty"`
	// Name defined by user
	Name string `json:"name"`
	Driver *string `json:"driver,omitempty"`
	// Type of snapshot trigger scheme. Manual, layer, gcode only for LINK cameras
	TriggerScheme *string `json:"trigger_scheme,omitempty"`
	Resolution *CameraResolution `json:"resolution,omitempty"`
	NetworkInfo *NetworkInfo `json:"network_info,omitempty"`
	// Firmware version of the camera
	Firmware *string `json:"firmware,omitempty"`
	// Manufacturer of the camera
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Model of the camera
	Model *string `json:"model,omitempty"`
}

type _CameraConfig CameraConfig

// NewCameraConfig instantiates a new CameraConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCameraConfig(name string) *CameraConfig {
	this := CameraConfig{}
	this.Name = name
	var triggerScheme string = "THIRTY_SEC"
	this.TriggerScheme = &triggerScheme
	return &this
}

// NewCameraConfigWithDefaults instantiates a new CameraConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCameraConfigWithDefaults() *CameraConfig {
	this := CameraConfig{}
	var triggerScheme string = "THIRTY_SEC"
	this.TriggerScheme = &triggerScheme
	return &this
}

// GetCameraId returns the CameraId field value if set, zero value otherwise.
func (o *CameraConfig) GetCameraId() string {
	if o == nil || IsNil(o.CameraId) {
		var ret string
		return ret
	}
	return *o.CameraId
}

// GetCameraIdOk returns a tuple with the CameraId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetCameraIdOk() (*string, bool) {
	if o == nil || IsNil(o.CameraId) {
		return nil, false
	}
	return o.CameraId, true
}

// HasCameraId returns a boolean if a field has been set.
func (o *CameraConfig) HasCameraId() bool {
	if o != nil && !IsNil(o.CameraId) {
		return true
	}

	return false
}

// SetCameraId gets a reference to the given string and assigns it to the CameraId field.
func (o *CameraConfig) SetCameraId(v string) {
	o.CameraId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CameraConfig) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CameraConfig) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CameraConfig) SetPath(v string) {
	o.Path = &v
}

// GetName returns the Name field value
func (o *CameraConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CameraConfig) SetName(v string) {
	o.Name = v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *CameraConfig) GetDriver() string {
	if o == nil || IsNil(o.Driver) {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetDriverOk() (*string, bool) {
	if o == nil || IsNil(o.Driver) {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *CameraConfig) HasDriver() bool {
	if o != nil && !IsNil(o.Driver) {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *CameraConfig) SetDriver(v string) {
	o.Driver = &v
}

// GetTriggerScheme returns the TriggerScheme field value if set, zero value otherwise.
func (o *CameraConfig) GetTriggerScheme() string {
	if o == nil || IsNil(o.TriggerScheme) {
		var ret string
		return ret
	}
	return *o.TriggerScheme
}

// GetTriggerSchemeOk returns a tuple with the TriggerScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetTriggerSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerScheme) {
		return nil, false
	}
	return o.TriggerScheme, true
}

// HasTriggerScheme returns a boolean if a field has been set.
func (o *CameraConfig) HasTriggerScheme() bool {
	if o != nil && !IsNil(o.TriggerScheme) {
		return true
	}

	return false
}

// SetTriggerScheme gets a reference to the given string and assigns it to the TriggerScheme field.
func (o *CameraConfig) SetTriggerScheme(v string) {
	o.TriggerScheme = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *CameraConfig) GetResolution() CameraResolution {
	if o == nil || IsNil(o.Resolution) {
		var ret CameraResolution
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetResolutionOk() (*CameraResolution, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *CameraConfig) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given CameraResolution and assigns it to the Resolution field.
func (o *CameraConfig) SetResolution(v CameraResolution) {
	o.Resolution = &v
}

// GetNetworkInfo returns the NetworkInfo field value if set, zero value otherwise.
func (o *CameraConfig) GetNetworkInfo() NetworkInfo {
	if o == nil || IsNil(o.NetworkInfo) {
		var ret NetworkInfo
		return ret
	}
	return *o.NetworkInfo
}

// GetNetworkInfoOk returns a tuple with the NetworkInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetNetworkInfoOk() (*NetworkInfo, bool) {
	if o == nil || IsNil(o.NetworkInfo) {
		return nil, false
	}
	return o.NetworkInfo, true
}

// HasNetworkInfo returns a boolean if a field has been set.
func (o *CameraConfig) HasNetworkInfo() bool {
	if o != nil && !IsNil(o.NetworkInfo) {
		return true
	}

	return false
}

// SetNetworkInfo gets a reference to the given NetworkInfo and assigns it to the NetworkInfo field.
func (o *CameraConfig) SetNetworkInfo(v NetworkInfo) {
	o.NetworkInfo = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *CameraConfig) GetFirmware() string {
	if o == nil || IsNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetFirmwareOk() (*string, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *CameraConfig) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *CameraConfig) SetFirmware(v string) {
	o.Firmware = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *CameraConfig) GetManufacturer() string {
	if o == nil || IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetManufacturerOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *CameraConfig) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *CameraConfig) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CameraConfig) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraConfig) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CameraConfig) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CameraConfig) SetModel(v string) {
	o.Model = &v
}

func (o CameraConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CameraConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CameraId) {
		toSerialize["camera_id"] = o.CameraId
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Driver) {
		toSerialize["driver"] = o.Driver
	}
	if !IsNil(o.TriggerScheme) {
		toSerialize["trigger_scheme"] = o.TriggerScheme
	}
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.NetworkInfo) {
		toSerialize["network_info"] = o.NetworkInfo
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return toSerialize, nil
}

func (o *CameraConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCameraConfig := _CameraConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCameraConfig)

	if err != nil {
		return err
	}

	*o = CameraConfig(varCameraConfig)

	return err
}

type NullableCameraConfig struct {
	value *CameraConfig
	isSet bool
}

func (v NullableCameraConfig) Get() *CameraConfig {
	return v.value
}

func (v *NullableCameraConfig) Set(val *CameraConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCameraConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCameraConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCameraConfig(val *CameraConfig) *NullableCameraConfig {
	return &NullableCameraConfig{value: val, isSet: true}
}

func (v NullableCameraConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCameraConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


