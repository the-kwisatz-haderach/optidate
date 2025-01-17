/*
Nager.Date API - V3

Nager.Date is open source software. If you would like to support the project you can award a GitHub star ⭐ or much better <a href='https://github.com/sponsors/nager'>start a sponsorship</a>

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dateapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CountryInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryInfoDto{}

// CountryInfoDto CountryInfo Dto
type CountryInfoDto struct {
	// CommonName
	CommonName NullableString `json:"commonName"`
	// OfficialName
	OfficialName NullableString `json:"officialName"`
	// CountryCode
	CountryCode NullableString `json:"countryCode"`
	// Region
	Region NullableString `json:"region"`
	// Country Borders
	Borders []CountryInfoDto `json:"borders,omitempty"`
}

type _CountryInfoDto CountryInfoDto

// NewCountryInfoDto instantiates a new CountryInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryInfoDto(commonName NullableString, officialName NullableString, countryCode NullableString, region NullableString) *CountryInfoDto {
	this := CountryInfoDto{}
	this.CommonName = commonName
	this.OfficialName = officialName
	this.CountryCode = countryCode
	this.Region = region
	return &this
}

// NewCountryInfoDtoWithDefaults instantiates a new CountryInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryInfoDtoWithDefaults() *CountryInfoDto {
	this := CountryInfoDto{}
	return &this
}

// GetCommonName returns the CommonName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CountryInfoDto) GetCommonName() string {
	if o == nil || o.CommonName.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommonName.Get()
}

// GetCommonNameOk returns a tuple with the CommonName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfoDto) GetCommonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommonName.Get(), o.CommonName.IsSet()
}

// SetCommonName sets field value
func (o *CountryInfoDto) SetCommonName(v string) {
	o.CommonName.Set(&v)
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CountryInfoDto) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfoDto) GetOfficialNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *CountryInfoDto) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetCountryCode returns the CountryCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CountryInfoDto) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfoDto) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// SetCountryCode sets field value
func (o *CountryInfoDto) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CountryInfoDto) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfoDto) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *CountryInfoDto) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetBorders returns the Borders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryInfoDto) GetBorders() []CountryInfoDto {
	if o == nil {
		var ret []CountryInfoDto
		return ret
	}
	return o.Borders
}

// GetBordersOk returns a tuple with the Borders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfoDto) GetBordersOk() ([]CountryInfoDto, bool) {
	if o == nil || IsNil(o.Borders) {
		return nil, false
	}
	return o.Borders, true
}

// HasBorders returns a boolean if a field has been set.
func (o *CountryInfoDto) HasBorders() bool {
	if o != nil && !IsNil(o.Borders) {
		return true
	}

	return false
}

// SetBorders gets a reference to the given []CountryInfoDto and assigns it to the Borders field.
func (o *CountryInfoDto) SetBorders(v []CountryInfoDto) {
	o.Borders = v
}

func (o CountryInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commonName"] = o.CommonName.Get()
	toSerialize["officialName"] = o.OfficialName.Get()
	toSerialize["countryCode"] = o.CountryCode.Get()
	toSerialize["region"] = o.Region.Get()
	if o.Borders != nil {
		toSerialize["borders"] = o.Borders
	}
	return toSerialize, nil
}

func (o *CountryInfoDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commonName",
		"officialName",
		"countryCode",
		"region",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCountryInfoDto := _CountryInfoDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCountryInfoDto)

	if err != nil {
		return err
	}

	*o = CountryInfoDto(varCountryInfoDto)

	return err
}

type NullableCountryInfoDto struct {
	value *CountryInfoDto
	isSet bool
}

func (v NullableCountryInfoDto) Get() *CountryInfoDto {
	return v.value
}

func (v *NullableCountryInfoDto) Set(val *CountryInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryInfoDto(val *CountryInfoDto) *NullableCountryInfoDto {
	return &NullableCountryInfoDto{value: val, isSet: true}
}

func (v NullableCountryInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


