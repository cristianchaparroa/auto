package schema

type IEntityRepository interface {
	FindByName(name string)
}
type EntityRepository struct {
}
