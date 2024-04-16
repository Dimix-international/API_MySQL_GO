package models

import (
	"context"
	"errors"
)

var (
	ErrValidation = errors.New("")
)

type CloseFunc func(context.Context) error
