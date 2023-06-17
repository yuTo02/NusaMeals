package usecase

import (
	"reglog/internal/dtos"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type PaymentUseCase interface {
	CreatePayment(paymentDTO *dtos.PaymentDTO) (*dtos.PaymentDTO, error)
	UpdatePayment(paymentID uint, paymentDTO *dtos.PaymentDTO) (*dtos.PaymentDTO, error)
	DeletePayment(paymentID uint) error
	GetPaymentByID(paymentID uint) (*dtos.PaymentDTO, error)
	GetAllPayments() ([]dtos.PaymentDTO, error)
}

type paymentUseCase struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentUseCase(paymentRepo repository.PaymentRepository) PaymentUseCase {
	return &paymentUseCase{
		paymentRepo: paymentRepo,
	}
}

func (u *paymentUseCase) CreatePayment(paymentDTO *dtos.PaymentDTO) (*dtos.PaymentDTO, error) {
	payment := &model.Payment{
		OrderID:     paymentDTO.OrderID,
		UserID:      paymentDTO.UserID,
		Amount:      paymentDTO.Amount,
		Method:      paymentDTO.Method,
		PaymentType: paymentDTO.PaymentType,
		// Set other fields accordingly
	}

	createdPayment, err := u.paymentRepo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	createdPaymentDTO := &dtos.PaymentDTO{
		ID:          createdPayment.ID,
		OrderID:     createdPayment.OrderID,
		UserID:      createdPayment.UserID,
		Amount:      createdPayment.Amount,
		Method:      createdPayment.Method,
		PaymentType: createdPayment.PaymentType,
		// Set other fields accordingly
	}

	return createdPaymentDTO, nil
}

func (u *paymentUseCase) UpdatePayment(paymentID uint, paymentDTO *dtos.PaymentDTO) (*dtos.PaymentDTO, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment.OrderID = paymentDTO.OrderID
	payment.UserID = paymentDTO.UserID
	payment.Amount = paymentDTO.Amount
	payment.Method = paymentDTO.Method
	payment.PaymentType = paymentDTO.PaymentType
	// Update other fields accordingly

	updatedPayment, err := u.paymentRepo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	updatedPaymentDTO := &dtos.PaymentDTO{
		ID:          updatedPayment.ID,
		OrderID:     updatedPayment.OrderID,
		UserID:      updatedPayment.UserID,
		Amount:      updatedPayment.Amount,
		Method:      updatedPayment.Method,
		PaymentType: updatedPayment.PaymentType,
		// Set other fields accordingly
	}

	return updatedPaymentDTO, nil
}

func (u *paymentUseCase) DeletePayment(paymentID uint) error {
	return u.paymentRepo.DeletePayment(paymentID)
}

func (u *paymentUseCase) GetPaymentByID(paymentID uint) (*dtos.PaymentDTO, error) {
	payment, err := u.paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	paymentDTO := &dtos.PaymentDTO{
		ID:          payment.ID,
		OrderID:     payment.OrderID,
		UserID:      payment.UserID,
		Amount:      payment.Amount,
		Method:      payment.Method,
		PaymentType: payment.PaymentType,
		// Set other fields accordingly
	}

	return paymentDTO, nil
}

func (u *paymentUseCase) GetAllPayments() ([]dtos.PaymentDTO, error) {
	payments, err := u.paymentRepo.GetAllPayments()
	if err != nil {
		return nil, err
	}

	paymentDTOs := make([]dtos.PaymentDTO, len(payments))
	for i, payment := range payments {
		paymentDTO := dtos.PaymentDTO{
			ID:          payment.ID,
			OrderID:     payment.OrderID,
			UserID:      payment.UserID,
			Amount:      payment.Amount,
			Method:      payment.Method,
			PaymentType: payment.PaymentType,
			// Set other fields accordingly
		}
		paymentDTOs[i] = paymentDTO
	}

	return paymentDTOs, nil
}
