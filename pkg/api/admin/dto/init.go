package dto

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v8"
	"strings"
)

func init() {

	// Register custom validate methods
}

// Bind : bind request dto and auto verify parameters
func Bind(c *gin.Context, obj interface{}) error {
	_ = c.ShouldBindUri(obj)
	if err := c.ShouldBind(obj); err != nil {
		if fieldErr, ok := err.(validator.ValidationErrors); ok {
			var tagErrorMsg []string
			for _, v := range fieldErr {
				if _, has := ValidateErrorMessage[v.Tag]; has {
					tagErrorMsg = append(tagErrorMsg, fmt.Sprintf(ValidateErrorMessage[v.Tag], v.Field, v.Value))
				} else {
					tagErrorMsg = append(tagErrorMsg, err.Error())
				}
			}
			return errors.New(strings.Join(tagErrorMsg, ","))
		}
	}
	return nil
}

//ValidateErrorMessage : customize error messages
var ValidateErrorMessage = map[string]string{
	"required": "%s is required,got empty %#v",
}
