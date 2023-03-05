package err

import (
	"fmt"
	"log"
)

type ProductProviderError struct {
	Message       string
	originalError error
}

func NewProductProviderError(msg string, originalError error) *ProductProviderError {
	return &ProductProviderError{
		Message:       msg,
		originalError: originalError,
	}
}
func (e ProductProviderError) Error() string {
	log.Println("Provider error occurred", e.originalError.Error())
	return fmt.Sprintf("An error occurred using the ProductProvider %s", e.Message)
}
