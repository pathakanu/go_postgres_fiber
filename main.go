package main

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pathakanu/go_postgres_fiber/models"
	"github.com/pathakanu/go_postgres_fiber/storage"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := &models.Book{}

	err := context.BodyParser(&book)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result := r.DB.Create(&book)
	if result.Error != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot create book",
		})
	}

	return context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book created successfully",
			"book":    book,
		},
	)
}

func (r *Repository) GetAllBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot get books",
		})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Books fetched successfully", "books": bookModels})
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "ID is required",
		})
	}

	bookModel := &models.Book{}

	err := r.DB.Delete(bookModel, id).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot delete book",
		})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book deleted successfully",
	})
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "ID is required",
		})
	}

	bookModel := &models.Book{}

	err := r.DB.Where("id=?", id).First(bookModel).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot find book",
		})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book fetched successfully",
		"book":    bookModel,
	})
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/get_all_books", r.GetAllBooks)
}

// var db *gorm.DB

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		panic(err)
	}

	err = models.MigrateBooks(db)
	if err != nil {
		panic(err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)

	app.Listen(":3000")

}
