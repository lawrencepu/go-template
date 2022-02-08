package backend

import "go-template/app/http/controllers"

var Ctrl = &controller{}

type controller struct {
	controllers.BaseController
}

