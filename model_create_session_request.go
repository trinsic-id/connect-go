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
	// Whether to immediately launch the identity provider, without invoking the Trinsic Widget UI.                Users will not be shown the Widget; therefore, reuse of credentials, selection of an identity provider, and saving a verification for future reuse  are not available to the end user in this mode.                Sessions created with this option enabled must be created with a `RedirectUrl` specified, and cannot be invoked using the frontend SDK at this time.
	LaunchProviderDirectly *bool `json:"launchProviderDirectly,omitempty"`
	// Whether to enable Trinsic's \"Remember Me\" feature, which allows users to save their credentials for future use.                This option is only relevant when `LaunchProviderDirectly` is unspecified or set to `false`.  If `LaunchProviderDirectly` is `true`, this field must be unspecified or set to `false`.                If this field is set to `true`, then:    - The user will be prompted to authenticate with their phone number at the start of the flow    - If the user has previously saved a verification for reuse with Trinsic, they will be offered the ability to reuse it    - After the user has verified their identity (and if the identity provider in question supports it), they will be prompted to save their credentials for future use                If this field is set to `false`, then:    - The user will not be prompted to authenticate with their phone number at the start of the flow.      - Instead, the user will be immediately shown the list of available providers    - The user will not be offered the ability to reuse a previously-saved Trinsic credential    - After the user has verified their identity, they will not be prompted to save their credentials for future use      - Instead, they will immediately return to your product
	EnableRememberMe *bool `json:"enableRememberMe,omitempty"`
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

// GetLaunchProviderDirectly returns the LaunchProviderDirectly field value if set, zero value otherwise.
func (o *CreateSessionRequest) GetLaunchProviderDirectly() bool {
	if o == nil || IsNil(o.LaunchProviderDirectly) {
		var ret bool
		return ret
	}
	return *o.LaunchProviderDirectly
}

// GetLaunchProviderDirectlyOk returns a tuple with the LaunchProviderDirectly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetLaunchProviderDirectlyOk() (*bool, bool) {
	if o == nil || IsNil(o.LaunchProviderDirectly) {
		return nil, false
	}
	return o.LaunchProviderDirectly, true
}

// HasLaunchProviderDirectly returns a boolean if a field has been set.
func (o *CreateSessionRequest) HasLaunchProviderDirectly() bool {
	if o != nil && !IsNil(o.LaunchProviderDirectly) {
		return true
	}

	return false
}

// SetLaunchProviderDirectly gets a reference to the given bool and assigns it to the LaunchProviderDirectly field.
func (o *CreateSessionRequest) SetLaunchProviderDirectly(v bool) {
	o.LaunchProviderDirectly = &v
}

// GetEnableRememberMe returns the EnableRememberMe field value if set, zero value otherwise.
func (o *CreateSessionRequest) GetEnableRememberMe() bool {
	if o == nil || IsNil(o.EnableRememberMe) {
		var ret bool
		return ret
	}
	return *o.EnableRememberMe
}

// GetEnableRememberMeOk returns a tuple with the EnableRememberMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetEnableRememberMeOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRememberMe) {
		return nil, false
	}
	return o.EnableRememberMe, true
}

// HasEnableRememberMe returns a boolean if a field has been set.
func (o *CreateSessionRequest) HasEnableRememberMe() bool {
	if o != nil && !IsNil(o.EnableRememberMe) {
		return true
	}

	return false
}

// SetEnableRememberMe gets a reference to the given bool and assigns it to the EnableRememberMe field.
func (o *CreateSessionRequest) SetEnableRememberMe(v bool) {
	o.EnableRememberMe = &v
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
	if !IsNil(o.LaunchProviderDirectly) {
		toSerialize["launchProviderDirectly"] = o.LaunchProviderDirectly
	}
	if !IsNil(o.EnableRememberMe) {
		toSerialize["enableRememberMe"] = o.EnableRememberMe
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
		delete(additionalProperties, "launchProviderDirectly")
		delete(additionalProperties, "enableRememberMe")
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


