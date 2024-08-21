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

// checks if the DisclosedFieldsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisclosedFieldsRequest{}

// DisclosedFieldsRequest struct for DisclosedFieldsRequest
type DisclosedFieldsRequest struct {
	IdNumber *bool `json:"idNumber,omitempty"`
	GivenName *bool `json:"givenName,omitempty"`
	FamilyName *bool `json:"familyName,omitempty"`
	Address *bool `json:"address,omitempty"`
	DateOfBirth *bool `json:"dateOfBirth,omitempty"`
	Country *bool `json:"country,omitempty"`
	IssueDate *bool `json:"issueDate,omitempty"`
	ExpirationDate *bool `json:"expirationDate,omitempty"`
	DocumentFront *bool `json:"documentFront,omitempty"`
	DocumentBack *bool `json:"documentBack,omitempty"`
	DocumentPortrait *bool `json:"documentPortrait,omitempty"`
	Selfie *bool `json:"selfie,omitempty"`
}

// NewDisclosedFieldsRequest instantiates a new DisclosedFieldsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisclosedFieldsRequest() *DisclosedFieldsRequest {
	this := DisclosedFieldsRequest{}
	return &this
}

// NewDisclosedFieldsRequestWithDefaults instantiates a new DisclosedFieldsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisclosedFieldsRequestWithDefaults() *DisclosedFieldsRequest {
	this := DisclosedFieldsRequest{}
	return &this
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetIdNumber() bool {
	if o == nil || IsNil(o.IdNumber) {
		var ret bool
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetIdNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.IdNumber) {
		return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasIdNumber() bool {
	if o != nil && !IsNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given bool and assigns it to the IdNumber field.
func (o *DisclosedFieldsRequest) SetIdNumber(v bool) {
	o.IdNumber = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetGivenName() bool {
	if o == nil || IsNil(o.GivenName) {
		var ret bool
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetGivenNameOk() (*bool, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given bool and assigns it to the GivenName field.
func (o *DisclosedFieldsRequest) SetGivenName(v bool) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetFamilyName() bool {
	if o == nil || IsNil(o.FamilyName) {
		var ret bool
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetFamilyNameOk() (*bool, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given bool and assigns it to the FamilyName field.
func (o *DisclosedFieldsRequest) SetFamilyName(v bool) {
	o.FamilyName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetAddress() bool {
	if o == nil || IsNil(o.Address) {
		var ret bool
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given bool and assigns it to the Address field.
func (o *DisclosedFieldsRequest) SetAddress(v bool) {
	o.Address = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetDateOfBirth() bool {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret bool
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetDateOfBirthOk() (*bool, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given bool and assigns it to the DateOfBirth field.
func (o *DisclosedFieldsRequest) SetDateOfBirth(v bool) {
	o.DateOfBirth = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetCountry() bool {
	if o == nil || IsNil(o.Country) {
		var ret bool
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetCountryOk() (*bool, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given bool and assigns it to the Country field.
func (o *DisclosedFieldsRequest) SetCountry(v bool) {
	o.Country = &v
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetIssueDate() bool {
	if o == nil || IsNil(o.IssueDate) {
		var ret bool
		return ret
	}
	return *o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetIssueDateOk() (*bool, bool) {
	if o == nil || IsNil(o.IssueDate) {
		return nil, false
	}
	return o.IssueDate, true
}

// HasIssueDate returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasIssueDate() bool {
	if o != nil && !IsNil(o.IssueDate) {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given bool and assigns it to the IssueDate field.
func (o *DisclosedFieldsRequest) SetIssueDate(v bool) {
	o.IssueDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetExpirationDate() bool {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret bool
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetExpirationDateOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given bool and assigns it to the ExpirationDate field.
func (o *DisclosedFieldsRequest) SetExpirationDate(v bool) {
	o.ExpirationDate = &v
}

// GetDocumentFront returns the DocumentFront field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetDocumentFront() bool {
	if o == nil || IsNil(o.DocumentFront) {
		var ret bool
		return ret
	}
	return *o.DocumentFront
}

// GetDocumentFrontOk returns a tuple with the DocumentFront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetDocumentFrontOk() (*bool, bool) {
	if o == nil || IsNil(o.DocumentFront) {
		return nil, false
	}
	return o.DocumentFront, true
}

// HasDocumentFront returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasDocumentFront() bool {
	if o != nil && !IsNil(o.DocumentFront) {
		return true
	}

	return false
}

// SetDocumentFront gets a reference to the given bool and assigns it to the DocumentFront field.
func (o *DisclosedFieldsRequest) SetDocumentFront(v bool) {
	o.DocumentFront = &v
}

// GetDocumentBack returns the DocumentBack field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetDocumentBack() bool {
	if o == nil || IsNil(o.DocumentBack) {
		var ret bool
		return ret
	}
	return *o.DocumentBack
}

// GetDocumentBackOk returns a tuple with the DocumentBack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetDocumentBackOk() (*bool, bool) {
	if o == nil || IsNil(o.DocumentBack) {
		return nil, false
	}
	return o.DocumentBack, true
}

// HasDocumentBack returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasDocumentBack() bool {
	if o != nil && !IsNil(o.DocumentBack) {
		return true
	}

	return false
}

// SetDocumentBack gets a reference to the given bool and assigns it to the DocumentBack field.
func (o *DisclosedFieldsRequest) SetDocumentBack(v bool) {
	o.DocumentBack = &v
}

// GetDocumentPortrait returns the DocumentPortrait field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetDocumentPortrait() bool {
	if o == nil || IsNil(o.DocumentPortrait) {
		var ret bool
		return ret
	}
	return *o.DocumentPortrait
}

// GetDocumentPortraitOk returns a tuple with the DocumentPortrait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetDocumentPortraitOk() (*bool, bool) {
	if o == nil || IsNil(o.DocumentPortrait) {
		return nil, false
	}
	return o.DocumentPortrait, true
}

// HasDocumentPortrait returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasDocumentPortrait() bool {
	if o != nil && !IsNil(o.DocumentPortrait) {
		return true
	}

	return false
}

// SetDocumentPortrait gets a reference to the given bool and assigns it to the DocumentPortrait field.
func (o *DisclosedFieldsRequest) SetDocumentPortrait(v bool) {
	o.DocumentPortrait = &v
}

// GetSelfie returns the Selfie field value if set, zero value otherwise.
func (o *DisclosedFieldsRequest) GetSelfie() bool {
	if o == nil || IsNil(o.Selfie) {
		var ret bool
		return ret
	}
	return *o.Selfie
}

// GetSelfieOk returns a tuple with the Selfie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosedFieldsRequest) GetSelfieOk() (*bool, bool) {
	if o == nil || IsNil(o.Selfie) {
		return nil, false
	}
	return o.Selfie, true
}

// HasSelfie returns a boolean if a field has been set.
func (o *DisclosedFieldsRequest) HasSelfie() bool {
	if o != nil && !IsNil(o.Selfie) {
		return true
	}

	return false
}

// SetSelfie gets a reference to the given bool and assigns it to the Selfie field.
func (o *DisclosedFieldsRequest) SetSelfie(v bool) {
	o.Selfie = &v
}

func (o DisclosedFieldsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisclosedFieldsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdNumber) {
		toSerialize["idNumber"] = o.IdNumber
	}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.IssueDate) {
		toSerialize["issueDate"] = o.IssueDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.DocumentFront) {
		toSerialize["documentFront"] = o.DocumentFront
	}
	if !IsNil(o.DocumentBack) {
		toSerialize["documentBack"] = o.DocumentBack
	}
	if !IsNil(o.DocumentPortrait) {
		toSerialize["documentPortrait"] = o.DocumentPortrait
	}
	if !IsNil(o.Selfie) {
		toSerialize["selfie"] = o.Selfie
	}
	return toSerialize, nil
}

type NullableDisclosedFieldsRequest struct {
	value *DisclosedFieldsRequest
	isSet bool
}

func (v NullableDisclosedFieldsRequest) Get() *DisclosedFieldsRequest {
	return v.value
}

func (v *NullableDisclosedFieldsRequest) Set(val *DisclosedFieldsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosedFieldsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosedFieldsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosedFieldsRequest(val *DisclosedFieldsRequest) *NullableDisclosedFieldsRequest {
	return &NullableDisclosedFieldsRequest{value: val, isSet: true}
}

func (v NullableDisclosedFieldsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosedFieldsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


