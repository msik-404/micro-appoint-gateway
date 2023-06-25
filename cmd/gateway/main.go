package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/msik-404/micro-appoint-gateway/internal/grpc"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/employees"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/services"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/users"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/middleware"
)

func main() {
    conns, err := grpc.New()
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

	r.GET("/companies", companies.GetCompanies(conns))
	r.GET("/companies/:company_id", companies.GetCompany(conns))
	r.GET("/companies/:company_id/services", services.GetServices(conns))
	r.GET("/companies/:company_id/employees", employees.GetEmployees(conns))
	r.GET("/companies/:company_id/employees/:employee_id", middleware.RequireOwnerAuth(conns), employees.GetEmployee(conns))
	r.GET("/owners/companies", middleware.RequireOwnerAuth(conns), companies.GetCompaniesByIds(conns))

	r.POST("/login/customers", middleware.Bind[middleware.User], users.LoginCustomer(conns))
	r.POST("/login/owners", middleware.Bind[middleware.User], users.LoginOwner(conns))
	r.POST("/companies", middleware.RequireOwnerAuth(conns), middleware.Bind[middleware.Company], companies.AddCompany(conns))
	r.POST("/companies/:company_id/services", middleware.RequireOwnerAuth(conns), middleware.Bind[middleware.Service], services.AddService(conns))
	r.POST("/companies/:company_id/employees", middleware.RequireOwnerAuth(conns), middleware.Bind[middleware.Employee], employees.AddEmployee(conns))
	r.POST("/customers", middleware.Bind[middleware.User], users.AddCustomer(conns))
	r.POST("/owners", middleware.Bind[middleware.User], users.AddOwner(conns))

	r.PUT("/companies/:company_id", middleware.RequireOwnerAuth(conns), middleware.Bind[middleware.Company], companies.UpdateCompany(conns))
	r.PUT("/companies/:company_id/services/:service_id", middleware.RequireOwnerAuth(conns), middleware.Bind[middleware.Service], services.UpdateService(conns))
	r.PUT("/companies/:company_id/employees/:employee_id", middleware.RequireOwnerAuth(conns), middleware.Bind[middleware.Employee], employees.UpdateEmployee(conns))
	r.PUT("/customers", middleware.RequireCustomerAuth(conns), middleware.Bind[middleware.UserUpdate], users.UpdateCustomer(conns))
	r.PUT("/owners", middleware.RequireOwnerAuth(conns), middleware.Bind[middleware.UserUpdate], users.UpdateOwner(conns))

	r.DELETE("/companies/:company_id", middleware.RequireOwnerAuth(conns), companies.DeleteCompany(conns))
	r.DELETE("/companies/:company_id/services/:service_id", middleware.RequireOwnerAuth(conns), services.DeleteService(conns))
	r.DELETE("/companies/:company_id/employees/:employee_id", middleware.RequireOwnerAuth(conns), employees.DeleteEmployee(conns))
	r.DELETE("/customers", middleware.RequireCustomerAuth(conns), users.DeleteCustomer(conns))
	r.DELETE("/owners", middleware.RequireOwnerAuth(conns), users.DeleteOwner(conns))

	r.Run() // listen and serve on 0.0.0.0:8080
}
