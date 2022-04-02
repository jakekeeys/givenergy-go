/*
GivEnergy API Documentation (v1.5.1)

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package givenergy

import (
	"encoding/json"
)

// InlineResponse20010 struct for InlineResponse20010
type InlineResponse20010 struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewInlineResponse20010 instantiates a new InlineResponse20010 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20010() *InlineResponse20010 {
	this := InlineResponse20010{}
	return &this
}

// NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20010WithDefaults() *InlineResponse20010 {
	this := InlineResponse20010{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20010) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20010) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20010) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *InlineResponse20010) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o InlineResponse20010) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20010 struct {
	value *InlineResponse20010
	isSet bool
}

func (v NullableInlineResponse20010) Get() *InlineResponse20010 {
	return v.value
}

func (v *NullableInlineResponse20010) Set(val *InlineResponse20010) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20010) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20010) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20010(val *InlineResponse20010) *NullableInlineResponse20010 {
	return &NullableInlineResponse20010{value: val, isSet: true}
}

func (v NullableInlineResponse20010) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20010) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

