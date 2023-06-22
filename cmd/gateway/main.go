package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/msik-404/micro-appoint-gateway/internal/auth"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/employees"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/services"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/users"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"content-type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/companies", companies.GetCompanies)
	r.GET("/companies/:company_id", companies.GetCompany)
	r.GET("/companies/:company_id/services", services.GetServices)
	r.GET("/companies/:company_id/employees", employees.GetEmployees)
	r.GET("/employees/:employee_id", employees.GetEmployee)
	r.GET("/owners/companies", middleware.BindAuth[auth.Token], middleware.RequireOwnerAuth, companies.GetCompaniesByIds)

	r.POST("/login/customers", middleware.Bind[middleware.User], users.LoginCustomer)
	r.POST("/login/owners", middleware.Bind[middleware.User], users.LoginOwner)
	r.POST("/companies", middleware.BindWithAuth[middleware.Company], middleware.RequireOwnerAuth, companies.AddCompany)
	r.POST("/companies/:company_id/services", middleware.BindWithAuth[middleware.Service], middleware.RequireOwnerAuth, services.AddService)
	r.POST("/companies/:company_id/employees", middleware.BindWithAuth[middleware.Employee], middleware.RequireOwnerAuth, employees.AddEmployee)
	r.POST("/customers", middleware.Bind[middleware.User], users.AddCustomer)
	r.POST("/owners", middleware.Bind[middleware.User], users.AddOwner)

	r.PUT("/companies/:company_id", middleware.BindWithAuth[middleware.Company], middleware.RequireOwnerAuth, companies.UpdateCompany)
	r.PUT("/companies/:company_id/services/:service_id", middleware.BindWithAuth[middleware.Service], middleware.RequireOwnerAuth, services.UpdateService)
	r.PUT("/companies/:company_id/employees/:employee_id", middleware.BindWithAuth[middleware.Employee], middleware.RequireOwnerAuth, employees.UpdateEmployee)
	r.PUT("/customers", middleware.BindWithAuth[middleware.UserUpdate], middleware.RequireCustomerAuth, users.UpdateCustomer)
	r.PUT("/owners", middleware.BindWithAuth[middleware.UserUpdate], middleware.RequireOwnerAuth, users.UpdateOwner)

	r.DELETE("/companies/:company_id", middleware.BindAuth[auth.Token], middleware.RequireOwnerAuth, companies.DeleteCompany)
	r.DELETE("/companies/:company_id/services/:service_id", middleware.BindAuth[auth.Token], middleware.RequireOwnerAuth, services.DeleteService)
	r.DELETE("/companies/:company_id/employees/:employee_id", middleware.BindAuth[auth.Token], middleware.RequireOwnerAuth, employees.DeleteEmployee)
	r.DELETE("/customers", middleware.BindAuth[auth.Token], middleware.RequireCustomerAuth, users.DeleteCustomer)
	r.DELETE("/owners", middleware.BindAuth[auth.Token], middleware.RequireOwnerAuth, users.DeleteOwner)

	r.Run() // listen and serve on 0.0.0.0:8080
}
