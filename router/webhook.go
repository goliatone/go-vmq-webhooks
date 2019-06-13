package routes

import (
	"github.com/goliatone/go-vmq-webhooks/controller"

	"github.com/gin-gonic/gin"
)

//WebHooks registers the following endpoints to handle vmq webhooks
//- POST /auth-on-register
//- POST /auth-on-publish
//- POST /auth-on-subscribe
//- POST /on-publish
func WebHooks(router *gin.RouterGroup) {
	router.POST("/auth-on-register", controller.AuthOnRegister)
	router.POST("/auth-on-publish", controller.AuthOnPublish)
	router.POST("/auth-on-subscribe", controller.AuthOnSubscribe)

	router.POST("/on-publish", controller.OnPublish)

	router.POST("/on-unsubscribe", controller.OnUnsubscribe)
	router.POST("/on-client-gone", controller.OnClientGone)
	router.POST("/on-client-offline", controller.OnClientOffline)
}
