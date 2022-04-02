# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | **string** | The serial number of the smart device. | 
**Alias** | Pointer to **string** | An alternate name for the smart device. | [optional] 
**ProductName** | Pointer to **string** | The model name of the product. | [optional] 
**Manufacturer** | Pointer to **string** | The name of the product&#39;s manufacturer. | [optional] 
**OtherData** | Pointer to [**Json**](json.md) |  | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2(serialNumber string, ) *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *InlineObject2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *InlineObject2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *InlineObject2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetAlias

`func (o *InlineObject2) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *InlineObject2) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *InlineObject2) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *InlineObject2) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetProductName

`func (o *InlineObject2) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *InlineObject2) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *InlineObject2) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *InlineObject2) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetManufacturer

`func (o *InlineObject2) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InlineObject2) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InlineObject2) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InlineObject2) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetOtherData

`func (o *InlineObject2) GetOtherData() Json`

GetOtherData returns the OtherData field if non-nil, zero value otherwise.

### GetOtherDataOk

`func (o *InlineObject2) GetOtherDataOk() (*Json, bool)`

GetOtherDataOk returns a tuple with the OtherData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherData

`func (o *InlineObject2) SetOtherData(v Json)`

SetOtherData sets OtherData field to given value.

### HasOtherData

`func (o *InlineObject2) HasOtherData() bool`

HasOtherData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


