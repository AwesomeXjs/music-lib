package controller

import (
	"fmt"
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
	response := helpers.ResponseHelper(ctx, http.StatusOK, helpers.SUCCESS, "Song created id: "+songId, ctx.Request().RequestURI, e.logger)

	/*
			Когда мы получаем ID записи в базе - мы запускаем горутину которая на фоне ищет данные на стороннем сервисе
			Если данные найдены - она обновляет нужные нам поля в базе
			Это сделано для того чтобы клиент не ждал ответа от стороннего сервиса а только получал ответ от базы при первой своей записи
			Так как сторонний сервис может не работать, либо долго отвечать либо записи там не будет найдено,
		    в таком случае мы запишем в базу "дефолтные" поля (Not found)

			на случай долгого ответа у кастомного клиента установлен timeout в 10 секунд
	*/
	go func(songId string, input model.SongCreate) {
		err = e.service.Song.FetchSongData(songId, input)
		if err != nil {
			e.logger.Info("Failed to fetch and update song details:", err.Error())
		}
	}(songId, input)

	e.logger.Info(helpers.APP_PREFIX, "Song created id: "+songId)
	return response
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
// @Param link query string false "Filter by link"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} []model.Song
// @Failure 400 {object} helpers.Response
// @Failure 422 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /v1/songs [get]
func (e *Controller) GetSongs(ctx echo.Context) error {
	params := ctx.QueryParams()

	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	songs, err := e.service.Song.GetSongs(
		strings.Trim(params.Get("group"), " "),
		strings.Trim(params.Get("song"), " "),
		strings.Trim(params.Get("releaseDate"), " "),
		strings.Trim(params.Get("text"), " "),
		strings.Trim(params.Get("link"), " "),
		page, limit)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_GET_ELEMENTS, err.Error(), ctx.Request().RequestURI, e.logger)
	}

	return ctx.JSON(http.StatusOK, songs)
}

// @Summary Get verse
// @Tags Song
// @Description get verse of song
// @ID get-verse
// @Accept  json
// @Produce  json
// @Param id path string false "Song id"
// @Param num query string false "Number of verse (номер куплета)"
// @Success 200 {object} helpers.Verse
// @Failure 400 {object} helpers.Response
// @Failure 422 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /v1/songs/verse/{id} [get]
func (e *Controller) GetVerse(ctx echo.Context) error {
	songId := ctx.Param("id")
	numberOfVerse, err := strconv.Atoi(ctx.QueryParam("num"))
	text, err := e.service.Song.GetVerse(songId)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FAILED_TO_GET_ELEMENTS, err.Error(), ctx.Request().RequestURI, e.logger)
	}
	fmt.Println(text)
	verse := strings.Split(text, "\n\n")

	if numberOfVerse > len(verse) {
		return helpers.ResponseHelper(ctx, http.StatusUnprocessableEntity, helpers.FAILED_TO_GET_ELEMENTS, "This song doesn't have that many verses", ctx.Request().RequestURI, e.logger)
	}
	if numberOfVerse < 1 {
		fullVerse := text
		return ctx.JSON(http.StatusOK, helpers.Verse{Verse: fullVerse})
	}

	return ctx.JSON(http.StatusOK, helpers.Verse{Verse: verse[numberOfVerse-1]})
}
