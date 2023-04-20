package http

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/auth")
	auth.POST("/sign-up", s.handler.User.SignUp)
	auth.POST("/sign-in", s.handler.User.Login)

	v1 := s.App.Group("/api", s.m.ValidateAuth)
	v1.GET("/my-orders", s.handler.User.GetOrders)
	v1.POST("/books/return/:orderID", s.handler.User.ReturnBook)
	v1.POST("/books/rent-book", s.handler.User.RentBook)
	v1.POST("/user/balance/:username/:balance", s.handler.User.ReplenishBalance)

	v2 := s.App.Group("/api")
	v2.GET("/orders", s.handler.Order.GetAllOrders)
	v2.GET("/orders/last-month-orders", s.handler.Order.GetLastMonthOrders)
	v1.GET("/all-users", s.handler.User.GetAllUsers)
	v2.GET("/books", s.handler.Book.GetAllBooks)
	v2.GET("/books/available", s.handler.Book.GetAvailable)
	v2.GET("/books/:id", s.handler.Book.GetByID)
	v2.POST("/books/create", s.handler.Book.Create)
}
