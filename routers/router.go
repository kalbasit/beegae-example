package routers

import (
	"github.com/eMxyzptlk/beegae-example/controllers"
	"github.com/astaxie/beegae"
)

func init() {
    beegae.Router("/", &controllers.MainController{})
}
