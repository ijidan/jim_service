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

func newDevice(db *gorm.DB) device {
	_device := device{}

	_device.deviceDo.UseDB(db)
	_device.deviceDo.UseModel(&gen_model.Device{})

	tableName := _device.deviceDo.TableName()
	_device.ALL = field.NewField(tableName, "*")
	_device.ID = field.NewUint64(tableName, "id")
	_device.DeviceID = field.NewInt64(tableName, "device_id")
	_device.AppID = field.NewInt64(tableName, "app_id")
	_device.UserID = field.NewInt64(tableName, "user_id")
	_device.Type = field.NewInt32(tableName, "type")
	_device.Brand = field.NewString(tableName, "brand")
	_device.Model = field.NewString(tableName, "model")
	_device.SystemVersion = field.NewString(tableName, "system_version")
	_device.SdkVersion = field.NewString(tableName, "sdk_version")
	_device.Status = field.NewInt32(tableName, "status")
	_device.ConnAddr = field.NewString(tableName, "conn_addr")
	_device.ConnFd = field.NewInt64(tableName, "conn_fd")
	_device.CreateTime = field.NewTime(tableName, "create_time")
	_device.UpdateTime = field.NewTime(tableName, "update_time")

	_device.fillFieldMap()

	return _device
}

type device struct {
	deviceDo deviceDo

	ALL           field.Field
	ID            field.Uint64
	DeviceID      field.Int64
	AppID         field.Int64
	UserID        field.Int64
	Type          field.Int32
	Brand         field.String
	Model         field.String
	SystemVersion field.String
	SdkVersion    field.String
	Status        field.Int32
	ConnAddr      field.String
	ConnFd        field.Int64
	CreateTime    field.Time
	UpdateTime    field.Time

	fieldMap map[string]field.Expr
}

func (d device) As(alias string) *device {
	d.deviceDo.DO = *(d.deviceDo.As(alias).(*gen.DO))

	d.ALL = field.NewField(alias, "*")
	d.ID = field.NewUint64(alias, "id")
	d.DeviceID = field.NewInt64(alias, "device_id")
	d.AppID = field.NewInt64(alias, "app_id")
	d.UserID = field.NewInt64(alias, "user_id")
	d.Type = field.NewInt32(alias, "type")
	d.Brand = field.NewString(alias, "brand")
	d.Model = field.NewString(alias, "model")
	d.SystemVersion = field.NewString(alias, "system_version")
	d.SdkVersion = field.NewString(alias, "sdk_version")
	d.Status = field.NewInt32(alias, "status")
	d.ConnAddr = field.NewString(alias, "conn_addr")
	d.ConnFd = field.NewInt64(alias, "conn_fd")
	d.CreateTime = field.NewTime(alias, "create_time")
	d.UpdateTime = field.NewTime(alias, "update_time")

	d.fillFieldMap()

	return &d
}

func (d *device) WithContext(ctx context.Context) *deviceDo { return d.deviceDo.WithContext(ctx) }

func (d device) TableName() string { return d.deviceDo.TableName() }

func (d *device) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := d.fieldMap[fieldName]
	return field, ok
}

func (d *device) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 14)
	d.fieldMap["id"] = d.ID
	d.fieldMap["device_id"] = d.DeviceID
	d.fieldMap["app_id"] = d.AppID
	d.fieldMap["user_id"] = d.UserID
	d.fieldMap["type"] = d.Type
	d.fieldMap["brand"] = d.Brand
	d.fieldMap["model"] = d.Model
	d.fieldMap["system_version"] = d.SystemVersion
	d.fieldMap["sdk_version"] = d.SdkVersion
	d.fieldMap["status"] = d.Status
	d.fieldMap["conn_addr"] = d.ConnAddr
	d.fieldMap["conn_fd"] = d.ConnFd
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["update_time"] = d.UpdateTime
}

func (d device) clone(db *gorm.DB) device {
	d.deviceDo.ReplaceDB(db)
	return d
}

type deviceDo struct{ gen.DO }

func (d deviceDo) Debug() *deviceDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceDo) WithContext(ctx context.Context) *deviceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceDo) Clauses(conds ...clause.Expression) *deviceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceDo) Not(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceDo) Or(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceDo) Select(conds ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceDo) Where(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceDo) Order(conds ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceDo) Distinct(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceDo) Omit(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceDo) Join(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceDo) RightJoin(table schema.Tabler, on ...field.Expr) *deviceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceDo) Group(cols ...field.Expr) *deviceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceDo) Having(conds ...gen.Condition) *deviceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceDo) Limit(limit int) *deviceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceDo) Offset(offset int) *deviceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *deviceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceDo) Unscoped() *deviceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceDo) Create(values ...*gen_model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceDo) CreateInBatches(values []*gen_model.Device, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceDo) Save(values ...*gen_model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceDo) First() (*gen_model.Device, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Device), nil
	}
}

func (d deviceDo) Take() (*gen_model.Device, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Device), nil
	}
}

func (d deviceDo) Last() (*gen_model.Device, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Device), nil
	}
}

func (d deviceDo) Find() ([]*gen_model.Device, error) {
	result, err := d.DO.Find()
	return result.([]*gen_model.Device), err
}

func (d deviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.Device, err error) {
	buf := make([]*gen_model.Device, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceDo) FindInBatches(result *[]*gen_model.Device, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceDo) Attrs(attrs ...field.AssignExpr) *deviceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceDo) Assign(attrs ...field.AssignExpr) *deviceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceDo) Joins(field field.RelationField) *deviceDo {
	return d.withDO(d.DO.Joins(field))
}

func (d deviceDo) Preload(field field.RelationField) *deviceDo {
	return d.withDO(d.DO.Preload(field))
}

func (d deviceDo) FirstOrInit() (*gen_model.Device, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Device), nil
	}
}

func (d deviceDo) FirstOrCreate() (*gen_model.Device, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.Device), nil
	}
}

func (d deviceDo) FindByPage(offset int, limit int) (result []*gen_model.Device, count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	result, err = d.Offset(offset).Limit(limit).Find()
	return
}

func (d deviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d *deviceDo) withDO(do gen.Dao) *deviceDo {
	d.DO = *do.(*gen.DO)
	return d
}
