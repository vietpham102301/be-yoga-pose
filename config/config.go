package config

const (
	DBUsername             = "root"
	DBPassword             = "viet1234"
	DBHost                 = "localhost"
	DBPort                 = "3306"
	DBName                 = "yoga_support"
	SubscriptionKey        = "2748de0fd11f4bb6b9e991223f0edccb"
	SavedImagePath         = "/Users/vietpham1023/Desktop/saved_frames/images/"
	SavedCroppedImagePath  = "/Users/vietpham1023/Desktop/saved_frames/cropped_images/"
	ExecuteTerminalType    = "zsh"
	ZshPath                = "/Users/vietpham1023/Desktop/zsh_yoga/cnn.zsh"
	AzureComputerVisionURL = "https://yoga-pose-europe.cognitiveservices.azure.com/computervision/imageanalysis:analyze"
	RequestParams          = "%s?api-version=%s&features=people"
	HostURL                = "http://localhost:8080"
	DataSourceName         = "root:viet1234@tcp(mysql:3306)/yoga_support?charset=utf8mb4&parseTime=True&loc=Local"
)

func GetMySQLURL() string {
	return DBUsername + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?parseTime=true"
}
