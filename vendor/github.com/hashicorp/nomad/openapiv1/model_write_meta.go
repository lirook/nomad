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

// WriteMeta WriteMeta allows a write response to include potentially useful metadata about the write
type WriteMeta struct {
	// This is the index associated with the write
	Index *int32 `json:"Index,omitempty"`
}

// NewWriteMeta instantiates a new WriteMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteMeta() *WriteMeta {
	this := WriteMeta{}
	return &this
}

// NewWriteMetaWithDefaults instantiates a new WriteMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteMetaWithDefaults() *WriteMeta {
	this := WriteMeta{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *WriteMeta) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteMeta) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *WriteMeta) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *WriteMeta) SetIndex(v int32) {
	o.Index = &v
}

func (o WriteMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

type NullableWriteMeta struct {
	value *WriteMeta
	isSet bool
}

func (v NullableWriteMeta) Get() *WriteMeta {
	return v.value
}

func (v *NullableWriteMeta) Set(val *WriteMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteMeta(val *WriteMeta) *NullableWriteMeta {
	return &NullableWriteMeta{value: val, isSet: true}
}

func (v NullableWriteMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


