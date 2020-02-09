package face

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

//ByteCounter 统计字符串长度
type ByteCounter int

//统计字符串的长度
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c = ByteCounter(len(p))
	fmt.Printf("%d: %s", *c, p)
	return len(p), nil
}

var c ByteCounter

//Run 主方法
func Run() {
	fmt.Fprintf(os.Stdout, "Hello: %s\n", "duanchao")
	// fmt.Fprintf(c, "Hello: %s\n", "duanchao")
	c.Write([]byte("sixther"))
	fmt.Println(c)
	c = 0
	fmt.Fprintf(&c, "Hello: %s\n", "duanchao")

}

//Dc 接口
type Dc interface {
	Say(string)
}

//Duanchao Dc类型
type Duanchao struct{}

//Say 方法
func (d *Duanchao) Say(s string) {
	fmt.Printf("%s\n", s)
}

//Flag 方法
func Flag() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Sleeping for %v ...\n", *period)
	// time.Sleep(*period)
	// fmt.Println()
	var d Dc
	//使用new进行实例化操作
	d = new(Duanchao)
	d.Say("simon")
	fmt.Printf("%T\n", d)
	var w io.Writer
	//任何一个接口的动态类型都是nil
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	if w == nil {
		fmt.Println("ok")
		return
	}
	w.Write([]byte("duanchao\n"))
}

//SortByString 按照字符串排序
type SortByString []string

func (a SortByString) Len() int           { return len(a) }
func (a SortByString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByString) Less(i, j int) bool { return a[i] < a[j] }

//SortByTrack 按照*Track排序
type SortByTrack []*Track

func (a SortByTrack) Len() int           { return len(a) }
func (a SortByTrack) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByTrack) Less(i, j int) bool { return a[i].Artist < a[j].Artist }

//Track 音乐记录
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

//RunSort 方法
func RunSort() {
	// var names = []string{"b", "a"}
	var names = strings.Split("duanchao", "")
	sort.Sort(SortByString(names))
	fmt.Printf("%s\n", names)

	var tracks = []*Track{
		{"GO", "Hua", "One", 1990, lengths("3m23s")},
		{"HO", "Pua", "One", 1995, lengths("3m23s")},
		{"LO", "Cua", "One", 1930, lengths("3m23s")},
		{"BO", "Sua", "One", 2990, lengths("3m23s")},
	}
	sort.Sort(SortByTrack(tracks))
	printTracks(tracks)
}

func lengths(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "---", "---", "---", "---", "---")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
