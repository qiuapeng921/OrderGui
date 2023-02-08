package gui

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"time"
)

func (g *Gui) Order() *widget.LayoutEle {
	layout := widget.NewLayoutEle(40, 65, 850, 480, g.rightLayoutBody.Handle)
	// layout.LayoutItem_SetMargin(10, 10, 10, 10)
	// layout.ShowLayoutFrame(true)

	headerLayout := widget.NewLayoutEle(1, 1, 850, 40, layout.Handle)
	// headerLayout.ShowLayoutFrame(true)

	searchText := widget.NewEdit(0, 0, 120, 30, headerLayout.Handle)
	searchText.SetDefaultText("供应商编码")
	searchBut := widget.NewButton(0, 0, 50, 30, "查询", headerLayout.Handle)
	searchBut.Event_BnClick(func(pbHandled *bool) int {
		g.App.Alert("搜索框", searchText.GetTextEx())
		return 0
	})

	syncOrderBut := widget.NewButton(0, 0, 80, 30, "同步三方订单", headerLayout.Handle)
	syncOrderBut.Event_BnClick(func(pbHandled *bool) int {
		g.App.Alert("同步", searchText.GetTextEx())
		return 0
	})

	addBut := widget.NewButton(0, 0, 70, 30, "＋新增订单", headerLayout.Handle)
	addBut.LayoutItem_SetMargin(500, 0, 0, 0)
	addBut.Event_BnClick(func(pbHandled *bool) int {
		addWin := window.NewModalWindow(500, 600, "新增订单", g.win.GetHWND(), xcc.Window_Style_Modal)
		// 设置边框
		addWin.SetBorderSize(1, 30, 1, 1)
		addWin.DoModal()
		return 0
	})

	bodyLayout := widget.NewLayoutEle(1, 1, 850, 440, layout.Handle)
	bodyLayout.ShowLayoutFrame(true)
	// 创建List
	ls := widget.NewList(10, 33, 840, 430, bodyLayout.Handle)
	// 创建表头数据适配器
	ls.CreateAdapterHeader()
	// 创建数据适配器: 1列
	ls.CreateAdapter(9)
	// 列表_置项默认高度
	ls.SetItemHeightDefault(28, 28)

	// 添加列
	ls.AddColumnText(120, "name1", "前台订单")
	ls.AddColumnText(50, "name2", "数量")
	ls.AddColumnText(50, "name3", "尺码")
	ls.AddColumnText(50, "name4", "颜色")
	ls.AddColumnText(100, "name5", "供方货号")
	ls.AddColumnText(120, "name6", "订单条码")
	ls.AddColumnText(120, "name7", "商家名称")
	ls.AddColumnText(120, "name8", "商家货号")
	ls.AddColumnText(120, "name9", "添加时间")
	ls.AddColumnText(100, "name9", "操作")
	// 循环添加数据
	for i := 0; i < 10; i++ {
		// 添加行
		index := ls.AddItemText("")
		// 置行数据
		ls.SetItemText(index, 0, "WB230112001173")
		ls.SetItemText(index, 1, "1")
		ls.SetItemText(index, 2, "43")
		ls.SetItemText(index, 3, "紫红色")
		ls.SetItemText(index, 4, "PD0522")
		ls.SetItemText(index, 5, "c7e7f8e47cc97d80")
		ls.SetItemText(index, 6, "火山鞋业")
		ls.SetItemText(index, 6, "861")
		ls.SetItemText(index, 7, time.Now().Format("2006-01-02 15:04:05"))
	}

	return layout
}
