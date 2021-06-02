/*
 * nomad
 *
 * Documentation of the Nomad v1 API.
 *
 * API version: 1.0.0
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv1

import (
	"encoding/json"
)

// RequestedDevice struct for RequestedDevice
type RequestedDevice struct {
	// Affinities are a set of affinites to apply when selecting the device to use.
	Affinities *[]Affinity `json:"Affinities,omitempty"`
	// Constraints are a set of constraints to apply when selecting the device to use.
	Constraints *[]Constraint `json:"Constraints,omitempty"`
	// Count is the number of requested devices
	Count *int32 `json:"Count,omitempty"`
	// Name is the request name. The possible values are as follows: <type>: A single value only specifies the type of request. <vendor>/<type>: A single slash delimiter assumes the vendor and type of device is specified. <vendor>/<type>/<name>: Two slash delimiters assume vendor, type and specific model are specified.  Examples are as follows: \"gpu\" \"nvidia/gpu\" \"nvidia/gpu/GTX2080Ti\"
	Name *string `json:"Name,omitempty"`
}

// NewRequestedDevice instantiates a new RequestedDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedDevice() *RequestedDevice {
	this := RequestedDevice{}
	return &this
}

// NewRequestedDeviceWithDefaults instantiates a new RequestedDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedDeviceWithDefaults() *RequestedDevice {
	this := RequestedDevice{}
	return &this
}

// GetAffinities returns the Affinities field value if set, zero value otherwise.
func (o *RequestedDevice) GetAffinities() []Affinity {
	if o == nil || o.Affinities == nil {
		var ret []Affinity
		return ret
	}
	return *o.Affinities
}

// GetAffinitiesOk returns a tuple with the Affinities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedDevice) GetAffinitiesOk() (*[]Affinity, bool) {
	if o == nil || o.Affinities == nil {
		return nil, false
	}
	return o.Affinities, true
}

// HasAffinities returns a boolean if a field has been set.
func (o *RequestedDevice) HasAffinities() bool {
	if o != nil && o.Affinities != nil {
		return true
	}

	return false
}

// SetAffinities gets a reference to the given []Affinity and assigns it to the Affinities field.
func (o *RequestedDevice) SetAffinities(v []Affinity) {
	o.Affinities = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *RequestedDevice) GetConstraints() []Constraint {
	if o == nil || o.Constraints == nil {
		var ret []Constraint
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedDevice) GetConstraintsOk() (*[]Constraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *RequestedDevice) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []Constraint and assigns it to the Constraints field.
func (o *RequestedDevice) SetConstraints(v []Constraint) {
	o.Constraints = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RequestedDevice) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedDevice) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RequestedDevice) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *RequestedDevice) SetCount(v int32) {
	o.Count = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestedDevice) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedDevice) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestedDevice) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestedDevice) SetName(v string) {
	o.Name = &v
}

func (o RequestedDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Affinities != nil {
		toSerialize["Affinities"] = o.Affinities
	}
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRequestedDevice struct {
	value *RequestedDevice
	isSet bool
}

func (v NullableRequestedDevice) Get() *RequestedDevice {
	return v.value
}

func (v *NullableRequestedDevice) Set(val *RequestedDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedDevice(val *RequestedDevice) *NullableRequestedDevice {
	return &NullableRequestedDevice{value: val, isSet: true}
}

func (v NullableRequestedDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


