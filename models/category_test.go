// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

func testCategories(t *testing.T) {
	t.Parallel()

	query := Categories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
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

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Categories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CategorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CategoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Category exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CategoryExists to return true, but got false.")
	}
}

func testCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	categoryFound, err := FindCategory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if categoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Categories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Categories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	categoryOne := &Category{}
	categoryTwo := &Category{}
	if err = randomize.Struct(seed, categoryOne, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTwo, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = categoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = categoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Categories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	categoryOne := &Category{}
	categoryTwo := &Category{}
	if err = randomize.Struct(seed, categoryOne, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTwo, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = categoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = categoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func categoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Category) error {
	*o = Category{}
	return nil
}

func testCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Category{}
	o := &Category{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, categoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Category object: %s", err)
	}

	AddCategoryHook(boil.BeforeInsertHook, categoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	categoryBeforeInsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterInsertHook, categoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	categoryAfterInsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterSelectHook, categoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	categoryAfterSelectHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeUpdateHook, categoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	categoryBeforeUpdateHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterUpdateHook, categoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	categoryAfterUpdateHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeDeleteHook, categoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	categoryBeforeDeleteHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterDeleteHook, categoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	categoryAfterDeleteHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeUpsertHook, categoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	categoryBeforeUpsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterUpsertHook, categoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	categoryAfterUpsertHooks = []CategoryHook{}
}

func testCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(categoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoryToManyCategoryTypeRefUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Category
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
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

	check, err := a.CategoryTypeRefUsers().All(ctx, tx)
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

	slice := CategorySlice{&a}
	if err = a.L.LoadCategoryTypeRefUsers(ctx, tx, false, (*[]*Category)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryTypeRefUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CategoryTypeRefUsers = nil
	if err = a.L.LoadCategoryTypeRefUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryTypeRefUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCategoryToManyAddOpCategoryTypeRefUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Category
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCategoryTypeRefUsers(ctx, tx, i != 0, x...)
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

		if a.R.CategoryTypeRefUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CategoryTypeRefUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CategoryTypeRefUsers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCategoryToOneCategoryTypeUsingCategoryTypeRef(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Category
	var foreign CategoryType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, categoryTypeDBTypes, false, categoryTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CategoryType struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CategoryTypeRefID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.CategoryTypeRef().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddCategoryTypeHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *CategoryType) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := CategorySlice{&local}
	if err = local.L.LoadCategoryTypeRef(ctx, tx, false, (*[]*Category)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CategoryTypeRef == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CategoryTypeRef = nil
	if err = local.L.LoadCategoryTypeRef(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CategoryTypeRef == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testCategoryToOneSetOpCategoryTypeUsingCategoryTypeRef(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Category
	var b, c CategoryType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, categoryTypeDBTypes, false, strmangle.SetComplement(categoryTypePrimaryKeyColumns, categoryTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, categoryTypeDBTypes, false, strmangle.SetComplement(categoryTypePrimaryKeyColumns, categoryTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CategoryType{&b, &c} {
		err = a.SetCategoryTypeRef(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CategoryTypeRef != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CategoryTypeRefCategories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CategoryTypeRefID != x.ID {
			t.Error("foreign key was wrong value", a.CategoryTypeRefID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CategoryTypeRefID))
		reflect.Indirect(reflect.ValueOf(&a.CategoryTypeRefID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CategoryTypeRefID != x.ID {
			t.Error("foreign key was wrong value", a.CategoryTypeRefID, x.ID)
		}
	}
}

func testCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
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

func testCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CategorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Categories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	categoryDBTypes = map[string]string{`ID`: `int`, `CreatedBy`: `int`, `CreatedDate`: `timestamp`, `LastUpdated`: `timestamp`, `UpdatedBy`: `int`, `IsActive`: `int`, `CategoryTypeRefID`: `int`, `Name`: `varchar`}
	_               = bytes.MinRead
)

func testCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(categoryAllColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(categoryAllColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Category{}
	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, categoryDBTypes, true, categoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(categoryAllColumns, categoryPrimaryKeyColumns) {
		fields = categoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			categoryAllColumns,
			categoryPrimaryKeyColumns,
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

	slice := CategorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(categoryAllColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCategoryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Category{}
	if err = randomize.Struct(seed, &o, categoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Category: %s", err)
	}

	count, err := Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, categoryDBTypes, false, categoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Category: %s", err)
	}

	count, err = Categories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
