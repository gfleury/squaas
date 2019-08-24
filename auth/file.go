package auth

import (
	"github.com/gfleury/dbquerybench/config"
	"github.com/gin-gonic/gin"
)

func ConfigBasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(getConfigCredentials())
}

func getConfigCredentials() map[string]string {
	return config.GetConfig().GetStringMapString("auth.config.users")
}
