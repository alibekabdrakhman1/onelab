package http

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/auth")
	auth.POST("/sign-up", s.handler.User.SignUp)
	auth.POST("/sign-in", s.handler.User.Login)

	v1 := s.App.Group("/api", s.m.ValidateAuth)
	v1.GET("/my-orders", s.handler.User.GetOrders)
	v1.POST("/orders/create", s.handler.Order.Create)
	v1.POST("/return-book/:orderID", s.handler.Order.ReturnBook)

	v2 := s.App.Group("/api")
	v2.GET("/orders", s.handler.Order.GetAllOrders)
	v2.GET("/orders/last-month-orders", s.handler.Order.GetLastMonthOrders)
	v1.GET("/all-users", s.handler.User.GetAllUsers)
	v2.GET("/users/:username", s.handler.User.GetByUsername)
	v2.GET("/books", s.handler.Book.GetAllBooks)
	v2.GET("/books/available", s.handler.Book.GetAvailable)
	v2.GET("/books/:author", s.handler.Book.GetByAuthor)
	v2.POST("/books/create", s.handler.Book.Create)
}
