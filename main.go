package main

import (
	"fmt"

	"github.com/CROWNIX/go-utils/apperror"
)

func main() {
	fmt.Println(apperror.InternalServerError("Internal Server Error"))
}