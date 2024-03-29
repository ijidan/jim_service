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

func newDeviceAck(db *gorm.DB) deviceAck {
	_deviceAck := deviceAck{}

	_deviceAck.deviceAckDo.UseDB(db)
	_deviceAck.deviceAckDo.UseModel(&gen_model.DeviceAck{})

	tableName := _deviceAck.deviceAckDo.TableName()
	_deviceAck.ALL = field.NewField(tableName, "*")
	_deviceAck.ID = field.NewUint64(tableName, "id")
	_deviceAck.DeviceID = field.NewInt64(tableName, "device_id")
	_deviceAck.Ack = field.NewInt64(tableName, "ack")
	_deviceAck.CreateTime = field.NewTime(tableName, "create_time")
	_deviceAck.UpdateTime = field.NewTime(tableName, "update_time")
	_deviceAck.Device = deviceAckBelongsToDevice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Device", "gen_model.Device"),
	}

	_deviceAck.fillFieldMap()

	return _deviceAck
}

type deviceAck struct {
	deviceAckDo deviceAckDo

	ALL        field.Field
	ID         field.Uint64
	DeviceID   field.Int64
	Ack        field.Int64
	CreateTime field.Time
	UpdateTime field.Time
	Device     deviceAckBelongsToDevice

	fieldMap map[string]field.Expr
}

func (d deviceAck) As(alias string) *deviceAck {
	d.deviceAckDo.DO = *(d.deviceAckDo.As(alias).(*gen.DO))

	d.ALL = field.NewField(alias, "*")
	d.ID = field.NewUint64(alias, "id")
	d.DeviceID = field.NewInt64(alias, "device_id")
	d.Ack = field.NewInt64(alias, "ack")
	d.CreateTime = field.NewTime(alias, "create_time")
	d.UpdateTime = field.NewTime(alias, "update_time")

	d.fillFieldMap()

	return &d
}

func (d *deviceAck) WithContext(ctx context.Context) *deviceAckDo {
	return d.deviceAckDo.WithContext(ctx)
}

func (d deviceAck) TableName() string { return d.deviceAckDo.TableName() }

func (d *deviceAck) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := d.fieldMap[fieldName]
	return field, ok
}

func (d *deviceAck) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 6)
	d.fieldMap["id"] = d.ID
	d.fieldMap["device_id"] = d.DeviceID
	d.fieldMap["ack"] = d.Ack
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["update_time"] = d.UpdateTime

}

func (d deviceAck) clone(db *gorm.DB) deviceAck {
	d.deviceAckDo.ReplaceDB(db)
	return d
}

type deviceAckBelongsToDevice struct {
	db *gorm.DB

	field.RelationField
}

func (a deviceAckBelongsToDevice) Where(conds ...field.Expr) *deviceAckBelongsToDevice {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a deviceAckBelongsToDevice) WithContext(ctx context.Context) *deviceAckBelongsToDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceAckBelongsToDevice) Model(m *gen_model.DeviceAck) *deviceAckBelongsToDeviceTx {
	return &deviceAckBelongsToDeviceTx{a.db.Model(m).Association(a.Name())}
}

type deviceAckBelongsToDeviceTx struct{ tx *gorm.Association }

func (a deviceAckBelongsToDeviceTx) Find() (result *gen_model.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceAckBelongsToDeviceTx) Append(values ...*gen_model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceAckBelongsToDeviceTx) Replace(values ...*gen_model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceAckBelongsToDeviceTx) Delete(values ...*gen_model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceAckBelongsToDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceAckBelongsToDeviceTx) Count() int64 {
	return a.tx.Count()
}

type deviceAckDo struct{ gen.DO }

func (d deviceAckDo) Debug() *deviceAckDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceAckDo) WithContext(ctx context.Context) *deviceAckDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceAckDo) Clauses(conds ...clause.Expression) *deviceAckDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceAckDo) Not(conds ...gen.Condition) *deviceAckDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceAckDo) Or(conds ...gen.Condition) *deviceAckDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceAckDo) Select(conds ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceAckDo) Where(conds ...gen.Condition) *deviceAckDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceAckDo) Order(conds ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceAckDo) Distinct(cols ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceAckDo) Omit(cols ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceAckDo) Join(table schema.Tabler, on ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceAckDo) LeftJoin(table schema.Tabler, on ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceAckDo) RightJoin(table schema.Tabler, on ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceAckDo) Group(cols ...field.Expr) *deviceAckDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceAckDo) Having(conds ...gen.Condition) *deviceAckDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceAckDo) Limit(limit int) *deviceAckDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceAckDo) Offset(offset int) *deviceAckDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceAckDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *deviceAckDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceAckDo) Unscoped() *deviceAckDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceAckDo) Create(values ...*gen_model.DeviceAck) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceAckDo) CreateInBatches(values []*gen_model.DeviceAck, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceAckDo) Save(values ...*gen_model.DeviceAck) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceAckDo) First() (*gen_model.DeviceAck, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.DeviceAck), nil
	}
}

func (d deviceAckDo) Take() (*gen_model.DeviceAck, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.DeviceAck), nil
	}
}

func (d deviceAckDo) Last() (*gen_model.DeviceAck, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.DeviceAck), nil
	}
}

func (d deviceAckDo) Find() ([]*gen_model.DeviceAck, error) {
	result, err := d.DO.Find()
	return result.([]*gen_model.DeviceAck), err
}

func (d deviceAckDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.DeviceAck, err error) {
	buf := make([]*gen_model.DeviceAck, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceAckDo) FindInBatches(result *[]*gen_model.DeviceAck, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceAckDo) Attrs(attrs ...field.AssignExpr) *deviceAckDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceAckDo) Assign(attrs ...field.AssignExpr) *deviceAckDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceAckDo) Joins(field field.RelationField) *deviceAckDo {
	return d.withDO(d.DO.Joins(field))
}

func (d deviceAckDo) Preload(field field.RelationField) *deviceAckDo {
	return d.withDO(d.DO.Preload(field))
}

func (d deviceAckDo) FirstOrInit() (*gen_model.DeviceAck, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.DeviceAck), nil
	}
}

func (d deviceAckDo) FirstOrCreate() (*gen_model.DeviceAck, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.DeviceAck), nil
	}
}

func (d deviceAckDo) FindByPage(offset int, limit int) (result []*gen_model.DeviceAck, count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	result, err = d.Offset(offset).Limit(limit).Find()
	return
}

func (d deviceAckDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d *deviceAckDo) withDO(do gen.Dao) *deviceAckDo {
	d.DO = *do.(*gen.DO)
	return d
}
