package gui

import (
	"github.com/twgh/xcgui/widget"
)

func (g *Gui) Merchant() *widget.LayoutEle {
	layout := widget.NewLayoutEle(40, 65, 850, 480, g.rightLayoutBody.Handle)
	// layout.LayoutItem_SetMargin(10, 10, 10, 10)
	// layout.ShowLayoutFrame(true)

	headerLayout := widget.NewLayoutEle(1, 1, 850, 40, layout.Handle)
	// headerLayout.ShowLayoutFrame(true)

	searchText := widget.NewEdit(0, 0, 120, 30, headerLayout.Handle)
	searchText.SetDefaultText("商家名称")
	searchBut := widget.NewButton(0, 0, 50, 30, "查询", headerLayout.Handle)
	searchBut.Event_BnClick(func(pbHandled *bool) int {
		g.App.Alert("商家", searchText.GetTextEx())
		return 0
	})

	addBut := widget.NewButton(0, 0, 70, 30, "＋新增商家", headerLayout.Handle)
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
	ls.CreateAdapter(9)
	// 列表_置项默认高度
	ls.SetItemHeightDefault(28, 28)

	// 添加列
	ls.AddColumnText(100, "name1", "商家名称")
	ls.AddColumnText(350, "name2", "商家地址")
	ls.AddColumnText(100, "name3", "操作")
	// 循环添加数据
	for i := 0; i < 13; i++ {
		// 添加行
		index := ls.AddItemText("")
		// 置行数据
		ls.SetItemText(index, 0, "青春鞋业")
		ls.SetItemText(index, 1, "义全街60号，幸福街口公交站这边（工商银行正对面）")
	}

	return layout
}
