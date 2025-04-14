package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
)

func (s *Server) getSorprotectionRoutes() []Route {
	return []Route{
		{
			Name:    "Index",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.String(http.StatusOK, "Hello nycu-ucr!")
			},
		},
		{
			Name:    "SupiUeSorPost",
			Method:  http.MethodPost,
			Pattern: "/:supi/ue-sor/generate-sor-data",
			APIFunc: s.HTTPSupiUeSorPost,
		},
	}
}

func (s *Server) HTTPSupiUeSorPost(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
