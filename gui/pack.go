package gui

import (
	"github.com/twgh/xcgui/widget"
	"time"
)

func (g *Gui) Pack() *widget.LayoutEle {
	layout := widget.NewLayoutEle(40, 65, 850, 480, g.rightLayoutBody.Handle)
	// layout.LayoutItem_SetMargin(10, 10, 10, 10)
	// layout.ShowLayoutFrame(true)

	headerLayout := widget.NewLayoutEle(1, 1, 850, 40, layout.Handle)
	// headerLayout.ShowLayoutFrame(true)

	searchText := widget.NewEdit(0, 0, 120, 30, headerLayout.Handle)
	searchText.SetDefaultText("包裹信息/封装信息")
	searchBut := widget.NewButton(0, 0, 50, 30, "查询", headerLayout.Handle)
	searchBut.Event_BnClick(func(pbHandled *bool) int {
		g.App.Alert("打包", searchText.GetTextEx())
		return 0
	})

	addBut := widget.NewButton(0, 0, 70, 30, "＋新增打包", headerLayout.Handle)
	addBut.LayoutItem_SetMargin(600, 0, 0, 0)
	addBut.Event_BnClick(func(pbHandled *bool) int {
		g.GoodsAdd()
		return 0
	})

	bodyLayout := widget.NewLayoutEle(1, 1, 850, 440, layout.Handle)
	// bodyLayout.ShowLayoutFrame(true)
	// 创建List
	ls := widget.NewList(10, 33, 840, 430, bodyLayout.Handle)
	// 创建表头数据适配器
	ls.CreateAdapterHeader()
	// 创建数据适配器: 1列
	ls.CreateAdapter(5)
	// 列表_置项默认高度
	ls.SetItemHeightDefault(28, 28)

	// 添加列
	ls.AddColumnText(130, "name1", "车辆信息")
	ls.AddColumnText(130, "name2", "封包信息")
	ls.AddColumnText(130, "name3", "包裹信息")
	ls.AddColumnText(150, "name4", "添加时间")
	ls.AddColumnText(230, "name5", "操作")
	// 循环添加数据
	for i := 0; i < 13; i++ {
		// 添加行
		index := ls.AddItemText("")
		// 置行数据
		ls.SetItemText(index, 0, "川A50038")
		ls.SetItemText(index, 1, "川A50038")
		ls.SetItemText(index, 2, "川A50038")
		ls.SetItemText(index, 3, time.Now().Format("2006-01-02 15:04:05"))
	}

	return layout
}
