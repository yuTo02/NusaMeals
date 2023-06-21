package usecase

import (
	"errors"
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type PaymentUseCase interface {
	CreatePayment(paymentRequest *request.Payment) (*response.Payment, error)
	UpdatePayment(paymentID uint, paymentRequest *request.Payment) (*response.Payment, error)
	UpdatePaymentByAdmin(paymentID uint, paymentRequest *request.PaymentUpdate) (*response.PaymentUpdate, error)
	DeletePayment(paymentID uint) error
	GetPaymentByID(paymentID uint) (*response.Payment, error)
	GetAllPayments() ([]response.Payment, error)
}

type paymentUseCase struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentUseCase(paymentRepo repository.PaymentRepository) PaymentUseCase {
	return &paymentUseCase{
		paymentRepo: paymentRepo,
	}
}

func (u *paymentUseCase) CreatePayment(paymentRequest *request.Payment) (*response.Payment, error) {
	payment := &model.Payment{
		OrderID:     paymentRequest.OrderID,
		UserID:      paymentRequest.UserID,
		Amount:      paymentRequest.Amount,
		Method:      paymentRequest.Method,
		PaymentType: paymentRequest.PaymentType,
		Status:      "on progress",
		// Set other fields accordingly
	}

	createdPayment, err := u.paymentRepo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	createdPaymentResponse := &response.Payment{
		ID:          createdPayment.ID,
		OrderID:     createdPayment.OrderID,
		UserID:      createdPayment.UserID,
		Amount:      createdPayment.Amount,
		Status:      createdPayment.Status,
		Method:      createdPayment.Method,
		PaymentType: createdPayment.PaymentType,
		// Set other fields accordingly
	}

	return createdPaymentResponse, nil
}

func (u *paymentUseCase) UpdatePayment(paymentID uint, paymentRequest *request.Payment) (*response.Payment, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment.Method = paymentRequest.Method
	payment.PaymentType = paymentRequest.PaymentType
	// Update other fields accordingly

	updatedPayment, err := u.paymentRepo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	updatedPaymentResponse := &response.Payment{
		ID:          updatedPayment.ID,
		OrderID:     updatedPayment.OrderID,
		UserID:      updatedPayment.UserID,
		Amount:      updatedPayment.Amount,
		Status:      updatedPayment.Status,
		Method:      updatedPayment.Method,
		PaymentType: updatedPayment.PaymentType,
		// Set other fields accordingly
	}

	return updatedPaymentResponse, nil
}

func (u *paymentUseCase) UpdatePaymentByAdmin(paymentID uint, paymentRequest *request.PaymentUpdate) (*response.PaymentUpdate, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}
	payment.Status = paymentRequest.Status
	// Update other fields accordingly

	updatedPayment, err := u.paymentRepo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	updatedPaymentResponse := &response.PaymentUpdate{
		ID:          updatedPayment.ID,
		OrderID:     updatedPayment.OrderID,
		UserID:      updatedPayment.UserID,
		Amount:      updatedPayment.Amount,
		Status:      updatedPayment.Status,
		Method:      updatedPayment.Method,
		PaymentType: updatedPayment.PaymentType,
		// Set other fields accordingly
	}

	return updatedPaymentResponse, nil
}

func (u *paymentUseCase) DeletePayment(paymentID uint) error {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return err
	}

	if payment == nil {
		return errors.New("payment not found")
	}

	return u.paymentRepo.DeletePayment(paymentID)
}

func (u *paymentUseCase) GetPaymentByID(paymentID uint) (*response.Payment, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	paymentResponse := &response.Payment{
		ID:          payment.ID,
		OrderID:     payment.OrderID,
		UserID:      payment.UserID,
		Amount:      payment.Amount,
		Status:      payment.Status,
		Method:      payment.Method,
		PaymentType: payment.PaymentType,
		// Set other fields accordingly
	}

	return paymentResponse, nil
}

func (u *paymentUseCase) GetAllPayments() ([]response.Payment, error) {
	payments, err := u.paymentRepo.GetAllPayments()
	if err != nil {
		return nil, err
	}

	paymentResponses := make([]response.Payment, len(payments))
	for i, payment := range payments {
		paymentResponse := response.Payment{
			ID:          payment.ID,
			OrderID:     payment.OrderID,
			UserID:      payment.UserID,
			Amount:      payment.Amount,
			Status:      payment.Status,
			Method:      payment.Method,
			PaymentType: payment.PaymentType,
			// Set other fields accordingly
		}
		paymentResponses[i] = paymentResponse
	}

	return paymentResponses, nil
}
