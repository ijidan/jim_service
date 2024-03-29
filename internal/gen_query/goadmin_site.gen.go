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

func newGoadminSite(db *gorm.DB) goadminSite {
	_goadminSite := goadminSite{}

	_goadminSite.goadminSiteDo.UseDB(db)
	_goadminSite.goadminSiteDo.UseModel(&gen_model.GoadminSite{})

	tableName := _goadminSite.goadminSiteDo.TableName()
	_goadminSite.ALL = field.NewField(tableName, "*")
	_goadminSite.ID = field.NewInt32(tableName, "id")
	_goadminSite.Key = field.NewString(tableName, "key")
	_goadminSite.Value = field.NewString(tableName, "value")
	_goadminSite.Description = field.NewString(tableName, "description")
	_goadminSite.State = field.NewInt32(tableName, "state")
	_goadminSite.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminSite.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminSite.fillFieldMap()

	return _goadminSite
}

type goadminSite struct {
	goadminSiteDo goadminSiteDo

	ALL         field.Field
	ID          field.Int32
	Key         field.String
	Value       field.String
	Description field.String
	State       field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (g goadminSite) As(alias string) *goadminSite {
	g.goadminSiteDo.DO = *(g.goadminSiteDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.ID = field.NewInt32(alias, "id")
	g.Key = field.NewString(alias, "key")
	g.Value = field.NewString(alias, "value")
	g.Description = field.NewString(alias, "description")
	g.State = field.NewInt32(alias, "state")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")

	g.fillFieldMap()

	return &g
}

func (g *goadminSite) WithContext(ctx context.Context) *goadminSiteDo {
	return g.goadminSiteDo.WithContext(ctx)
}

func (g goadminSite) TableName() string { return g.goadminSiteDo.TableName() }

func (g *goadminSite) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *goadminSite) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["key"] = g.Key
	g.fieldMap["value"] = g.Value
	g.fieldMap["description"] = g.Description
	g.fieldMap["state"] = g.State
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminSite) clone(db *gorm.DB) goadminSite {
	g.goadminSiteDo.ReplaceDB(db)
	return g
}

type goadminSiteDo struct{ gen.DO }

func (g goadminSiteDo) Debug() *goadminSiteDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminSiteDo) WithContext(ctx context.Context) *goadminSiteDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminSiteDo) Clauses(conds ...clause.Expression) *goadminSiteDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminSiteDo) Not(conds ...gen.Condition) *goadminSiteDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminSiteDo) Or(conds ...gen.Condition) *goadminSiteDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminSiteDo) Select(conds ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminSiteDo) Where(conds ...gen.Condition) *goadminSiteDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminSiteDo) Order(conds ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminSiteDo) Distinct(cols ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminSiteDo) Omit(cols ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminSiteDo) Join(table schema.Tabler, on ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminSiteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminSiteDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminSiteDo) Group(cols ...field.Expr) *goadminSiteDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminSiteDo) Having(conds ...gen.Condition) *goadminSiteDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminSiteDo) Limit(limit int) *goadminSiteDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminSiteDo) Offset(offset int) *goadminSiteDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminSiteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminSiteDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminSiteDo) Unscoped() *goadminSiteDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminSiteDo) Create(values ...*gen_model.GoadminSite) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminSiteDo) CreateInBatches(values []*gen_model.GoadminSite, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminSiteDo) Save(values ...*gen_model.GoadminSite) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminSiteDo) First() (*gen_model.GoadminSite, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminSite), nil
	}
}

func (g goadminSiteDo) Take() (*gen_model.GoadminSite, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminSite), nil
	}
}

func (g goadminSiteDo) Last() (*gen_model.GoadminSite, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminSite), nil
	}
}

func (g goadminSiteDo) Find() ([]*gen_model.GoadminSite, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.GoadminSite), err
}

func (g goadminSiteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.GoadminSite, err error) {
	buf := make([]*gen_model.GoadminSite, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminSiteDo) FindInBatches(result *[]*gen_model.GoadminSite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminSiteDo) Attrs(attrs ...field.AssignExpr) *goadminSiteDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminSiteDo) Assign(attrs ...field.AssignExpr) *goadminSiteDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminSiteDo) Joins(field field.RelationField) *goadminSiteDo {
	return g.withDO(g.DO.Joins(field))
}

func (g goadminSiteDo) Preload(field field.RelationField) *goadminSiteDo {
	return g.withDO(g.DO.Preload(field))
}

func (g goadminSiteDo) FirstOrInit() (*gen_model.GoadminSite, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminSite), nil
	}
}

func (g goadminSiteDo) FirstOrCreate() (*gen_model.GoadminSite, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminSite), nil
	}
}

func (g goadminSiteDo) FindByPage(offset int, limit int) (result []*gen_model.GoadminSite, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g goadminSiteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *goadminSiteDo) withDO(do gen.Dao) *goadminSiteDo {
	g.DO = *do.(*gen.DO)
	return g
}
