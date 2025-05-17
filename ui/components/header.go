package components

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func Header(th *material.Theme, gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.
		Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx,
					material.Body1(th, "Refresh").Layout,
				)
			}))
}
