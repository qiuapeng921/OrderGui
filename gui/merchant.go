package gui

import (
	"github.com/twgh/xcgui/widget"
)

func (g *Gui) Merchant() *widget.LayoutEle {
	layout := widget.NewLayoutEle(30, 65, 860, 480, g.rightLayoutBody.Handle)
	layout.ShowLayoutFrame(true)

	headerLayout := widget.NewLayoutEle(0, 2, 860, 40, layout.Handle)
	headerLayout.ShowLayoutFrame(true)

	widget.NewShapeText(0, 0, 500, 30, "商家", headerLayout.Handle)

	return layout
}
