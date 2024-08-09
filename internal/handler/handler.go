package handler

import (
	"github.com/Oriseer/testruns/internal/middleware"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)
		r.Get("/money", GetMoneyBalance)
		r.Get("/getAccount", GetAccount)
	})

	// New route for updating account
	r.Route("/newAccount", func(r chi.Router) {
		r.Use(middleware.AdminAuth)
		r.Post("/add_account", AddAccount)
		r.Post("/add_money", AddUserMoney)
		r.Post("/update_account", UpdateAccount)
	})

	// New Route for delete account
	r.Route("/deleteAccount", func(r chi.Router) {
		r.Post("/", DeleteAccount)
	})
}
