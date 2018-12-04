package model

import (
	"os"
)

// Secret is signing key of JWT.
var Secret = []byte(os.Getenv("SECRET"))

