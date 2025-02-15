package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

/**
 * Function extractInterval, takes one parameter which must be of type "xy".
 * Where x is a number and y is a letter (s - second, m - month, h - hour, d - day).
 * The function extracts the number and the letter in two different variables numberPart and intervalPart.
 */
func extractInterval(interval string) (numberPart string, intervalPart string) {
	for index, character := range interval {
		if unicode.IsLetter(character) {
			numberPart = interval[:index]
			intervalPart = interval[index:]
			break
		}
	}
	return
}

/**
 * Function computeExpirationInterval, start with two variables number and unitPart.
 * The function purpose is to set the interval in a int64 type, based on the JWT_EXPIRATION_INTERVAL.
 * We switch through the unitPart, were all of the scenarios are covered. As a default value we set the interval to 1 hour.
 */
func computeExpirationInterval(number int, unitPart string) int64 {
	switch unitPart {
	case "h":
		return time.Now().Add(time.Duration(number) * time.Hour).Unix()
	case "m":
		return time.Now().Add(time.Duration(number) * time.Minute).Unix()
	case "s":
		return time.Now().Add(time.Duration(number) * time.Second).Unix()
	case "d":
		return time.Now().Add(time.Duration(number*24) * time.Hour).Unix()
	default:
		return time.Now().Add(time.Hour * 1).Unix()
	}
}

/**
 * Function IntervalGenerator - The function can be accessed from other modules.
 * The JWT_EXPIRE_INTERVAL is read, extracted and computed so it can be used in the jwt.MapClaims.
 */
func IntervalGenerator(jwtInterval string) int64 {
	interval := jwtInterval
	if interval == "" {
		interval = os.Getenv("JWT_EXPIRE_INTERVAL")
	}

	numberPart, unitPart := extractInterval(interval)

	var number int
	var err error
	// Convert string number to int
	if number, err = strconv.Atoi(numberPart); err != nil {
		fmt.Println("Error converting string to int:", err)
		return -1
	}

	return computeExpirationInterval(number, unitPart)
}
