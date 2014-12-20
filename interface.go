// An interface for entropy estimation of a byte stream

package entropy

import (
	"io"
)

type Estimator interface {
	io.Writer
	Value() float64
}
