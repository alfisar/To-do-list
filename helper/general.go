package helper

import "os"

func DeleteImage(data string) (err error) {
	err = os.Remove(data)
	return
}
