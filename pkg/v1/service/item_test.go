package service_test

import (
	"context"
	"github.com/libmonsoon-dev/shop/pkg/v1"
	"testing"

	"github.com/golang/mock/gomock"

	mock_repository "github.com/libmonsoon-dev/shop/pkg/v1/mock/repository"
	mock_unitofwork "github.com/libmonsoon-dev/shop/pkg/v1/mock/unitofwork"
	"github.com/libmonsoon-dev/shop/pkg/v1/repository"
	"github.com/libmonsoon-dev/shop/pkg/v1/service"
	"github.com/libmonsoon-dev/shop/pkg/v1/unitofwork"
)

func TestItemServiceCreate(t *testing.T) {
	defaultGetItemRepository := func() func(ctrl *gomock.Controller) repository.ItemRepository {
		return func(ctrl *gomock.Controller) repository.ItemRepository {
			mock := mock_repository.NewMockItemRepository(ctrl)
			mock.EXPECT().Create(gomock.Eq(context.Background()), gomock.Any()).DoAndReturn(func(ctx context.Context, item *shopv1.Item) error {
				item.ID = 1
				return nil
			})

			return mock
		}
	}

	defaultGetAttributeRepository := func() func(ctrl *gomock.Controller) repository.AttributeRepository {
		var lastID shopv1.ID

		return func(ctrl *gomock.Controller) repository.AttributeRepository {
			mock := mock_repository.NewMockAttributeRepository(ctrl)
			mock.EXPECT().Create(gomock.Eq(context.Background()), gomock.Any()).DoAndReturn(
				func(ctx context.Context, valueType *shopv1.Attribute) error {
					lastID++
					valueType.ID = lastID
					return nil
				},
			).AnyTimes()

			return mock
		}
	}

	defaultGetValueTypeRepository := func() func(ctrl *gomock.Controller) repository.ValueTypeRepository {
		var lastID shopv1.ID

		return func(ctrl *gomock.Controller) repository.ValueTypeRepository {
			mock := mock_repository.NewMockValueTypeRepository(ctrl)
			mock.EXPECT().Save(gomock.Eq(context.Background()), gomock.Any()).DoAndReturn(
				func(ctx context.Context, valueType *shopv1.ValueType) error {
					lastID++
					valueType.ID = lastID
					return nil
				},
			).AnyTimes()

			return mock
		}
	}

	defaultGetUnitOfWorkFactory := func(ctrl *gomock.Controller, unitOfWork unitofwork.ItemCreator) unitofwork.ItemCreatorFactory {
		mock := mock_unitofwork.NewMockItemCreatorFactory(ctrl)
		mock.EXPECT().CreateContext(gomock.Any()).Return(unitOfWork, nil)
		return mock
	}

	defaultGetItemCreator := func(
		ctrl *gomock.Controller,
		itemRepository repository.ItemRepository,
		attributeRepository repository.AttributeRepository,
		valueTypeRepository repository.ValueTypeRepository,
	) unitofwork.ItemCreator {
		mock := mock_unitofwork.NewMockItemCreator(ctrl)
		mock.EXPECT().ItemRepository().Return(itemRepository)
		mock.EXPECT().AttributeRepository().Return(attributeRepository).AnyTimes()
		mock.EXPECT().ValueTypeRepository().Return(valueTypeRepository).AnyTimes()

		mock.EXPECT().Complete(gomock.Any())

		return mock
	}

	tests := []struct {
		name                   string
		getItemRepository      func(*gomock.Controller) repository.ItemRepository
		getAttributeRepository func(*gomock.Controller) repository.AttributeRepository
		getValueTypeRepository func(*gomock.Controller) repository.ValueTypeRepository
		getItemCreator         func(
			*gomock.Controller,
			repository.ItemRepository,
			repository.AttributeRepository,
			repository.ValueTypeRepository,
		) unitofwork.ItemCreator
		getUnitOfWorkFactory func(*gomock.Controller, unitofwork.ItemCreator) unitofwork.ItemCreatorFactory
		ctx     context.Context
		item    shopv1.Item
		wantErr bool
	}{
		{
			name: "Regular create",
			ctx:  context.Background(),
			item: shopv1.Item{
				Name: "Ball",
				Attributes: []shopv1.Attribute{
					{
						ValueType: &shopv1.ValueType{
							Name: "color",
						},
						Value: "red",
					},
					{
						ValueType: &shopv1.ValueType{
							Name: "sport",
						},
						Value: "football",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		if tt.getItemRepository == nil {
			tt.getItemRepository = defaultGetItemRepository()
		}

		if tt.getAttributeRepository == nil {
			tt.getAttributeRepository = defaultGetAttributeRepository()
		}

		if tt.getValueTypeRepository == nil {
			tt.getValueTypeRepository = defaultGetValueTypeRepository()
		}

		if tt.getUnitOfWorkFactory == nil {
			tt.getUnitOfWorkFactory = defaultGetUnitOfWorkFactory
		}

		if tt.getItemCreator == nil {
			tt.getItemCreator = defaultGetItemCreator
		}

		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			itemRepository := tt.getItemRepository(mockCtrl)
			attributeRepository := tt.getAttributeRepository(mockCtrl)
			valueTypeRepository := tt.getValueTypeRepository(mockCtrl)
			unitOfWork := tt.getItemCreator(mockCtrl, itemRepository, attributeRepository, valueTypeRepository)
			unitOfWorkFactory := tt.getUnitOfWorkFactory(mockCtrl, unitOfWork)
			s := service.NewItemService(unitOfWorkFactory)

			item := tt.item
			if err := s.Create(tt.ctx, &item); err != nil {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			for i := range item.Attributes {
				if item.Attributes[i].ItemID != item.ID {
					t.Errorf("item.Attributes[%d].ItemID != item.ID (%v != %v)", i, item.Attributes[i].ItemID, item.ID)
				}

				if item.Attributes[i].ValueTypeID != item.Attributes[i].ValueType.ID {
					t.Errorf(
						"item.Attributes[%d].ValueTypeID != item.Attributes[%d].ValueType.ID (%v != %v)",
						i, i,
						item.Attributes[i].ValueTypeID,
						item.Attributes[i].ValueType.ID,
					)
				}
			}
		})
	}
}
