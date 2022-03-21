/*
Terraform Private Registry

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tfclient

import (
	"encoding/json"
)

// GpgKeyData struct for GpgKeyData
type GpgKeyData struct {
	Attributes *GpgKeyDataAttributes `json:"attributes,omitempty"`
}

// NewGpgKeyData instantiates a new GpgKeyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpgKeyData() *GpgKeyData {
	this := GpgKeyData{}
	return &this
}

// NewGpgKeyDataWithDefaults instantiates a new GpgKeyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpgKeyDataWithDefaults() *GpgKeyData {
	this := GpgKeyData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GpgKeyData) GetAttributes() GpgKeyDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret GpgKeyDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpgKeyData) GetAttributesOk() (*GpgKeyDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GpgKeyData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GpgKeyDataAttributes and assigns it to the Attributes field.
func (o *GpgKeyData) SetAttributes(v GpgKeyDataAttributes) {
	o.Attributes = &v
}

func (o GpgKeyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableGpgKeyData struct {
	value *GpgKeyData
	isSet bool
}

func (v NullableGpgKeyData) Get() *GpgKeyData {
	return v.value
}

func (v *NullableGpgKeyData) Set(val *GpgKeyData) {
	v.value = val
	v.isSet = true
}

func (v NullableGpgKeyData) IsSet() bool {
	return v.isSet
}

func (v *NullableGpgKeyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpgKeyData(val *GpgKeyData) *NullableGpgKeyData {
	return &NullableGpgKeyData{value: val, isSet: true}
}

func (v NullableGpgKeyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpgKeyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
