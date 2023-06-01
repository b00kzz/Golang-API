package httpserv

import (
	"ticket/goapi/infrastructure"
	"ticket/goapi/internal/adapter/handler"
	"ticket/goapi/internal/adapter/repo"
	"ticket/goapi/internal/core/service"
	"ticket/goapi/middleware"

	"github.com/gin-gonic/gin"
)

func bindBill(g gin.Engine) {
	repo := repo.NewBillRepo(infrastructure.DB)
	svc := service.NewBillSvc(repo)
	hdl := handler.NewBillHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/bills", hdl.GetBills)
		v1.GET("/bill/:BillId", hdl.GetBill)
		v1.POST("/bill", hdl.AddBill)
		v1.PUT("/bill/:BillId", hdl.UpdateBill)
		v1.DELETE("/bill/:BillId", hdl.DeleteBill)
	}

}

func bindCustomer(g gin.Engine) {
	repo := repo.NewCustomerRepo(infrastructure.DB)
	svc := service.NewCustomerSvc(repo)
	hdl := handler.NewCustomerHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/customers", hdl.GetCustomers)
		v1.GET("/customers/:customerID", hdl.GetCustomer)
		v1.POST("/customers", hdl.AddCustomer)
		v1.PUT("/customers/:customerID", hdl.UpdateCustomer)
		v1.DELETE("/customers/:customerID", hdl.DeleteCustomer)
	}

}

func bindPayment(g gin.Engine) {
	repo := repo.NewPaymentRepo(infrastructure.DB)
	svc := service.NewPaymentSvc(repo)
	hdl := handler.NewPaymentHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/payments", hdl.GetPayments)
		v1.GET("/payment/:PaymentId", hdl.GetPayment)
		v1.POST("/payment", hdl.AddPayment)
		v1.PUT("/payment/:PaymentId", hdl.UpdatePayment)
		v1.DELETE("/payment/:PaymentId", hdl.DeletePayment)
	}

}

func bindUser(g gin.Engine) {
	repo := repo.NewRegisterRepo(infrastructure.DB)
	svc := service.NewRegisterSvc(repo)
	hdl := handler.NewRegisterHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/users", hdl.GetUsers)
		v1.GET("/user/:ID", hdl.GetUser)
		v1.POST("/register", hdl.AddUser)
		v1.PUT("/user/:ID", hdl.UpdateUser)
		v1.DELETE("/user/:ID", hdl.DeleteUser)
		v1.POST("/login", hdl.Login)
		v1.PUT("user/role/:ID", hdl.UpdateRole)
		v1.PUT("user/status/:ID", hdl.UpdateStatus)
		v1.GET("/user", middleware.DeserializeUser(repo), hdl.GetUsers)
		v1.GET("/profile", middleware.DeserializeUser(repo), hdl.GetProfile)
	}
}

func bindReview(g gin.Engine) {
	repo := repo.NewReviewRepo(infrastructure.DB)
	svc := service.NewReviewSvc(repo)
	hdl := handler.NewReviewHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/reviews", hdl.GetReviews)
		v1.GET("/review/:ReviewId", hdl.GetReview)
		v1.POST("/review", hdl.AddReview)
		v1.PUT("/review/:ReviewId", hdl.UpdateReview)
		v1.DELETE("/review/:ReviewId", hdl.DeleteReview)
	}

}

func bindRole(g gin.Engine) {
	repo := repo.NewRoleRepo(infrastructure.DB)
	svc := service.NewRoleSvc(repo)
	hdl := handler.NewRoleHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/roles", hdl.GetRoles)
		v1.GET("/role/:RoleId", hdl.GetRole)
		v1.POST("/role", hdl.AddRole)
		v1.PUT("/role/:RoleId", hdl.UpdateRole)
		v1.DELETE("/role/:RoleId", hdl.DeleteRole)
	}

}

func bindTicket(g gin.Engine) {
	repo := repo.NewTicketRepo(infrastructure.DB)
	svc := service.NewTicketSvc(repo)
	hdl := handler.NewTicketHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/tickets", hdl.GetTickets)
		v1.GET("/ticket/:TicketId", hdl.GetTicket)
		v1.POST("/ticket", hdl.AddTicket)
		v1.PUT("/ticket/:TicketId", hdl.UpdateTicket)
		v1.PUT("/ticket/status/:TicketId", hdl.UpdateStatusTicket)
		v1.DELETE("/ticket/:TicketId", hdl.DeleteTicket)
		v1.GET("/search/:TicketName", hdl.Search)
	}

}

func bindUserDetail(g gin.Engine) {
	repo := repo.NewUserDetailRepo(infrastructure.DB)
	svc := service.NewUserDetailSvc(repo)
	hdl := handler.NewUserDetailHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/userdetails", hdl.GetUserDetails)
		v1.GET("/userdetail/:UserDetailId", hdl.GetUserDetail)
		v1.POST("/userdetail", hdl.AddUserDetail)
		v1.PUT("/userdetail/:UserDetailId", hdl.UpdateUserDetail)
		v1.DELETE("/userdetail/:UserDetailId", hdl.DeleteUserDetail)
	}

}

func bindImage(g gin.Engine) {
	v1 := g.Group("/v1")
	{
		v1.POST("/image", handler.FileUpload())
		v1.POST("/remote", handler.RemoteUpload())
	}
}
