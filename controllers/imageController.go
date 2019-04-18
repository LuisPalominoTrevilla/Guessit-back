package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/LuisPalominoTrevilla/Guessit-back/modules"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"

	auth "github.com/LuisPalominoTrevilla/Guessit-back/authentication"
	database "github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/redis"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/LuisPalominoTrevilla/Guessit-back/authentication"
)

// ImageController wraps the ImageDB inside the controller
type ImageController struct {
	imageDB        *database.ImageDB
	redisClient    *redis.Client
	authMiddleware *auth.Middleware
}

// Get godoc
// @Summary Retrieve all images. If user is logged in, his/her images are not returned.
// @Description Get all images
// @ID get-images-endpoint
// @Produce json
// @Success 200 {object} models.ImagesResponse
// @Failure 500 {string} Server error
// @Router /Image/ [get]
func (controller *ImageController) Get(w http.ResponseWriter, r *http.Request) {
	var userID string
	// var ratedImages []primitive.ObjectID
	logedIn := false

	auth := strings.Fields(r.Header.Get("Authorization"))
	if len(auth) > 1 && auth[0] == "Bearer" {
		claims, err := authentication.VerifyJWT(auth[1])

		if err == nil {
			userID = claims["userId"].(string)
			logedIn = true
		}
	}

	filter := bson.D{}

	if !logedIn {
		uid, _ := primitive.ObjectIDFromHex(userID)

		filter = bson.D{{
			"userId",
			bson.D{{
				"$ne",
				uid,
			}},
		}}

		ratedImageIds := modules.RetrieveRatedFromCookie("ratedPics", r)

		fmt.Println("Rated images are ", ratedImageIds)
	}

	images, err := controller.imageDB.Get(filter)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error trying to retrieve images from db")
		return
	}

	response := models.ImagesResponse{
		Images: images,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// UploadImage godoc
// @Summary Let a user upload images
// @Description Upload an image
// @ID upload-image-endpoint
// @Accept mpfd
// @Produce plain
// @Param image formData file true "Image to be uploaded"
// @Param age formData string true "Age that corresponts to the person in the image"
// @Security Bearer
// @Success 200 {object} models.Image
// @Failure 400 {string} Bad request
// @Failure 401 {string} Authentication error
// @Failure 413 {string} File too large
// @Failure 500 {string} Server error
// @Router /Image/UploadImage [post]
func (controller *ImageController) UploadImage(w http.ResponseWriter, r *http.Request) {
	var maxBytes int64 = 64 * 1024 * 1024
	validImageFormats := map[string]bool{
		"image/png":  true,
		"image/jpeg": true,
	}

	// Parse multipart form data
	err := r.ParseMultipartForm(maxBytes)
	if err != nil {
		println(err.Error())
		w.WriteHeader(500)
		fmt.Fprint(w, "Error parsing multiform data")
		return
	}

	// Ensure that both the image and age are contained in multipartform
	if len(r.MultipartForm.File["image"]) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "Missing image")
		return
	}
	if len(r.MultipartForm.Value["age"]) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "Missing age")
		return
	}
	// get userId from header
	userID := r.Header.Get("uid")

	// get age from image
	age, err := strconv.Atoi(r.MultipartForm.Value["age"][0])
	if err != nil {
		println(err.Error())
		w.WriteHeader(400)
		fmt.Fprint(w, "Age is not a number")
		return
	}

	// get image file header
	imFileHeader := r.MultipartForm.File["image"][0]
	im, err := imFileHeader.Open()
	defer im.Close()
	if err != nil {
		println(err.Error())
		w.WriteHeader(500)
		fmt.Fprint(w, "Error opening image file")
		return
	}

	if _, exists := validImageFormats[imFileHeader.Header["Content-Type"][0]]; !exists {
		w.WriteHeader(400)
		fmt.Fprint(w, "File uploaded does not have a valid image format")
		return
	}

	if imFileHeader.Size/1000000 > 5 {
		w.WriteHeader(413)
		fmt.Fprint(w, "Image uploaded is more than 5 MB")
		return
	}

	imageURL := "/" + userID

	// ensure dir exists and create final file
	os.MkdirAll("/static"+imageURL, os.ModePerm)
	imageURL += "/"
	filename := imFileHeader.Filename
	filename = strings.Replace(filename, " ", "", -1)

	additionalNum := ""
	for fileExists("/static" + imageURL + additionalNum + filename) {
		if additionalNum == "" {
			additionalNum = "1"
		} else {
			num, _ := strconv.Atoi(additionalNum)
			num++
			additionalNum = strconv.Itoa(num)
		}
	}
	imageURL += additionalNum + filename

	file, err := os.Create("/static" + imageURL)
	defer file.Close()
	if err != nil {
		println(err.Error())
		w.WriteHeader(500)
		fmt.Fprint(w, "Error creating file")
		return
	}

	// write image file to dir
	_, err = io.Copy(file, im)
	if err != nil {
		println(err.Error())
		w.WriteHeader(500)
		fmt.Fprint(w, "Error copying image file")
		return
	}

	oid, _ := primitive.ObjectIDFromHex(userID)

	image := models.Image{
		URL:       "/images" + imageURL,
		Age:       age,
		Owner:     oid,
		CreatedAt: time.Now(),
	}

	result, err := controller.imageDB.Insert(image)

	if err != nil {
		println(err.Error())
		_ = os.Remove(("/static" + imageURL))
		w.WriteHeader(500)
		fmt.Fprint(w, "Error trying to insert image into db")
		return
	}
	image.ID = result.InsertedID.(primitive.ObjectID)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(image)
}

// GetUserImages godoc
// @Summary Let a user get its images
// @Description Get user images
// @ID user-images-endpoint
// @Produce json
// @Security Bearer
// @Success 200 {object} models.ImagesResponse
// @Failure 401 {string} Authentication error
// @Failure 500 {string} Server error
// @Router /Image/FromUser [get]
func (controller *ImageController) GetUserImages(w http.ResponseWriter, r *http.Request) {
	userID, _ := primitive.ObjectIDFromHex(r.Header.Get("uid"))
	filter := bson.D{{"userId", userID}}
	images, err := controller.imageDB.Get(filter)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error trying to retrieve images from db")
		return
	}

	response := models.ImagesResponse{
		Images: images,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// InitializeController initializes the routes
func (controller *ImageController) InitializeController(r *mux.Router) {
	r.HandleFunc("/", controller.Get).Methods(http.MethodGet)
	r.Handle("/UploadImage", controller.authMiddleware.AccessControl(controller.UploadImage)).Methods(http.MethodPost)
	r.Handle("/FromUser", controller.authMiddleware.AccessControl(controller.GetUserImages)).Methods(http.MethodGet)
}

// SetImageController creates the ImageController and wraps the user collection into ImageDB
func SetImageController(r *mux.Router, db *mongo.Database, redisClient *redis.Client) {
	image := database.ImageDB{Images: db.Collection("images")}
	ImageController := ImageController{
		imageDB:     &image,
		redisClient: redisClient,
		authMiddleware: &auth.Middleware{
			RedisClient: redisClient,
		},
	}
	ImageController.InitializeController(r)
}
