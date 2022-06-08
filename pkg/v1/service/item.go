package service

import (
	"context"
	"fmt"

	shopapi "github.com/libmonsoon-dev/shop/pkg/v1/api"
	"github.com/libmonsoon-dev/shop/pkg/v1/unitofwork"
)

func NewItemService(unitOfWorkFactory unitofwork.ItemCreatorFactory) *ItemService {
	return &ItemService{
		unitOfWorkFactory: unitOfWorkFactory,
	}
}

type ItemService struct {
	unitOfWorkFactory unitofwork.ItemCreatorFactory
}

func (s *ItemService) Create(ctx context.Context, item *shopapi.Item) (err error) {
	unitOfWork, err := s.unitOfWorkFactory.CreateContext(ctx)
	if err != nil {
		return fmt.Errorf("create unit of work: %w", err)
	}

	defer unitOfWork.Complete(&err)

	err = unitOfWork.ItemRepository().Create(ctx, item)
	if err != nil {
		return fmt.Errorf("create item %+v: %w", item, err)
	}

	for i := range item.Attributes {
		err = unitOfWork.ValueTypeRepository().Save(ctx, item.Attributes[i].ValueType)
		if err != nil {
			return fmt.Errorf("save value type %+v: %w", item.Attributes[i].ValueType, err)
		}

		item.Attributes[i].ItemID = item.ID
		item.Attributes[i].ValueTypeID = item.Attributes[i].ValueType.ID

		err = unitOfWork.AttributeRepository().Create(ctx, &item.Attributes[i])
		if err != nil {
			return fmt.Errorf("create attribute %+v: %w", item.Attributes[i], err)
		}
	}

	return
}
