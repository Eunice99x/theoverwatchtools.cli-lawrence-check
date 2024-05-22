// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CategoryType is an object representing the database table.
type CategoryType struct {
	ID          int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedBy   null.Int  `boil:"created_by" json:"created_by,omitempty" toml:"created_by" yaml:"created_by,omitempty"`
	CreatedDate null.Time `boil:"created_date" json:"created_date,omitempty" toml:"created_date" yaml:"created_date,omitempty"`
	LastUpdated null.Time `boil:"last_updated" json:"last_updated,omitempty" toml:"last_updated" yaml:"last_updated,omitempty"`
	UpdatedBy   null.Int  `boil:"updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
	IsActive    null.Int  `boil:"is_active" json:"is_active,omitempty" toml:"is_active" yaml:"is_active,omitempty"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *categoryTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L categoryTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CategoryTypeColumns = struct {
	ID          string
	CreatedBy   string
	CreatedDate string
	LastUpdated string
	UpdatedBy   string
	IsActive    string
	Name        string
}{
	ID:          "id",
	CreatedBy:   "created_by",
	CreatedDate: "created_date",
	LastUpdated: "last_updated",
	UpdatedBy:   "updated_by",
	IsActive:    "is_active",
	Name:        "name",
}

var CategoryTypeTableColumns = struct {
	ID          string
	CreatedBy   string
	CreatedDate string
	LastUpdated string
	UpdatedBy   string
	IsActive    string
	Name        string
}{
	ID:          "category_type.id",
	CreatedBy:   "category_type.created_by",
	CreatedDate: "category_type.created_date",
	LastUpdated: "category_type.last_updated",
	UpdatedBy:   "category_type.updated_by",
	IsActive:    "category_type.is_active",
	Name:        "category_type.name",
}

// Generated where

var CategoryTypeWhere = struct {
	ID          whereHelperint
	CreatedBy   whereHelpernull_Int
	CreatedDate whereHelpernull_Time
	LastUpdated whereHelpernull_Time
	UpdatedBy   whereHelpernull_Int
	IsActive    whereHelpernull_Int
	Name        whereHelperstring
}{
	ID:          whereHelperint{field: "`category_type`.`id`"},
	CreatedBy:   whereHelpernull_Int{field: "`category_type`.`created_by`"},
	CreatedDate: whereHelpernull_Time{field: "`category_type`.`created_date`"},
	LastUpdated: whereHelpernull_Time{field: "`category_type`.`last_updated`"},
	UpdatedBy:   whereHelpernull_Int{field: "`category_type`.`updated_by`"},
	IsActive:    whereHelpernull_Int{field: "`category_type`.`is_active`"},
	Name:        whereHelperstring{field: "`category_type`.`name`"},
}

// CategoryTypeRels is where relationship names are stored.
var CategoryTypeRels = struct {
	CategoryTypeRefCategories string
}{
	CategoryTypeRefCategories: "CategoryTypeRefCategories",
}

// categoryTypeR is where relationships are stored.
type categoryTypeR struct {
	CategoryTypeRefCategories CategorySlice `boil:"CategoryTypeRefCategories" json:"CategoryTypeRefCategories" toml:"CategoryTypeRefCategories" yaml:"CategoryTypeRefCategories"`
}

// NewStruct creates a new relationship struct
func (*categoryTypeR) NewStruct() *categoryTypeR {
	return &categoryTypeR{}
}

func (r *categoryTypeR) GetCategoryTypeRefCategories() CategorySlice {
	if r == nil {
		return nil
	}
	return r.CategoryTypeRefCategories
}

// categoryTypeL is where Load methods for each relationship are stored.
type categoryTypeL struct{}

var (
	categoryTypeAllColumns            = []string{"id", "created_by", "created_date", "last_updated", "updated_by", "is_active", "name"}
	categoryTypeColumnsWithoutDefault = []string{"created_by", "last_updated", "updated_by", "name"}
	categoryTypeColumnsWithDefault    = []string{"id", "created_date", "is_active"}
	categoryTypePrimaryKeyColumns     = []string{"id"}
	categoryTypeGeneratedColumns      = []string{}
)

type (
	// CategoryTypeSlice is an alias for a slice of pointers to CategoryType.
	// This should almost always be used instead of []CategoryType.
	CategoryTypeSlice []*CategoryType
	// CategoryTypeHook is the signature for custom CategoryType hook methods
	CategoryTypeHook func(context.Context, boil.ContextExecutor, *CategoryType) error

	categoryTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	categoryTypeType                 = reflect.TypeOf(&CategoryType{})
	categoryTypeMapping              = queries.MakeStructMapping(categoryTypeType)
	categoryTypePrimaryKeyMapping, _ = queries.BindMapping(categoryTypeType, categoryTypeMapping, categoryTypePrimaryKeyColumns)
	categoryTypeInsertCacheMut       sync.RWMutex
	categoryTypeInsertCache          = make(map[string]insertCache)
	categoryTypeUpdateCacheMut       sync.RWMutex
	categoryTypeUpdateCache          = make(map[string]updateCache)
	categoryTypeUpsertCacheMut       sync.RWMutex
	categoryTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var categoryTypeAfterSelectMu sync.Mutex
var categoryTypeAfterSelectHooks []CategoryTypeHook

var categoryTypeBeforeInsertMu sync.Mutex
var categoryTypeBeforeInsertHooks []CategoryTypeHook
var categoryTypeAfterInsertMu sync.Mutex
var categoryTypeAfterInsertHooks []CategoryTypeHook

var categoryTypeBeforeUpdateMu sync.Mutex
var categoryTypeBeforeUpdateHooks []CategoryTypeHook
var categoryTypeAfterUpdateMu sync.Mutex
var categoryTypeAfterUpdateHooks []CategoryTypeHook

var categoryTypeBeforeDeleteMu sync.Mutex
var categoryTypeBeforeDeleteHooks []CategoryTypeHook
var categoryTypeAfterDeleteMu sync.Mutex
var categoryTypeAfterDeleteHooks []CategoryTypeHook

var categoryTypeBeforeUpsertMu sync.Mutex
var categoryTypeBeforeUpsertHooks []CategoryTypeHook
var categoryTypeAfterUpsertMu sync.Mutex
var categoryTypeAfterUpsertHooks []CategoryTypeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CategoryType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CategoryType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CategoryType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CategoryType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CategoryType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CategoryType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CategoryType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CategoryType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CategoryType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCategoryTypeHook registers your hook function for all future operations.
func AddCategoryTypeHook(hookPoint boil.HookPoint, categoryTypeHook CategoryTypeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		categoryTypeAfterSelectMu.Lock()
		categoryTypeAfterSelectHooks = append(categoryTypeAfterSelectHooks, categoryTypeHook)
		categoryTypeAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		categoryTypeBeforeInsertMu.Lock()
		categoryTypeBeforeInsertHooks = append(categoryTypeBeforeInsertHooks, categoryTypeHook)
		categoryTypeBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		categoryTypeAfterInsertMu.Lock()
		categoryTypeAfterInsertHooks = append(categoryTypeAfterInsertHooks, categoryTypeHook)
		categoryTypeAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		categoryTypeBeforeUpdateMu.Lock()
		categoryTypeBeforeUpdateHooks = append(categoryTypeBeforeUpdateHooks, categoryTypeHook)
		categoryTypeBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		categoryTypeAfterUpdateMu.Lock()
		categoryTypeAfterUpdateHooks = append(categoryTypeAfterUpdateHooks, categoryTypeHook)
		categoryTypeAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		categoryTypeBeforeDeleteMu.Lock()
		categoryTypeBeforeDeleteHooks = append(categoryTypeBeforeDeleteHooks, categoryTypeHook)
		categoryTypeBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		categoryTypeAfterDeleteMu.Lock()
		categoryTypeAfterDeleteHooks = append(categoryTypeAfterDeleteHooks, categoryTypeHook)
		categoryTypeAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		categoryTypeBeforeUpsertMu.Lock()
		categoryTypeBeforeUpsertHooks = append(categoryTypeBeforeUpsertHooks, categoryTypeHook)
		categoryTypeBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		categoryTypeAfterUpsertMu.Lock()
		categoryTypeAfterUpsertHooks = append(categoryTypeAfterUpsertHooks, categoryTypeHook)
		categoryTypeAfterUpsertMu.Unlock()
	}
}

// One returns a single categoryType record from the query.
func (q categoryTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CategoryType, error) {
	o := &CategoryType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for category_type")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CategoryType records from the query.
func (q categoryTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (CategoryTypeSlice, error) {
	var o []*CategoryType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CategoryType slice")
	}

	if len(categoryTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CategoryType records in the query.
func (q categoryTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count category_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q categoryTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if category_type exists")
	}

	return count > 0, nil
}

// CategoryTypeRefCategories retrieves all the category's Categories with an executor via category_type_ref_id column.
func (o *CategoryType) CategoryTypeRefCategories(mods ...qm.QueryMod) categoryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`category`.`category_type_ref_id`=?", o.ID),
	)

	return Categories(queryMods...)
}

// LoadCategoryTypeRefCategories allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (categoryTypeL) LoadCategoryTypeRefCategories(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCategoryType interface{}, mods queries.Applicator) error {
	var slice []*CategoryType
	var object *CategoryType

	if singular {
		var ok bool
		object, ok = maybeCategoryType.(*CategoryType)
		if !ok {
			object = new(CategoryType)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCategoryType)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCategoryType))
			}
		}
	} else {
		s, ok := maybeCategoryType.(*[]*CategoryType)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCategoryType)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCategoryType))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &categoryTypeR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &categoryTypeR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`category`),
		qm.WhereIn(`category.category_type_ref_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load category")
	}

	var resultSlice []*Category
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on category")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for category")
	}

	if len(categoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CategoryTypeRefCategories = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &categoryR{}
			}
			foreign.R.CategoryTypeRef = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CategoryTypeRefID {
				local.R.CategoryTypeRefCategories = append(local.R.CategoryTypeRefCategories, foreign)
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.CategoryTypeRef = local
				break
			}
		}
	}

	return nil
}

// AddCategoryTypeRefCategories adds the given related objects to the existing relationships
// of the category_type, optionally inserting them as new records.
// Appends related to o.R.CategoryTypeRefCategories.
// Sets related.R.CategoryTypeRef appropriately.
func (o *CategoryType) AddCategoryTypeRefCategories(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Category) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CategoryTypeRefID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `category` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"category_type_ref_id"}),
				strmangle.WhereClause("`", "`", 0, categoryPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CategoryTypeRefID = o.ID
		}
	}

	if o.R == nil {
		o.R = &categoryTypeR{
			CategoryTypeRefCategories: related,
		}
	} else {
		o.R.CategoryTypeRefCategories = append(o.R.CategoryTypeRefCategories, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &categoryR{
				CategoryTypeRef: o,
			}
		} else {
			rel.R.CategoryTypeRef = o
		}
	}
	return nil
}

// CategoryTypes retrieves all the records using an executor.
func CategoryTypes(mods ...qm.QueryMod) categoryTypeQuery {
	mods = append(mods, qm.From("`category_type`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`category_type`.*"})
	}

	return categoryTypeQuery{q}
}

// FindCategoryType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCategoryType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CategoryType, error) {
	categoryTypeObj := &CategoryType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `category_type` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, categoryTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from category_type")
	}

	if err = categoryTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return categoryTypeObj, err
	}

	return categoryTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CategoryType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no category_type provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(categoryTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	categoryTypeInsertCacheMut.RLock()
	cache, cached := categoryTypeInsertCache[key]
	categoryTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			categoryTypeAllColumns,
			categoryTypeColumnsWithDefault,
			categoryTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(categoryTypeType, categoryTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(categoryTypeType, categoryTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `category_type` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `category_type` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `category_type` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, categoryTypePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into category_type")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == categoryTypeMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for category_type")
	}

CacheNoHooks:
	if !cached {
		categoryTypeInsertCacheMut.Lock()
		categoryTypeInsertCache[key] = cache
		categoryTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CategoryType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CategoryType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	categoryTypeUpdateCacheMut.RLock()
	cache, cached := categoryTypeUpdateCache[key]
	categoryTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			categoryTypeAllColumns,
			categoryTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update category_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `category_type` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, categoryTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(categoryTypeType, categoryTypeMapping, append(wl, categoryTypePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update category_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for category_type")
	}

	if !cached {
		categoryTypeUpdateCacheMut.Lock()
		categoryTypeUpdateCache[key] = cache
		categoryTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q categoryTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for category_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for category_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CategoryTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), categoryTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `category_type` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, categoryTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in categoryType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all categoryType")
	}
	return rowsAff, nil
}

var mySQLCategoryTypeUniqueColumns = []string{
	"id",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CategoryType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no category_type provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(categoryTypeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCategoryTypeUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	categoryTypeUpsertCacheMut.RLock()
	cache, cached := categoryTypeUpsertCache[key]
	categoryTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			categoryTypeAllColumns,
			categoryTypeColumnsWithDefault,
			categoryTypeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			categoryTypeAllColumns,
			categoryTypePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert category_type, could not build update column list")
		}

		ret := strmangle.SetComplement(categoryTypeAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`category_type`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `category_type` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(categoryTypeType, categoryTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(categoryTypeType, categoryTypeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for category_type")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == categoryTypeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(categoryTypeType, categoryTypeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for category_type")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for category_type")
	}

CacheNoHooks:
	if !cached {
		categoryTypeUpsertCacheMut.Lock()
		categoryTypeUpsertCache[key] = cache
		categoryTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CategoryType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CategoryType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CategoryType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), categoryTypePrimaryKeyMapping)
	sql := "DELETE FROM `category_type` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from category_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for category_type")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q categoryTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no categoryTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from category_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for category_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CategoryTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(categoryTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), categoryTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `category_type` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, categoryTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from categoryType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for category_type")
	}

	if len(categoryTypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CategoryType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCategoryType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CategoryTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CategoryTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), categoryTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `category_type`.* FROM `category_type` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, categoryTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CategoryTypeSlice")
	}

	*o = slice

	return nil
}

// CategoryTypeExists checks if the CategoryType row exists.
func CategoryTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `category_type` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if category_type exists")
	}

	return exists, nil
}

// Exists checks if the CategoryType row exists.
func (o *CategoryType) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CategoryTypeExists(ctx, exec, o.ID)
}
