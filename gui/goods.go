package gui

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"time"
)

func (g *Gui) Goods() *widget.LayoutEle {
	layout := widget.NewLayoutEle(40, 65, 850, 480, g.rightLayoutBody.Handle)
	// layout.LayoutItem_SetMargin(10, 10, 10, 10)
	// layout.ShowLayoutFrame(true)

	headerLayout := widget.NewLayoutEle(1, 1, 850, 40, layout.Handle)
	// headerLayout.ShowLayoutFrame(true)

	searchText := widget.NewEdit(0, 0, 120, 30, headerLayout.Handle)
	searchText.SetDefaultText("供货方号/商家货号")
	searchBut := widget.NewButton(0, 0, 50, 30, "查询", headerLayout.Handle)
	searchBut.Event_BnClick(func(pbHandled *bool) int {
		g.App.Alert("商品", searchText.GetTextEx())
		return 0
	})

	addBut := widget.NewButton(0, 0, 70, 30, "＋新增商品", headerLayout.Handle)
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
	ls.AddColumnText(100, "name1", "供方货号")
	ls.AddColumnText(100, "name2", "商家名称")
	ls.AddColumnText(100, "name3", "商家货号")
	ls.AddColumnText(120, "name4", "商家颜色")
	ls.AddColumnText(60, "name5", "商家价格")
	ls.AddColumnText(60, "name6", "前台售价")
	ls.AddColumnText(200, "name7", "商家地址")
	ls.AddColumnText(130, "name8", "添加时间")
	ls.AddColumnText(100, "name9", "操作")
	// 循环添加数据
	for i := 0; i < 13; i++ {
		// 添加行
		index := ls.AddItemText("")
		// 置行数据
		ls.SetItemText(index, 0, "PD0326")
		ls.SetItemText(index, 1, "青春鞋业")
		ls.SetItemText(index, 2, "D172")
		ls.SetItemText(index, 3, "黑色*黄棕色*蓝色")
		ls.SetItemText(index, 4, "33")
		ls.SetItemText(index, 5, "58")
		ls.SetItemText(index, 6, "义全街60号，幸福街口公交站这边（工商银行正对面）")
		ls.SetItemText(index, 7, time.Now().Format("2006-01-02 15:04:05"))
	}

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
