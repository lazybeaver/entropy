// Estimate shannon entropy of the standard input byte stream

package main

import (
	"bufio"
	"fmt"
	"github.com/lazybeaver/entropy"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	estimator := entropy.NewShannonEstimator()
	if _, err := io.Copy(estimator, reader); err != nil {
		fmt.Printf("IO Error: %s", err)
		return
	}
	fmt.Println(estimator.Value())
}
