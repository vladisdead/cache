package server

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"strconv"
	"task4/model"
)

func (s *Server) GetAllCacheHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	body, err := json.Marshal(s.users.GetAllCache())
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации пользователей")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}

func (s *Server) GetAllUsersHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	users, err := s.users.GetAllUsers()
	if err != nil {
		log.Err(err).Msg("Невозможно показать всех пользователей")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать всех пользователей")

		return
	}

	body, err := json.Marshal(users)
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации пользователей")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}

func (s *Server) GetUsersWithMinAgeHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	byteAge := ctx.QueryArgs().Peek("minAge")
	age, err := strconv.Atoi(string(byteAge))
	if err != nil {
		log.Err(err).Msg("Невозможно показать пользователя")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать пользователя")

		return
	}

	users, err := s.users.GetUsersWithMinAge(age)
	if err != nil {
		log.Err(err).Msg("Невозможно показать пользователей")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать пользователей")

		return
	}

	body, err := json.Marshal(users)
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации пользователей")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}

func (s *Server) GetUserByIdHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Msg("Невозможно показать пользователя")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать пользователя")

		return
	}

	user, err := s.users.GetUserById(id)
	if err != nil {
		log.Err(err).Msg("Невозможно показать пользователя")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать пользователя")

		return
	}
	body, err := json.Marshal(user)
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации пользователя")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)

}

func (s *Server) CreateUserHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var (
		err  error
		user model.Users
	)

	if err = json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации о пользователе")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверный JSON")
	}

	if err = s.users.CreateUser(user); err != nil {
		log.Err(err).Msg("Ошибка при добавлении пользователя")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Ошибка при добавлении пользователя")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (s *Server) UpdateUserHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var user model.Users

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Msg("Невозможно показать пользователя")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать пользователя")

		return
	}

	if err = json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации о пользователе")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверный JSON")
	}

	if err = s.users.UpdateUser(id, user); err != nil {
		log.Err(err).Msg("Невозможно обновить информацию о пользователе")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно обновить информацию о пользователе")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (s *Server) DeleteUserHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Msg("Невозможно считать id пользователя")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	err = s.users.DeleteUser(id)
	if err != nil {
		log.Err(err).Msg("Ошибка при удалении пользователя")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Ошибка")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
