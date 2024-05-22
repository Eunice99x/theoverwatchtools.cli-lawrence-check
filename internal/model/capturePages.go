package model

import (
	"fmt"
	"github.com/dembygenesis/local.tools/internal/utilities/validationutils"
)

// CapturePagesFilters contains the capture pages filters.
type CapturePagesFilters struct {
	CapturePagesNameIn     []string `query:"CapturePages_name_in" json:"CapturePages_name_in"`
	CapturePagesTypeNameIn []string `query:"CapturePages_type_name_in" json:"CapturePages_type_name_in"`
	CapturePagesTypeIdIn   []int    `query:"CapturePages_type_id_in" json:"CapturePages_type_id_in"`
	CapturePagesIsActive   []int    `query:"is_active" json:"is_active"`
	IdsIn                  []uint   `query:"ids_in" json:"ids_in"`
	PaginationQueryFilters `swaggerignore:"true"`
}

type CapturePage struct {
	Id                    int    `json:"id" boil:"id"`
	CapturePagesTypeRefId int    `json:"capturePages_type_ref_id" boil:"capturePages_type_ref_id" swaggerignore:"true"`
	Name                  string `json:"name" boil:"name"`
	CapturePagesType      string `json:"capturePages_type" boil:"capturePages_type"`
}

func (c *CapturePagesFilters) Validate() error {
	if err := c.ValidatePagination(); err != nil {
		return fmt.Errorf("pagination: %v", err)
	}
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("CapturePagesFilters filters: %v", err)
	}

	return nil
}

type CapturePages []CapturePage

type PaginatedCapturePages struct {
	CapturePages []CapturePage `json:"capturePages"`
	Pagination   *Pagination   `json:"pagination"`
}
