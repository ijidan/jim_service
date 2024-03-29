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

func newGoadminRoleMenu(db *gorm.DB) goadminRoleMenu {
	_goadminRoleMenu := goadminRoleMenu{}

	_goadminRoleMenu.goadminRoleMenuDo.UseDB(db)
	_goadminRoleMenu.goadminRoleMenuDo.UseModel(&gen_model.GoadminRoleMenu{})

	tableName := _goadminRoleMenu.goadminRoleMenuDo.TableName()
	_goadminRoleMenu.ALL = field.NewField(tableName, "*")
	_goadminRoleMenu.RoleID = field.NewInt32(tableName, "role_id")
	_goadminRoleMenu.MenuID = field.NewInt32(tableName, "menu_id")
	_goadminRoleMenu.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminRoleMenu.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminRoleMenu.fillFieldMap()

	return _goadminRoleMenu
}

type goadminRoleMenu struct {
	goadminRoleMenuDo goadminRoleMenuDo

	ALL       field.Field
	RoleID    field.Int32
	MenuID    field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (g goadminRoleMenu) As(alias string) *goadminRoleMenu {
	g.goadminRoleMenuDo.DO = *(g.goadminRoleMenuDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.RoleID = field.NewInt32(alias, "role_id")
	g.MenuID = field.NewInt32(alias, "menu_id")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")

	g.fillFieldMap()

	return &g
}

func (g *goadminRoleMenu) WithContext(ctx context.Context) *goadminRoleMenuDo {
	return g.goadminRoleMenuDo.WithContext(ctx)
}

func (g goadminRoleMenu) TableName() string { return g.goadminRoleMenuDo.TableName() }

func (g *goadminRoleMenu) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *goadminRoleMenu) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 4)
	g.fieldMap["role_id"] = g.RoleID
	g.fieldMap["menu_id"] = g.MenuID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminRoleMenu) clone(db *gorm.DB) goadminRoleMenu {
	g.goadminRoleMenuDo.ReplaceDB(db)
	return g
}

type goadminRoleMenuDo struct{ gen.DO }

func (g goadminRoleMenuDo) Debug() *goadminRoleMenuDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminRoleMenuDo) WithContext(ctx context.Context) *goadminRoleMenuDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminRoleMenuDo) Clauses(conds ...clause.Expression) *goadminRoleMenuDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminRoleMenuDo) Not(conds ...gen.Condition) *goadminRoleMenuDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminRoleMenuDo) Or(conds ...gen.Condition) *goadminRoleMenuDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminRoleMenuDo) Select(conds ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminRoleMenuDo) Where(conds ...gen.Condition) *goadminRoleMenuDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminRoleMenuDo) Order(conds ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminRoleMenuDo) Distinct(cols ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminRoleMenuDo) Omit(cols ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminRoleMenuDo) Join(table schema.Tabler, on ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminRoleMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminRoleMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminRoleMenuDo) Group(cols ...field.Expr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminRoleMenuDo) Having(conds ...gen.Condition) *goadminRoleMenuDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminRoleMenuDo) Limit(limit int) *goadminRoleMenuDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminRoleMenuDo) Offset(offset int) *goadminRoleMenuDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminRoleMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminRoleMenuDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminRoleMenuDo) Unscoped() *goadminRoleMenuDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminRoleMenuDo) Create(values ...*gen_model.GoadminRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminRoleMenuDo) CreateInBatches(values []*gen_model.GoadminRoleMenu, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminRoleMenuDo) Save(values ...*gen_model.GoadminRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminRoleMenuDo) First() (*gen_model.GoadminRoleMenu, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoleMenu), nil
	}
}

func (g goadminRoleMenuDo) Take() (*gen_model.GoadminRoleMenu, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoleMenu), nil
	}
}

func (g goadminRoleMenuDo) Last() (*gen_model.GoadminRoleMenu, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoleMenu), nil
	}
}

func (g goadminRoleMenuDo) Find() ([]*gen_model.GoadminRoleMenu, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.GoadminRoleMenu), err
}

func (g goadminRoleMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.GoadminRoleMenu, err error) {
	buf := make([]*gen_model.GoadminRoleMenu, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminRoleMenuDo) FindInBatches(result *[]*gen_model.GoadminRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminRoleMenuDo) Attrs(attrs ...field.AssignExpr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminRoleMenuDo) Assign(attrs ...field.AssignExpr) *goadminRoleMenuDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminRoleMenuDo) Joins(field field.RelationField) *goadminRoleMenuDo {
	return g.withDO(g.DO.Joins(field))
}

func (g goadminRoleMenuDo) Preload(field field.RelationField) *goadminRoleMenuDo {
	return g.withDO(g.DO.Preload(field))
}

func (g goadminRoleMenuDo) FirstOrInit() (*gen_model.GoadminRoleMenu, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoleMenu), nil
	}
}

func (g goadminRoleMenuDo) FirstOrCreate() (*gen_model.GoadminRoleMenu, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminRoleMenu), nil
	}
}

func (g goadminRoleMenuDo) FindByPage(offset int, limit int) (result []*gen_model.GoadminRoleMenu, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g goadminRoleMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *goadminRoleMenuDo) withDO(do gen.Dao) *goadminRoleMenuDo {
	g.DO = *do.(*gen.DO)
	return g
}
