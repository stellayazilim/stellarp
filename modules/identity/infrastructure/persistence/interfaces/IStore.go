package interfaces

type IStore[TEntity interface{}] interface {
	FindRealms() []TEntity
	FindOneRealm(TEntity) TEntity
	CreateRealm(TEntity) error
	UpdateRealm(TEntity) error
	DeleteRealm(TEntity) error
}
