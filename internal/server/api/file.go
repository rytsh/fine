package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/rytsh/fine/internal/fs"
	"github.com/rytsh/fine/internal/server/msg"
)

// GetFile
//
// @Summary     Get File
// @Description Get File
// @Param       path query string false "file path in server"
// @Router      /file [get]
// @Tags		file
// @Success     200 {file}   binary
// @Failure     400 {object} msg.WebApiError{}
// @Failure     404 {object} msg.WebApiError{}
// @Failure     500 {object} msg.WebApiError{}
func fileGet(c echo.Context) error {
	filePath := c.QueryParam("path")
	if filePath == "" {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: "path is required"})
	}

	return c.File(fs.AddPath(filePath))
}

// PutFile
//
// @Summary     Put File
// @Description Put File
// @Accept      multipart/form-data
// @Produce		json
// @Param       file formData file true "this is a test file"
// @Param       path query string false "file path in server"
// @Router      /file [put]
// @Tags		file
// @Success     200 {object} msg.WebApiSuccess{}
// @Failure     400 {object} msg.WebApiError{}
// @Failure     500 {object} msg.WebApiError{}
func filePut(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg.WebApiError{Err: err.Error()})
	}

	// Source
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: err.Error()})
	}
	defer src.Close()

	if err := fs.Save(file.Filename, src); err != nil {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: err.Error()})
	}

	message := fmt.Sprintf("file %s transfered/updated", file.Filename)

	log.Info().Msg(message)

	return c.JSON(http.StatusOK, msg.WebApiSuccess{Msg: message})
}

// PostFile
//
// @Summary     Post File
// @Description Post File
// @Accept      multipart/form-data
// @Produce		json
// @Param       file formData file true "this is a test file"
// @Param       path query string false "file path in server"
// @Router      /file [post]
// @Tags		file
// @Success     200 {object} msg.WebApiSuccess{}
// @Failure     400 {object} msg.WebApiError{}
// @Failure     500 {object} msg.WebApiError{}
func filePost(c echo.Context) error {
	filePath := c.QueryParam("path")
	if filePath == "" {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: "path is required"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg.WebApiError{Err: err.Error()})
	}

	// check exist
	if fs.IsExist(filePath) {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: fmt.Sprintf("file [%v] already exists", filePath)})
	}

	// Source
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: err.Error()})
	}
	defer src.Close()

	if err := fs.Save(file.Filename, src); err != nil {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: err.Error()})
	}

	message := fmt.Sprintf("file [%s] transfered", file.Filename)

	log.Info().Msg(message)

	return c.JSON(http.StatusOK, msg.WebApiSuccess{Msg: message})
}

// DeleteFile
//
// @Summary     Delete File
// @Description Delete File
// @Produce		json
// @Param       path query string false "file path in server"
// @Router      /file [delete]
// @Tags		file
// @Success     200 {object} msg.WebApiSuccess{}
// @Failure     400 {object} msg.WebApiError{}
// @Failure     500 {object} msg.WebApiError{}
func fileDelete(c echo.Context) error {
	// get the file name from path
	filePath := c.QueryParam("path")

	if err := fs.Delete(filePath); err != nil {
		return c.JSON(http.StatusBadRequest, msg.WebApiError{Err: err.Error()})
	}

	message := fmt.Sprintf("file [%s] deleted", filePath)

	log.Info().Msg(message)

	return c.JSON(http.StatusOK, msg.WebApiSuccess{Msg: message})
}

func File(e *echo.Group) {
	e.GET("/file", fileGet)
	e.PUT("/file", filePut)
	e.POST("/file", filePost)
	e.DELETE("/file", fileDelete)
}
