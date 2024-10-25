package interfaces

import "context"

type IStore[TEntity interface{}] interface {
	Find(ctx context.Context) []TEntity
	FindOne(ctx context.Context, entity TEntity) TEntity
	Create(ctx context.Context, entity TEntity) error
	Update(ctx context.Context, entity TEntity) error
	Delete(ctx context.Context, entity TEntity) error
}
