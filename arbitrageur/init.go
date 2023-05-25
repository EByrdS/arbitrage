package arbitrageur

import (
	golog "log"
	"os"
)

var logger *golog.Logger

func init() {
	logger = golog.New(os.Stdout, "Arbitrageur: ", 0)
}
