package shopapi

type Item struct {
	ID             ID
	Name           string
	Description    string
	Image          string
	QuantityRemain int

	Attributes []Attribute
}
