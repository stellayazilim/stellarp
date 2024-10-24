package interfaces

import "context"

type IStore[TEntity interface{}] interface {
	Find(ctx context.Context) []TEntity
	FindOne(context.Context, TEntity) TEntity
	Create(context.Context, TEntity) error
	Update(context.Context, TEntity) error
	Delete(context.Context, TEntity) error
}
