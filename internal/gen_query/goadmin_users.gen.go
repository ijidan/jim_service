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

func newGoadminUsers(db *gorm.DB) goadminUsers {
	_goadminUsers := goadminUsers{}

	_goadminUsers.goadminUsersDo.UseDB(db)
	_goadminUsers.goadminUsersDo.UseModel(&gen_model.GoadminUsers{})

	tableName := _goadminUsers.goadminUsersDo.TableName()
	_goadminUsers.ALL = field.NewField(tableName, "*")
	_goadminUsers.ID = field.NewInt32(tableName, "id")
	_goadminUsers.Username = field.NewString(tableName, "username")
	_goadminUsers.Password = field.NewString(tableName, "password")
	_goadminUsers.Name = field.NewString(tableName, "name")
	_goadminUsers.Avatar = field.NewString(tableName, "avatar")
	_goadminUsers.RememberToken = field.NewString(tableName, "remember_token")
	_goadminUsers.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminUsers.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminUsers.fillFieldMap()

	return _goadminUsers
}

type goadminUsers struct {
	goadminUsersDo goadminUsersDo

	ALL           field.Field
	ID            field.Int32
	Username      field.String
	Password      field.String
	Name          field.String
	Avatar        field.String
	RememberToken field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (g goadminUsers) As(alias string) *goadminUsers {
	g.goadminUsersDo.DO = *(g.goadminUsersDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.ID = field.NewInt32(alias, "id")
	g.Username = field.NewString(alias, "username")
	g.Password = field.NewString(alias, "password")
	g.Name = field.NewString(alias, "name")
	g.Avatar = field.NewString(alias, "avatar")
	g.RememberToken = field.NewString(alias, "remember_token")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")

	g.fillFieldMap()

	return &g
}

func (g *goadminUsers) WithContext(ctx context.Context) *goadminUsersDo {
	return g.goadminUsersDo.WithContext(ctx)
}

func (g goadminUsers) TableName() string { return g.goadminUsersDo.TableName() }

func (g *goadminUsers) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *goadminUsers) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 8)
	g.fieldMap["id"] = g.ID
	g.fieldMap["username"] = g.Username
	g.fieldMap["password"] = g.Password
	g.fieldMap["name"] = g.Name
	g.fieldMap["avatar"] = g.Avatar
	g.fieldMap["remember_token"] = g.RememberToken
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminUsers) clone(db *gorm.DB) goadminUsers {
	g.goadminUsersDo.ReplaceDB(db)
	return g
}

type goadminUsersDo struct{ gen.DO }

func (g goadminUsersDo) Debug() *goadminUsersDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminUsersDo) WithContext(ctx context.Context) *goadminUsersDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminUsersDo) Clauses(conds ...clause.Expression) *goadminUsersDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminUsersDo) Not(conds ...gen.Condition) *goadminUsersDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminUsersDo) Or(conds ...gen.Condition) *goadminUsersDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminUsersDo) Select(conds ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminUsersDo) Where(conds ...gen.Condition) *goadminUsersDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminUsersDo) Order(conds ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminUsersDo) Distinct(cols ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminUsersDo) Omit(cols ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminUsersDo) Join(table schema.Tabler, on ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminUsersDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminUsersDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminUsersDo) Group(cols ...field.Expr) *goadminUsersDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminUsersDo) Having(conds ...gen.Condition) *goadminUsersDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminUsersDo) Limit(limit int) *goadminUsersDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminUsersDo) Offset(offset int) *goadminUsersDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminUsersDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminUsersDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminUsersDo) Unscoped() *goadminUsersDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminUsersDo) Create(values ...*gen_model.GoadminUsers) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminUsersDo) CreateInBatches(values []*gen_model.GoadminUsers, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminUsersDo) Save(values ...*gen_model.GoadminUsers) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminUsersDo) First() (*gen_model.GoadminUsers, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUsers), nil
	}
}

func (g goadminUsersDo) Take() (*gen_model.GoadminUsers, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUsers), nil
	}
}

func (g goadminUsersDo) Last() (*gen_model.GoadminUsers, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUsers), nil
	}
}

func (g goadminUsersDo) Find() ([]*gen_model.GoadminUsers, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.GoadminUsers), err
}

func (g goadminUsersDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.GoadminUsers, err error) {
	buf := make([]*gen_model.GoadminUsers, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminUsersDo) FindInBatches(result *[]*gen_model.GoadminUsers, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminUsersDo) Attrs(attrs ...field.AssignExpr) *goadminUsersDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminUsersDo) Assign(attrs ...field.AssignExpr) *goadminUsersDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminUsersDo) Joins(field field.RelationField) *goadminUsersDo {
	return g.withDO(g.DO.Joins(field))
}

func (g goadminUsersDo) Preload(field field.RelationField) *goadminUsersDo {
	return g.withDO(g.DO.Preload(field))
}

func (g goadminUsersDo) FirstOrInit() (*gen_model.GoadminUsers, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUsers), nil
	}
}

func (g goadminUsersDo) FirstOrCreate() (*gen_model.GoadminUsers, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUsers), nil
	}
}

func (g goadminUsersDo) FindByPage(offset int, limit int) (result []*gen_model.GoadminUsers, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g goadminUsersDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *goadminUsersDo) withDO(do gen.Dao) *goadminUsersDo {
	g.DO = *do.(*gen.DO)
	return g
}
