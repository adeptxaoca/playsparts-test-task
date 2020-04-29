package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetType(t *testing.T) {
	err := ValidationError.New("error")
	assert.Equal(t, ValidationError, GetType(err))
	err = errors.New("error")
	assert.Equal(t, InternalError, GetType(err))
}

func TestWrap(t *testing.T) {
	err := errors.New("some kind of error")
	err = Wrap(err, "comment on error")
	assert.Equal(t, InternalError, GetType(err))
	assert.Equal(t, "comment on error: some kind of error", err.Error())
}

func TestCause(t *testing.T) {
	err := errors.New("original")
	err = Wrap(err, "second")
	err = Wrap(err, "another")
	assert.Equal(t, "original", Cause(err).Error())
}

func TestAddErrorContext(t *testing.T) {
	err := errors.New("error")
	ctx := GetErrorContext(err)
	assert.Nil(t, ctx)
	err = AddErrorContext(err, "name", "invalid")
	ctx = GetErrorContext(err)
	assert.Equal(t, "invalid", ctx["name"])
	err = AddErrorContext(err, "status", "invalid")
	ctx = GetErrorContext(err)
	assert.Equal(t, "invalid", ctx["status"])
}

func TestGrpcError(t *testing.T) {
	grpcErr := GrpcError(ValidationError.New("invalid id"))
	assert.Equal(t, codes.InvalidArgument, status.Code(grpcErr))
	assert.Equal(t, "rpc error: code = InvalidArgument desc = invalid id", grpcErr.Error())
	grpcErr = GrpcError(NotFound.New("user was not found"))
	assert.Equal(t, codes.NotFound, status.Code(grpcErr))
	assert.Equal(t, "rpc error: code = NotFound desc = user was not found", grpcErr.Error())
}
