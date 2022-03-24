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

func newGoadminPermissions(db *gorm.DB) goadminPermissions {
	_goadminPermissions := goadminPermissions{}

	_goadminPermissions.goadminPermissionsDo.UseDB(db)
	_goadminPermissions.goadminPermissionsDo.UseModel(&gen_model.GoadminPermissions{})

	tableName := _goadminPermissions.goadminPermissionsDo.TableName()
	_goadminPermissions.ALL = field.NewField(tableName, "*")
	_goadminPermissions.ID = field.NewInt32(tableName, "id")
	_goadminPermissions.Name = field.NewString(tableName, "name")
	_goadminPermissions.Slug = field.NewString(tableName, "slug")
	_goadminPermissions.HTTPMethod = field.NewString(tableName, "http_method")
	_goadminPermissions.HTTPPath = field.NewString(tableName, "http_path")
	_goadminPermissions.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminPermissions.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminPermissions.fillFieldMap()

	return _goadminPermissions
}

type goadminPermissions struct {
	goadminPermissionsDo goadminPermissionsDo

	ALL        field.Field
	ID         field.Int32
	Name       field.String
	Slug       field.String
	HTTPMethod field.String
	HTTPPath   field.String
	CreatedAt  field.Time
	UpdatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (g goadminPermissions) As(alias string) *goadminPermissions {
	g.goadminPermissionsDo.DO = *(g.goadminPermissionsDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.ID = field.NewInt32(alias, "id")
	g.Name = field.NewString(alias, "name")
	g.Slug = field.NewString(alias, "slug")
	g.HTTPMethod = field.NewString(alias, "http_method")
	g.HTTPPath = field.NewString(alias, "http_path")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")

	g.fillFieldMap()

	return &g
}

func (g *goadminPermissions) WithContext(ctx context.Context) *goadminPermissionsDo {
	return g.goadminPermissionsDo.WithContext(ctx)
}

func (g goadminPermissions) TableName() string { return g.goadminPermissionsDo.TableName() }

func (g *goadminPermissions) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *goadminPermissions) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["name"] = g.Name
	g.fieldMap["slug"] = g.Slug
	g.fieldMap["http_method"] = g.HTTPMethod
	g.fieldMap["http_path"] = g.HTTPPath
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminPermissions) clone(db *gorm.DB) goadminPermissions {
	g.goadminPermissionsDo.ReplaceDB(db)
	return g
}

type goadminPermissionsDo struct{ gen.DO }

func (g goadminPermissionsDo) Debug() *goadminPermissionsDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminPermissionsDo) WithContext(ctx context.Context) *goadminPermissionsDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminPermissionsDo) Clauses(conds ...clause.Expression) *goadminPermissionsDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminPermissionsDo) Not(conds ...gen.Condition) *goadminPermissionsDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminPermissionsDo) Or(conds ...gen.Condition) *goadminPermissionsDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminPermissionsDo) Select(conds ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminPermissionsDo) Where(conds ...gen.Condition) *goadminPermissionsDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminPermissionsDo) Order(conds ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminPermissionsDo) Distinct(cols ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminPermissionsDo) Omit(cols ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminPermissionsDo) Join(table schema.Tabler, on ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminPermissionsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminPermissionsDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminPermissionsDo) Group(cols ...field.Expr) *goadminPermissionsDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminPermissionsDo) Having(conds ...gen.Condition) *goadminPermissionsDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminPermissionsDo) Limit(limit int) *goadminPermissionsDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminPermissionsDo) Offset(offset int) *goadminPermissionsDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminPermissionsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminPermissionsDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminPermissionsDo) Unscoped() *goadminPermissionsDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminPermissionsDo) Create(values ...*gen_model.GoadminPermissions) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminPermissionsDo) CreateInBatches(values []*gen_model.GoadminPermissions, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminPermissionsDo) Save(values ...*gen_model.GoadminPermissions) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminPermissionsDo) First() (*gen_model.GoadminPermissions, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminPermissions), nil
	}
}

func (g goadminPermissionsDo) Take() (*gen_model.GoadminPermissions, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminPermissions), nil
	}
}

func (g goadminPermissionsDo) Last() (*gen_model.GoadminPermissions, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminPermissions), nil
	}
}

func (g goadminPermissionsDo) Find() ([]*gen_model.GoadminPermissions, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.GoadminPermissions), err
}

func (g goadminPermissionsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.GoadminPermissions, err error) {
	buf := make([]*gen_model.GoadminPermissions, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminPermissionsDo) FindInBatches(result *[]*gen_model.GoadminPermissions, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminPermissionsDo) Attrs(attrs ...field.AssignExpr) *goadminPermissionsDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminPermissionsDo) Assign(attrs ...field.AssignExpr) *goadminPermissionsDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminPermissionsDo) Joins(field field.RelationField) *goadminPermissionsDo {
	return g.withDO(g.DO.Joins(field))
}

func (g goadminPermissionsDo) Preload(field field.RelationField) *goadminPermissionsDo {
	return g.withDO(g.DO.Preload(field))
}

func (g goadminPermissionsDo) FirstOrInit() (*gen_model.GoadminPermissions, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminPermissions), nil
	}
}

func (g goadminPermissionsDo) FirstOrCreate() (*gen_model.GoadminPermissions, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminPermissions), nil
	}
}

func (g goadminPermissionsDo) FindByPage(offset int, limit int) (result []*gen_model.GoadminPermissions, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g goadminPermissionsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *goadminPermissionsDo) withDO(do gen.Dao) *goadminPermissionsDo {
	g.DO = *do.(*gen.DO)
	return g
}
