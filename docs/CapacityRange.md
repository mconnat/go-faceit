# CapacityRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitBytes** | Pointer to **int64** | LimitBytes specifies that a volume must not be bigger than this. The value of 0 indicates an unspecified maximum | [optional] 
**RequiredBytes** | Pointer to **int64** | RequiredBytes specifies that a volume must be at least this big. The value of 0 indicates an unspecified minimum. | [optional] 

## Methods

### NewCapacityRange

`func NewCapacityRange() *CapacityRange`

NewCapacityRange instantiates a new CapacityRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityRangeWithDefaults

`func NewCapacityRangeWithDefaults() *CapacityRange`

NewCapacityRangeWithDefaults instantiates a new CapacityRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitBytes

`func (o *CapacityRange) GetLimitBytes() int64`

GetLimitBytes returns the LimitBytes field if non-nil, zero value otherwise.

### GetLimitBytesOk

`func (o *CapacityRange) GetLimitBytesOk() (*int64, bool)`

GetLimitBytesOk returns a tuple with the LimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBytes

`func (o *CapacityRange) SetLimitBytes(v int64)`

SetLimitBytes sets LimitBytes field to given value.

### HasLimitBytes

`func (o *CapacityRange) HasLimitBytes() bool`

HasLimitBytes returns a boolean if a field has been set.

### GetRequiredBytes

`func (o *CapacityRange) GetRequiredBytes() int64`

GetRequiredBytes returns the RequiredBytes field if non-nil, zero value otherwise.

### GetRequiredBytesOk

`func (o *CapacityRange) GetRequiredBytesOk() (*int64, bool)`

GetRequiredBytesOk returns a tuple with the RequiredBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBytes

`func (o *CapacityRange) SetRequiredBytes(v int64)`

SetRequiredBytes sets RequiredBytes field to given value.

### HasRequiredBytes

`func (o *CapacityRange) HasRequiredBytes() bool`

HasRequiredBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


