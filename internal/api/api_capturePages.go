package api

import (
	"github.com/dembygenesis/local.tools/internal/model"
	"github.com/dembygenesis/local.tools/internal/utilities/errs"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (a *Api) ListCapturePages(ctx *fiber.Ctx) error {

	filter := model.CapturePagesFilters{
		CapturePagesIsActive: []int{1},
	}

	if err := ctx.QueryParser(&filter); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errs.ToArr(err))
	}

	if err := filter.Validate(); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errs.ToArr(err))
	}

	filter.SetPaginationDefaults()

	capturePages, err := a.cfg.CapturePagesService.ListCapturePages(ctx.Context(), &filter)
	return a.WriteResponse(ctx, http.StatusOK, capturePages, err)
}
