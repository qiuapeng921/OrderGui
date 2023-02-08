package gui

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func (g *Gui) Goods() *widget.LayoutEle {
	layout := widget.NewLayoutEle(30, 65, 860, 480, g.rightLayoutBody.Handle)
	layout.ShowLayoutFrame(true)

	headerLayout := widget.NewLayoutEle(0, 2, 860, 40, layout.Handle)
	headerLayout.ShowLayoutFrame(true)

	// // 页面中放置内容
	// widget.NewShapeText(0, 0, 500, 30, "商品", headerPage.Handle)

	searchText := widget.NewEdit(0, 10, 200, 30, headerLayout.Handle)
	searchText.SetDefaultText("商品编号")
	searchBut := widget.NewButton(220, 10, 60, 30, "搜索", headerLayout.Handle)
	searchBut.Event_BnClick(func(pbHandled *bool) int {
		g.App.Alert("111", searchText.GetTextEx())
		return 0
	})

	// addBut := widget.NewButton(0, 10, 60, 30, "新增", headerLayout.Handle)
	// addBut.Event_BnClick1(func(hEle int, pbHandled *bool) int {
	// 	g.GoodsAdd()
	// 	return 0
	// })

	return layout
}

func (g *Gui) GoodsAdd() {
	// 创建模态窗口
	winModal := window.NewModalWindow(500, 600, "商品创建", g.win.GetHWND(), xcc.Window_Style_Modal)
	// 设置边框
	winModal.SetBorderSize(1, 30, 1, 1)
	// 显示模态窗口
	winModal.DoModal()

}
