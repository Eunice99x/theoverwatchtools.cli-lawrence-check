// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package mysqlmodel

import "testing"

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CapturePageToCapturePageSetUsingCapturePageSet", testCapturePageToOneCapturePageSetUsingCapturePageSet)
	t.Run("CategoryToCategoryTypeUsingCategoryTypeRef", testCategoryToOneCategoryTypeUsingCategoryTypeRef)
	t.Run("OrganizationToOrganizationTypeUsingOrganizationTypeRef", testOrganizationToOneOrganizationTypeUsingOrganizationTypeRef)
	t.Run("UserToCategoryUsingCategoryTypeRef", testUserToOneCategoryUsingCategoryTypeRef)
	t.Run("UserToUserUsingCreatedByUser", testUserToOneUserUsingCreatedByUser)
	t.Run("UserToUserUsingLastUpdatedByUser", testUserToOneUserUsingLastUpdatedByUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CapturePageSetToCapturePages", testCapturePageSetToManyCapturePages)
	t.Run("CategoryToCategoryTypeRefUsers", testCategoryToManyCategoryTypeRefUsers)
	t.Run("CategoryTypeToCategoryTypeRefCategories", testCategoryTypeToManyCategoryTypeRefCategories)
	t.Run("OrganizationTypeToOrganizationTypeRefOrganizations", testOrganizationTypeToManyOrganizationTypeRefOrganizations)
	t.Run("UserToCreatedByUsers", testUserToManyCreatedByUsers)
	t.Run("UserToLastUpdatedByUsers", testUserToManyLastUpdatedByUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CapturePageToCapturePageSetUsingCapturePages", testCapturePageToOneSetOpCapturePageSetUsingCapturePageSet)
	t.Run("CategoryToCategoryTypeUsingCategoryTypeRefCategories", testCategoryToOneSetOpCategoryTypeUsingCategoryTypeRef)
	t.Run("OrganizationToOrganizationTypeUsingOrganizationTypeRefOrganizations", testOrganizationToOneSetOpOrganizationTypeUsingOrganizationTypeRef)
	t.Run("UserToCategoryUsingCategoryTypeRefUsers", testUserToOneSetOpCategoryUsingCategoryTypeRef)
	t.Run("UserToUserUsingCreatedByUsers", testUserToOneSetOpUserUsingCreatedByUser)
	t.Run("UserToUserUsingLastUpdatedByUsers", testUserToOneSetOpUserUsingLastUpdatedByUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("UserToUserUsingCreatedByUsers", testUserToOneRemoveOpUserUsingCreatedByUser)
	t.Run("UserToUserUsingLastUpdatedByUsers", testUserToOneRemoveOpUserUsingLastUpdatedByUser)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CapturePageSetToCapturePages", testCapturePageSetToManyAddOpCapturePages)
	t.Run("CategoryToCategoryTypeRefUsers", testCategoryToManyAddOpCategoryTypeRefUsers)
	t.Run("CategoryTypeToCategoryTypeRefCategories", testCategoryTypeToManyAddOpCategoryTypeRefCategories)
	t.Run("OrganizationTypeToOrganizationTypeRefOrganizations", testOrganizationTypeToManyAddOpOrganizationTypeRefOrganizations)
	t.Run("UserToCreatedByUsers", testUserToManyAddOpCreatedByUsers)
	t.Run("UserToLastUpdatedByUsers", testUserToManyAddOpLastUpdatedByUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("UserToCreatedByUsers", testUserToManySetOpCreatedByUsers)
	t.Run("UserToLastUpdatedByUsers", testUserToManySetOpLastUpdatedByUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("UserToCreatedByUsers", testUserToManyRemoveOpCreatedByUsers)
	t.Run("UserToLastUpdatedByUsers", testUserToManyRemoveOpLastUpdatedByUsers)
}