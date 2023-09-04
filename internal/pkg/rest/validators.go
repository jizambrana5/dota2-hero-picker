package rest

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/errors"
)

const (
	required  = "required"
	dive      = "dive"
	attribute = "attribute"
	role      = "role"
	rankname  = "rank_name"
)

type CustomValidator struct {
	v *validator.Validate
}

func (cv *CustomValidator) ValidateStruct(obj interface{}) error {
	return cv.v.Struct(obj)
}

func (cv *CustomValidator) Engine() interface{} {
	return cv.v
}

func SetupValidators() error {
	// Create a new validator instance
	v := validator.New()

	// Register custom validation functions for roles, rank names, and attributes
	err := v.RegisterValidation("role", validateRole)
	if err != nil {
		return errors.Validators
	}
	err = v.RegisterValidation("rank_name", validateRankName)
	if err != nil {
		return errors.Validators
	}
	err = v.RegisterValidation("attribute", validateAttribute)
	if err != nil {
		return errors.Validators
	}

	// Set the custom validator for Gin
	binding.Validator = &CustomValidator{v}
	return nil
}

func validateRole(fl validator.FieldLevel) bool {
	// Convert the field value to the Role type
	roleValue := domain.Role(fl.Field().String())

	// Check if the roleValue is one of the valid rolesreturn roleValue.IsValid()
	return roleValue.IsValid()
}

func validateRankName(fl validator.FieldLevel) bool {
	// Convert the field value to the Role type
	rankName := domain.RankName(fl.Field().String())

	// Check if the rankName is one of the valid rank name return roleValue.IsValid()
	return rankName.IsValid()
}

func validateAttribute(fl validator.FieldLevel) bool {
	// Convert the field value to the Role type
	attr := domain.Attribute(fl.Field().String())

	// Check if the attr value is one of the valid attributes roleValue.IsValid()
	return attr.IsValid()
}

func handleShouldBindJSONErrors(validationErrors validator.ValidationErrors) map[string]string {
	errorMsgs := make(map[string]string)
	for _, e := range validationErrors {
		// You can customize the error messages based on the field and tag
		switch e.Tag() {
		case required:
			errorMsgs[e.Field()] = fmt.Sprintf("%s is required", e.Field())
		case dive:
			// Handle errors for elements inside an array, if needed
		case attribute:
			errorMsgs[e.Field()] = fmt.Sprintf("%s is not a valid attribute", e.Field())
		case role:
			errorMsgs[e.Field()] = fmt.Sprintf("%s is not a valid role", e.Field())
		case rankname:
			errorMsgs[e.Field()] = fmt.Sprintf("%s is not a valid rank name", e.Field())
		}
	}
	return errorMsgs
}
