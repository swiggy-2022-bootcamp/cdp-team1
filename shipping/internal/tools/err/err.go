package err

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//LoggerError ..
func LoggerError(err error, file *os.File) {
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}
}

//LogFatalError ..
func LogFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
