package book

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Book struct {
		ID        uuid.UUID
		Title     string
		Author    string
		CreatedAt time.Time
		UpdatedAt time.Time
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

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		uuid := uuid.New()
		b.ID = uuid
	}
	return nil
}

// add book
func (h *Handler) NewBook(c *fiber.Ctx) error {

	book := Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err := h.db.Create(&book).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := map[string]interface{}{
		"success": true,
		"id":      book.ID,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// get book with specific id
func (h *Handler) GetBookByID(c *fiber.Ctx) error {

	id := c.Params("id")
	book := Book{}

	err := h.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(book)
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

	input := Book{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	h.db.Model(&book).Updates(input)

	return c.Status(http.StatusOK).JSON(book)
}

// Delete Book
func (h *Handler) DeleteBook(c *fiber.Ctx) error {

	id := c.Params("id")
	book := Book{}

	err := h.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	h.db.Delete(&book)

	return c.JSON(map[string]interface{}{
		"status": "success",
		"id":     id,
	})

}
