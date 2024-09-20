/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ContainerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerConfig{}

// ContainerConfig struct for ContainerConfig
type ContainerConfig struct {
	Image string `json:"image"`
	User  string `json:"user"`
}

type _ContainerConfig ContainerConfig

// NewContainerConfig instantiates a new ContainerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerConfig(image string, user string) *ContainerConfig {
	this := ContainerConfig{}
	this.Image = image
	this.User = user
	return &this
}

// NewContainerConfigWithDefaults instantiates a new ContainerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerConfigWithDefaults() *ContainerConfig {
	this := ContainerConfig{}
	return &this
}

// GetImage returns the Image field value
func (o *ContainerConfig) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ContainerConfig) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ContainerConfig) SetImage(v string) {
	o.Image = v
}

// GetUser returns the User field value
func (o *ContainerConfig) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ContainerConfig) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ContainerConfig) SetUser(v string) {
	o.User = v
}

func (o ContainerConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *ContainerConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
		"user",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContainerConfig := _ContainerConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerConfig)

	if err != nil {
		return err
	}

	*o = ContainerConfig(varContainerConfig)

	return err
}

type NullableContainerConfig struct {
	value *ContainerConfig
	isSet bool
}

func (v NullableContainerConfig) Get() *ContainerConfig {
	return v.value
}

func (v *NullableContainerConfig) Set(val *ContainerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerConfig(val *ContainerConfig) *NullableContainerConfig {
	return &NullableContainerConfig{value: val, isSet: true}
}

func (v NullableContainerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
