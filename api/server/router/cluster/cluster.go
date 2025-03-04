/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"github.com/gin-gonic/gin"

	"github.com/caoyingjunz/pixiu/cmd/app/options"
	"github.com/caoyingjunz/pixiu/pkg/controller"
)

// clusterRouter is a router to talk with the cluster controller
type clusterRouter struct {
	c controller.PixiuInterface
}

// NewRouter initializes a new cluster router
func NewRouter(o *options.Options) {
	s := &clusterRouter{
		c: o.Controller,
	}
	s.initRoutes(o.HttpEngine)
}

func (cr *clusterRouter) initRoutes(httpEngine *gin.Engine) {
	clusterRoute := httpEngine.Group("/pixiu/clusters")

	{
		clusterRoute.POST("", cr.createCluster)
		clusterRoute.PUT("/:clusterId", cr.updateCluster)
		clusterRoute.DELETE("/:clusterId", cr.deleteCluster)
		clusterRoute.GET("/:clusterId", cr.getCluster)
		clusterRoute.GET("", cr.listClusters)

		// 检查 kubernetes 的连通性
		clusterRoute.POST("/ping", cr.pingCluster)
	}
}
