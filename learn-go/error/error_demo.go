package error

import "errors"

func test() (error error) {
	return errors.New("wow")
}
