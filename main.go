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

// Repository is a struct that holds a pointer to a gorm.DB instance.
// This is used to interact with the database.
type Repository struct {
	DB *gorm.DB
}

// CreateBook is a handler function that creates a new book in the database.
// It expects a JSON body with the book details.
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := &models.Book{}

	// Parse the request body into the book struct
	err := context.BodyParser(&book)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Create the book in the database
	result := r.DB.Create(&book)
	if result.Error != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot create book",
		})
	}

	// Return a success message with the created book
	return context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book created successfully",
			"book":    book,
		},
	)
}

// GetAllBooks is a handler function that retrieves all books from the database.
func (r *Repository) GetAllBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	// Find all books in the database
	err := r.DB.Find(bookModels).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot get books",
		})
	}

	// Return a success message with the list of books
	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Books fetched successfully", "books": bookModels})
}

// DeleteBook is a handler function that deletes a book from the database by its ID.
func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "ID is required",
		})
	}

	bookModel := &models.Book{}

	// Delete the book from the database
	err := r.DB.Delete(bookModel, id).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot delete book",
		})
	}

	// Return a success message
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book deleted successfully",
	})
}

// GetBookByID is a handler function that retrieves a single book from the database by its ID.
func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "ID is required",
		})
	}

	bookModel := &models.Book{}

	// Find the book in the database by its ID
	err := r.DB.Where("id=?", id).First(bookModel).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Cannot find book",
		})
	}

	// Return a success message with the book details
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book fetched successfully",
		"book":    bookModel,
	})
}

// SetupRoutes defines the API routes for the application.
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/get_all_books", r.GetAllBooks)
}

// var db *gorm.DB

// main is the entry point of the application.
func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// Create a new database configuration from environment variables
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Establish a new database connection
	db, err := storage.NewConnection(config)

	if err != nil {
		panic(err)
	}

	// Migrate the Book model to the database
	err = models.MigrateBooks(db)
	if err != nil {
		panic(err)
	}

	// Create a new repository with the database connection
	r := Repository{
		DB: db,
	}
	// Create a new Fiber application
	app := fiber.New()
	// Setup the API routes
	r.SetupRoutes(app)

	// Start the Fiber application on port 3000
	app.Listen(":3000")

}
