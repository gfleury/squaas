package log

import (
	"fmt"
	"log"
	"time"
)

func Printf(format string, v ...interface{}) {
	fmt.Printf(fmt.Sprintf("[INFO] %v | %s",
		time.Now().Format("2006/01/02 - 15:04:05"), format), v)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(fmt.Sprintf("[FATAL] %v | %s",
		time.Now().Format("2006/01/02 - 15:04:05"), format), v)
}

func Fatal(v ...interface{}) {
	log.Fatal(fmt.Sprintf("[FATAL] %v | ",
		time.Now().Format("2006/01/02 - 15:04:05")), v)
}
