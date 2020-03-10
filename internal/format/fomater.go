package format

import (
	"log"
	"os/exec"
)

var (
	protoStyle = `{
BasedOnStyle: google,
AlignConsecutiveAssignments: true,
AlignConsecutiveDeclarations: true,
AllowShortFunctionsOnASingleLine: None,
ColumnLimit: 0,
IndentWidth: 4,
Language: Proto,
}`
)

type Formatter struct {
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Format(absFile []string) {
	args := []string{"-style", protoStyle, "-i"}
	args = append(args, absFile...)
	cmd := exec.Command("clang-format", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal("compile Error:", err)
	}
}
