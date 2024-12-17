package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	start := time.Now() // 记录开始时间
	// 设置小票的尺寸：宽度较小，符合实际小票的大小（例如80mm宽，297mm长的打印纸）
	const W = 300 // 宽度
	const H = 600 // 高度
	dc := gg.NewContext(W, H)

	// 设置背景颜色为白色
	dc.SetColor(color.White)
	dc.Clear()

	// 设置文本颜色为黑色
	dc.SetRGB(0, 0, 0)

	// 加载字体，使用中国常见的小票字体
	if err := dc.LoadFontFace("fonts/MiSans-Normal.ttf", 14); err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	// 绘制商家信息（例如：商店名称、电话、地址等）
	drawTextCenter(dc, "商店名称: 示例店铺", W/2, 20)
	drawTextCenter(dc, "电话: 123-45678901", W/2, 40)
	drawTextCenter(dc, "地址: 北京市某某路123号", W/2, 60)

	// 绘制日期和交易编号
	drawTextCenter(dc, "日期: 2024年12月17日", W/2, 100)
	drawTextCenter(dc, "交易编号: 123456789", W/2, 120)

	// 绘制商品列表
	drawTextCenter(dc, "商品列表", W/2, 160)

	// 绘制商品列表的上方虚线
	drawDashedLine(dc, 10, 140, W-10, 140)

	// 商品数据
	items := []struct {
		name   string
		price  float64
		amount int
	}{
		{"商品一", 9.99, 2},
		{"商品二", 4.50, 1},
		{"商品三", 3.00, 3},
	}

	yOffset := 180.0
	dc.DrawStringAnchored(fmt.Sprintf("%-15s ¥%-6s  %-6s ¥%-4s", "商品名", "单价", "数量", "总价"), 10, yOffset, 0, 0)
	yOffset += 20
	for _, item := range items {
		// 商品名、单价、数量、总价
		dc.DrawStringAnchored(fmt.Sprintf("%-15s ¥%-8.2f  %-10d ¥%-4.2f", item.name, item.price, item.amount, item.price*float64(item.amount)), 10, yOffset, 0, 0)

		yOffset += 20
	}

	// 绘制总计
	total := 0.0
	for _, item := range items {
		total += item.price * float64(item.amount)
	}
	drawTextCenter(dc, fmt.Sprintf("总计: ¥%.2f", total), W/2, yOffset+10)

	// 绘制商品列表的下方虚线
	drawDashedLine(dc, 10, yOffset+60, W-10, yOffset+60)

	// 绘制商家提醒（例如：感谢光临）
	drawTextCenter(dc, "感谢光临，欢迎下次光临！", W/2, yOffset+80)

	// 保存为PNG图像
	err := dc.SavePNG("receipt.png")
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("小票已生成，保存为 receipt.png")

	elapsed := time.Since(start) // 计算耗时
	fmt.Printf("耗时: %s\n", elapsed)
}

// 绘制虚线的函数
func drawDashedLine(dc *gg.Context, x1, y1, x2, y2 float64) {
	dashLength := 2.0 // 每段线的长度
	gapLength := 2.0  // 线段之间的间隙
	for i := 0.0; i < x2-x1; i += dashLength + gapLength {
		dc.DrawLine(x1+i, y1, x1+i+dashLength, y2)
		dc.Stroke()
	}
}

// 绘制居中的文本
func drawTextCenter(dc *gg.Context, text string, x, y float64) {
	// MeasureString is only used for calculating position
	dc.DrawStringAnchored(text, x, y, 0.5, 0)
}
