package app

import (
	"github.com/gin-gonic/gin"

	"gateway/internal/config"
	"gateway/internal/utils"
	log "gateway/internal/utils"
)

func App() *gin.Engine {
	log.Info(connect_services) 

	config.NewPostgreSql()

	if err := utils.LoadJwtSecretKey(); err != nil {
		panic(err)
	}

	log.Info(services_connected) 

	log.Info(starting_http_server)

	gin.SetMode(gin.ReleaseMode)

	return gin.New()
}

func Run(serve *gin.Engine, addr string){
	log.Info(http_server_started, addr)
	if err := serve.Run(addr); err != nil {
		log.Fatalf(err.Error())
	}
}

const(
	connect_services 		= "...CONNECTING TO DEPEND EXTERNAL I/O SERVICES"
	services_connected 		= "✅ ALL EXTERNAL DEPEND SERVICES CONNECTED SUCCESSFULLY"
	starting_http_server 	= "...Starting HTTP Server Using GIN"
	http_server_started 	= "✅ HTTP Web Server Started On Address: "
)
