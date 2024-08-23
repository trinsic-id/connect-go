/*
Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trinsic_api

import (
	"encoding/json"
)

// checks if the CreateSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSessionRequest{}

// CreateSessionRequest struct for CreateSessionRequest
type CreateSessionRequest struct {
	// Whether to immediately launch the identity provider, without invoking the Trinsic Connect Widget UI.                Users will not be shown the Connect Widget; therefore, reuse of Connect credentials, selection of an identity provider, and saving a verification for future reuse  are not available to the end user in this mode.                Sessions created with this option enabled must be created with a `RedirectUrl` specified, and cannot be invoked using the frontend SDK at this time.
	LaunchMethodDirectly *bool `json:"launchMethodDirectly,omitempty"`
	// The list of allowed identity providers. If not specified, all available providers will be allowed.                If `LaunchMethodDirectly` is `true`, this field must be set, and must have only a single entry.  If `LaunchMethodDirectly` is not specified or is `false`, this field may have any number of entries.
	Providers []string `json:"providers,omitempty"`
	// Specific identity attributes to request. If not provided, all available attributes will be requested.
	DisclosedFields *DisclosedFieldsRequest `json:"disclosedFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSessionRequest CreateSessionRequest

// NewCreateSessionRequest instantiates a new CreateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSessionRequest() *CreateSessionRequest {
	this := CreateSessionRequest{}
	return &this
}

// NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSessionRequestWithDefaults() *CreateSessionRequest {
	this := CreateSessionRequest{}
	return &this
}

// GetLaunchMethodDirectly returns the LaunchMethodDirectly field value if set, zero value otherwise.
func (o *CreateSessionRequest) GetLaunchMethodDirectly() bool {
	if o == nil || IsNil(o.LaunchMethodDirectly) {
		var ret bool
		return ret
	}
	return *o.LaunchMethodDirectly
}

// GetLaunchMethodDirectlyOk returns a tuple with the LaunchMethodDirectly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetLaunchMethodDirectlyOk() (*bool, bool) {
	if o == nil || IsNil(o.LaunchMethodDirectly) {
		return nil, false
	}
	return o.LaunchMethodDirectly, true
}

// HasLaunchMethodDirectly returns a boolean if a field has been set.
func (o *CreateSessionRequest) HasLaunchMethodDirectly() bool {
	if o != nil && !IsNil(o.LaunchMethodDirectly) {
		return true
	}

	return false
}

// SetLaunchMethodDirectly gets a reference to the given bool and assigns it to the LaunchMethodDirectly field.
func (o *CreateSessionRequest) SetLaunchMethodDirectly(v bool) {
	o.LaunchMethodDirectly = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *CreateSessionRequest) GetProviders() []string {
	if o == nil || IsNil(o.Providers) {
		var ret []string
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *CreateSessionRequest) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []string and assigns it to the Providers field.
func (o *CreateSessionRequest) SetProviders(v []string) {
	o.Providers = v
}

// GetDisclosedFields returns the DisclosedFields field value if set, zero value otherwise.
func (o *CreateSessionRequest) GetDisclosedFields() DisclosedFieldsRequest {
	if o == nil || IsNil(o.DisclosedFields) {
		var ret DisclosedFieldsRequest
		return ret
	}
	return *o.DisclosedFields
}

// GetDisclosedFieldsOk returns a tuple with the DisclosedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetDisclosedFieldsOk() (*DisclosedFieldsRequest, bool) {
	if o == nil || IsNil(o.DisclosedFields) {
		return nil, false
	}
	return o.DisclosedFields, true
}

// HasDisclosedFields returns a boolean if a field has been set.
func (o *CreateSessionRequest) HasDisclosedFields() bool {
	if o != nil && !IsNil(o.DisclosedFields) {
		return true
	}

	return false
}

// SetDisclosedFields gets a reference to the given DisclosedFieldsRequest and assigns it to the DisclosedFields field.
func (o *CreateSessionRequest) SetDisclosedFields(v DisclosedFieldsRequest) {
	o.DisclosedFields = &v
}

func (o CreateSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LaunchMethodDirectly) {
		toSerialize["launchMethodDirectly"] = o.LaunchMethodDirectly
	}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}
	if !IsNil(o.DisclosedFields) {
		toSerialize["disclosedFields"] = o.DisclosedFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSessionRequest) UnmarshalJSON(data []byte) (err error) {
	varCreateSessionRequest := _CreateSessionRequest{}

	err = json.Unmarshal(data, &varCreateSessionRequest)

	if err != nil {
		return err
	}

	*o = CreateSessionRequest(varCreateSessionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "launchMethodDirectly")
		delete(additionalProperties, "providers")
		delete(additionalProperties, "disclosedFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSessionRequest struct {
	value *CreateSessionRequest
	isSet bool
}

func (v NullableCreateSessionRequest) Get() *CreateSessionRequest {
	return v.value
}

func (v *NullableCreateSessionRequest) Set(val *CreateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSessionRequest(val *CreateSessionRequest) *NullableCreateSessionRequest {
	return &NullableCreateSessionRequest{value: val, isSet: true}
}

func (v NullableCreateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


