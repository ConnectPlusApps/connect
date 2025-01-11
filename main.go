package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/connectplus/models"
	"github.com/connectplus/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	userRepo repositories.UserRepository
	profileRepo repositories.ProfileRepository
	matchRepo repositories.MatchRepository
	messageRepo repositories.MessageRepository
	preferenceRepo repositories.PreferenceRepository
)

// User represents a Connect+ user profile
// @swagger:model
type User struct {
	models.User `gorm:"embedded"`
}

// CreateUserRequest represents the request payload for creating a user
// @swagger:model
type CreateUserRequest struct {
	// Username for the new account
	// required: true
	// minLength: 3
	// maxLength: 20
	// example: john_doe
	Username string `json:"username"`
	
	// Email for the new account
	// required: true
	// example: john@example.com
	Email string `json:"email"`
	
	// Password for the new account
	// required: true
	// minLength: 8
	// example: securePassword123!
	Password string `json:"password"`
}

// CreateUserResponse represents the response after creating a user
// @swagger:model
type CreateUserResponse struct {
	// ID of the created user
	// example: 1
	ID       int    `json:"id"`
	
	// Username of the created user
	// example: john_doe
	Username string `json:"username"`
	
	// Email of the created user
	// example: john@example.com
	Email    string `json:"email"`
	
	// JWT token for authentication
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	Token    string `json:"token"`
}

func initDB() error {
	connStr := "postgres://postgres:G3%7C4gC%3B-%5BH%5EH%60C@34.129.224.98/connect-plus-app?sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Initialize repositories
	userRepo = repositories.NewUserRepository(db)
	profileRepo = repositories.NewProfileRepository(db)
	matchRepo = repositories.NewMatchRepository(db)
	messageRepo = repositories.NewMessageRepository(db)
	preferenceRepo = repositories.NewPreferenceRepository(db)

	// Auto migrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Match{},
		&models.Message{},
		&models.Preference{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto migrate models: %w", err)
	}

	return nil
}

// @title Connect+ API
// @version 1.0
// @description This is the API documentation for Connect+ dating app
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@connectplus.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// rootHandler godoc
// @Summary Show API status
// @Description get API status
// @Tags general
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router / [get]
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to Connect+! The API is running!",
	})
}

// userHandler godoc
// @Summary Get user details
// @Description Get details of the authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} User
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user [get]
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get user ID from JWT token
	tokenString := r.Header.Get("Authorization")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	// Query database for user
	var user models.User
	result := db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// createUserHandler godoc
// @Summary Create a new user
// @Description Create a new user account
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body CreateUserRequest true "User creation details"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/create [post]
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Validate username
	if len(req.Username) < 3 || len(req.Username) > 20 {
		http.Error(w, "Username must be between 3 and 20 characters", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Validate password strength
	if len(req.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	// Check if email already exists
	var count int64
	result := db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count)
	if result.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Create user in database
	user := &models.User{
		Email:    req.Email,
	}
	
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = string(hashedPassword)

	// Create user
	result = db.Create(user)
	if result.Error != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Generate token
	token, err := generateToken(int(user.ID))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{
		ID:       int(user.ID),
		Username: req.Username,
		Email:    req.Email,
		Token:    token,
	})
}

func generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// authMiddleware verifies JWT tokens for protected routes
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// corsMiddleware handles CORS requests
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// UpdateProfileRequest represents the request payload for updating a user profile
// @swagger:model
type UpdateProfileRequest struct {
	// First name
	// example: John
	FirstName        string `json:"first_name"`
	
	// Last name
	// example: Doe
	LastName         string `json:"last_name"`
	
	// User bio
	// example: Love hiking and photography
	Bio              string `json:"bio"`
	
	// Gender identity
	// example: Male
	GenderIdentity   string `json:"gender_identity"`
	
	// Sexual orientation
	// example: Heterosexual
	SexualOrientation string `json:"sexual_orientation"`
	
	// Profile picture URL
	// example: https://example.com/profile.jpg
	ProfilePictureURL string `json:"profile_picture_url"`
	
	// Date of birth (YYYY-MM-DD)
	// example: 1990-01-01
	DateOfBirth      string `json:"date_of_birth"`
	
	// Location
	// example: San Francisco, CA
	Location         string `json:"location"`
}

// updateProfileHandler godoc
// @Summary Update user profile
// @Description Update the authenticated user's profile information
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param profile body UpdateProfileRequest true "Profile update details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/profile [put]
func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from JWT token
	tokenString := r.Header.Get("Authorization")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate date of birth format if provided
	if req.DateOfBirth != "" {
		_, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
	}

	// Update profile in database
	result := db.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"first_name":         req.FirstName,
		"last_name":          req.LastName,
		"bio":                req.Bio,
		"gender_identity":    req.GenderIdentity,
		"sexual_orientation": req.SexualOrientation,
		"profile_picture_url": req.ProfilePictureURL,
		"date_of_birth":      req.DateOfBirth,
		"location":           req.Location,
	})

	if result.Error != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profile updated successfully",
	})
}

// LoginRequest represents the login credentials
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Find user by email
	var user models.User
	result := db.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Verify password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate token
	token, err := generateToken(int(user.ID))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{
		Token: token,
		User: User{User: user},
	})
}

func main() {
	// Initialize logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	// Initialize database connection
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}

	// Create a new ServeMux to handle routes
	mux := http.NewServeMux()

	// Public routes with logging and CORS
	mux.HandleFunc("/", corsMiddleware(loggingMiddleware(rootHandler)))
	mux.HandleFunc("/user/create", corsMiddleware(loggingMiddleware(createUserHandler)))
	mux.HandleFunc("/user/login", corsMiddleware(loggingMiddleware(loginHandler)))
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Protected routes with logging and CORS
	mux.HandleFunc("/user", corsMiddleware(loggingMiddleware(authMiddleware(userHandler))))
	mux.HandleFunc("/user/profile", corsMiddleware(loggingMiddleware(authMiddleware(updateProfileHandler))))
	
	fmt.Println("Server starting on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
