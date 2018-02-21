package main

import (
	"crypto/rand"
	"fmt"

	"github.com/oklog/ulid"
)

func main() {
	fmt.Println(ulid.MustNew(ulid.Now(), rand.Reader).String())
}
