package httpserv

import (
	"fmt"
	"ticket/goapi/app"
	"ticket/goapi/middleware"

	"github.com/spf13/viper"
)

func Run() {
	a := app.NewApp()
	m := middleware.New()
	a.UseMiddleware(m.Logger)
	a.UseMiddleware(m.ErrorHandler)
	a.UseMiddleware(m.CORS)
	bindBill(a.GinEngine())
	bindCustomer(a.GinEngine())
	bindPayment(a.GinEngine())
	bindUser(a.GinEngine())
	bindReview(a.GinEngine())
	bindRole(a.GinEngine())
	bindTicket(a.GinEngine())
	bindUserDetail(a.GinEngine())

	port := fmt.Sprintf(":%v", viper.GetInt("app.port"))
	a.Start(port)
}
