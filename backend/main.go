package main

import (
	"log"
	"net/http"

	"book-recommendation-system/backend/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"book-recommendation-system/backend/handlers"
	"book-recommendation-system/backend/repositories"
	"book-recommendation-system/backend/services"
)

func main() {
	// Database configuration
	dbConfig := database.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "user",     // Replace with your DB user
		Password: "password", // Replace with your DB password
		DBName:   "bookrecsys",
	}

	db, err := database.ConnectDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseDB(db)
	log.Println("Successfully connected to the database!")

	// Initialize repositories
	bookRepo := repositories.NewBookRepository(db)
	authorRepo := repositories.NewAuthorRepository(db)
	genreRepo := repositories.NewGenreRepository(db)
	libraryRepo := repositories.NewLibraryRepository(db)
	userInteractionRepo := repositories.NewUserInteractionRepository(db)
	recommendationRepo := repositories.NewRecommendationRepository(db)

	// Initialize services
	bookService := services.NewBookService(bookRepo)
	authorService := services.NewAuthorService(authorRepo)
	genreService := services.NewGenreService(genreRepo)
	libraryService := services.NewLibraryService(libraryRepo)
	userInteractionService := services.NewUserInteractionService(userInteractionRepo)
	recommendationService := services.NewRecommendationService(recommendationRepo)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookService)
	authorHandler := handlers.NewAuthorHandler(authorService)
	genreHandler := handlers.NewGenreHandler(genreService)
	libraryHandler := handlers.NewLibraryHandler(libraryService)
	userInteractionHandler := handlers.NewUserInteractionHandler(userInteractionService)
	recommendationHandler := handlers.NewRecommendationHandler(recommendationService)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	// Setup routes
	setupRoutes(r, bookHandler, authorHandler, genreHandler, libraryHandler, userInteractionHandler, recommendationHandler)

	fmt.Println("Server starting on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// setupRoutes configures all the API routes.
func setupRoutes(r *chi.Mux, bookH *handlers.BookHandler, authorH *handlers.AuthorHandler, genreH *handlers.GenreHandler, libraryH *handlers.LibraryHandler, userInteractionH *handlers.UserInteractionHandler, recommendationH *handlers.RecommendationHandler) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Book Recommendation System Backend!"))
	})

	r.Route("/books", func(r chi.Router) {
		r.Post("/", bookH.CreateBook)
		r.Get("/", bookH.GetAllBooks)
		r.Get("/{id}", bookH.GetBookByID)
		r.Put("/{id}", bookH.UpdateBook)
		r.Delete("/{id}", bookH.DeleteBook)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Post("/", authorH.CreateAuthor)
		r.Get("/", authorH.GetAllAuthors)
		r.Get("/{id}", authorH.GetAuthorByID)
		r.Put("/{id}", authorH.UpdateAuthor)
		r.Delete("/{id}", authorH.DeleteAuthor)
	})

	r.Route("/genres", func(r chi.Router) {
		r.Post("/", genreH.CreateGenre)
		r.Get("/", genreH.GetAllGenres)
		r.Get("/{id}", genreH.GetGenreByID)
		r.Put("/{id}", genreH.UpdateGenre)
		r.Delete("/{id}", genreH.DeleteGenre) // Added missing delete
	})

	r.Route("/libraries", func(r chi.Router) {
		r.Post("/", libraryH.CreateLibrary)
		r.Get("/", libraryH.GetAllLibraries)
		r.Get("/{id}", libraryH.GetLibraryByID)
		r.Put("/{id}", libraryH.UpdateLibrary)
		r.Delete("/{id}", libraryH.DeleteLibrary) // Added missing delete
	})

	r.Route("/user-interactions", func(r chi.Router) {
		r.Post("/", userInteractionH.CreateUserInteraction)
		r.Get("/", userInteractionH.GetAllUserInteractions) // Changed to GetAll
		r.Get("/{id}", userInteractionH.GetUserInteractionByID)
		r.Get("/user/{userID}", userInteractionH.GetUserInteractionsByUserID)
		r.Put("/{id}", userInteractionH.UpdateUserInteraction)
		r.Delete("/{id}", userInteractionH.DeleteUserInteraction)
	})

	r.Route("/recommendations", func(r chi.Router) {
		r.Post("/", recommendationH.CreateRecommendation)
		r.Get("/", recommendationH.GetAllRecommendations) // Changed to GetAll
		r.Get("/{id}", recommendationH.GetRecommendationByID)
		r.Get("/user/{userID}", recommendationH.GetRecommendationsByUserID)
		r.Put("/{id}", recommendationH.UpdateRecommendation)
		r.Delete("/{id}", recommendationH.DeleteRecommendation)
	})
}
