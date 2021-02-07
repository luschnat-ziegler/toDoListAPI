/*
 * package: logger
 * --------------------
 * Includes a custom logger (using go.uber.org/zap)
 */

package logger

import "go.uber.org/zap"

var log *zap.Logger

/*
 * Function: init
 * --------------------
 * Initiates a logger instance and refers caller info one step up the chain.
 *
 * returns: nothing
 */

func init() {
	var err error
	log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

/*
 * Function: Fatal
 * --------------------
 * Creates a Fatal log output using the initiated logger instance
 * Terminates the program by calling os.Exit(1)
 *
 * message: A string with logging info
 * fields: Any number of zap.Field
 *
 * returns: nothing
 */

func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

/*
 * Function: Info
 * --------------------
 * Creates an Info log output using the initiated logger instance
 *
 * message: A string with logging info
 * fields: Any number of zap.Field
 *
 * returns: nothing
 */

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

/*
 * Function: Error
 * --------------------
 * Creates an Error log output using the initiated logger instance
 *
 * message: A string with logging info
 * fields: Any number of zap.Field
 *
 * returns: nothing
 */

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
