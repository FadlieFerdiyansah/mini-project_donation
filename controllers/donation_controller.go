package controllers

import (
	"math/rand"
	"mini_project/config"
	"mini_project/helpers"
	"mini_project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Donation(c echo.Context) error {
	user_id, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return err
	}
	program_id, err := strconv.Atoi(c.FormValue("program_id"))
	if err != nil {
		return err
	}
	payment_method_id, err := strconv.Atoi(c.FormValue("payment_method_id"))
	if err != nil {
		return err
	}
	amount, err := strconv.Atoi(c.FormValue("amount"))
	if err != nil {
		return err
	}

	donation := &models.Donation{
		UserID:          uint64(user_id),
		ProgramID:       uint64(program_id),
		PaymentMethodId: uint64(payment_method_id),
		UniqueId:        randSeq(10),
		Amount:          uint64(amount),
		Status:          "pending",
	}

	config.InitDB().Create(donation)

	// urlMidtransSand := "https://app.sandbox.midtrans.com/snap/v1/transactions"
	// c.Request().Header.Set("Accept", "application/json")
	// c.Request().Header.Set("Content-Type", "application/json")
	// c.Request().Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("SB-Mid-server-9V2-7cdPsfeiAfR3EGxvq4QV")))
	// var params = map[string]interface{}{
	// 	"enabled_payments": [1]string{"va"}[0],
	// 	"transaction_details": map[string]interface{}{
	// 		"order_id":     donation.UniqueId,
	// 		"gross_amount": donation.Amount,
	// 	},
	// 	"costumer_details": "Jhon Doe",
	// 	"expiry": map[string]interface{}{
	// 		"start_time": time.Now().AddDate(13, 0, -2),
	// 		"unit":       "day",
	// 		"duration":   1,
	// 	},
	// }
	// http.Post(urlMidtransSand, "application/json", params)
	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Waiting for payment", donation))
}

var letters = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
