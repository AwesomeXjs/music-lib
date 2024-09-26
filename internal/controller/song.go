package controller

import (
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
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
// @Router /v1/songs [post]
func (e *Controller) CreateSong(ctx echo.Context) error {
	var input model.SongCreate

	if err := ctx.Bind(&input); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusBadRequest, helpers.JSON_PARSE_ERROR, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusUnprocessableEntity, helpers.JSON_PARSE_ERROR, err.Error(), ctx.Request().RequestURI, e.logger)
	}

	songId, err := e.service.Song.CreateSong(input)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_CREATE_ELEMENT, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	return helpers.ResponseHelper(ctx, http.StatusOK, helpers.SUCCESS, "Song created id: "+songId, ctx.Request().RequestURI, e.logger)
}

// @Summary Update song
// @Tags Song
// @Description Update song
// @ID update-song
// @Accept  json
// @Produce  json
// @Param id path string true "update by id"
// @Param input body model.SongUpdate true "song info"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 422 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /v1/songs/{id} [put]
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

// @Summary Delete song
// @Tags Song
// @Description delete song from library
// @ID delete-song
// @Accept  json
// @Produce  json
// @Param id path string true "delete by id"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 422 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /v1/songs/{id} [delete]
func (e *Controller) DeleteSong(ctx echo.Context) error {
	songId := ctx.Param("id")
	if err := e.service.Song.DeleteSong(songId); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_DELETE_ELEMENT, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	return helpers.ResponseHelper(ctx, http.StatusOK, helpers.SUCCESS, "Song deleted", ctx.Request().RequestURI, e.logger)
}

// @Summary Get songs
// @Tags Song
// @Description get songs from library
// @ID get-song
// @Accept  json
// @Produce  json
// @Param group query string false "Filter by group"
// @Param song query string false "Filter by song"
// @Param releaseDate query string false "Filter by created_at"
// @Param text query string false "Filter by text"
// @Param patronymic query string false "Filter by patronymic"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} []model.Song
// @Failure 400 {object} helpers.Response
// @Failure 422 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /v1/songs [get]
func (e *Controller) GetSongs(ctx echo.Context) error {
	group := strings.Trim(ctx.QueryParam("group"), " ")
	song := strings.Trim(ctx.QueryParam("song"), " ")
	createdAt := strings.Trim(ctx.QueryParam("releaseDate"), " ")
	text := strings.Trim(ctx.QueryParam("text"), " ")
	patronymic := strings.Trim(ctx.QueryParam("patronymic"), " ")

	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	songs, err := e.service.Song.GetSongs(group, song, createdAt, text, patronymic, page, limit)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_GET_ELEMENTS, err.Error(), ctx.Request().RequestURI, e.logger)
	}

	return ctx.JSON(http.StatusOK, songs)
}
