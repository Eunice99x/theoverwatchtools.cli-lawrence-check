// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Categories", testCategories)
	t.Run("CategoryTypes", testCategoryTypes)
	t.Run("Organizations", testOrganizations)
	t.Run("OrganizationTypes", testOrganizationTypes)
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Categories", testCategoriesDelete)
	t.Run("CategoryTypes", testCategoryTypesDelete)
	t.Run("Organizations", testOrganizationsDelete)
	t.Run("OrganizationTypes", testOrganizationTypesDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Categories", testCategoriesQueryDeleteAll)
	t.Run("CategoryTypes", testCategoryTypesQueryDeleteAll)
	t.Run("Organizations", testOrganizationsQueryDeleteAll)
	t.Run("OrganizationTypes", testOrganizationTypesQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Categories", testCategoriesSliceDeleteAll)
	t.Run("CategoryTypes", testCategoryTypesSliceDeleteAll)
	t.Run("Organizations", testOrganizationsSliceDeleteAll)
	t.Run("OrganizationTypes", testOrganizationTypesSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Categories", testCategoriesExists)
	t.Run("CategoryTypes", testCategoryTypesExists)
	t.Run("Organizations", testOrganizationsExists)
	t.Run("OrganizationTypes", testOrganizationTypesExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Categories", testCategoriesFind)
	t.Run("CategoryTypes", testCategoryTypesFind)
	t.Run("Organizations", testOrganizationsFind)
	t.Run("OrganizationTypes", testOrganizationTypesFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Categories", testCategoriesBind)
	t.Run("CategoryTypes", testCategoryTypesBind)
	t.Run("Organizations", testOrganizationsBind)
	t.Run("OrganizationTypes", testOrganizationTypesBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Categories", testCategoriesOne)
	t.Run("CategoryTypes", testCategoryTypesOne)
	t.Run("Organizations", testOrganizationsOne)
	t.Run("OrganizationTypes", testOrganizationTypesOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Categories", testCategoriesAll)
	t.Run("CategoryTypes", testCategoryTypesAll)
	t.Run("Organizations", testOrganizationsAll)
	t.Run("OrganizationTypes", testOrganizationTypesAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Categories", testCategoriesCount)
	t.Run("CategoryTypes", testCategoryTypesCount)
	t.Run("Organizations", testOrganizationsCount)
	t.Run("OrganizationTypes", testOrganizationTypesCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Categories", testCategoriesHooks)
	t.Run("CategoryTypes", testCategoryTypesHooks)
	t.Run("Organizations", testOrganizationsHooks)
	t.Run("OrganizationTypes", testOrganizationTypesHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Categories", testCategoriesInsert)
	t.Run("Categories", testCategoriesInsertWhitelist)
	t.Run("CategoryTypes", testCategoryTypesInsert)
	t.Run("CategoryTypes", testCategoryTypesInsertWhitelist)
	t.Run("Organizations", testOrganizationsInsert)
	t.Run("Organizations", testOrganizationsInsertWhitelist)
	t.Run("OrganizationTypes", testOrganizationTypesInsert)
	t.Run("OrganizationTypes", testOrganizationTypesInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Categories", testCategoriesReload)
	t.Run("CategoryTypes", testCategoryTypesReload)
	t.Run("Organizations", testOrganizationsReload)
	t.Run("OrganizationTypes", testOrganizationTypesReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Categories", testCategoriesReloadAll)
	t.Run("CategoryTypes", testCategoryTypesReloadAll)
	t.Run("Organizations", testOrganizationsReloadAll)
	t.Run("OrganizationTypes", testOrganizationTypesReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Categories", testCategoriesSelect)
	t.Run("CategoryTypes", testCategoryTypesSelect)
	t.Run("Organizations", testOrganizationsSelect)
	t.Run("OrganizationTypes", testOrganizationTypesSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Categories", testCategoriesUpdate)
	t.Run("CategoryTypes", testCategoryTypesUpdate)
	t.Run("Organizations", testOrganizationsUpdate)
	t.Run("OrganizationTypes", testOrganizationTypesUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Categories", testCategoriesSliceUpdateAll)
	t.Run("CategoryTypes", testCategoryTypesSliceUpdateAll)
	t.Run("Organizations", testOrganizationsSliceUpdateAll)
	t.Run("OrganizationTypes", testOrganizationTypesSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
