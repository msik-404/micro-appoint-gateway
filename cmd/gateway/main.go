package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/rabbit"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/employees"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/orders"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/services"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/users"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func main() {
	connsGRPC, err := grpc.New()
	if err != nil {
		panic(err)
	}
    connsRabbit, err := rabbit.New()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"content-type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/companies", companies.GetCompanies(connsGRPC))
	r.GET("/companies/:company_id", companies.GetCompany(connsGRPC))
	r.GET("/companies/:company_id/services", services.GetServices(connsGRPC))
	r.GET("/companies/:company_id/employees", employees.GetEmployees(connsGRPC))
	r.GET("/companies/:company_id/employees/:employee_id", middleware.RequireOwnerAuth(connsGRPC), employees.GetEmployee(connsGRPC))
	r.GET("/owners/companies", middleware.RequireOwnerAuth(connsGRPC), companies.GetCompaniesByIds(connsGRPC))
	r.GET("/companies/:company_id/orders", middleware.RequireOwnerAuth(connsGRPC), orders.GetCompanyOrders(connsGRPC))
	r.GET("/customers/orders", middleware.RequireCustomerAuth(connsGRPC), orders.GetCustomerOrders(connsGRPC))
	r.GET("/avaliable-dates/companies/:company_id/services/:service_id", middleware.RequireCustomerAuth(connsGRPC), orders.GetAvaliableDates(connsGRPC))

	r.POST("/login/customers", middleware.Bind[middleware.User], users.LoginCustomer(connsGRPC))
	r.POST("/login/owners", middleware.Bind[middleware.User], users.LoginOwner(connsGRPC))
	r.POST("/companies", middleware.RequireOwnerAuth(connsGRPC), middleware.Bind[middleware.Company], companies.AddCompany(connsGRPC))
	r.POST("/companies/:company_id/services", middleware.RequireOwnerAuth(connsGRPC), middleware.Bind[middleware.Service], services.AddService(connsGRPC))
	r.POST("/companies/:company_id/employees", middleware.RequireOwnerAuth(connsGRPC), middleware.Bind[middleware.Employee], employees.AddEmployee(connsGRPC))
	r.POST("/customers", middleware.Bind[middleware.User], users.AddCustomer(connsGRPC))
	r.POST("/owners", middleware.Bind[middleware.User], users.AddOwner(connsGRPC))
	r.POST("/orders/cancel", middleware.RequireCustomerAuth(connsGRPC), middleware.Bind[middleware.CancelRequest], orders.CancelOrder(connsRabbit))
	r.POST("/orders", middleware.RequireCustomerAuth(connsGRPC), middleware.Bind[middleware.Order], orders.AddOrder(connsGRPC))

	r.PUT("/companies/:company_id", middleware.RequireOwnerAuth(connsGRPC), middleware.Bind[middleware.Company], companies.UpdateCompany(connsGRPC))
	r.PUT("/companies/:company_id/services/:service_id", middleware.RequireOwnerAuth(connsGRPC), middleware.Bind[middleware.Service], services.UpdateService(connsGRPC))
	r.PUT("/companies/:company_id/employees/:employee_id", middleware.RequireOwnerAuth(connsGRPC), middleware.Bind[middleware.Employee], employees.UpdateEmployee(connsGRPC))
	r.PUT("/customers", middleware.RequireCustomerAuth(connsGRPC), middleware.Bind[middleware.UserUpdate], users.UpdateCustomer(connsGRPC))
	r.PUT("/owners", middleware.RequireOwnerAuth(connsGRPC), middleware.Bind[middleware.UserUpdate], users.UpdateOwner(connsGRPC))

	r.DELETE("/companies/:company_id", middleware.RequireOwnerAuth(connsGRPC), companies.DeleteCompany(connsGRPC))
	r.DELETE("/companies/:company_id/services/:service_id", middleware.RequireOwnerAuth(connsGRPC), services.DeleteService(connsGRPC))
	r.DELETE("/companies/:company_id/employees/:employee_id", middleware.RequireOwnerAuth(connsGRPC), employees.DeleteEmployee(connsGRPC))
	r.DELETE("/customers", middleware.RequireCustomerAuth(connsGRPC), users.DeleteCustomer(connsGRPC))
	r.DELETE("/owners", middleware.RequireOwnerAuth(connsGRPC), users.DeleteOwner(connsGRPC))

	r.Run() // listen and serve on 0.0.0.0:8080
}
