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

// RegistryProvidersResponse struct for RegistryProvidersResponse
type RegistryProvidersResponse struct {
	Data []CreateRegistryProviderResponseData `json:"data,omitempty"`
}

// NewRegistryProvidersResponse instantiates a new RegistryProvidersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryProvidersResponse() *RegistryProvidersResponse {
	this := RegistryProvidersResponse{}
	return &this
}

// NewRegistryProvidersResponseWithDefaults instantiates a new RegistryProvidersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryProvidersResponseWithDefaults() *RegistryProvidersResponse {
	this := RegistryProvidersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RegistryProvidersResponse) GetData() []CreateRegistryProviderResponseData {
	if o == nil || o.Data == nil {
		var ret []CreateRegistryProviderResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryProvidersResponse) GetDataOk() ([]CreateRegistryProviderResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RegistryProvidersResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []CreateRegistryProviderResponseData and assigns it to the Data field.
func (o *RegistryProvidersResponse) SetData(v []CreateRegistryProviderResponseData) {
	o.Data = v
}

func (o RegistryProvidersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryProvidersResponse struct {
	value *RegistryProvidersResponse
	isSet bool
}

func (v NullableRegistryProvidersResponse) Get() *RegistryProvidersResponse {
	return v.value
}

func (v *NullableRegistryProvidersResponse) Set(val *RegistryProvidersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryProvidersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryProvidersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryProvidersResponse(val *RegistryProvidersResponse) *NullableRegistryProvidersResponse {
	return &NullableRegistryProvidersResponse{value: val, isSet: true}
}

func (v NullableRegistryProvidersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryProvidersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
