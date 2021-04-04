package wordtable

import (
	"testing"
)

func TestNewWordTable(t *testing.T) {

	t.Run("TestNewWordTable", func(t *testing.T) {

		headers := []WordTableHeader{
			WordTableHeader{ Title: "AAAAAbbbcc", TextAlign: Left, Padding: 2 },
			WordTableHeader{ Title: "运行状态", TextAlign: Center, Padding: 2 },
			WordTableHeader{ Title: "域名", TextAlign: Right, Padding: 2 },
		}

		body := [][]string{
			[]string{
				"AAAAAbbbcc", "正在启动", "http://5671.xxxxxxx.cn",
			},
			[]string{
				"BBB", "运行中", "http://",
			},
			[]string{
				"cccccc", "已停止", "http://asf3333.bbbbbbbb.cn",
			},
			[]string{
				"dddddddd", "重启中", "http://3cccccc.cccccccccccccc.cn",
			},
		}

		wordTable := NewWordTable(headers, body)

		wordTable.Println()

		t.Log("\n" + wordTable.String() + "\n")

	})
}