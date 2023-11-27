package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"yoga-pose-backend/config"
)

type BoundingBox struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type Person struct {
	BoundingBox BoundingBox `json:"boundingBox"`
	Confidence  float64     `json:"confidence"`
}

type PeopleResult struct {
	Values []Person `json:"values"`
}

type ImageAnalysisResult struct {
	ModelVersion string `json:"modelVersion"`
	Metadata     struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"metadata"`
	PeopleResult PeopleResult `json:"peopleResult"`
}

type PredictedResponse struct {
	Status         string  `json:"status"`
	PredictedClass string  `json:"predicted_class"`
	Confidence     float64 `json:"confidence"`
	ExecutionTime  string  `json:"execution_time"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleClientConnection(c *gin.Context) {

	fmt.Printf("start handleClientConnection \n")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	for {
		start := time.Now()
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if messageType != websocket.BinaryMessage {
			log.Println("Received non-binary message, skipping")
			continue
		}

		filePath, err := handleImage(data)
		if err != nil {
			log.Println("Error handling image:", err)
		}

		analysisResult, err := localization(*filePath, config.SubscriptionKey)
		if err != nil {
			log.Println("Error analyzing image:", err)
			return
		}

		bounds := analysisResult.PeopleResult.Values[0].BoundingBox
		croppedImagePath := config.SavedCroppedImagePath + generateRandomFileName("_crop.png")
		if err := cropAndSaveImage(*filePath, croppedImagePath, bounds); err != nil {
			log.Println("Error cropping image:", err)
			return
		}

		cmd := exec.Command(config.ExecuteTerminalType, config.ZshPath, croppedImagePath)
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			println(err.Error())
			return
		}
		println(string(stdout))
		stdoutStr := string(stdout)

		predictedClass := ""
		confidence := 0.0
		executionTime := ""

		lines := strings.Split(stdoutStr, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Predicted Class:") {
				predictedClass = strings.TrimSpace(strings.TrimPrefix(line, "Predicted Class:"))
			} else if strings.HasPrefix(line, "Confidence:") {
				confidenceStr := strings.TrimSpace(strings.TrimPrefix(line, "Confidence:"))
				confidence, _ = strconv.ParseFloat(confidenceStr, 64)
			} else if strings.HasPrefix(line, "Execution time:") {
				executionTime = strings.TrimSpace(strings.TrimPrefix(line, "Execution time:"))
			}
		}

		response := PredictedResponse{
			Status:         "success",
			PredictedClass: predictedClass,
			Confidence:     confidence,
			ExecutionTime:  executionTime,
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			log.Println("Error marshaling response to JSON:", err)
			return
		}

		// Send the JSON response back to the client
		if err := conn.WriteMessage(messageType, responseJSON); err != nil {
			log.Println(err)
			return
		}

		if err := os.Remove(*filePath); err != nil {
			log.Println("Error deleting image:", err)
		}
		// Delete the cropped image
		//if err := os.Remove(croppedImagePath); err != nil {
		//	log.Println("Error deleting cropped image:", err)
		//}
		elapsed := time.Since(start)
		fmt.Printf("execute-time: %s \n", elapsed)
	}
	
}

func handleImage(data []byte) (*string, error) {
	reader := bytes.NewReader(data)

	// Decode the image using the imaging library
	img, err := imaging.Decode(reader)
	if err != nil {
		log.Println("Error decoding image:", err)
		return nil, err
	}

	width := 800
	height := 600
	img = imaging.Resize(img, width, height, imaging.Lanczos)

	fileName := config.SavedImagePath + generateRandomFileName(".png")

	if err := imaging.Save(img, fileName); err != nil {
		log.Println("Error saving image:", err)
		return nil, err
	}

	return &fileName, nil
}

func generateRandomFileName(extension string) string {
	rand.NewSource(time.Now().UnixNano())
	randomString := fmt.Sprintf("%d", rand.Intn(10000)) // Random number
	timestamp := time.Now().Unix()                      // Current timestamp

	fileName := fmt.Sprintf("%s_%d%s", randomString, timestamp, extension)

	return fileName
}

func cropAndSaveImage(inputImagePath, outputImagePath string, bounds BoundingBox) error {
	img, err := gg.LoadImage(inputImagePath)
	if err != nil {
		return err
	}
	dc := gg.NewContext(bounds.W, bounds.H)
	dc.DrawImage(img, -bounds.X, -bounds.Y)
	err = dc.SavePNG(outputImagePath)
	if err != nil {
		return err
	}

	return nil
}

// localization performs image analysis using Azure Cognitive Services Computer Vision API
func localization(imagePath, subscriptionKey string) (*ImageAnalysisResult, error) {
	apiURL := config.AzureComputerVisionURL
	apiVersion := "2023-02-01-preview"

	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("error reading the image: %v", err)
	}

	url := fmt.Sprintf(config.RequestParams, apiURL, apiVersion)

	body := bytes.NewReader(imageData)

	client := &http.Client{}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Ocp-Apim-Subscription-Key", subscriptionKey)

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}(response.Body)

	// Read the response body
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if response.StatusCode == http.StatusOK {
		var result ImageAnalysisResult
		if err := json.Unmarshal(responseData, &result); err != nil {
			return nil, fmt.Errorf("error parsing JSON: %v", err)
		}
		return &result, nil
	}

	return nil, fmt.Errorf("error: %d - %s", response.StatusCode, responseData)
}
