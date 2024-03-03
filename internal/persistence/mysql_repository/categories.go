package mysql_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dembygenesis/local.tools/internal/models"
	sqlboiler_autogenerated2 "github.com/dembygenesis/local.tools/internal/persistence/mysql_repository/sqlboiler_autogenerated"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strings"
)

func (m *repository) upsertCategory(ctx context.Context, exec boil.ContextExecutor, category *models.Category) (*sqlboiler_autogenerated2.Category, error) {
	var (
		entry *sqlboiler_autogenerated2.Category
		err   error
	)

	if category.CategoryTypeRefId < 1 {
		return nil, fmt.Errorf("invalid/category empty")
	}

	if strings.TrimSpace(category.Name) == "" {
		return nil, models.ErrEmptyName
	}

	if _, err = sqlboiler_autogenerated2.FindCategoryType(ctx, exec, category.CategoryTypeRefId); err != nil {
		return nil, fmt.Errorf("category type id: %v", err)
	}

	exists := false

	entry, err = sqlboiler_autogenerated2.Categories(
		sqlboiler_autogenerated2.CategoryWhere.Name.EQ(category.Name),
	).One(ctx, exec)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("check entry: %v", err)
		}
	}

	if entry != nil {
		exists = true
	}

	ctx, cancel := context.WithTimeout(ctx, m.cfg.QueryTimeout)
	defer cancel()

	if exists {
		entry.Name = category.Name
		if _, err = entry.Update(ctx, exec, boil.Infer()); err != nil {
			return nil, fmt.Errorf("insert: %v", err)
		}
	} else {
		entry = &sqlboiler_autogenerated2.Category{
			CategoryTypeRefID: category.CategoryTypeRefId,
			Name:              category.Name,
		}

		if err = entry.Insert(ctx, exec, boil.Infer()); err != nil {
			return nil, fmt.Errorf("insert: %v", err)
		}
	}

	return entry, nil
}

func (m *repository) GetCategories(ctx context.Context, t Transaction, f *models.CategoryFilter) (models.Categories, error) {
	var (
		categories models.Categories
		err        error
	)

	h, err := m.getTxHandler(t)
	if err != nil {
		return nil, fmt.Errorf("tx handler: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, m.cfg.QueryTimeout)
	defer cancel()

	queryMods := make([]qm.QueryMod, 0)

	if len(f.IdsIn) > 0 {
		queryMods = append(queryMods, sqlboiler_autogenerated2.CategoryWhere.ID.IN(f.IdsIn))
	}

	if err = sqlboiler_autogenerated2.Categories(queryMods...).Bind(ctx, h.getExec(), &categories); err != nil {
		return nil, fmt.Errorf("get categories: %v", err)
	}

	return categories, nil
}

func (m *repository) UpsertCategories(ctx context.Context, t Transaction, categories []models.Category) (models.Categories, error) {
	h, err := m.getTxHandler(t)
	if err != nil {
		return nil, fmt.Errorf("tx handler: %v", err)
	}

	for i, category := range categories {
		entry, err := m.upsertCategory(ctx, h.getExec(), &category)
		if err != nil {
			return nil, fmt.Errorf("upsert category: %v", err)
		}
		categories[i].Id = entry.ID
	}

	return categories, nil
}

// getTxHandler attempts to parse `t` as a txHandler.
func (m *repository) getTxHandler(t Transaction) (*txHandler, error) {
	h, ok := t.(*txHandler)
	if !ok {
		return nil, models.ErrNotATransactionHandler
	}

	return h, nil
}