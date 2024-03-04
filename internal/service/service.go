package service

import (
	"fmt"
	"sldt-test-task/internal/models"
	"time"
)

type VerificationService struct {
}

func NewVerficationService() *VerificationService {
	return &VerificationService{}
}

func (v *VerificationService) Verify(card models.CreditCard) error {
	err := v.verifyExpMonth(card)
	if err != nil {
		return err
	}

	err = v.verifyExpYear(card)
	if err != nil {
		return err
	}

	err = v.verifyCardNum(card.CardNum)
	if err != nil {
		return err
	}

	return nil
}

func (v *VerificationService) verifyExpYear(card models.CreditCard) error {
	currYear := time.Now().Year()
	currMonth := int(time.Now().Month())

	if card.ExpYear < currYear {
		return fmt.Errorf("Your card is expired")
	}

	if card.ExpYear == currYear && card.ExpMonth < currMonth {
		return fmt.Errorf("Your card is expired")
	}

	return nil
}

func (v *VerificationService) verifyExpMonth(card models.CreditCard) error {
	if card.ExpMonth > 12 || card.ExpMonth < 1 {
		return fmt.Errorf("Invalid expiration month")
	}

	return nil
}

func (v *VerificationService) verifyCardNum(number int) error {
	var luhn int

	for i := 1; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur -= 9
			}
		}

		luhn += cur
		number = number / 10
	}

	if luhn%10 != 0 {
		return fmt.Errorf("Invalid card number")
	}

	return nil
}
