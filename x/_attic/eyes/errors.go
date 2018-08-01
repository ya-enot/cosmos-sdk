package eyes

import (
	"fmt"

	abci "github.com/ya-enot/abci/types"

	"github.com/ya-enot/cosmos-sdk/errors"
)

var (
	errMissingData = fmt.Errorf("All tx fields must be filled")

	malformed = abci.CodeType_EncodingError
)

//nolint
func ErrMissingData() errors.TMError {
	return errors.WithCode(errMissingData, malformed)
}
func IsMissingDataErr(err error) bool {
	return errors.IsSameError(errMissingData, err)
}
