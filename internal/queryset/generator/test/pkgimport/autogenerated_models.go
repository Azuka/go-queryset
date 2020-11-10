// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	stderrors "errors"
	"strings"

	forex "github.com/jirfag/go-queryset/internal/queryset/generator/test/pkgimport/forex/v1"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ExampleQuerySet

// ExampleQuerySet is an queryset type for Example
type ExampleQuerySet struct {
	db *gorm.DB
}

// NewExampleQuerySet constructs new ExampleQuerySet
func NewExampleQuerySet(db *gorm.DB) ExampleQuerySet {
	return ExampleQuerySet{
		db: db.Model(&Example{}),
	}
}

func (qs ExampleQuerySet) w(db *gorm.DB) ExampleQuerySet {
	return NewExampleQuerySet(db)
}

func (qs ExampleQuerySet) Select(fields ...ExampleDBSchemaField) ExampleQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *Example) Create(db *gorm.DB) error {
	return errors.WithStack(db.Create(o).Error)
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Example) Delete(db *gorm.DB) error {
	return errors.WithStack(db.Delete(o).Error)
}

// All is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) All(ret *[]Example) error {
	return errors.WithStack(qs.db.Find(ret).Error)
}

// Count is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Count() (int64, error) {
	var count int64
	err := qs.db.Count(&count).Error
	return count, errors.WithStack(err)
}

// Currency1Eq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Eq(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("currency1 = ?", currency1))
}

// Currency1Gt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Gt(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("currency1 > ?", currency1))
}

// Currency1Gte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Gte(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("currency1 >= ?", currency1))
}

// Currency1In is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1In(currency1 ...forex.Currency1) ExampleQuerySet {
	if len(currency1) == 0 {
		qs.db.AddError(errors.New("must at least pass one currency1 in Currency1In"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("currency1 IN (?)", currency1))
}

// Currency1Lt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Lt(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("currency1 < ?", currency1))
}

// Currency1Lte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Lte(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("currency1 <= ?", currency1))
}

// Currency1Ne is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1Ne(currency1 forex.Currency1) ExampleQuerySet {
	return qs.w(qs.db.Where("currency1 != ?", currency1))
}

// Currency1NotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency1NotIn(currency1 ...forex.Currency1) ExampleQuerySet {
	if len(currency1) == 0 {
		qs.db.AddError(errors.New("must at least pass one currency1 in Currency1NotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("currency1 NOT IN (?)", currency1))
}

// Currency2Eq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Eq(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 = ?", currency2))
}

// Currency2Gt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Gt(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 > ?", currency2))
}

// Currency2Gte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Gte(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 >= ?", currency2))
}

// Currency2In is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2In(currency2 ...forex.Currency2) ExampleQuerySet {
	if len(currency2) == 0 {
		qs.db.AddError(errors.New("must at least pass one currency2 in Currency2In"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("currency2 IN (?)", currency2))
}

// Currency2Like is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Like(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 LIKE ?", currency2))
}

// Currency2Lt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Lt(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 < ?", currency2))
}

// Currency2Lte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Lte(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 <= ?", currency2))
}

// Currency2Ne is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Ne(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 != ?", currency2))
}

// Currency2NotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2NotIn(currency2 ...forex.Currency2) ExampleQuerySet {
	if len(currency2) == 0 {
		qs.db.AddError(errors.New("must at least pass one currency2 in Currency2NotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("currency2 NOT IN (?)", currency2))
}

// Currency2Notlike is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency2Notlike(currency2 forex.Currency2) ExampleQuerySet {
	return qs.w(qs.db.Where("currency2 NOT LIKE ?", currency2))
}

// Currency3Eq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Eq(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 = ?", currency3))
}

// Currency3Gt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Gt(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 > ?", currency3))
}

// Currency3Gte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Gte(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 >= ?", currency3))
}

// Currency3In is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3In(currency3 ...forex.Currency3) ExampleQuerySet {
	if len(currency3) == 0 {
		qs.db.AddError(errors.New("must at least pass one currency3 in Currency3In"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("currency3 IN (?)", currency3))
}

// Currency3Like is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Like(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 LIKE ?", currency3))
}

// Currency3Lt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Lt(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 < ?", currency3))
}

// Currency3Lte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Lte(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 <= ?", currency3))
}

// Currency3Ne is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Ne(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 != ?", currency3))
}

// Currency3NotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3NotIn(currency3 ...forex.Currency3) ExampleQuerySet {
	if len(currency3) == 0 {
		qs.db.AddError(errors.New("must at least pass one currency3 in Currency3NotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("currency3 NOT IN (?)", currency3))
}

// Currency3Notlike is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Currency3Notlike(currency3 forex.Currency3) ExampleQuerySet {
	return qs.w(qs.db.Where("currency3 NOT LIKE ?", currency3))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Delete() error {
	return errors.WithStack(qs.db.Delete(Example{}).Error)
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Example{})
	return db.RowsAffected, errors.WithStack(db.Error)
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Example{})
	return db.RowsAffected, errors.WithStack(db.Error)
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) GetUpdater() ExampleUpdater {
	return NewExampleUpdater(qs.db)
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Limit(limit int) ExampleQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) Offset(offset int) ExampleQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ExampleQuerySet) One(ret *Example) error {
	return errors.WithStack(qs.db.First(ret).Error)
}

// OrderAscByCurrency1 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderAscByCurrency1() ExampleQuerySet {
	return qs.w(qs.db.Order("currency1 ASC"))
}

// OrderAscByCurrency2 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderAscByCurrency2() ExampleQuerySet {
	return qs.w(qs.db.Order("currency2 ASC"))
}

// OrderAscByCurrency3 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderAscByCurrency3() ExampleQuerySet {
	return qs.w(qs.db.Order("currency3 ASC"))
}

// OrderAscByPriceID is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderAscByPriceID() ExampleQuerySet {
	return qs.w(qs.db.Order("price_id ASC"))
}

// OrderDescByCurrency1 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderDescByCurrency1() ExampleQuerySet {
	return qs.w(qs.db.Order("currency1 DESC"))
}

// OrderDescByCurrency2 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderDescByCurrency2() ExampleQuerySet {
	return qs.w(qs.db.Order("currency2 DESC"))
}

// OrderDescByCurrency3 is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderDescByCurrency3() ExampleQuerySet {
	return qs.w(qs.db.Order("currency3 DESC"))
}

// OrderDescByPriceID is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) OrderDescByPriceID() ExampleQuerySet {
	return qs.w(qs.db.Order("price_id DESC"))
}

// PriceIDEq is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDEq(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("price_id = ?", priceID))
}

// PriceIDGt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDGt(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("price_id > ?", priceID))
}

// PriceIDGte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDGte(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("price_id >= ?", priceID))
}

// PriceIDIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDIn(priceID ...int64) ExampleQuerySet {
	if len(priceID) == 0 {
		qs.db.AddError(errors.New("must at least pass one priceID in PriceIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("price_id IN (?)", priceID))
}

// PriceIDLt is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDLt(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("price_id < ?", priceID))
}

// PriceIDLte is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDLte(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("price_id <= ?", priceID))
}

// PriceIDNe is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDNe(priceID int64) ExampleQuerySet {
	return qs.w(qs.db.Where("price_id != ?", priceID))
}

// PriceIDNotIn is an autogenerated method
// nolint: dupl
func (qs ExampleQuerySet) PriceIDNotIn(priceID ...int64) ExampleQuerySet {
	if len(priceID) == 0 {
		qs.db.AddError(errors.New("must at least pass one priceID in PriceIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("price_id NOT IN (?)", priceID))
}

// SetCurrency1 is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetCurrency1(currency1 forex.Currency1) ExampleUpdater {
	u.fields[string(ExampleDBSchema.Currency1)] = currency1
	return u
}

// SetCurrency2 is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetCurrency2(currency2 forex.Currency2) ExampleUpdater {
	u.fields[string(ExampleDBSchema.Currency2)] = currency2
	return u
}

// SetCurrency3 is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetCurrency3(currency3 forex.Currency3) ExampleUpdater {
	u.fields[string(ExampleDBSchema.Currency3)] = currency3
	return u
}

// SetPriceID is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) SetPriceID(priceID int64) ExampleUpdater {
	u.fields[string(ExampleDBSchema.PriceID)] = priceID
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) Update() error {
	return errors.WithStack(u.db.Updates(u.fields).Error)
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ExampleUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, errors.WithStack(db.Error)
}

// ===== END of query set ExampleQuerySet

// ===== BEGIN of Example modifiers

// ExampleDBSchemaField describes database schema field. It requires for method 'Update'
type ExampleDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ExampleDBSchemaField) String() string {
	return string(f)
}

// ExampleDBSchema stores db field names of Example
var ExampleDBSchema = struct {
	PriceID   ExampleDBSchemaField
	Currency1 ExampleDBSchemaField
	Currency2 ExampleDBSchemaField
	Currency3 ExampleDBSchemaField
}{

	PriceID:   ExampleDBSchemaField("price_id"),
	Currency1: ExampleDBSchemaField("currency1"),
	Currency2: ExampleDBSchemaField("currency2"),
	Currency3: ExampleDBSchemaField("currency3"),
}

// Update updates Example fields by primary key
// nolint: dupl
func (o *Example) Update(db *gorm.DB, fields ...ExampleDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"price_id":  o.PriceID,
		"currency1": o.Currency1,
		"currency2": o.Currency2,
		"currency3": o.Currency3,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithStack(err)
		}

		return errors.Wrapf(err, "can't update Example %v fields %v",
			o, fields)
	}

	return nil
}

// ExampleUpdater is an Example updates manager
type ExampleUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewExampleUpdater creates new Example updater
// nolint: dupl
func NewExampleUpdater(db *gorm.DB) ExampleUpdater {
	return ExampleUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Example{}),
	}
}

// ===== END of Example modifiers

// ===== END of all query sets
