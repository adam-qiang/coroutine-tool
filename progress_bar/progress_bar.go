/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-23 00:45:14
 */

package progress_bar

import (
	"fmt"
	"strconv"
)

// Bar
// @Description: 进度条结构体
type Bar struct {
	percent int64  //百分比
	cur     int64  //当前进度位置
	total   int64  //总进度
	rate    string //进度条
	graph   string //显示符号
}

// NewBar 实例化一个进度条
// @receiver bar *Bar
// @param start int64
// @param total int64
func (bar *Bar) NewBar(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		//初始化进度条位置
		bar.rate += bar.graph
	}
}

// NewBarWithGraph 实例化一个进度条（可自定义进度条图形）
// @receiver bar *Bar
// @param start int64
// @param total int64
// @param graph string
func (bar *Bar) NewBarWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewBar(start, total)
}

// Run 执行进度条展示
// @receiver bar *Bar
// @param cur int64
func (bar *Bar) Run(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r %-"+strconv.FormatInt(bar.total/2, 10)+"s%3d%%  %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

// 获取进度百分比
// @receiver bar *Bar
// @return int64
func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}
