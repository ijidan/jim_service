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

func newGid(db *gorm.DB) gid {
	_gid := gid{}

	_gid.gidDo.UseDB(db)
	_gid.gidDo.UseModel(&gen_model.Gid{})

	tableName := _gid.gidDo.TableName()
	_gid.ALL = field.NewField(tableName, "*")
	_gid.ID = field.NewInt64(tableName, "id")
	_gid.BusinessID = field.NewString(tableName, "business_id")
	_gid.MaxID = field.NewInt64(tableName, "max_id")
	_gid.Step = field.NewInt32(tableName, "step")
	_gid.Description = field.NewString(tableName, "description")
	_gid.CreateTime = field.NewTime(tableName, "create_time")
	_gid.UpdateTime = field.NewTime(tableName, "update_time")

	_gid.fillFieldMap()

	return _gid
}

type gid struct {
	gidDo gidDo

	ALL         field.Field
	ID          field.Int64
	BusinessID  field.String
	MaxID       field.Int64
	Step        field.Int32
	Description field.String
	CreateTime  field.Time
	UpdateTime  field.Time

	fieldMap map[string]field.Expr
}

func (g gid) As(alias string) *gid {
	g.gidDo.DO = *(g.gidDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.ID = field.NewInt64(alias, "id")
	g.BusinessID = field.NewString(alias, "business_id")
	g.MaxID = field.NewInt64(alias, "max_id")
	g.Step = field.NewInt32(alias, "step")
	g.Description = field.NewString(alias, "description")
	g.CreateTime = field.NewTime(alias, "create_time")
	g.UpdateTime = field.NewTime(alias, "update_time")

	g.fillFieldMap()

	return &g
}

func (g *gid) WithContext(ctx context.Context) *gidDo { return g.gidDo.WithContext(ctx) }

func (g gid) TableName() string { return g.gidDo.TableName() }

func (g *gid) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *gid) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["business_id"] = g.BusinessID
	g.fieldMap["max_id"] = g.MaxID
	g.fieldMap["step"] = g.Step
	g.fieldMap["description"] = g.Description
	g.fieldMap["create_time"] = g.CreateTime
	g.fieldMap["update_time"] = g.UpdateTime
}

func (g gid) clone(db *gorm.DB) gid {
	g.gidDo.ReplaceDB(db)
	return g
}

type gidDo struct{ gen.DO }

func (g gidDo) Debug() *gidDo {
	return g.withDO(g.DO.Debug())
}

func (g gidDo) WithContext(ctx context.Context) *gidDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gidDo) Clauses(conds ...clause.Expression) *gidDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gidDo) Not(conds ...gen.Condition) *gidDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gidDo) Or(conds ...gen.Condition) *gidDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gidDo) Select(conds ...field.Expr) *gidDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gidDo) Where(conds ...gen.Condition) *gidDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gidDo) Order(conds ...field.Expr) *gidDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gidDo) Distinct(cols ...field.Expr) *gidDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gidDo) Omit(cols ...field.Expr) *gidDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gidDo) Join(table schema.Tabler, on ...field.Expr) *gidDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gidDo) LeftJoin(table schema.Tabler, on ...field.Expr) *gidDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gidDo) RightJoin(table schema.Tabler, on ...field.Expr) *gidDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gidDo) Group(cols ...field.Expr) *gidDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gidDo) Having(conds ...gen.Condition) *gidDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gidDo) Limit(limit int) *gidDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gidDo) Offset(offset int) *gidDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gidDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *gidDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gidDo) Unscoped() *gidDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gidDo) Create(values ...*gen_model.Gid) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gidDo) CreateInBatches(values []*gen_model.Gid, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gidDo) Save(values ...*gen_model.Gid) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gidDo) First() (*gen_model.Gid, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Gid), nil
	}
}

func (g gidDo) Take() (*gen_model.Gid, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Gid), nil
	}
}

func (g gidDo) Last() (*gen_model.Gid, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Gid), nil
	}
}

func (g gidDo) Find() ([]*gen_model.Gid, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.Gid), err
}

func (g gidDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.Gid, err error) {
	buf := make([]*gen_model.Gid, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gidDo) FindInBatches(result *[]*gen_model.Gid, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gidDo) Attrs(attrs ...field.AssignExpr) *gidDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gidDo) Assign(attrs ...field.AssignExpr) *gidDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gidDo) Joins(field field.RelationField) *gidDo {
	return g.withDO(g.DO.Joins(field))
}

func (g gidDo) Preload(field field.RelationField) *gidDo {
	return g.withDO(g.DO.Preload(field))
}

func (g gidDo) FirstOrInit() (*gen_model.Gid, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Gid), nil
	}
}

func (g gidDo) FirstOrCreate() (*gen_model.Gid, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Gid), nil
	}
}

func (g gidDo) FindByPage(offset int, limit int) (result []*gen_model.Gid, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g gidDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *gidDo) withDO(do gen.Dao) *gidDo {
	g.DO = *do.(*gen.DO)
	return g
}
