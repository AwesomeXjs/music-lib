package controller

import (
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary Create song
// @Tags Song
// @Description add song to library
// @ID create-song
// @Accept  json
// @Produce  json
// @Param input body model.SongCreate true "song info"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 422 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /v1/songs/ [post]
func (e *Controller) CreateSong(ctx echo.Context) error {
	var input model.SongCreate

	if err := ctx.Bind(&input); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusBadRequest, helpers.JSON_PARSE_ERROR, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusUnprocessableEntity, helpers.JSON_PARSE_ERROR, err.Error(), ctx.Request().RequestURI, e.logger)
	}

	if err = e.service.Song.CreateSong(input); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_CREATE_ELEMENT, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	return helpers.ResponseHelper(ctx, http.StatusOK, helpers.SUCCESS, "Song created", ctx.Request().RequestURI, e.logger)
}

func (e *Controller) UpdateSong(ctx echo.Context) error {
	songId := ctx.Param("id")

	var input model.SongUpdate
	if err := ctx.Bind(&input); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusBadRequest, helpers.JSON_PARSE_ERROR, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusUnprocessableEntity, helpers.JSON_PARSE_ERROR, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	if err = e.service.Song.UpdateSong(songId, input); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_UPDATE_ELEMENT, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	return helpers.ResponseHelper(ctx, http.StatusOK, helpers.SUCCESS, "Song updated", ctx.Request().RequestURI, e.logger)
}

func (e *Controller) DeleteSong(ctx echo.Context) error {
	songId := ctx.Param("id")
	if err := e.service.Song.DeleteSong(songId); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_DELETE_ELEMENT, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	return helpers.ResponseHelper(ctx, http.StatusOK, helpers.SUCCESS, "Song deleted", ctx.Request().RequestURI, e.logger)
}
