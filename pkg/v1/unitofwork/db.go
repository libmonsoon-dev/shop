package unitofwork

import "github.com/libmonsoon-dev/shop/pkg/v1/factory"

type DB = factory.ContextInterface[Tx]
