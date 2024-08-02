# CountryInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | **NullableString** | CommonName | 
**OfficialName** | **NullableString** | OfficialName | 
**CountryCode** | **NullableString** | CountryCode | 
**Region** | **NullableString** | Region | 
**Borders** | Pointer to [**[]CountryInfoDto**](CountryInfoDto.md) | Country Borders | [optional] 

## Methods

### NewCountryInfoDto

`func NewCountryInfoDto(commonName NullableString, officialName NullableString, countryCode NullableString, region NullableString, ) *CountryInfoDto`

NewCountryInfoDto instantiates a new CountryInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryInfoDtoWithDefaults

`func NewCountryInfoDtoWithDefaults() *CountryInfoDto`

NewCountryInfoDtoWithDefaults instantiates a new CountryInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *CountryInfoDto) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CountryInfoDto) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CountryInfoDto) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### SetCommonNameNil

`func (o *CountryInfoDto) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CountryInfoDto) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetOfficialName

`func (o *CountryInfoDto) GetOfficialName() string`

GetOfficialName returns the OfficialName field if non-nil, zero value otherwise.

### GetOfficialNameOk

`func (o *CountryInfoDto) GetOfficialNameOk() (*string, bool)`

GetOfficialNameOk returns a tuple with the OfficialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialName

`func (o *CountryInfoDto) SetOfficialName(v string)`

SetOfficialName sets OfficialName field to given value.


### SetOfficialNameNil

`func (o *CountryInfoDto) SetOfficialNameNil(b bool)`

 SetOfficialNameNil sets the value for OfficialName to be an explicit nil

### UnsetOfficialName
`func (o *CountryInfoDto) UnsetOfficialName()`

UnsetOfficialName ensures that no value is present for OfficialName, not even an explicit nil
### GetCountryCode

`func (o *CountryInfoDto) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CountryInfoDto) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CountryInfoDto) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *CountryInfoDto) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *CountryInfoDto) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetRegion

`func (o *CountryInfoDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CountryInfoDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CountryInfoDto) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *CountryInfoDto) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CountryInfoDto) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetBorders

`func (o *CountryInfoDto) GetBorders() []CountryInfoDto`

GetBorders returns the Borders field if non-nil, zero value otherwise.

### GetBordersOk

`func (o *CountryInfoDto) GetBordersOk() (*[]CountryInfoDto, bool)`

GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorders

`func (o *CountryInfoDto) SetBorders(v []CountryInfoDto)`

SetBorders sets Borders field to given value.

### HasBorders

`func (o *CountryInfoDto) HasBorders() bool`

HasBorders returns a boolean if a field has been set.

### SetBordersNil

`func (o *CountryInfoDto) SetBordersNil(b bool)`

 SetBordersNil sets the value for Borders to be an explicit nil

### UnsetBorders
`func (o *CountryInfoDto) UnsetBorders()`

UnsetBorders ensures that no value is present for Borders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


