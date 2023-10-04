package errors

import (
	"fmt"
	"net/http"
)

type CustomError interface {
	error
	HTTPCode() int
	InternalCode() string
}

type AppError struct {
	Err          error
	httpCode     int
	internalCode string
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func (e AppError) HTTPCode() int {
	return e.httpCode
}

func (e AppError) InternalCode() string {
	return e.internalCode
}

var (
	HeroGet         = AppError{Err: fmt.Errorf("error getting hero"), httpCode: http.StatusInternalServerError, internalCode: "error_getting_hero"}
	HeroNotFound    = AppError{Err: fmt.Errorf("hero not found"), httpCode: http.StatusNotFound, internalCode: "hero_not_found"}
	HeroGetAll      = AppError{Err: fmt.Errorf("error getting all heroes"), httpCode: http.StatusInternalServerError, internalCode: "error_getting_all_heroes"}
	HeroSave        = AppError{Err: fmt.Errorf("error saving hero"), httpCode: http.StatusInternalServerError, internalCode: "error_saving_hero"}
	SaveAllHeroes   = AppError{Err: fmt.Errorf("error saving all heroes"), httpCode: http.StatusInternalServerError, internalCode: "error_saving_all_heroes"}
	ErrInvalidInput = AppError{Err: fmt.Errorf("invalid input"), httpCode: http.StatusBadRequest, internalCode: "invalid_input"}
	GetDataSet      = AppError{Err: fmt.Errorf("error getting dataset"), httpCode: http.StatusInternalServerError, internalCode: "error_getting_dataset"}
	WinRate         = AppError{Err: fmt.Errorf("error building winrates"), httpCode: http.StatusInternalServerError, internalCode: "error_building_winrates"}
	Validators      = AppError{Err: fmt.Errorf("error building validators"), httpCode: http.StatusInternalServerError, internalCode: "error_building_validators"}
)
