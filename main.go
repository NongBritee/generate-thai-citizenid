package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math/rand"
)

func main() {
	// use echo expose port 3000
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "The server is running.")
	})
	e.GET("/generate", GetThaiCitizenID)
	e.Logger.Fatal(e.Start(":3000"))
}

func GetThaiCitizenID(c echo.Context) error {
	// generate thai citizen id
	citizenID := GenerateThaiCitizenID()
	resp := Response{
		CitizenID: citizenID,
	}
	return c.JSON(200, resp)
}

func GenerateThaiCitizenID() string {
	// random 1-9
	Digit13 := RandomNumber(1, 9)
	Digit12 := RandomNumber(0, 9)
	Digit11 := RandomNumber(0, 9)
	Digit10 := RandomNumber(0, 9)
	Digit9 := RandomNumber(0, 9)
	Digit8 := RandomNumber(0, 9)
	Digit7 := RandomNumber(0, 9)
	Digit6 := RandomNumber(0, 9)
	Digit5 := RandomNumber(0, 9)
	Digit4 := RandomNumber(0, 9)
	Digit3 := RandomNumber(0, 9)
	Digit2 := RandomNumber(0, 9)
	sum := (Digit13 * 13) + (Digit12 * 12) + (Digit11 * 11) + (Digit10 * 10) + (Digit9 * 9) + (Digit8 * 8) + (Digit7 * 7) + (Digit6 * 6) + (Digit5 * 5) + (Digit4 * 4) + (Digit3 * 3) + (Digit2 * 2)
	Digit1 := (11 - (sum % 11)) % 10
	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d", Digit13, Digit12, Digit11, Digit10, Digit9, Digit8, Digit7, Digit6, Digit5, Digit4, Digit3, Digit2, Digit1)
}

func RandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

type Response struct {
	CitizenID string `json:"citizen_id"`
}
