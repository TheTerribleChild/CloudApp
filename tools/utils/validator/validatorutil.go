package validator

import (
    "context"
    "log"
)


func Validate(ctx context.Context, valiators ...Validator) error {
    for _, validator := range valiators {
        if err := validator.Validate(); err != nil {
            log.Printf("Validation failed: %s", err.Error())
            return err
        }
    }
    return nil
}