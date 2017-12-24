package template

import (
	"errors"
	"fmt"
)

//Template returns true
func Template(name string) string {
	return fmt.Sprintf("this is some template project called %s", name)
}

//Error returns some string and error
func Error() (string, error) {
	return "ups", errors.New("something bad happend")
}
