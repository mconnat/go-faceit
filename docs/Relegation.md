# Relegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsecutiveLosses** | Pointer to **int64** | Consecutive losses needed for a player to get relegated to the previous tier. | [optional] 

## Methods

### NewRelegation

`func NewRelegation() *Relegation`

NewRelegation instantiates a new Relegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelegationWithDefaults

`func NewRelegationWithDefaults() *Relegation`

NewRelegationWithDefaults instantiates a new Relegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsecutiveLosses

`func (o *Relegation) GetConsecutiveLosses() int64`

GetConsecutiveLosses returns the ConsecutiveLosses field if non-nil, zero value otherwise.

### GetConsecutiveLossesOk

`func (o *Relegation) GetConsecutiveLossesOk() (*int64, bool)`

GetConsecutiveLossesOk returns a tuple with the ConsecutiveLosses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveLosses

`func (o *Relegation) SetConsecutiveLosses(v int64)`

SetConsecutiveLosses sets ConsecutiveLosses field to given value.

### HasConsecutiveLosses

`func (o *Relegation) HasConsecutiveLosses() bool`

HasConsecutiveLosses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


