package book

import (
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


// db handler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

// add book
func (h *Handler) NewBook(c *fiber.Ctx) error {

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

// get book with specific id
func (h *Handler) GetBookByID(c *fiber.Ctx) error {
	
	id := c.Params("id")
	book := Book{}

	err := h.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	
	return c.JSON(book)
}

// get all books
func (h *Handler) GetBooks(c *fiber.Ctx) error {

	books := []Book{}

	err := h.db.Find(&books).Error
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(books)
}

// Update book
func (h *Handler) UpdateBook(c *fiber.Ctx) error {

	id := c.Params("id")
	book := Book{}

	err := h.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	input:= Book{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	
	h.db.Model(&book)
	
	return c.Status(http.StatusOK).JSON(book)
}

// Delete Book
// func (h *Handler) DeleteBook(c *fiber.Ctx) error {

// 	id := c.Params("id")
// 	book := Book {}

// 	err := h.db.Where("id = ?", id).First(&book).Error
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(err.Error())
// 	}

// 	del := book

// }


