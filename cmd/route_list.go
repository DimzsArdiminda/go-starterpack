package cmd

import (
	"fmt"
	routes "golang-backend/Routes"
	"sort"

	"github.com/gin-gonic/gin"
)

func RouteList() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	routes.Register(router, nil)
	routes := router.Routes()
	sort.Slice(routes, func(i, j int) bool {
		if routes[i].Path == routes[j].Path {
			return routes[i].Method < routes[j].Method
		}
		return routes[i].Path < routes[j].Path
	})
	for _, route := range routes {
		fmt.Printf("%-7s %s\n", route.Method, route.Path)
	}
}
