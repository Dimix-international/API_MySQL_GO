package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidationError(errs validator.ValidationErrors, model interface{}) string {
	var errMsgs []string

	for i := range errs {
		structField, _ := reflect.TypeOf(model).FieldByName(errs[i].Field())
		jsonFieldName := strings.Split(structField.Tag.Get("json"), ",")[0]

		switch errs[i].ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required", jsonFieldName))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", jsonFieldName))
		}
	}

	return strings.Join(errMsgs, ", ")
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
