package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

// CreateSong - create new song
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
		return helpers.ResponseHelper(ctx, http.StatusBadRequest, helpers.JSONParseError, err.Error(), e.logger)
	}

	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusUnprocessableEntity, helpers.JSONParseError, err.Error(), e.logger)
	}

	songID, err := e.service.Song.CreateSong(input)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FailedToCreateElement, err.Error(), e.logger)
	}
	err = helpers.ResponseHelper(ctx, http.StatusOK, helpers.Success, "Song created id: "+songID, e.logger)

	/*
			Когда мы получаем ID записи в базе - мы запускаем горутину которая на фоне ищет данные на стороннем сервисе
			Если данные найдены - она обновляет нужные нам поля в базе
			Это сделано для того чтобы клиент не ждал ответа от стороннего сервиса а только получал ответ от базы при первой своей записи
			Так как сторонний сервис может не работать, либо долго отвечать либо там не будет найдено ничего,
		    в таком случае мы запишем в базу "дефолтные" поля (Not found)

			на случай долгого ответа у кастомного клиента установлен timeout в 10 секунд
	*/

	go func(songId string, input model.SongCreate) {
		err = e.service.Song.FetchSongData(songId, input)
		if err != nil {
			e.logger.Info("Failed to fetch and update song details:", err.Error())
		}
	}(songID, input)

	e.logger.Info(helpers.AppPrefix, "Song created id: "+songID)
	return err
}

// UpdateSong - update song
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
	songID := ctx.Param("id")
	var input model.SongUpdate
	if err := ctx.Bind(&input); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusBadRequest, helpers.JSONParseError, err.Error(), e.logger)
	}

	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusUnprocessableEntity, helpers.JSONParseError, err.Error(), e.logger)
	}

	if err = e.service.Song.UpdateSong(songID, input); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FailedToCreateElement, err.Error(), e.logger)
	}

	return helpers.ResponseHelper(ctx, http.StatusOK, helpers.Success, "Song updated", e.logger)
}

// DeleteSong - delete song
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
	songID := ctx.Param("id")

	if err := e.service.Song.DeleteSong(songID); err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FailedToDeleteElement, err.Error(), e.logger)
	}
	return helpers.ResponseHelper(ctx, http.StatusOK, helpers.Success, "Song deleted", e.logger)
}

// GetSongs - get songs
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
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FailedToGetElements, err.Error(), e.logger)
	}

	return ctx.JSON(http.StatusOK, songs)
}

// GetVerse - get verse
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
	songID := ctx.Param("id")
	numberOfVerse, err := strconv.Atoi(ctx.QueryParam("num"))

	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusBadRequest, helpers.FailedToGetElements, err.Error(), e.logger)
	}

	text, err := e.service.Song.GetVerse(songID)
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FailedToGetElements, err.Error(), e.logger)
	}

	verse := strings.Split(text, "\n\n")
	if numberOfVerse > len(verse) {
		return helpers.ResponseHelper(ctx, http.StatusUnprocessableEntity, helpers.FailedToGetElements, "This song doesn't have that many verses", e.logger)
	}
	if numberOfVerse < 1 {
		fullVerse := text
		return ctx.JSON(http.StatusOK, helpers.Verse{Verse: fullVerse})
	}

	return ctx.JSON(http.StatusOK, helpers.Verse{Verse: verse[numberOfVerse-1]})
}

// GetAllFromMockService - get all data from mock service
// @Summary Get All from mock service
// @Tags MockServer
// @Description Посмотреть все доступные песни с данными
// @ID get-all
// @Produce  json
// @Success 200 {object} []model.Song
// @Failure 400 {object} helpers.Response
// @Failure 422 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /v1/all [get]
func (e *Controller) GetAllFromMockService(ctx echo.Context) error {
	songs, err := e.service.Song.GetAllFromMockService()
	if err != nil {
		return helpers.ResponseHelper(ctx, http.StatusInternalServerError, helpers.FailedToGetElements, err.Error(), e.logger)
	}
	return ctx.JSON(http.StatusOK, songs)
}
