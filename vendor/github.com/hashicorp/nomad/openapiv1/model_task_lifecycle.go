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

// TaskLifecycle struct for TaskLifecycle
type TaskLifecycle struct {
	Hook *string `json:"Hook,omitempty"`
	Sidecar *bool `json:"Sidecar,omitempty"`
}

// NewTaskLifecycle instantiates a new TaskLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskLifecycle() *TaskLifecycle {
	this := TaskLifecycle{}
	return &this
}

// NewTaskLifecycleWithDefaults instantiates a new TaskLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskLifecycleWithDefaults() *TaskLifecycle {
	this := TaskLifecycle{}
	return &this
}

// GetHook returns the Hook field value if set, zero value otherwise.
func (o *TaskLifecycle) GetHook() string {
	if o == nil || o.Hook == nil {
		var ret string
		return ret
	}
	return *o.Hook
}

// GetHookOk returns a tuple with the Hook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLifecycle) GetHookOk() (*string, bool) {
	if o == nil || o.Hook == nil {
		return nil, false
	}
	return o.Hook, true
}

// HasHook returns a boolean if a field has been set.
func (o *TaskLifecycle) HasHook() bool {
	if o != nil && o.Hook != nil {
		return true
	}

	return false
}

// SetHook gets a reference to the given string and assigns it to the Hook field.
func (o *TaskLifecycle) SetHook(v string) {
	o.Hook = &v
}

// GetSidecar returns the Sidecar field value if set, zero value otherwise.
func (o *TaskLifecycle) GetSidecar() bool {
	if o == nil || o.Sidecar == nil {
		var ret bool
		return ret
	}
	return *o.Sidecar
}

// GetSidecarOk returns a tuple with the Sidecar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLifecycle) GetSidecarOk() (*bool, bool) {
	if o == nil || o.Sidecar == nil {
		return nil, false
	}
	return o.Sidecar, true
}

// HasSidecar returns a boolean if a field has been set.
func (o *TaskLifecycle) HasSidecar() bool {
	if o != nil && o.Sidecar != nil {
		return true
	}

	return false
}

// SetSidecar gets a reference to the given bool and assigns it to the Sidecar field.
func (o *TaskLifecycle) SetSidecar(v bool) {
	o.Sidecar = &v
}

func (o TaskLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hook != nil {
		toSerialize["Hook"] = o.Hook
	}
	if o.Sidecar != nil {
		toSerialize["Sidecar"] = o.Sidecar
	}
	return json.Marshal(toSerialize)
}

type NullableTaskLifecycle struct {
	value *TaskLifecycle
	isSet bool
}

func (v NullableTaskLifecycle) Get() *TaskLifecycle {
	return v.value
}

func (v *NullableTaskLifecycle) Set(val *TaskLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskLifecycle(val *TaskLifecycle) *NullableTaskLifecycle {
	return &NullableTaskLifecycle{value: val, isSet: true}
}

func (v NullableTaskLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


