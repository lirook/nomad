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

// JobRegisterResponse JobRegisterResponse is used to respond to a job registration
type JobRegisterResponse struct {
	EvalCreateIndex *int32 `json:"EvalCreateIndex,omitempty"`
	EvalID *string `json:"EvalID,omitempty"`
	JobModifyIndex *int32 `json:"JobModifyIndex,omitempty"`
	// Is there a known leader
	KnownLeader *bool `json:"KnownLeader,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	LastContact *int64 `json:"LastContact,omitempty"`
	// LastIndex. This can be used as a WaitIndex to perform a blocking query
	LastIndex *int32 `json:"LastIndex,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	RequestTime *int64 `json:"RequestTime,omitempty"`
	// Warnings contains any warnings about the given job. These may include deprecation warnings.
	Warnings *string `json:"Warnings,omitempty"`
}

// NewJobRegisterResponse instantiates a new JobRegisterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRegisterResponse() *JobRegisterResponse {
	this := JobRegisterResponse{}
	return &this
}

// NewJobRegisterResponseWithDefaults instantiates a new JobRegisterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRegisterResponseWithDefaults() *JobRegisterResponse {
	this := JobRegisterResponse{}
	return &this
}

// GetEvalCreateIndex returns the EvalCreateIndex field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetEvalCreateIndex() int32 {
	if o == nil || o.EvalCreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.EvalCreateIndex
}

// GetEvalCreateIndexOk returns a tuple with the EvalCreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetEvalCreateIndexOk() (*int32, bool) {
	if o == nil || o.EvalCreateIndex == nil {
		return nil, false
	}
	return o.EvalCreateIndex, true
}

// HasEvalCreateIndex returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasEvalCreateIndex() bool {
	if o != nil && o.EvalCreateIndex != nil {
		return true
	}

	return false
}

// SetEvalCreateIndex gets a reference to the given int32 and assigns it to the EvalCreateIndex field.
func (o *JobRegisterResponse) SetEvalCreateIndex(v int32) {
	o.EvalCreateIndex = &v
}

// GetEvalID returns the EvalID field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetEvalID() string {
	if o == nil || o.EvalID == nil {
		var ret string
		return ret
	}
	return *o.EvalID
}

// GetEvalIDOk returns a tuple with the EvalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetEvalIDOk() (*string, bool) {
	if o == nil || o.EvalID == nil {
		return nil, false
	}
	return o.EvalID, true
}

// HasEvalID returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasEvalID() bool {
	if o != nil && o.EvalID != nil {
		return true
	}

	return false
}

// SetEvalID gets a reference to the given string and assigns it to the EvalID field.
func (o *JobRegisterResponse) SetEvalID(v string) {
	o.EvalID = &v
}

// GetJobModifyIndex returns the JobModifyIndex field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetJobModifyIndex() int32 {
	if o == nil || o.JobModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.JobModifyIndex
}

// GetJobModifyIndexOk returns a tuple with the JobModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetJobModifyIndexOk() (*int32, bool) {
	if o == nil || o.JobModifyIndex == nil {
		return nil, false
	}
	return o.JobModifyIndex, true
}

// HasJobModifyIndex returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasJobModifyIndex() bool {
	if o != nil && o.JobModifyIndex != nil {
		return true
	}

	return false
}

// SetJobModifyIndex gets a reference to the given int32 and assigns it to the JobModifyIndex field.
func (o *JobRegisterResponse) SetJobModifyIndex(v int32) {
	o.JobModifyIndex = &v
}

// GetKnownLeader returns the KnownLeader field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetKnownLeader() bool {
	if o == nil || o.KnownLeader == nil {
		var ret bool
		return ret
	}
	return *o.KnownLeader
}

// GetKnownLeaderOk returns a tuple with the KnownLeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetKnownLeaderOk() (*bool, bool) {
	if o == nil || o.KnownLeader == nil {
		return nil, false
	}
	return o.KnownLeader, true
}

// HasKnownLeader returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasKnownLeader() bool {
	if o != nil && o.KnownLeader != nil {
		return true
	}

	return false
}

// SetKnownLeader gets a reference to the given bool and assigns it to the KnownLeader field.
func (o *JobRegisterResponse) SetKnownLeader(v bool) {
	o.KnownLeader = &v
}

// GetLastContact returns the LastContact field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetLastContact() int64 {
	if o == nil || o.LastContact == nil {
		var ret int64
		return ret
	}
	return *o.LastContact
}

// GetLastContactOk returns a tuple with the LastContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetLastContactOk() (*int64, bool) {
	if o == nil || o.LastContact == nil {
		return nil, false
	}
	return o.LastContact, true
}

// HasLastContact returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasLastContact() bool {
	if o != nil && o.LastContact != nil {
		return true
	}

	return false
}

// SetLastContact gets a reference to the given int64 and assigns it to the LastContact field.
func (o *JobRegisterResponse) SetLastContact(v int64) {
	o.LastContact = &v
}

// GetLastIndex returns the LastIndex field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetLastIndex() int32 {
	if o == nil || o.LastIndex == nil {
		var ret int32
		return ret
	}
	return *o.LastIndex
}

// GetLastIndexOk returns a tuple with the LastIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetLastIndexOk() (*int32, bool) {
	if o == nil || o.LastIndex == nil {
		return nil, false
	}
	return o.LastIndex, true
}

// HasLastIndex returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasLastIndex() bool {
	if o != nil && o.LastIndex != nil {
		return true
	}

	return false
}

// SetLastIndex gets a reference to the given int32 and assigns it to the LastIndex field.
func (o *JobRegisterResponse) SetLastIndex(v int32) {
	o.LastIndex = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetRequestTime() int64 {
	if o == nil || o.RequestTime == nil {
		var ret int64
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetRequestTimeOk() (*int64, bool) {
	if o == nil || o.RequestTime == nil {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasRequestTime() bool {
	if o != nil && o.RequestTime != nil {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given int64 and assigns it to the RequestTime field.
func (o *JobRegisterResponse) SetRequestTime(v int64) {
	o.RequestTime = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *JobRegisterResponse) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRegisterResponse) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *JobRegisterResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *JobRegisterResponse) SetWarnings(v string) {
	o.Warnings = &v
}

func (o JobRegisterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EvalCreateIndex != nil {
		toSerialize["EvalCreateIndex"] = o.EvalCreateIndex
	}
	if o.EvalID != nil {
		toSerialize["EvalID"] = o.EvalID
	}
	if o.JobModifyIndex != nil {
		toSerialize["JobModifyIndex"] = o.JobModifyIndex
	}
	if o.KnownLeader != nil {
		toSerialize["KnownLeader"] = o.KnownLeader
	}
	if o.LastContact != nil {
		toSerialize["LastContact"] = o.LastContact
	}
	if o.LastIndex != nil {
		toSerialize["LastIndex"] = o.LastIndex
	}
	if o.RequestTime != nil {
		toSerialize["RequestTime"] = o.RequestTime
	}
	if o.Warnings != nil {
		toSerialize["Warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableJobRegisterResponse struct {
	value *JobRegisterResponse
	isSet bool
}

func (v NullableJobRegisterResponse) Get() *JobRegisterResponse {
	return v.value
}

func (v *NullableJobRegisterResponse) Set(val *JobRegisterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRegisterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRegisterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRegisterResponse(val *JobRegisterResponse) *NullableJobRegisterResponse {
	return &NullableJobRegisterResponse{value: val, isSet: true}
}

func (v NullableJobRegisterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRegisterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


