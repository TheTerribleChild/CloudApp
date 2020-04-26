package validator

import (
    "fmt"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type Validator interface {
    Validate() error
}

type StringLengthValidator struct {
    MaxLength uint
    MinLength uint
    ArgValue  string
    ArgName   string
}

const defaultArgMaxLength = ^uint(0)

func (instance *StringLengthValidator) Validate() error {
    if instance.MaxLength == 0 {
        instance.MaxLength = defaultArgMaxLength
    }
    strLength := uint(len(instance.ArgValue))
    if instance.MinLength <= strLength &&  strLength <= instance.MaxLength {
        return nil
    }
    if len(instance.ArgName) > 0 {
        instance.ArgName = "Argument"
    }
    if instance.MinLength > 0 && instance.MaxLength != defaultArgMaxLength {
        return status.Error(codes.InvalidArgument, fmt.Sprintf("Length of %s must be between %d and %d", instance.ArgName, instance.MinLength, instance.MaxLength))
    } else if instance.MinLength > 0 {
        return status.Error(codes.InvalidArgument, fmt.Sprintf("Length of %s must be at least %d", instance.ArgName, instance.MinLength))
    } else if instance.MaxLength != defaultArgMaxLength  {
        return status.Error(codes.InvalidArgument, fmt.Sprintf("Length of %s must be at most %d", instance.ArgName, instance.MaxLength))
    }
    return nil
}
