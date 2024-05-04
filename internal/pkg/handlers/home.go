package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/nickduskey/flowbite-go/internal/pkg/pages/Home"
	"github.com/nickduskey/flowbite-go/internal/pkg/utils"
)

func homeHandler(c echo.Context) error {
	cmp := home.Home()
	return utils.RenderView(c, cmp)
}
