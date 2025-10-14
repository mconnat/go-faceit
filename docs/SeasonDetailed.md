# SeasonDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Divisions** | Pointer to [**[]Division**](Division.md) | The divisions of the given season. | [optional] 
**Season** | Pointer to [**Season**](Season.md) |  | [optional] 

## Methods

### NewSeasonDetailed

`func NewSeasonDetailed() *SeasonDetailed`

NewSeasonDetailed instantiates a new SeasonDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonDetailedWithDefaults

`func NewSeasonDetailedWithDefaults() *SeasonDetailed`

NewSeasonDetailedWithDefaults instantiates a new SeasonDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisions

`func (o *SeasonDetailed) GetDivisions() []Division`

GetDivisions returns the Divisions field if non-nil, zero value otherwise.

### GetDivisionsOk

`func (o *SeasonDetailed) GetDivisionsOk() (*[]Division, bool)`

GetDivisionsOk returns a tuple with the Divisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisions

`func (o *SeasonDetailed) SetDivisions(v []Division)`

SetDivisions sets Divisions field to given value.

### HasDivisions

`func (o *SeasonDetailed) HasDivisions() bool`

HasDivisions returns a boolean if a field has been set.

### GetSeason

`func (o *SeasonDetailed) GetSeason() Season`

GetSeason returns the Season field if non-nil, zero value otherwise.

### GetSeasonOk

`func (o *SeasonDetailed) GetSeasonOk() (*Season, bool)`

GetSeasonOk returns a tuple with the Season field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeason

`func (o *SeasonDetailed) SetSeason(v Season)`

SetSeason sets Season field to given value.

### HasSeason

`func (o *SeasonDetailed) HasSeason() bool`

HasSeason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


