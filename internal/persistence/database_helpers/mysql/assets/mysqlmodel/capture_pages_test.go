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

func testCapturePages(t *testing.T) {
	t.Parallel()

	query := CapturePages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCapturePagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
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

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCapturePagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CapturePages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCapturePagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CapturePageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCapturePagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CapturePageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CapturePage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CapturePageExists to return true, but got false.")
	}
}

func testCapturePagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	capturePageFound, err := FindCapturePage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if capturePageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCapturePagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CapturePages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCapturePagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CapturePages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCapturePagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	capturePageOne := &CapturePage{}
	capturePageTwo := &CapturePage{}
	if err = randomize.Struct(seed, capturePageOne, capturePageDBTypes, false, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}
	if err = randomize.Struct(seed, capturePageTwo, capturePageDBTypes, false, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = capturePageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = capturePageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CapturePages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCapturePagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	capturePageOne := &CapturePage{}
	capturePageTwo := &CapturePage{}
	if err = randomize.Struct(seed, capturePageOne, capturePageDBTypes, false, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}
	if err = randomize.Struct(seed, capturePageTwo, capturePageDBTypes, false, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = capturePageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = capturePageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func capturePageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func capturePageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CapturePage) error {
	*o = CapturePage{}
	return nil
}

func testCapturePagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CapturePage{}
	o := &CapturePage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, capturePageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CapturePage object: %s", err)
	}

	AddCapturePageHook(boil.BeforeInsertHook, capturePageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	capturePageBeforeInsertHooks = []CapturePageHook{}

	AddCapturePageHook(boil.AfterInsertHook, capturePageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	capturePageAfterInsertHooks = []CapturePageHook{}

	AddCapturePageHook(boil.AfterSelectHook, capturePageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	capturePageAfterSelectHooks = []CapturePageHook{}

	AddCapturePageHook(boil.BeforeUpdateHook, capturePageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	capturePageBeforeUpdateHooks = []CapturePageHook{}

	AddCapturePageHook(boil.AfterUpdateHook, capturePageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	capturePageAfterUpdateHooks = []CapturePageHook{}

	AddCapturePageHook(boil.BeforeDeleteHook, capturePageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	capturePageBeforeDeleteHooks = []CapturePageHook{}

	AddCapturePageHook(boil.AfterDeleteHook, capturePageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	capturePageAfterDeleteHooks = []CapturePageHook{}

	AddCapturePageHook(boil.BeforeUpsertHook, capturePageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	capturePageBeforeUpsertHooks = []CapturePageHook{}

	AddCapturePageHook(boil.AfterUpsertHook, capturePageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	capturePageAfterUpsertHooks = []CapturePageHook{}
}

func testCapturePagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCapturePagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(capturePageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCapturePageToOneCapturePageSetUsingCapturePageSet(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CapturePage
	var foreign CapturePageSet

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, capturePageDBTypes, false, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, capturePageSetDBTypes, false, capturePageSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePageSet struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CapturePageSetID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.CapturePageSet().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddCapturePageSetHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *CapturePageSet) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := CapturePageSlice{&local}
	if err = local.L.LoadCapturePageSet(ctx, tx, false, (*[]*CapturePage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CapturePageSet == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CapturePageSet = nil
	if err = local.L.LoadCapturePageSet(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CapturePageSet == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testCapturePageToOneSetOpCapturePageSetUsingCapturePageSet(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CapturePage
	var b, c CapturePageSet

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, capturePageDBTypes, false, strmangle.SetComplement(capturePagePrimaryKeyColumns, capturePageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, capturePageSetDBTypes, false, strmangle.SetComplement(capturePageSetPrimaryKeyColumns, capturePageSetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, capturePageSetDBTypes, false, strmangle.SetComplement(capturePageSetPrimaryKeyColumns, capturePageSetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CapturePageSet{&b, &c} {
		err = a.SetCapturePageSet(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CapturePageSet != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CapturePages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CapturePageSetID != x.ID {
			t.Error("foreign key was wrong value", a.CapturePageSetID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CapturePageSetID))
		reflect.Indirect(reflect.ValueOf(&a.CapturePageSetID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CapturePageSetID != x.ID {
			t.Error("foreign key was wrong value", a.CapturePageSetID, x.ID)
		}
	}
}

func testCapturePagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
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

func testCapturePagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CapturePageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCapturePagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CapturePages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	capturePageDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `HTML`: `longtext`, `Clicks`: `int`, `IsControl`: `tinyint`, `CapturePageSetID`: `bigint`, `CreatedBy`: `int`, `UpdatedBy`: `int`, `LastImpressionAt`: `timestamp`, `Impressions`: `int`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                  = bytes.MinRead
)

func testCapturePagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(capturePagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(capturePageAllColumns) == len(capturePagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCapturePagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(capturePageAllColumns) == len(capturePagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CapturePage{}
	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, capturePageDBTypes, true, capturePagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(capturePageAllColumns, capturePagePrimaryKeyColumns) {
		fields = capturePageAllColumns
	} else {
		fields = strmangle.SetComplement(
			capturePageAllColumns,
			capturePagePrimaryKeyColumns,
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

	slice := CapturePageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCapturePagesUpsert(t *testing.T) {
	t.Parallel()

	if len(capturePageAllColumns) == len(capturePagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCapturePageUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CapturePage{}
	if err = randomize.Struct(seed, &o, capturePageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CapturePage: %s", err)
	}

	count, err := CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, capturePageDBTypes, false, capturePagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CapturePage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CapturePage: %s", err)
	}

	count, err = CapturePages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
