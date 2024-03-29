package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	"github.com/core-go/search"
	"github.com/gorilla/mux"

	"go-service/internal/user"
	"go-service/internal/user/entity"
)

const InternalServerError = "Internal Server Error"

type UserHandler struct {
	service user.UserService
	*search.SearchHandler
	validate func(context.Context, interface{}) ([]core.ErrorMessage, error)
}

func NewUserHandler(
	find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
	service user.UserService,
	validate func(context.Context, interface{}) ([]core.ErrorMessage, error),
	logError func(context.Context, string, ...map[string]interface{}),
) *UserHandler {
	filterType := reflect.TypeOf(entity.UserFilter{})
	modelType := reflect.TypeOf(entity.User{})
	searchHandler := search.NewSearchHandler(find, modelType, filterType, logError, nil)
	return &UserHandler{service: service, SearchHandler: searchHandler, validate: validate}
}

func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := h.service.Load(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	errors, er2 := h.validate(r.Context(), &user)
	if er2 != nil {
		h.LogError(r.Context(), er2.Error())
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.Create(r.Context(), &user)
	if er3 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, res)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	errors, er2 := h.validate(r.Context(), &user)
	if er2 != nil {
		h.LogError(r.Context(), er2.Error())
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.Update(r.Context(), &user)
	if er3 != nil {
		http.Error(w, er3.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *UserHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	var user entity.User
	userType := reflect.TypeOf(user)
	_, jsonMap, _ := core.BuildMapField(userType)
	body, er1 := core.BuildMapAndStruct(r, &user)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	json, er2 := core.BodyToJsonMap(r, user, body, []string{"id"}, jsonMap)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "method", "patch"))
	errors, er3 := h.validate(r.Context(), &user)
	if er3 != nil {
		h.LogError(r.Context(), er3.Error())
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er4 := h.service.Patch(r.Context(), json)
	if er4 != nil {
		http.Error(w, er4.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
