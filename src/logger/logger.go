package logger

import (
	"fmt"
	// "os"
	log "github.com/sirupsen/logrus"
)

type Fields = log.Fields

var standardFields Fields

func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{})
	// for production, use JSON instead of text
	// log.SetFormatter(&log.JSONFormatter{})

	// Log to STDOUT (default is STDERR)
	// log.SetOutput(os.Stdout)
	
	// OR log to File
	// file, err := os.OpenFile("service.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    // if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
    // }
	// log.SetOutput(file)
	// log.Println("Service log re-opened")
	
	// log.SetReportCaller(true) // isn't great if you have a wrapper interface like here

	// log.SetLevel(log.DebugLevel)

	standardFields = Fields{
	// 	"hostAddress": "http://localhost:3000",
		"appName":  "go-error-logger",
	}
}

func Info(message string, fields Fields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Info(message)
}

func Error(message string, err error, fields Fields) {
	log.WithFields(standardFields).
		WithFields(fields).

		// Error(fmt.Sprintf("%s (%+v)", message, err))
		// A database-related error occured (SQL_ERROR <HUGE_STACK_TRACE> Repo Failure <STACK_TRACE>)
		// stack trace is very large (26 lines, 2 wraps), perhaps only useful in a log file?
		
		Error(fmt.Sprintf("%s (%v)", message, err))
		// A database-related error occured (Repo Failure: database entry already exists: SQL_ERROR)
		
		// Error(fmt.Sprintf("%s (%w)", message, err))
		// A database-related error occured (%!w(*errors.withStack=&{0xc000276660 0xc000276680}))
}