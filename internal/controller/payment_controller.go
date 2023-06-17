package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reglog/internal/dtos"
	"reglog/internal/usecase"
	"strconv"
)

type PaymentController struct {
	paymentUseCase usecase.PaymentUseCase
}

func NewPaymentController(paymentUseCase usecase.PaymentUseCase) *PaymentController {
	return &PaymentController{
		paymentUseCase: paymentUseCase,
	}
}

func (c *PaymentController) CreatePayment(ctx echo.Context) error {
	var paymentDTO dtos.PaymentDTO
	if err := ctx.Bind(&paymentDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	createdPayment, err := c.paymentUseCase.CreatePayment(&paymentDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, createdPayment)
}

func (c *PaymentController) UpdatePayment(ctx echo.Context) error {
	paymentID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var paymentDTO dtos.PaymentDTO
	if err := ctx.Bind(&paymentDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updatedPayment, err := c.paymentUseCase.UpdatePayment(uint(paymentID), &paymentDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updatedPayment)
}

func (c *PaymentController) DeletePayment(ctx echo.Context) error {
	paymentID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	err = c.paymentUseCase.DeletePayment(uint(paymentID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *PaymentController) GetPaymentByID(ctx echo.Context) error {
	paymentID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	payment, err := c.paymentUseCase.GetPaymentByID(uint(paymentID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, payment)
}

func (c *PaymentController) GetAllPayments(ctx echo.Context) error {
	payments, err := c.paymentUseCase.GetAllPayments()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, payments)
}
