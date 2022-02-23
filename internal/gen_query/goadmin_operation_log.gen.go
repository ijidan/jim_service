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

func newGoadminOperationLog(db *gorm.DB) goadminOperationLog {
	_goadminOperationLog := goadminOperationLog{}

	_goadminOperationLog.goadminOperationLogDo.UseDB(db)
	_goadminOperationLog.goadminOperationLogDo.UseModel(&gen_model.GoadminOperationLog{})

	tableName := _goadminOperationLog.goadminOperationLogDo.TableName()
	_goadminOperationLog.ALL = field.NewField(tableName, "*")
	_goadminOperationLog.ID = field.NewInt32(tableName, "id")
	_goadminOperationLog.UserID = field.NewInt32(tableName, "user_id")
	_goadminOperationLog.Path = field.NewString(tableName, "path")
	_goadminOperationLog.Method = field.NewString(tableName, "method")
	_goadminOperationLog.IP = field.NewString(tableName, "ip")
	_goadminOperationLog.Input = field.NewString(tableName, "input")
	_goadminOperationLog.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminOperationLog.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminOperationLog.fillFieldMap()

	return _goadminOperationLog
}

type goadminOperationLog struct {
	goadminOperationLogDo goadminOperationLogDo

	ALL       field.Field
	ID        field.Int32
	UserID    field.Int32
	Path      field.String
	Method    field.String
	IP        field.String
	Input     field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (g goadminOperationLog) As(alias string) *goadminOperationLog {
	g.goadminOperationLogDo.DO = *(g.goadminOperationLogDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.ID = field.NewInt32(alias, "id")
	g.UserID = field.NewInt32(alias, "user_id")
	g.Path = field.NewString(alias, "path")
	g.Method = field.NewString(alias, "method")
	g.IP = field.NewString(alias, "ip")
	g.Input = field.NewString(alias, "input")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")

	g.fillFieldMap()

	return &g
}

func (g *goadminOperationLog) WithContext(ctx context.Context) *goadminOperationLogDo {
	return g.goadminOperationLogDo.WithContext(ctx)
}

func (g goadminOperationLog) TableName() string { return g.goadminOperationLogDo.TableName() }

func (g *goadminOperationLog) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *goadminOperationLog) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 8)
	g.fieldMap["id"] = g.ID
	g.fieldMap["user_id"] = g.UserID
	g.fieldMap["path"] = g.Path
	g.fieldMap["method"] = g.Method
	g.fieldMap["ip"] = g.IP
	g.fieldMap["input"] = g.Input
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminOperationLog) clone(db *gorm.DB) goadminOperationLog {
	g.goadminOperationLogDo.ReplaceDB(db)
	return g
}

type goadminOperationLogDo struct{ gen.DO }

func (g goadminOperationLogDo) Debug() *goadminOperationLogDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminOperationLogDo) WithContext(ctx context.Context) *goadminOperationLogDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminOperationLogDo) Clauses(conds ...clause.Expression) *goadminOperationLogDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminOperationLogDo) Not(conds ...gen.Condition) *goadminOperationLogDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminOperationLogDo) Or(conds ...gen.Condition) *goadminOperationLogDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminOperationLogDo) Select(conds ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminOperationLogDo) Where(conds ...gen.Condition) *goadminOperationLogDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminOperationLogDo) Order(conds ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminOperationLogDo) Distinct(cols ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminOperationLogDo) Omit(cols ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminOperationLogDo) Join(table schema.Tabler, on ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminOperationLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminOperationLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminOperationLogDo) Group(cols ...field.Expr) *goadminOperationLogDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminOperationLogDo) Having(conds ...gen.Condition) *goadminOperationLogDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminOperationLogDo) Limit(limit int) *goadminOperationLogDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminOperationLogDo) Offset(offset int) *goadminOperationLogDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminOperationLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminOperationLogDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminOperationLogDo) Unscoped() *goadminOperationLogDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminOperationLogDo) Create(values ...*gen_model.GoadminOperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminOperationLogDo) CreateInBatches(values []*gen_model.GoadminOperationLog, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminOperationLogDo) Save(values ...*gen_model.GoadminOperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminOperationLogDo) First() (*gen_model.GoadminOperationLog, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) Take() (*gen_model.GoadminOperationLog, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) Last() (*gen_model.GoadminOperationLog, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) Find() ([]*gen_model.GoadminOperationLog, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.GoadminOperationLog), err
}

func (g goadminOperationLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.GoadminOperationLog, err error) {
	buf := make([]*gen_model.GoadminOperationLog, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminOperationLogDo) FindInBatches(result *[]*gen_model.GoadminOperationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminOperationLogDo) Attrs(attrs ...field.AssignExpr) *goadminOperationLogDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminOperationLogDo) Assign(attrs ...field.AssignExpr) *goadminOperationLogDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminOperationLogDo) Joins(field field.RelationField) *goadminOperationLogDo {
	return g.withDO(g.DO.Joins(field))
}

func (g goadminOperationLogDo) Preload(field field.RelationField) *goadminOperationLogDo {
	return g.withDO(g.DO.Preload(field))
}

func (g goadminOperationLogDo) FirstOrInit() (*gen_model.GoadminOperationLog, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) FirstOrCreate() (*gen_model.GoadminOperationLog, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) FindByPage(offset int, limit int) (result []*gen_model.GoadminOperationLog, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g goadminOperationLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *goadminOperationLogDo) withDO(do gen.Dao) *goadminOperationLogDo {
	g.DO = *do.(*gen.DO)
	return g
}
