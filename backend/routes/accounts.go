package routes

import "backend/api"

func CreateRoutesAccount(server *api.Server) {
	//Create Account
	server.Router.POST("/accounts", server.CreateAccount)
	//Get information of a specific account
	server.Router.GET("/accounts/:id", server.GetAccount)
	//Get the list ofaccounts
	server.Router.GET("/accounts", server.ListAccount)
	//Delete Account
	server.Router.DELETE("/accounts/:id", server.DeleteAccount)
	//Update an account
	server.Router.PATCH("/accounts/:id", server.UpdateAccount)
}
