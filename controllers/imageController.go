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
)

// ImageController wraps the ImageDB inside the controller
type ImageController struct {
	imageDB        *database.ImageDB
	rateDB         *database.RateDB
	userDB         *database.UserDB
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
	var ratedImages map[primitive.ObjectID]bool = make(map[primitive.ObjectID]bool)
	var ratedIds []string

	loggedIn, userID := modules.IsAuthed(r)

	filter := bson.D{}

	if loggedIn {
		var user models.User
		uid, _ := primitive.ObjectIDFromHex(userID)

		filter = bson.D{{
			"userId",
			bson.D{{
				"$ne",
				uid,
			}},
		}}

		err := controller.userDB.GetOne(bson.D{{"_id", uid}}, &user)

		if err != nil {
			ratedIds = []string{}
		} else {
			ratedIds = user.RatedImages
		}
	} else {
		ratedIds = modules.RetrieveRatedFromCookie(os.Getenv("RATED_COOKIE"), r)
	}

	for _, id := range ratedIds {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		ratedImages[oid] = true
	}

	filteredImages := []*models.Image{}
	images, err := controller.imageDB.Get(filter)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error consiguiendo las imágenes de la base de datos.")
		return
	}

	for _, img := range images {
		if !ratedImages[img.ID] {
			filteredImages = append(filteredImages, img)
		}
	}

	response := models.ImagesResponse{
		Images: filteredImages,
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
		fmt.Fprint(w, "Error interpretando los datos.")
		return
	}

	// Ensure that both the image and age are contained in multipartform
	if len(r.MultipartForm.File["image"]) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "Falta la imagen.")
		return
	}
	if len(r.MultipartForm.Value["age"]) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "Falta la edad.")
		return
	}
	// get userId from header
	userID := r.Header.Get("uid")

	// get age from image
	age, err := strconv.Atoi(r.MultipartForm.Value["age"][0])
	if err != nil {
		println(err.Error())
		w.WriteHeader(400)
		fmt.Fprint(w, "La edad proporcionada no es un número.")
		return
	}

	// get image file header
	imFileHeader := r.MultipartForm.File["image"][0]
	im, err := imFileHeader.Open()
	defer im.Close()
	if err != nil {
		println(err.Error())
		w.WriteHeader(500)
		fmt.Fprint(w, "Error abriendo archivo de imagen.")
		return
	}

	if _, exists := validImageFormats[imFileHeader.Header["Content-Type"][0]]; !exists {
		w.WriteHeader(400)
		fmt.Fprint(w, "El archivo seleccionado no tiene formato de imagen válido.")
		return
	}

	if imFileHeader.Size/1000000 > 5 {
		w.WriteHeader(413)
		fmt.Fprint(w, "La imagen pesa más de 5 MB.")
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
		fmt.Fprint(w, "Error creando archivo.")
		return
	}

	// write image file to dir
	_, err = io.Copy(file, im)
	if err != nil {
		println(err.Error())
		w.WriteHeader(500)
		fmt.Fprint(w, "Error copiando el archivo de imagen.")
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
		fmt.Fprint(w, "Error al momento de subir la imagen a la base de datos.")
		return
	}
	image.ID = result.InsertedID.(primitive.ObjectID)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(image)
}

// GetUserImages godoc
// @Summary Let a user get its images with statistics
// @Description Get user images
// @ID user-images-endpoint
// @Produce json
// @Security Bearer
// @Success 200 {object} models.UserImagesResponse
// @Failure 401 {string} Authentication error
// @Failure 500 {string} Server error
// @Router /Image/FromUser [get]
func (controller *ImageController) GetUserImages(w http.ResponseWriter, r *http.Request) {
	userID, _ := primitive.ObjectIDFromHex(r.Header.Get("uid"))
	filter := bson.D{{"userId", userID}}
	images, err := controller.imageDB.Get(filter)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error intentando conseguir imágenes de la base de datos.")
		return
	}
	var imagesResponse []*models.StatisticalImage
	for _, image := range images {
		var imageStats *models.StatisticalImage
		registeredGuess := &models.ImageGuess{
			Quantity: 0,
			Correct:  0,
		}
		unregisteredGuess := &models.ImageGuess{
			Quantity: 0,
			Correct:  0,
		}

		ratesFilter := bson.D{{"imageid", image.ID}}
		rates, err := controller.rateDB.Get(ratesFilter)
		if err == nil {
			for _, rate := range rates {
				if rate.FromAuth {
					registeredGuess.Quantity++
					if rate.GuessedAge == image.Age {
						registeredGuess.Correct++
					}
				} else {
					unregisteredGuess.Quantity++
					if rate.GuessedAge == image.Age {
						unregisteredGuess.Correct++
					}
				}
			}
			imageStats = &models.StatisticalImage{
				ID:                  image.ID,
				URL:                 image.URL,
				Age:                 image.Age,
				CreatedAt:           image.CreatedAt,
				RegisteredGuesses:   registeredGuess,
				UnregisteredGuesses: unregisteredGuess,
			}
			imagesResponse = append(imagesResponse, imageStats)
		}
	}

	response := models.UserImagesResponse{
		Images: imagesResponse,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// RateImage godoc
// @Summary Allows a user to rate an image
// @ID rate-image-endpoint
// @Accept json
// @Produce json
// @Param guess body models.AgeGuess true "Guess attempt from the user"
// @Param id path string true "ID of the image that needs to be rated"
// @Success 200 {object} models.GuessResponse
// @Failure 400 {string} Bad request
// @Failure 404 {string} Image not found
// @Failure 409 {string} Rate conflict
// @Failure 500 {string} Server error
// @Router /Image/{id}/Rate [post]
func (controller *ImageController) RateImage(w http.ResponseWriter, r *http.Request) {
	var image models.Image
	var guess models.AgeGuess
	var err error

	decoder := json.NewDecoder(r.Body)
	imageID := mux.Vars(r)["id"]
	loggedIn, userID := modules.IsAuthed(r)
	err = decoder.Decode(&guess)

	if err != nil || guess.Age == nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Faltan parámetros")
		return
	}

	iid, err := primitive.ObjectIDFromHex(imageID)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Formato de imágen incorrecto")
		return
	}

	err = controller.imageDB.GetOne(bson.D{{"_id", iid}}, &image)

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "No se encontró la imagen")
		return
	}

	if !loggedIn {
		ratedImages := modules.RetrieveRatedFromCookie(os.Getenv("RATED_COOKIE"), r)

		if modules.Contains(ratedImages, imageID) {
			w.WriteHeader(409)
			fmt.Fprintf(w, "La imágen ya ha sido calificada")
			return
		}
		rate := models.Rate{
			ImageID:    iid,
			FromAuth:   false,
			GuessedAge: *guess.Age,
		}

		_, err = controller.rateDB.Insert(rate)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Ocurrió un error al calificar la imágen")
			return
		}
		modules.AddCookieValue(os.Getenv("RATED_COOKIE"), imageID, w, r)
	} else {
		var user models.User
		uid, _ := primitive.ObjectIDFromHex(userID)
		userFilter := bson.D{{"_id", uid}}

		err = controller.userDB.GetOne(userFilter, &user)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Ocurrió un error al calificar la imágen")
			return
		}

		ratedImages := user.RatedImages

		if modules.Contains(ratedImages, imageID) {
			w.WriteHeader(409)
			fmt.Fprintf(w, "La imágen ya ha sido calificada")
			return
		}
		rate := models.Rate{
			ImageID:    iid,
			FromAuth:   true,
			GuessedAge: *guess.Age,
		}

		_, err = controller.rateDB.Insert(rate)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Ocurrió un error al calificar la imágen")
			return
		}

		ratedImages = append(ratedImages, imageID)
		updatedUser := bson.D{
			{"$set",
				bson.D{
					{"ratedImages", ratedImages},
				}},
		}
		_, err = controller.userDB.UpdateOne(userFilter, updatedUser)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	correctGuess, responseMessage := modules.CalculateAgeGuessResponse(image.Age, *guess.Age)
	response := models.GuessResponse{
		Correct: correctGuess,
		Message: responseMessage,
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
	r.HandleFunc("/{id}/Rate", controller.RateImage).Methods(http.MethodPost)
}

// SetImageController creates the ImageController and wraps the user collection into ImageDB
func SetImageController(r *mux.Router, db *mongo.Database, redisClient *redis.Client) {
	image := database.ImageDB{Images: db.Collection("images")}
	rate := database.RateDB{Rates: db.Collection("rates")}
	user := database.UserDB{Users: db.Collection("users")}

	ImageController := ImageController{
		imageDB:     &image,
		rateDB:      &rate,
		userDB:      &user,
		redisClient: redisClient,
		authMiddleware: &auth.Middleware{
			RedisClient: redisClient,
		},
	}
	ImageController.InitializeController(r)
}
