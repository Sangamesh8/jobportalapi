package handlers

import (
	"encoding/json"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) addJobsById(c *gin.Context) {
	ctx := c.Request.Context()
	TraceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("TraceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	compId := c.Param("ID")
	var jobs []models.Job
	err := json.NewDecoder(c.Request.Body).Decode(&jobs)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("Trace Id", TraceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	jobData, err := h.s.AddJobs(jobs, compId)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", TraceId).Msg("Add Job by companyId problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Job creation failed"})
		return
	}
	c.JSON(http.StatusCreated, jobData)
}

func (h *handler) jobsByCompanyById(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	companyId := c.Param("companyId")
	listOfJobs, err := h.s.FetchJobByCompanyId(ctx, companyId)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing list of company by ID"})
		return
	}
	c.JSON(http.StatusOK, listOfJobs)
}

func (h *handler) fetchJobById(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	jobId := c.Param("ID")
	job, err := h.s.GetJobById(ctx, jobId)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing list of company by ID"})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (h *handler) getAllJobs(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	job, err := h.s.GetAllJobs(ctx)
	if err != nil {
		log.Error().Err(err).Str("Tracker Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing list of company by ID"})
		return
	}
	c.JSON(http.StatusOK, job)
}
