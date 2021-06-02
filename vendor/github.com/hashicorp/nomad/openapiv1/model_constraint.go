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

// Constraint struct for Constraint
type Constraint struct {
	LTarget *string `json:"LTarget,omitempty"`
	Operand *string `json:"Operand,omitempty"`
	RTarget *string `json:"RTarget,omitempty"`
}

// NewConstraint instantiates a new Constraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraint() *Constraint {
	this := Constraint{}
	return &this
}

// NewConstraintWithDefaults instantiates a new Constraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintWithDefaults() *Constraint {
	this := Constraint{}
	return &this
}

// GetLTarget returns the LTarget field value if set, zero value otherwise.
func (o *Constraint) GetLTarget() string {
	if o == nil || o.LTarget == nil {
		var ret string
		return ret
	}
	return *o.LTarget
}

// GetLTargetOk returns a tuple with the LTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetLTargetOk() (*string, bool) {
	if o == nil || o.LTarget == nil {
		return nil, false
	}
	return o.LTarget, true
}

// HasLTarget returns a boolean if a field has been set.
func (o *Constraint) HasLTarget() bool {
	if o != nil && o.LTarget != nil {
		return true
	}

	return false
}

// SetLTarget gets a reference to the given string and assigns it to the LTarget field.
func (o *Constraint) SetLTarget(v string) {
	o.LTarget = &v
}

// GetOperand returns the Operand field value if set, zero value otherwise.
func (o *Constraint) GetOperand() string {
	if o == nil || o.Operand == nil {
		var ret string
		return ret
	}
	return *o.Operand
}

// GetOperandOk returns a tuple with the Operand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetOperandOk() (*string, bool) {
	if o == nil || o.Operand == nil {
		return nil, false
	}
	return o.Operand, true
}

// HasOperand returns a boolean if a field has been set.
func (o *Constraint) HasOperand() bool {
	if o != nil && o.Operand != nil {
		return true
	}

	return false
}

// SetOperand gets a reference to the given string and assigns it to the Operand field.
func (o *Constraint) SetOperand(v string) {
	o.Operand = &v
}

// GetRTarget returns the RTarget field value if set, zero value otherwise.
func (o *Constraint) GetRTarget() string {
	if o == nil || o.RTarget == nil {
		var ret string
		return ret
	}
	return *o.RTarget
}

// GetRTargetOk returns a tuple with the RTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetRTargetOk() (*string, bool) {
	if o == nil || o.RTarget == nil {
		return nil, false
	}
	return o.RTarget, true
}

// HasRTarget returns a boolean if a field has been set.
func (o *Constraint) HasRTarget() bool {
	if o != nil && o.RTarget != nil {
		return true
	}

	return false
}

// SetRTarget gets a reference to the given string and assigns it to the RTarget field.
func (o *Constraint) SetRTarget(v string) {
	o.RTarget = &v
}

func (o Constraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LTarget != nil {
		toSerialize["LTarget"] = o.LTarget
	}
	if o.Operand != nil {
		toSerialize["Operand"] = o.Operand
	}
	if o.RTarget != nil {
		toSerialize["RTarget"] = o.RTarget
	}
	return json.Marshal(toSerialize)
}

type NullableConstraint struct {
	value *Constraint
	isSet bool
}

func (v NullableConstraint) Get() *Constraint {
	return v.value
}

func (v *NullableConstraint) Set(val *Constraint) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraint(val *Constraint) *NullableConstraint {
	return &NullableConstraint{value: val, isSet: true}
}

func (v NullableConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


