package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"ecom_go_api/internal/adapters/postgresql"
	"ecom_go_api/internal/categories"
	"ecom_go_api/internal/models"
	"ecom_go_api/internal/orders"
	"ecom_go_api/internal/products"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := postgresql.NewClient(databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Auto Migrate
	if err := db.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderItem{}, &models.Category{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	productService := products.NewService(db)
	orderService := orders.NewService(db)
	categoryService := categories.NewService(db)

	productHandler := products.NewHandler(productService)
	orderHandler := orders.NewHandler(orderService)
	categoryHandler := categories.NewHandler(categoryService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ecom_go_api!"))
	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", categoryHandler.ListCategories)
		r.Post("/", categoryHandler.CreateCategory)
	})

	r.Route("/products", func(r chi.Router) {
		r.Get("/", productHandler.ListProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", orderHandler.PlaceOrder)
		r.Get("/{id}", orderHandler.GetOrder)
		r.Get("/", orderHandler.ListOrders)
		r.Put("/{id}", orderHandler.UpdateOrder)
		r.Delete("/{id}", orderHandler.DeleteOrder)
	})

	log.Println("Server executing on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
