package common

import (
	"os"

	"github.com/aki36-an/cloudrun-demo/app/util"
)

var TemplatePathInfo = func() util.FileInfo {
	filePath := os.Getenv("APP_TEMPLATE_FILE")
	return util.GetAppFilePath(filePath)
}

var ProjectId = func() string {
	return os.Getenv("PROJECT_ID")
}
