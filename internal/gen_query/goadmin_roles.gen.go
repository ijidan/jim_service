// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_query

import (
	"context"
	"jim_service/internal/gen_model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newGoadminRoles(db *gorm.DB) goadminRoles {
	_goadminRoles := goadminRoles{}

	_goadminRoles.goadminRolesDo.UseDB(db)
	_goadminRoles.goadminRolesDo.UseModel(&gen_model.GoadminRoles{})

	tableName := _goadminRoles.goadminRolesDo.TableName()
	_goadminRoles.ALL = field.NewField(tableName, "*")
	_goadminRoles.ID = field.NewInt32(tableName, "id")
	_goadminRoles.Name = field.NewString(tableName, "name")
	_goadminRoles.Slug = field.NewString(tableName, "slug")
	_goadminRoles.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminRoles.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminRoles.fillFieldMap()

	return _goadminRoles
}

type goadminRoles struct {
	goadminRolesDo goadminRolesDo

	ALL       field.Field
	ID        field.Int32
	Name      field.String
	Slug      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (g goadminRoles) As(alias string) *goadminRoles {
	g.goadminRolesDo.DO = *(g.goadminRolesDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.ID = field.NewInt32(alias, "id")
	g.Name = field.NewString(alias, "name")
	g.Slug = field.NewString(alias, "slug")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")

	g.fillFieldMap()

	return &g
}

func (g *goadminRoles) WithContext(ctx context.Context) *goadminRolesDo {
	return g.goadminRolesDo.WithContext(ctx)
}

func (g goadminRoles) TableName() string { return g.goadminRolesDo.TableName() }

func (g *goadminRoles) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *goadminRoles) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 5)
	g.fieldMap["id"] = g.ID
	g.fieldMap["name"] = g.Name
	g.fieldMap["slug"] = g.Slug
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminRoles) clone(db *gorm.DB) goadminRoles {
	g.goadminRolesDo.ReplaceDB(db)
	return g
}

type goadminRolesDo struct{ gen.DO }

func (g goadminRolesDo) Debug() *goadminRolesDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminRolesDo) WithContext(ctx context.Context) *goadminRolesDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminRolesDo) Clauses(conds ...clause.Expression) *goadminRolesDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminRolesDo) Not(conds ...gen.Condition) *goadminRolesDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminRolesDo) Or(conds ...gen.Condition) *goadminRolesDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminRolesDo) Select(conds ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminRolesDo) Where(conds ...gen.Condition) *goadminRolesDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminRolesDo) Order(conds ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminRolesDo) Distinct(cols ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminRolesDo) Omit(cols ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminRolesDo) Join(table schema.Tabler, on ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminRolesDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminRolesDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminRolesDo) Group(cols ...field.Expr) *goadminRolesDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminRolesDo) Having(conds ...gen.Condition) *goadminRolesDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminRolesDo) Limit(limit int) *goadminRolesDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminRolesDo) Offset(offset int) *goadminRolesDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminRolesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminRolesDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminRolesDo) Unscoped() *goadminRolesDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminRolesDo) Create(values ...*gen_model.GoadminRoles) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminRolesDo) CreateInBatches(values []*gen_model.GoadminRoles, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminRolesDo) Save(values ...*gen_model.GoadminRoles) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminRolesDo) First() (*gen_model.GoadminRoles, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoles), nil
	}
}

func (g goadminRolesDo) Take() (*gen_model.GoadminRoles, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoles), nil
	}
}

func (g goadminRolesDo) Last() (*gen_model.GoadminRoles, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoles), nil
	}
}

func (g goadminRolesDo) Find() ([]*gen_model.GoadminRoles, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.GoadminRoles), err
}

func (g goadminRolesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.GoadminRoles, err error) {
	buf := make([]*gen_model.GoadminRoles, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminRolesDo) FindInBatches(result *[]*gen_model.GoadminRoles, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminRolesDo) Attrs(attrs ...field.AssignExpr) *goadminRolesDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminRolesDo) Assign(attrs ...field.AssignExpr) *goadminRolesDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminRolesDo) Joins(field field.RelationField) *goadminRolesDo {
	return g.withDO(g.DO.Joins(field))
}

func (g goadminRolesDo) Preload(field field.RelationField) *goadminRolesDo {
	return g.withDO(g.DO.Preload(field))
}

func (g goadminRolesDo) FirstOrInit() (*gen_model.GoadminRoles, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoles), nil
	}
}

func (g goadminRolesDo) FirstOrCreate() (*gen_model.GoadminRoles, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoles), nil
	}
}

func (g goadminRolesDo) FindByPage(offset int, limit int) (result []*gen_model.GoadminRoles, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g goadminRolesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *goadminRolesDo) withDO(do gen.Dao) *goadminRolesDo {
	g.DO = *do.(*gen.DO)
	return g
}
