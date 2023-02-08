package gui

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func (g *Gui) LoginView() {
	w := window.New(0, 0, 400, 300, appTitle, 0, xcc.Window_Style_Modal)
	// 设置窗口图标
	w.SetIcon(xc.XImage_LoadSvgStringW(svgIcon))
	// 禁止改变大小
	w.EnableDragBorder(false)
	// 设置边框
	w.SetBorderSize(0, 30, 0, 0)

	widget.NewShapeText(60, 70, 100, 30, "账号:", w.Handle)
	username := widget.NewEdit(100, 70, 200, 30, w.Handle)
	username.SetText("admin")

	widget.NewShapeText(60, 130, 100, 30, "密码:", w.Handle)
	password := widget.NewEdit(100, 130, 200, 30, w.Handle)
	password.EnablePassword(true)
	password.SetText("123456")

	widget.NewShapeText(60, 180, 100, 30, "密码:", w.Handle)
	database := widget.NewEdit(100, 180, 200, 30, w.Handle)
	database.SetText("order")

	loginBut := widget.NewButton(135, 230, 86, 26, "登录", w.Handle)
	loginBut.Event_BnClick(func(pbHandled *bool) int {
		if username.GetTextEx() != "admin" {
			g.App.Alert("信息:", "账号错误")
			return 0
		}
		if password.GetTextEx() != "123456" {
			g.App.Alert("信息:", "密码错误")
			return 0
		}
		w.CloseWindow()
		g.Dashboard()
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
}
