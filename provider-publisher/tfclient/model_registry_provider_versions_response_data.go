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

// RegistryProviderVersionsResponseData struct for RegistryProviderVersionsResponseData
type RegistryProviderVersionsResponseData struct {
	Id         *string                                     `json:"id,omitempty"`
	Type       *string                                     `json:"type,omitempty"`
	Attributes *RegistryProviderVersionsResponseAttributes `json:"attributes,omitempty"`
}

// NewRegistryProviderVersionsResponseData instantiates a new RegistryProviderVersionsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryProviderVersionsResponseData() *RegistryProviderVersionsResponseData {
	this := RegistryProviderVersionsResponseData{}
	return &this
}

// NewRegistryProviderVersionsResponseDataWithDefaults instantiates a new RegistryProviderVersionsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryProviderVersionsResponseDataWithDefaults() *RegistryProviderVersionsResponseData {
	this := RegistryProviderVersionsResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistryProviderVersionsResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryProviderVersionsResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistryProviderVersionsResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistryProviderVersionsResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegistryProviderVersionsResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryProviderVersionsResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegistryProviderVersionsResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegistryProviderVersionsResponseData) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RegistryProviderVersionsResponseData) GetAttributes() RegistryProviderVersionsResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret RegistryProviderVersionsResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryProviderVersionsResponseData) GetAttributesOk() (*RegistryProviderVersionsResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RegistryProviderVersionsResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given RegistryProviderVersionsResponseAttributes and assigns it to the Attributes field.
func (o *RegistryProviderVersionsResponseData) SetAttributes(v RegistryProviderVersionsResponseAttributes) {
	o.Attributes = &v
}

func (o RegistryProviderVersionsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryProviderVersionsResponseData struct {
	value *RegistryProviderVersionsResponseData
	isSet bool
}

func (v NullableRegistryProviderVersionsResponseData) Get() *RegistryProviderVersionsResponseData {
	return v.value
}

func (v *NullableRegistryProviderVersionsResponseData) Set(val *RegistryProviderVersionsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryProviderVersionsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryProviderVersionsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryProviderVersionsResponseData(val *RegistryProviderVersionsResponseData) *NullableRegistryProviderVersionsResponseData {
	return &NullableRegistryProviderVersionsResponseData{value: val, isSet: true}
}

func (v NullableRegistryProviderVersionsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryProviderVersionsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
