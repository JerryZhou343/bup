package format

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
)

type Formatter struct {
	transformer Transformer
}

func NewFormatter() *Formatter {

	ret := &Formatter{}
	ret.transformer = NewTransformer()
	return ret
}

func (f *Formatter) FormatSignFile(absPath string) error {
	_, err := f.format(absPath)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (f *Formatter) FormatMultiFile(absPaths []string) error {
	for _, itr := range absPaths {
		_, err := f.format(itr)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (f *Formatter) format(absPath string) (bool, error) {
	input, err := ioutil.ReadFile(absPath)
	if err != nil {
		return false, err
	}
	data, failures, err := f.transformer.Transform(absPath, input)
	if err != nil {
		return false, err
	}
	if len(failures) > 0 {
		log.Printf("format failed [%+v]", failures)
		return false, errors.WithStack(err)
	}
	if !bytes.Equal(input, data) {
		// 0 exit code in overwrite case
		return true, ioutil.WriteFile(absPath, data, os.ModePerm)
	}

	return true, nil
}
