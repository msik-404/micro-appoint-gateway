package main

import (
	"github.com/gin-gonic/gin"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/companies"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/services"
	"github.com/msik-404/micro-appoint-gateway/internal/rest/controllers/employees"
)

func main() {
	r := gin.Default()

	r.GET("/companies", companies.GetCompanies())
	r.GET("/companies/:id", companies.GetCompany())
    r.GET("/employees/:id", employees.GetEmployee())
	r.GET("/companies/services/:id", services.GetServices())
    r.GET("/companies/employees/:id", employees.GetEmployees())

	r.POST("/companies", companies.AddCompany())
	r.POST("/companies/services/:id", services.AddService())
	r.POST("/companies/employees/:id", employees.AddEmployee())

	r.PUT("/companies/:id", companies.UpdateCompany())
	r.PUT("/services/:id", services.UpdateService())
	r.PUT("/employees/:id", employees.UpdateEmployee())

	r.DELETE("/companies/:id", companies.DeleteCompany())
	r.DELETE("/services/:id", services.DeleteService())
	r.DELETE("/employees/:id", employees.DeleteEmployee())

	r.Run() // listen and serve on 0.0.0.0:8080
}
