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

// InlineResponse20011 struct for InlineResponse20011
type InlineResponse20011 struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewInlineResponse20011 instantiates a new InlineResponse20011 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// NewInlineResponse20011WithDefaults instantiates a new InlineResponse20011 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011WithDefaults() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20011) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20011) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *InlineResponse20011) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o InlineResponse20011) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011 struct {
	value *InlineResponse20011
	isSet bool
}

func (v NullableInlineResponse20011) Get() *InlineResponse20011 {
	return v.value
}

func (v *NullableInlineResponse20011) Set(val *InlineResponse20011) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011(val *InlineResponse20011) *NullableInlineResponse20011 {
	return &NullableInlineResponse20011{value: val, isSet: true}
}

func (v NullableInlineResponse20011) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


