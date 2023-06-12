package docs

import (
	"path"

	"github.com/swaggo/swag"
)

func SetVersion(appName, appVersion, basePath string) {
	if spec, ok := swag.GetSwagger("swagger").(*swag.Spec); ok {
		spec.Title = appName
		spec.Version = appVersion
		spec.BasePath = path.Join("/", basePath, spec.BasePath)
	}
}
