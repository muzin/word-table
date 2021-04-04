package wordtable

import (
	"fmt"
	"math"
	"strings"
)

const (
	WordTable_T11 string = "+"
	WordTable_T12 string = "+"
	WordTable_T13 string = "+"
	WordTable_T21 string = "+"
	WordTable_T22 string = "+"
	WordTable_T23 string = "+"
	WordTable_T31 string = "+"
	WordTable_T32 string = "+"
	WordTable_T33 string = "+"
	WordTable_TH  string = "-" //"─"
	WordTable_TV  string = "|"
)

type WordTable struct {

	// 表头
	header []WordTableHeader

	// 数据
	body [][]string

	stringTable []string

	// 最大列
	maxCols int

	// 是否更新过
	updated bool

}

type WordTableHeader struct {

	// 标题
	Title string

	// 宽度	默认为 标题长度 + 内边距 * 2
	Width int

	// 内边距  默认为 2
	Padding int

	// 自动宽度	如果width为0，自动调整宽度，默认 列最大值 + 4
	AutoWidth bool

	// 文本对齐方式	center: 居中 left: 居左 right: 居右  默认：居中
	TextAlign string
}

// 文本居中
var Center = "center"

// 文本居左
var Left = "left"

// 文本居右
var Right = "right"

func NewWordTable(header []WordTableHeader, body [][]string) *WordTable {
	w := &WordTable{header: header, body: body}
	w.updated = true
	w.checkAutoWidth()
	return w
}

// 添加 Body
func (this *WordTable) SetHeader(header []WordTableHeader) {
	this.header = header
	this.updated = true
	this.checkAutoWidth()
}

// 检查是否开启 自动宽度
func (this *WordTable) checkAutoWidth(){
	headers := this.header
	if headers != nil {
		for i, header := range headers {
			if header.Width == 0 {
				this.header[i].AutoWidth = true
			}
			if header.Padding < 0 {
				this.header[i].Padding = 2
			}
		}
	}
}

// 添加 Body
func (this *WordTable) AppendBody(item []string) {
	this.body = append(this.body, item)
	this.updated = true
}

// 添加 Body
func (this *WordTable) SetBody(body [][]string) {
	this.body = body
	this.updated = true
}

func (this *WordTable) Reset(){
	this.updated = true
	this.header = make([]WordTableHeader, 0)
	this.body = make([][]string, 0)
}

// 打印
func (this *WordTable) String() string{
	return printTable(this.header, this.body)
}

// 打印
func (this *WordTable) Println(){
	fmt.Println(this.String());
}

// 计算 表 宽度
func computTableWidth(headers *[]WordTableHeader, body *[][]string){

	headersLen := len(*headers)

	// 设置 header 的 width
	for hi, hv := range *headers {
		//(*headers)[hi].width = len(strings.TrimSpace(hv.title))
		(*headers)[hi].Width = len(strings.TrimSpace(hv.Title)) + hv.Padding * 2
	}

	// 比较最大长度，给header 的 width 设置 最大长度
	for _, bv := range *body {
		// bci body column index
		// bcv body column value
		for bci, bcv := range bv {
			if bci < headersLen {
				headerWidth := (*headers)[bci].Width
				bcvLen := len(bcv)
				if bcvLen > headerWidth {
					//(*headers)[bci].width = bcvLen
					(*headers)[bci].Width = bcvLen + (*headers)[bci].Padding * 2
				}
			}
		}
	}

}

func computTableAlign(headers *[]WordTableHeader, body *[][]string) {
	// 表头处理  居中
	//for hi, hv := range *headers {
	//	totalWidth := hv.width
	//	spaceWidth := totalWidth - len(hv.title)
	//	floor := int(math.Floor(float64(spaceWidth / 2)))
	//	ceil := int(math.Ceil(float64(spaceWidth / 2)))
	//	(*headers)[hi].title = fmt.Sprintf("% " + strconv.Itoa(floor) + "s", "") +
	//		strings.TrimSpace((*headers)[hi].title) +
	//		fmt.Sprintf("% " + strconv.Itoa(ceil) + "s", "")
	//}

	//for bodyRowIdx, bodyRowVal := range *body {
	//
	//	// 每一行 循环
	//	// bci 每一行中的每一列
	//	for bodyColumnIdx, bodyColumnVal := range bodyRowVal {
	//		// 表头以及表头的索引
	//		//headerIdx := bodyColumnIdx
	//		headerVal := (*headers)[bodyColumnIdx]
	//
	//		totalWidth := headerVal.Width
	//		spaceWidth := totalWidth - len(bodyColumnVal)
	//		if Left == headerVal.TextAlign {
	//			(*body)[bodyRowIdx][bodyColumnIdx] = strings.TrimSpace(bodyColumnVal) +
	//				fmt.Sprintf("% " + strconv.Itoa(spaceWidth) + "s", "")
	//		}else if Right == headerVal.TextAlign {
	//			(*body)[bodyRowIdx][bodyColumnIdx] = fmt.Sprintf("% " + strconv.Itoa(spaceWidth) + "s", "") +
	//				strings.TrimSpace(bodyColumnVal)
	//		}else{
	//			floor := int(math.Floor(float64(spaceWidth / 2)))
	//			// ceil := int(math.Ceil(float64(spaceWidth / 2)))
	//			ceil := spaceWidth - floor
	//			(*body)[bodyRowIdx][bodyColumnIdx] = fmt.Sprintf("% " + strconv.Itoa(floor) + "s", "") +
	//				strings.TrimSpace(bodyColumnVal) +
	//				fmt.Sprintf("% " + strconv.Itoa(ceil) + "s", "")
	//		}
	//	}
	//}
}

func printTable(headers []WordTableHeader, body [][]string) string {

	computTableWidth(&headers, &body)

	computTableAlign(&headers, &body)

	var tablestr = ""
	tablestr += printHeader(headers) + "\n"
	for _, v := range body {
		tablestr += printBody(v, headers) + "\n"
		tablestr += printButtom(headers) + "\n"
	}
	return tablestr
}

func printHeader(headers []WordTableHeader) string {
	fleng := len(headers)
	//printer top line
	var topLineStr string
	for i, header := range headers {

		switch i {
		case 0:
			var midstr string
			for x := 1; x <= header.Width; x++ {
				midstr = midstr + WordTable_TH
			}
			topLineStr = topLineStr + WordTable_T11 + midstr
			//如果只有一个字段
			if fleng == 1 {
				topLineStr = topLineStr + WordTable_T13
			}

		case fleng - 1:
			var midstr string
			for x := 1; x <= header.Width; x++ {
				midstr = midstr + WordTable_TH
			}
			topLineStr = topLineStr + WordTable_T12 + midstr + WordTable_T13
		default:
			var midstr string
			for x := 1; x <= header.Width; x++ {
				midstr = midstr + WordTable_TH
			}
			topLineStr = topLineStr + WordTable_T12 + midstr
		}
	}


	//print feild label
	var labstr string
	for fj, header := range headers {
		tr := []rune(header.Title)
		trLen := len(tr)
		var spLen int
		spLen = header.Width - trLen - (len(header.Title)-trLen)/2

		//var midStr string
		//midStr := fillStr(spLen, " ")

		floor := int(math.Floor(float64(spLen / 2)))
		// ceil := int(math.Ceil(float64(spLen / 2)))
		ceil := spLen - floor
		str := fillStr(floor, " ") + header.Title + fillStr(ceil, " ")
		labstr = labstr + WordTable_TV + str
		if fj == fleng-1 {
			labstr = labstr + WordTable_TV
		}

	}

	//print buttom of head
	var butStr string
	for bi, bv := range headers {
		var midstr string
		for x := 1; x <= bv.Width; x++ {
			midstr = midstr + WordTable_TH
		}
		switch bi {
		case 0:
			butStr = WordTable_T21 + midstr
			if fleng == 1 {
				butStr = butStr + WordTable_T23
			}
		case fleng - 1:
			butStr = butStr + WordTable_T22 + midstr + WordTable_T23
		default:
			butStr = butStr + WordTable_T22 + midstr
		}
	}
	return topLineStr + "\n" + labstr + "\n" + butStr
}
func printBody(body []string, headers []WordTableHeader) string {
	bodyLen := len(body)
	headersLen := len(headers)
	if bodyLen > headersLen {
		return ""
	} else {
		var fstr string
		for i, v := range body {
			// body [x] string
			ts := []rune(v)
			// body [x] string length
			tsLen := len(ts)
			stLen := len(v)
			spLen := headers[i].Width - tsLen - (stLen-tsLen)/2
			var midstr string = fillStr(spLen, " ")

			headerVal := headers[i]
			if Left == headerVal.TextAlign {
				fstr = fstr + WordTable_TV + v + midstr
			}else if Right == headerVal.TextAlign {
				fstr = fstr + WordTable_TV + midstr + v
			}else{
				floor := int(math.Floor(float64(spLen / 2)))
				//ceil := int(math.Ceil(float64(spLen / 2)))
				ceil := spLen - floor
				fstr = fstr + WordTable_TV + fillStr(floor, " ") + v + fillStr(ceil, " ")
			}

			if i == bodyLen-1 {
				fstr = fstr + WordTable_TV
			}
		}
		return fstr
	}
}
func printButtom(headers []WordTableHeader) string {
	var fleng int = len(headers)
	var butStr string
	for bi, bv := range headers {
		var midstr string
		for x := 1; x <= bv.Width; x++ {
			midstr = midstr + WordTable_TH
		}
		switch bi {
		case 0:
			butStr = WordTable_T31 + midstr
			if fleng == 1 {
				butStr = butStr + WordTable_T33
			}
		case fleng - 1:
			butStr = butStr + WordTable_T32 + midstr + WordTable_T33
		default:
			butStr = butStr + WordTable_T32 + midstr
		}
	}
	return butStr
}

func fillStr(l int, c string) string {
	var str = ""
	for i := 0; i < l; i++ {
		str += c
	}
	return str
}
