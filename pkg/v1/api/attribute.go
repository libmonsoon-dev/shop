package shopapi

type Attribute struct {
	ID          ID
	ItemID      ID
	ValueTypeID ID
	Value       string

	ValueType *ValueType
}
