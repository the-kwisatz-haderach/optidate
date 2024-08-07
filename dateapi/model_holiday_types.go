/*
Nager.Date API - V3

Nager.Date is open source software. If you would like to support the project you can award a GitHub star ⭐ or much better <a href='https://github.com/sponsors/nager'>start a sponsorship</a>

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dateapi

import (
	"encoding/json"
	"fmt"
)

// HolidayTypes the model 'HolidayTypes'
type HolidayTypes string

// List of HolidayTypes
const (
	PUBLIC HolidayTypes = "Public"
	BANK HolidayTypes = "Bank"
	SCHOOL HolidayTypes = "School"
	AUTHORITIES HolidayTypes = "Authorities"
	OPTIONAL HolidayTypes = "Optional"
	OBSERVANCE HolidayTypes = "Observance"
)

// All allowed values of HolidayTypes enum
var AllowedHolidayTypesEnumValues = []HolidayTypes{
	"Public",
	"Bank",
	"School",
	"Authorities",
	"Optional",
	"Observance",
}

func (v *HolidayTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HolidayTypes(value)
	for _, existing := range AllowedHolidayTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HolidayTypes", value)
}

// NewHolidayTypesFromValue returns a pointer to a valid HolidayTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHolidayTypesFromValue(v string) (*HolidayTypes, error) {
	ev := HolidayTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HolidayTypes: valid values are %v", v, AllowedHolidayTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HolidayTypes) IsValid() bool {
	for _, existing := range AllowedHolidayTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HolidayTypes value
func (v HolidayTypes) Ptr() *HolidayTypes {
	return &v
}

type NullableHolidayTypes struct {
	value *HolidayTypes
	isSet bool
}

func (v NullableHolidayTypes) Get() *HolidayTypes {
	return v.value
}

func (v *NullableHolidayTypes) Set(val *HolidayTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableHolidayTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableHolidayTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHolidayTypes(val *HolidayTypes) *NullableHolidayTypes {
	return &NullableHolidayTypes{value: val, isSet: true}
}

func (v NullableHolidayTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHolidayTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

