package routes

import (
	"echo_shop/app/context"
	"echo_shop/app/controllers/page"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/session"
	"net/http"
	"time"

	mymiddleware "echo_shop/routes/middleware"
	"echo_shop/routes/wrapper"

	"github.com/labstack/echo/v4"

	"echo_shop/app/controllers/auth/login"
	"echo_shop/app/controllers/auth/password"
	"echo_shop/app/controllers/auth/register"
	"echo_shop/app/controllers/auth/verification"
	"echo_shop/pkg/flash"
)

func registerWeb(e *echo.Echo) {
	ee := e.Group("")

	// session
	ee.Use(mymiddleware.Session())
	// csrf
	ee.Use(mymiddleware.Csrf())
	// old value
	ee.Use(flash.OldValueFlashMiddleware())

	ee.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// session 存储请求 url
			if c.Request().Method == http.MethodGet {
				session.Set(c, constants.PreviousURL, c.Request().URL.String())
			}

			return next(c)
		}
	})

	context.RegisterHandler(ee.GET, "/", page.Root, mymiddleware.AuthIfLoginCheckActived).Name = "root"
	context.RegisterHandler(ee.POST, "/", page.Root)

	// ------------------------------------- Auth -------------------------------------
	// +++++++++++++++ 用户身份验证相关的路由 +++++++++++++++
	// 展示登录页面
	context.RegisterHandler(ee.GET, "login", login.Show, mymiddleware.Guest).Name = "login.show"
	// 登录
	context.RegisterHandler(ee.POST, "login", login.Login, mymiddleware.Guest).Name = "login"
	// 登出
	context.RegisterHandler(ee.DELETE, "logout", login.Logout, mymiddleware.AuthAndNoCheckActived).Name = "logout"

	// +++++++++++++++ 用户注册相关路由 +++++++++++++++
	// 展示注册页面
	context.RegisterHandler(ee.GET, "register", register.Show, mymiddleware.Guest).Name = "register.show"
	// 注册
	context.RegisterHandler(ee.POST, "register", register.Register, mymiddleware.Guest).Name = "register"

	// +++++++++++++++ 密码重置相关路由 +++++++++++++++
	pwdRouter := ee.Group("/password", mymiddleware.Guest)
	{
		// 展示发送重置密码链接 email 的页面
		context.RegisterHandler(pwdRouter.GET, "/reset", password.ShowLinkForm).Name = "password.show_link_form"
		// 发送重置密码链接的 email
		context.RegisterHandler(pwdRouter.POST, "/email", password.Email).Name = "password.email"
		// 展示重置密码的页面
		context.RegisterHandler(pwdRouter.GET, "/reset/:token", password.ShowResetForm).Name = "password.show_reset_form"
		// 重置密码
		context.RegisterHandler(pwdRouter.POST, "/reset", password.Update).Name = "password.update"
	}

	// +++++++++++++++ Email 认证相关路由 +++++++++++++++
	verificationRouter := ee.Group("/verification")
	{
		// 展示发送激活用户链接邮件的页面
		context.RegisterHandler(verificationRouter.GET, "/verify",
			wrapper.User(verification.ShowLinkForm),
			mymiddleware.AuthAndNoCheckActived,
		).Name = "verification.show_link_form"
		// 激活用户
		context.RegisterHandler(verificationRouter.GET, "/verify/:token",
			verification.Verify,
			NewRateLimiter(1*time.Minute, 6),
		).Name = "verification.verify"
		// 重新发送激活用户链接
		context.RegisterHandler(verificationRouter.GET, "/resend",
			wrapper.User(verification.Resend),
			mymiddleware.Auth,
		).Name = "verification.resend"
	}
}
