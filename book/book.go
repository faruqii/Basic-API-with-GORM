package book

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Book struct {
		ID     uuid.UUID
		Title  string
		Author string
	}

	Handler struct {
		db *gorm.DB
	}
)

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetBookByID(c *fiber.Ctx) error {
	
	id := c.Params("id")
	log.Println(id)
	book := Book{}

	err := h.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	
	return c.JSON(book)
}

func (h *Handler) GetBooks(c *fiber.Ctx) error {

	books := []Book{}

	err := h.db.Find(&books).Error
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(books)
}

func (h *Handler) NewBooks(c *fiber.Ctx) error {

	book := Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err := h.db.Create(book).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON("SUCCESS")
}
