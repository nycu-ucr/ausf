package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
)

func (s *Server) getUpuprotectionRoutes() []Route {
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
			Name:    "SupiUeUpuPost",
			Method:  http.MethodPost,
			Pattern: "/:supi/ue-upu/generate-upu-data",
			APIFunc: s.HTTPSupiUeUpuPost,
		},
	}
}

func (s *Server) HTTPSupiUeUpuPost(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
