package handlers

import (
	"fmt"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/services"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth, c *services.Conn) *gin.Engine {
	r := gin.New()
	m, err := middleware.NewMid(a)
	store := services.NewStore(c)
	if err != nil {
		log.Panic().Msg("middleware not set up")

	}
	h := handler{
		A: a,
		s: store,
	}

	r.Use(m.Log(), gin.Recovery())

	r.GET("/check", check)
	r.POST("/userSignup", h.userSignup)
	r.POST("/userLogin", h.userLogin)
	r.POST("/createCompany", h.CreateCompany)
	r.GET("/viewCompany", h.viewCompanies)
	r.GET("/getCompanyByID/:ID", h.companyById)
	r.POST("/addJobs/:ID", h.addJobsById)
	r.GET("/fetchJob/:ID", h.fetchJobById)
	r.GET("/jobByCompany/:companyId", h.jobsByCompanyById)
	r.GET("/getAllJobs", h.getAllJobs)
	return r

}

func check(c *gin.Context) {
	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

	}

}
