// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package mysqlmodel

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCategoryTypes(t *testing.T) {
	t.Parallel()

	query := CategoryTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCategoryTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoryTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CategoryTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoryTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CategoryTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoryTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CategoryTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CategoryType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CategoryTypeExists to return true, but got false.")
	}
}

func testCategoryTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	categoryTypeFound, err := FindCategoryType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if categoryTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCategoryTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CategoryTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCategoryTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CategoryTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCategoryTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	categoryTypeOne := &CategoryType{}
	categoryTypeTwo := &CategoryType{}
	if err = randomize.Struct(seed, categoryTypeOne, categoryTypeDBTypes, false, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTypeTwo, categoryTypeDBTypes, false, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = categoryTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = categoryTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CategoryTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCategoryTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	categoryTypeOne := &CategoryType{}
	categoryTypeTwo := &CategoryType{}
	if err = randomize.Struct(seed, categoryTypeOne, categoryTypeDBTypes, false, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTypeTwo, categoryTypeDBTypes, false, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = categoryTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = categoryTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func categoryTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func categoryTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
	*o = CategoryType{}
	return nil
}

func testCategoryTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CategoryType{}
	o := &CategoryType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CategoryType object: %s", err)
	}

	AddCategoryTypeHook(boil.BeforeInsertHook, categoryTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	categoryTypeBeforeInsertHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.AfterInsertHook, categoryTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	categoryTypeAfterInsertHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.AfterSelectHook, categoryTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	categoryTypeAfterSelectHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.BeforeUpdateHook, categoryTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	categoryTypeBeforeUpdateHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.AfterUpdateHook, categoryTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	categoryTypeAfterUpdateHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.BeforeDeleteHook, categoryTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	categoryTypeBeforeDeleteHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.AfterDeleteHook, categoryTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	categoryTypeAfterDeleteHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.BeforeUpsertHook, categoryTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	categoryTypeBeforeUpsertHooks = []CategoryTypeHook{}

	AddCategoryTypeHook(boil.AfterUpsertHook, categoryTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	categoryTypeAfterUpsertHooks = []CategoryTypeHook{}
}

func testCategoryTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoryTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(categoryTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoryTypeToManyCategoryTypeRefCategories(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CategoryType
	var b, c Category

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CategoryTypeRefID = a.ID
	c.CategoryTypeRefID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CategoryTypeRefCategories().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.CategoryTypeRefID == b.CategoryTypeRefID {
			bFound = true
		}
		if v.CategoryTypeRefID == c.CategoryTypeRefID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CategoryTypeSlice{&a}
	if err = a.L.LoadCategoryTypeRefCategories(ctx, tx, false, (*[]*CategoryType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryTypeRefCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CategoryTypeRefCategories = nil
	if err = a.L.LoadCategoryTypeRefCategories(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryTypeRefCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCategoryTypeToManyAddOpCategoryTypeRefCategories(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CategoryType
	var b, c, d, e Category

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, categoryTypeDBTypes, false, strmangle.SetComplement(categoryTypePrimaryKeyColumns, categoryTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Category{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Category{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCategoryTypeRefCategories(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CategoryTypeRefID {
			t.Error("foreign key was wrong value", a.ID, first.CategoryTypeRefID)
		}
		if a.ID != second.CategoryTypeRefID {
			t.Error("foreign key was wrong value", a.ID, second.CategoryTypeRefID)
		}

		if first.R.CategoryTypeRef != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.CategoryTypeRef != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CategoryTypeRefCategories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CategoryTypeRefCategories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CategoryTypeRefCategories().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCategoryTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCategoryTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CategoryTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCategoryTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CategoryTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	categoryTypeDBTypes = map[string]string{`ID`: `int`, `CreatedBy`: `int`, `CreatedDate`: `timestamp`, `LastUpdated`: `timestamp`, `UpdatedBy`: `int`, `IsActive`: `int`, `Name`: `varchar`}
	_                   = bytes.MinRead
)

func testCategoryTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(categoryTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(categoryTypeAllColumns) == len(categoryTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCategoryTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(categoryTypeAllColumns) == len(categoryTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CategoryType{}
	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, categoryTypeDBTypes, true, categoryTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(categoryTypeAllColumns, categoryTypePrimaryKeyColumns) {
		fields = categoryTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			categoryTypeAllColumns,
			categoryTypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := CategoryTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCategoryTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(categoryTypeAllColumns) == len(categoryTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCategoryTypeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CategoryType{}
	if err = randomize.Struct(seed, &o, categoryTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CategoryType: %s", err)
	}

	count, err := CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, categoryTypeDBTypes, false, categoryTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CategoryType: %s", err)
	}

	count, err = CategoryTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
