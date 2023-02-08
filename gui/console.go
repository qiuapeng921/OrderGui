package gui

import (
	"fmt"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"time"
)

func (g *Gui) Console() *widget.LayoutEle {
	// 主体布局
	layout := widget.NewLayoutEle(40, 65, 850, 480, g.rightLayoutBody.Handle)
	// layout.LayoutItem_SetMargin(10, 10, 10, 10)
	// layout.ShowLayoutFrame(true)

	// TODO 顶部布局
	headerLayout := widget.NewLayoutEle(1, 1, 850, 40, layout.Handle)
	// headerLayout.ShowLayoutFrame(true)

	startTime := widget.NewDateTime(0, 20, 100, 30, headerLayout.Handle)
	endTime := widget.NewDateTime(120, 20, 100, 30, headerLayout.Handle)
	endTime.SetDate(time.Now().Year(), time.Now().Day(), time.Now().Day())

	searchButton := widget.NewButton(0, 20, 50, 30, "搜索", headerLayout.Handle)
	searchButton.LayoutItem_SetMargin(5, 0, 0, 0)
	searchButton.Event_BnClick(func(pbHandled *bool) int {
		var startYear, startMonth, startDay int
		var endYear, endMonth, endDay int
		startTime.GetDate(&startYear, &startMonth, &startDay)
		endTime.GetDate(&endYear, &endMonth, &endDay)
		g.App.Alert("开始时间", fmt.Sprintf("%d-%d-%d", startYear, startMonth, startDay))
		g.App.Alert("结束时间", fmt.Sprintf("%d-%d-%d", endYear, endMonth, endDay))
		return 0
	})

	widget.NewShapeText(0, 0, 60, 30, "商品查询:", headerLayout.Handle).LayoutItem_SetMargin(20, 0, 0, 0)
	goodsSearchText := widget.NewEdit(0, 0, 100, 30, headerLayout.Handle)
	goodsSearchText.SetDefaultText("供应商编码")
	goodsSearchText.Event_EDIT_CHANGED(func(pbHandled *bool) int {
		g.App.Alert("搜索框", goodsSearchText.GetTextEx())
		goodsSearchText.SetText("")
		return 0
	})

	widget.NewShapeText(0, 0, 60, 30, "订单出货:", headerLayout.Handle).LayoutItem_SetMargin(20, 0, 0, 0)
	warehouseOutText := widget.NewEdit(0, 0, 100, 30, headerLayout.Handle)
	warehouseOutText.SetDefaultText("订单条码")
	warehouseOutText.Event_EDIT_CHANGED(func(pbHandled *bool) int {
		g.App.Alert("搜索框", warehouseOutText.GetTextEx())
		warehouseOutText.SetText("")
		return 0
	})

	widget.NewShapeText(0, 0, 60, 30, "订单退货:", headerLayout.Handle).LayoutItem_SetMargin(20, 0, 0, 0)
	warehouseInputText := widget.NewEdit(0, 0, 100, 30, headerLayout.Handle)
	warehouseInputText.SetDefaultText("订单条码")
	warehouseInputText.Event_EDIT_CHANGED(func(pbHandled *bool) int {
		g.App.Alert("搜索框", warehouseInputText.GetTextEx())
		warehouseInputText.SetText("")
		return 0
	})

	// TODO 中间布局
	bodyLayout := widget.NewLayoutEle(0, 0, 850, 70, layout.Handle)
	bodyLayout.LayoutItem_SetMargin(0, 5, 0, 0)
	// bodyLayout.ShowLayoutFrame(true)

	// 创建List
	ls := widget.NewList(10, 33, 850, 60, bodyLayout.Handle)
	// 创建表头数据适配器
	ls.CreateAdapterHeader()
	// 创建数据适配器: 2列
	ls.CreateAdapter(4)
	// 列表_置项默认高度
	ls.SetItemHeightDefault(28, 28)

	// 添加列
	ls.AddColumnText(210, "name1", "成本")
	ls.AddColumnText(210, "name2", "销量")
	ls.AddColumnText(210, "name3", "利润")
	ls.AddColumnText(210, "name4", "入库")
	// 循环添加数据
	for i := 0; i < 1; i++ {
		// 添加行
		index := ls.AddItemText("")
		// 置行数据
		ls.SetItemText(index, 0, "1500")
		ls.SetItemText(index, 1, "300")
		ls.SetItemText(index, 2, "1000")
		ls.SetItemText(index, 3, "100")
	}

	// TODO 底部布局
	footerLayout := widget.NewLayoutEle(0, 0, 850, 350, layout.Handle)
	// footerLayout.ShowLayoutFrame(true)
	footerLayout.LayoutItem_SetMargin(0, 10, 0, 0)

	// TODO 底部左边布局
	footerLeftLayout := widget.NewLayoutEle(0, 0, 330, 340, footerLayout.Handle)
	// footerLeftLayout.ShowLayoutFrame(true)

	costAddBut := widget.NewButton(0, 0, 70, 30, "＋添加成本", footerLeftLayout.Handle)
	costAddBut.Event_BnClick(func(pbHandled *bool) int {
		g.costAdd()
		return 0
	})
	costExportBut := widget.NewButton(0, 0, 50, 30, "↓导出", footerLeftLayout.Handle)
	costExportBut.LayoutItem_SetMargin(200, 0, 0, 0)
	costExportBut.Event_BnClick1(func(hEle int, pbHandled *bool) int {
		g.App.Alert("成本", "导出")
		return 0
	})
	// 创建List
	costList := widget.NewList(10, 33, 320, 300, footerLeftLayout.Handle)
	costList.LayoutItem_SetMargin(0, 0, 0, 0)
	// 创建表头数据适配器
	costList.CreateAdapterHeader()
	// 创建数据适配器: 2列
	costList.CreateAdapter(10)
	// 列表_置项默认高度
	costList.SetItemHeightDefault(28, 28)

	// 添加列
	costList.AddColumnText(60, "name1", "支出金额")
	costList.AddColumnText(80, "name2", "支出项目")
	costList.AddColumnText(80, "name3", "支出成员")
	costList.AddColumnText(70, "name4", "支出时间")

	// 循环添加数据
	for i := 0; i < 10; i++ {
		// 添加行
		index := costList.AddItemText("")
		// 置行数据
		costList.SetItemText(index, 0, "100")
		costList.SetItemText(index, 1, "入库")
		costList.SetItemText(index, 2, "admin")
		costList.SetItemText(index, 3, time.Now().Format("2006-01-02"))
	}

	// TODO 底部右边布局
	footerRightLayout := widget.NewLayoutEle(0, 0, 510, 340, footerLayout.Handle)
	// footerRightLayout.ShowLayoutFrame(true)

	widget.NewShapeText(0, 0, 60, 30, "退货出库:", footerRightLayout.Handle)
	goodsReturnInputText := widget.NewEdit(0, 0, 100, 30, footerRightLayout.Handle)
	goodsReturnInputText.SetDefaultText("供货方号")
	goodsReturnInputText.Event_EDIT_CHANGED(func(pbHandled *bool) int {
		g.App.Alert("退供", goodsReturnInputText.GetTextEx())
		goodsReturnInputText.SetText("")
		return 0
	})

	goodsReturnExportBut := widget.NewButton(0, 0, 70, 30, "↓退供导出", footerRightLayout.Handle)
	goodsReturnExportBut.LayoutItem_SetMargin(280, 0, 0, 0)
	goodsReturnExportBut.Event_BnClick1(func(hEle int, pbHandled *bool) int {
		g.App.Alert("退供", "导出")
		return 0
	})

	// 创建List
	goodsReturnList := widget.NewList(10, 33, 500, 300, footerRightLayout.Handle)
	goodsReturnList.LayoutItem_SetMargin(0, 0, 0, 0)
	// 创建表头数据适配器
	goodsReturnList.CreateAdapterHeader()
	// 创建数据适配器: 2列
	goodsReturnList.CreateAdapter(10)
	// 列表_置项默认高度
	goodsReturnList.SetItemHeightDefault(28, 28)

	// 添加列
	goodsReturnList.AddColumnText(100, "name1", "商家名称")
	goodsReturnList.AddColumnText(80, "name2", "商家货号")
	goodsReturnList.AddColumnText(70, "name3", "货物数量")
	goodsReturnList.AddColumnText(70, "name4", "商家价格")
	goodsReturnList.AddColumnText(200, "name4", "商家地址")

	// 循环添加数据
	for i := 0; i < 10; i++ {
		// 添加行
		index := goodsReturnList.AddItemText("")
		// 置行数据
		goodsReturnList.SetItemText(index, 0, "转角")
		goodsReturnList.SetItemText(index, 1, "PD0435")
		goodsReturnList.SetItemText(index, 2, "7")
		goodsReturnList.SetItemText(index, 3, "35")
		goodsReturnList.SetItemText(index, 4, "天后路158号 艺兴堂1号门斜对面卖咸饭牛肉面隔壁")
	}

	return layout
}

func (g *Gui) costAdd() {
	// 创建模态窗口
	winModal := window.NewModalWindow(260, 200, "添加成本", g.win.GetHWND(), xcc.Window_Style_Modal)
	// 设置边框
	winModal.SetBorderSize(1, 30, 1, 1)
	// // 显示模态窗口
	// winModal.DoModal()
	winModal.ShowWindow(xcc.SW_SHOW)

	widget.NewShapeText(20, 50, 50, 30, "金额:", winModal.Handle)
	widget.NewShapeText(20, 90, 50, 30, "项目:", winModal.Handle)
	money := widget.NewEdit(60, 50, 150, 30, winModal.Handle)
	project := widget.NewEdit(60, 90, 150, 30, winModal.Handle)
	addBut := widget.NewButton(160, 140, 50, 30, "提交", winModal.Handle)
	addBut.Event_BnClick(func(pbHandled *bool) int {
		if money.GetTextEx() == "" {
			g.App.Alert("成本", "金额不能为空")
			return 0
		}
		if project.GetTextEx() == "" {
			g.App.Alert("成本", "项目不能为空")
			return 0
		}
		return 0
	})
}
