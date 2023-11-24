package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type usersHandler struct {
	usecase IUsersUsecase
}

func NewUsersHandler(usecase IUsersUsecase) IUsersHandler {
	return &usersHandler{
		usecase: usecase,
	}
}

func (h *usersHandler) CreateUser(c *fiber.Ctx) error {
	req := new(CreateUserReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	res, err := h.usecase.CreateUser(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *usersHandler) FindUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	if err := c.BodyParser(email); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	user, err := h.usecase.FindUserByEmail(email)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *usersHandler) FindUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)

	if err := c.BodyParser(id); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	user, err := h.usecase.FindUserByID(idInt)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *usersHandler) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)

	if err := c.BodyParser(id); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	err := h.usecase.DeleteUserByID(idInt)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}

	return nil
}