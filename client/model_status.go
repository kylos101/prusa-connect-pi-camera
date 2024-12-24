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

// checks if the Status type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Status{}

// Status struct for Status
type Status struct {
	// Error code
	Code string `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}

type _Status Status

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus(code string, message string) *Status {
	this := Status{}
	this.Code = code
	this.Message = message
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetCode returns the Code field value
func (o *Status) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Status) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Status) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *Status) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Status) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Status) SetMessage(v string) {
	o.Message = v
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Status) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *Status) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varStatus := _Status{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatus)

	if err != nil {
		return err
	}

	*o = Status(varStatus)

	return err
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


