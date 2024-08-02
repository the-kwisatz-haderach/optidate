# PublicHolidayV3Dto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date | [optional] 
**LocalName** | **NullableString** | Local name | 
**Name** | **NullableString** | English name | 
**CountryCode** | **NullableString** | ISO 3166-1 alpha-2 | 
**Fixed** | Pointer to **bool** | Is this public holiday every year on the same date | [optional] 
**Global** | Pointer to **bool** | Is this public holiday in every county (federal state) | [optional] 
**Counties** | Pointer to **[]string** | ISO-3166-2 - Federal states | [optional] 
**LaunchYear** | Pointer to **NullableInt32** | The launch year of the public holiday | [optional] 
**Types** | [**[]HolidayTypes**](HolidayTypes.md) | A list of types the public holiday it is valid | 

## Methods

### NewPublicHolidayV3Dto

`func NewPublicHolidayV3Dto(localName NullableString, name NullableString, countryCode NullableString, types []HolidayTypes, ) *PublicHolidayV3Dto`

NewPublicHolidayV3Dto instantiates a new PublicHolidayV3Dto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicHolidayV3DtoWithDefaults

`func NewPublicHolidayV3DtoWithDefaults() *PublicHolidayV3Dto`

NewPublicHolidayV3DtoWithDefaults instantiates a new PublicHolidayV3Dto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PublicHolidayV3Dto) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PublicHolidayV3Dto) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PublicHolidayV3Dto) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PublicHolidayV3Dto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLocalName

`func (o *PublicHolidayV3Dto) GetLocalName() string`

GetLocalName returns the LocalName field if non-nil, zero value otherwise.

### GetLocalNameOk

`func (o *PublicHolidayV3Dto) GetLocalNameOk() (*string, bool)`

GetLocalNameOk returns a tuple with the LocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalName

`func (o *PublicHolidayV3Dto) SetLocalName(v string)`

SetLocalName sets LocalName field to given value.


### SetLocalNameNil

`func (o *PublicHolidayV3Dto) SetLocalNameNil(b bool)`

 SetLocalNameNil sets the value for LocalName to be an explicit nil

### UnsetLocalName
`func (o *PublicHolidayV3Dto) UnsetLocalName()`

UnsetLocalName ensures that no value is present for LocalName, not even an explicit nil
### GetName

`func (o *PublicHolidayV3Dto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicHolidayV3Dto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicHolidayV3Dto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PublicHolidayV3Dto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PublicHolidayV3Dto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCountryCode

`func (o *PublicHolidayV3Dto) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PublicHolidayV3Dto) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PublicHolidayV3Dto) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *PublicHolidayV3Dto) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *PublicHolidayV3Dto) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetFixed

`func (o *PublicHolidayV3Dto) GetFixed() bool`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *PublicHolidayV3Dto) GetFixedOk() (*bool, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *PublicHolidayV3Dto) SetFixed(v bool)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *PublicHolidayV3Dto) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetGlobal

`func (o *PublicHolidayV3Dto) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *PublicHolidayV3Dto) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *PublicHolidayV3Dto) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *PublicHolidayV3Dto) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetCounties

`func (o *PublicHolidayV3Dto) GetCounties() []string`

GetCounties returns the Counties field if non-nil, zero value otherwise.

### GetCountiesOk

`func (o *PublicHolidayV3Dto) GetCountiesOk() (*[]string, bool)`

GetCountiesOk returns a tuple with the Counties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounties

`func (o *PublicHolidayV3Dto) SetCounties(v []string)`

SetCounties sets Counties field to given value.

### HasCounties

`func (o *PublicHolidayV3Dto) HasCounties() bool`

HasCounties returns a boolean if a field has been set.

### SetCountiesNil

`func (o *PublicHolidayV3Dto) SetCountiesNil(b bool)`

 SetCountiesNil sets the value for Counties to be an explicit nil

### UnsetCounties
`func (o *PublicHolidayV3Dto) UnsetCounties()`

UnsetCounties ensures that no value is present for Counties, not even an explicit nil
### GetLaunchYear

`func (o *PublicHolidayV3Dto) GetLaunchYear() int32`

GetLaunchYear returns the LaunchYear field if non-nil, zero value otherwise.

### GetLaunchYearOk

`func (o *PublicHolidayV3Dto) GetLaunchYearOk() (*int32, bool)`

GetLaunchYearOk returns a tuple with the LaunchYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchYear

`func (o *PublicHolidayV3Dto) SetLaunchYear(v int32)`

SetLaunchYear sets LaunchYear field to given value.

### HasLaunchYear

`func (o *PublicHolidayV3Dto) HasLaunchYear() bool`

HasLaunchYear returns a boolean if a field has been set.

### SetLaunchYearNil

`func (o *PublicHolidayV3Dto) SetLaunchYearNil(b bool)`

 SetLaunchYearNil sets the value for LaunchYear to be an explicit nil

### UnsetLaunchYear
`func (o *PublicHolidayV3Dto) UnsetLaunchYear()`

UnsetLaunchYear ensures that no value is present for LaunchYear, not even an explicit nil
### GetTypes

`func (o *PublicHolidayV3Dto) GetTypes() []HolidayTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PublicHolidayV3Dto) GetTypesOk() (*[]HolidayTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PublicHolidayV3Dto) SetTypes(v []HolidayTypes)`

SetTypes sets Types field to given value.


### SetTypesNil

`func (o *PublicHolidayV3Dto) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *PublicHolidayV3Dto) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


