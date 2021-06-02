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

// TaskLifecycleConfig struct for TaskLifecycleConfig
type TaskLifecycleConfig struct {
	Hook *string `json:"Hook,omitempty"`
	Sidecar *bool `json:"Sidecar,omitempty"`
}

// NewTaskLifecycleConfig instantiates a new TaskLifecycleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskLifecycleConfig() *TaskLifecycleConfig {
	this := TaskLifecycleConfig{}
	return &this
}

// NewTaskLifecycleConfigWithDefaults instantiates a new TaskLifecycleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskLifecycleConfigWithDefaults() *TaskLifecycleConfig {
	this := TaskLifecycleConfig{}
	return &this
}

// GetHook returns the Hook field value if set, zero value otherwise.
func (o *TaskLifecycleConfig) GetHook() string {
	if o == nil || o.Hook == nil {
		var ret string
		return ret
	}
	return *o.Hook
}

// GetHookOk returns a tuple with the Hook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLifecycleConfig) GetHookOk() (*string, bool) {
	if o == nil || o.Hook == nil {
		return nil, false
	}
	return o.Hook, true
}

// HasHook returns a boolean if a field has been set.
func (o *TaskLifecycleConfig) HasHook() bool {
	if o != nil && o.Hook != nil {
		return true
	}

	return false
}

// SetHook gets a reference to the given string and assigns it to the Hook field.
func (o *TaskLifecycleConfig) SetHook(v string) {
	o.Hook = &v
}

// GetSidecar returns the Sidecar field value if set, zero value otherwise.
func (o *TaskLifecycleConfig) GetSidecar() bool {
	if o == nil || o.Sidecar == nil {
		var ret bool
		return ret
	}
	return *o.Sidecar
}

// GetSidecarOk returns a tuple with the Sidecar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskLifecycleConfig) GetSidecarOk() (*bool, bool) {
	if o == nil || o.Sidecar == nil {
		return nil, false
	}
	return o.Sidecar, true
}

// HasSidecar returns a boolean if a field has been set.
func (o *TaskLifecycleConfig) HasSidecar() bool {
	if o != nil && o.Sidecar != nil {
		return true
	}

	return false
}

// SetSidecar gets a reference to the given bool and assigns it to the Sidecar field.
func (o *TaskLifecycleConfig) SetSidecar(v bool) {
	o.Sidecar = &v
}

func (o TaskLifecycleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hook != nil {
		toSerialize["Hook"] = o.Hook
	}
	if o.Sidecar != nil {
		toSerialize["Sidecar"] = o.Sidecar
	}
	return json.Marshal(toSerialize)
}

type NullableTaskLifecycleConfig struct {
	value *TaskLifecycleConfig
	isSet bool
}

func (v NullableTaskLifecycleConfig) Get() *TaskLifecycleConfig {
	return v.value
}

func (v *NullableTaskLifecycleConfig) Set(val *TaskLifecycleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskLifecycleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskLifecycleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskLifecycleConfig(val *TaskLifecycleConfig) *NullableTaskLifecycleConfig {
	return &NullableTaskLifecycleConfig{value: val, isSet: true}
}

func (v NullableTaskLifecycleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskLifecycleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


