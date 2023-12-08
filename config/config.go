package config

const (
	DBUsername             = "root"
	DBPassword             = "viet1234"
	DBHost                 = "localhost"
	DBPort                 = "3306"
	DBName                 = "yoga_support"
	SubscriptionKey        = "2748de0fd11f4bb6b9e991223f0edccb"
	SavedImagePath         = "/app/saved_frames/images/"
	SavedCroppedImagePath  = "/app/saved_frames/cropped_images/"
	ExecuteTerminalType    = "python"
	ZshPath                = "/app/python/production_run.py"
	AzureComputerVisionURL = "https://yoga-pose-europe.cognitiveservices.azure.com/computervision/imageanalysis:analyze"
	RequestParams          = "%s?api-version=%s&features=people"
	HostURL                = "http://35.247.188.88:80"
	DataSourceName         = "root:viet1234@tcp(10.121.243.54:3307)/yoga_support?charset=utf8mb4&parseTime=True&loc=Local"
)

func GetMySQLURL() string {
	return DBUsername + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?parseTime=true"
}
