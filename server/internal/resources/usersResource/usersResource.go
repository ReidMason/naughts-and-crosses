package usersResource

import (
	"context"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ReidMason/naughts-and-crosses/server/internal/database"
	httpHelper "github.com/ReidMason/naughts-and-crosses/server/internal/helpers"
	"github.com/go-chi/chi"
)

type usersResource struct {
	userService UserService
}

type UserService interface {
	CreateUser(name string) (database.User, error)
	GetUser(id int32) (database.User, error)
	GetUserByToken(token string) (database.User, error)
}

func New(userService UserService) *usersResource {
	return &usersResource{
		userService: userService,
	}
}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.create)
	r.Get("/", rs.getCurrentUser)
	r.Route("/{userId}", func(r chi.Router) {
		r.Use(rs.userCtx)
		r.Get("/", rs.get)
	})

	return r
}

func (ur usersResource) userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 32)
		if err != nil {
			slog.Info("Failed to convert userId", err)
			httpHelper.SendResponse[interface{}](w, nil, false, "User not found", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", int32(userId))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs usersResource) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	accessToken := ""
	for _, cookie := range r.Cookies() {
		if cookie.Name == "accessToken" {
			accessToken = cookie.Value
		}
	}

	if accessToken == "" {
		httpHelper.SendResponse[interface{}](w, nil, false, "User not yet authenticated", http.StatusUnauthorized)
		return
	}

	_, err := rs.userService.GetUserByToken(accessToken)
	if err != nil {
		httpHelper.SendResponse[interface{}](w, nil, false, "User not yet authenticated", http.StatusUnauthorized)
		return
	}

	w.Header().Set("HX-Redirect", "/login-successful")
}

func (rs usersResource) create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := strings.TrimSpace(r.Form.Get("username"))

	if username == "" {
		slog.Error("Name field missing from request body")
		httpHelper.SendResponse[interface{}](w, nil, false, "The 'name' field is required", http.StatusBadRequest)
		return
	}

	newUser, err := rs.userService.CreateUser(username)
	if err != nil {
		slog.Error("Error creating user", err)
		httpHelper.SendResponse[interface{}](w, nil, false, "Error creating user", http.StatusInternalServerError)
		return
	}

	expiration := time.Now().Add(10 * 365 * 24 * time.Hour)
	log.Println(expiration)
	cookie := http.Cookie{Name: "accessToken", Path: "/", Value: newUser.Token, Expires: expiration}
	http.SetCookie(w, &cookie)

	templ := template.Must(template.ParseFiles("internal/templates/partials/registrationSuccess.html"))
	templ.Execute(w, nil)
}

type UserDTO struct {
	DateCreated time.Time
	Name        string
	Wins        int64
	Losses      int64
	ID          int32
}

func (ur usersResource) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user").(int32)

	user, err := ur.userService.GetUser(userId)
	if err != nil {
		httpHelper.SendResponse[interface{}](w, nil, false, "User not found", http.StatusNotFound)
		return
	}

	userResponse := UserDTO{
		Name:   user.Name,
		ID:     user.ID,
		Wins:   user.Wins,
		Losses: user.Losses,
	}
	httpHelper.SendResponse(w, &userResponse, true, "User found", http.StatusOK)
}
