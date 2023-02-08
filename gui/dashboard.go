package gui

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func (g *Gui) Dashboard() {
	g.win = window.New(0, 0, 1000, 560, appTitle, 0, xcc.Window_Style_Modal)
	// 设置窗口图标
	g.win.SetIcon(xc.XImage_LoadSvgStringW(svgIcon))
	// 禁止改变大小
	g.win.EnableDragBorder(false)
	// 设置边框
	g.win.SetBorderSize(1, 30, 1, 1)

	// 菜单主体部分, 页面都放进这里面
	g.leftLayoutBody = widget.NewLayoutEle(10, 40, 100, 500, g.win.Handle)
	// g.leftLayoutBody.ShowLayoutFrame(true)

	// 放在水平布局元素中的组件, x, y绝对坐标是无效的
	dashBoardTabBtn := widget.NewButton(0, 0, 100, 40, "控制台", g.leftLayoutBody.Handle)
	orderTabBtn := widget.NewButton(0, 30, 100, 40, "订单管理", g.leftLayoutBody.Handle)
	goodsTabBtn := widget.NewButton(0, 60, 100, 40, "商品管理", g.leftLayoutBody.Handle)
	merchantTabBtn := widget.NewButton(0, 90, 100, 40, "商家管理", g.leftLayoutBody.Handle)
	packTabBtn := widget.NewButton(0, 120, 100, 40, "打包管理", g.leftLayoutBody.Handle)
	// 设为单选类型按钮
	dashBoardTabBtn.SetTypeEx(xcc.Button_Type_Radio)
	orderTabBtn.SetTypeEx(xcc.Button_Type_Radio)
	goodsTabBtn.SetTypeEx(xcc.Button_Type_Radio)
	merchantTabBtn.SetTypeEx(xcc.Button_Type_Radio)
	packTabBtn.SetTypeEx(xcc.Button_Type_Radio)
	// 设置为默认按钮样式, 就不会是单选按钮样式了
	dashBoardTabBtn.SetStyle(xcc.Button_Style_Icon)
	orderTabBtn.SetStyle(xcc.Button_Style_Icon)
	goodsTabBtn.SetStyle(xcc.Button_Style_Icon)
	merchantTabBtn.SetStyle(xcc.Button_Style_Icon)
	packTabBtn.SetStyle(xcc.Button_Style_Icon)
	// 默认选中第一个
	dashBoardTabBtn.SetCheck(true)

	// 右边页面主体部分, 页面都放进这里面
	g.rightLayoutBody = widget.NewLayoutEle(10+100+10, 40, 870, 510, g.win.Handle)
	// g.rightLayoutBody.ShowLayoutFrame(true)

	ConsolePage := g.Console()
	orderPage := g.Order()
	goodsPage := g.Goods()
	merchantPage := g.Merchant()
	packPage := g.Pack()
	// 只让第一页显示, 其他页都设为不显示
	orderPage.Show(false)
	goodsPage.Show(false)
	merchantPage.Show(false)
	packPage.Show(false)

	// 给按钮绑定页面
	dashBoardTabBtn.SetBindEle(ConsolePage.Handle)
	orderTabBtn.SetBindEle(orderPage.Handle)
	goodsTabBtn.SetBindEle(goodsPage.Handle)
	merchantTabBtn.SetBindEle(merchantPage.Handle)
	packTabBtn.SetBindEle(packPage.Handle)

	g.win.ShowWindow(xcc.SW_SHOW)
}
