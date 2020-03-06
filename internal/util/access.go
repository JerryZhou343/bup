package util

import (
	"io"
	"os"
)

func Access (name string) (io.ReadCloser, error) {
		return os.Open(name)
}
