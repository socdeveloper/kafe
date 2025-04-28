package orderSubmitHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

type Order struct {
	Amount        int    `json:"amount"`
	AmountDisplay string `json:"amountDisplay"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Comment       string `json:"comment"`
	Promo         string `json:"promo"`
	City          string `json:"city"`
	CardNumber    string `json:"cardNumber"`
}

func sendEmail(order Order) error {
	smtpHost := "smtp.mail.ru"
	smtpPort := "587"
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
		smtpHost,
	)

	userMsg := []byte(
		"To: " + order.Email + "\r\n" +
			"Subject: Ваш заказ подарочной карты\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			fmt.Sprintf(
				"Уважаемый %s,\n\nСпасибо за заказ подарочной карты на сумму %d руб.\n\n"+
					"Номер вашего заказа: #%d\n"+
					"Мы свяжемся с вами в ближайшее время.\n\n"+
					"С уважением,\nКоманда принцип",
				order.Name,
				order.Amount,
				time.Now().Unix(),
			),
	)

	err1 := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		os.Getenv("SMTP_USER"),
		[]string{order.Email},
		userMsg,
	)

	if err1 != nil {
		return fmt.Errorf("ошибка отправки: пользователь=%v", err1)
	}
	return nil
}

func SubmitOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if order.Email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}

	if err := sendEmail(order); err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
