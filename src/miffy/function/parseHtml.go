package function

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

//用户美化dom控制缩进的初始变量
var depth int

//ParseMyBlog 解析个人博客主页里面的所有超链接
func ParseMyBlog() {
	var links []string
	htmlContent := bytes.NewBufferString(getBlogContent())
	doc, err := html.Parse(htmlContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error: %v \n", err)
	}
	links = getLinksFromHTML(links, doc)
	for _, link := range links {
		fmt.Println(link)
	}
	// fmt.Printf("%v\n", doc.Type)
	// fmt.Printf("%v\n", doc.FirstChild.Type)
	// fmt.Println((doc.Type == html.DocumentNode))
}

//ParseStructOfMyBlog 解析个人博客的dom结构
func ParseStructOfMyBlog() {
	htmlContent := bytes.NewBufferString(getBlogContent())
	doc, err := html.Parse(htmlContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error: %v \n", err)
	}
	outline(nil, doc)
}

//BeautyStructOfMyBlog 美化输出博客的dom结构
func BeautyStructOfMyBlog() {
	htmlContent := bytes.NewBufferString(getBlogContent())
	doc, err := html.Parse(htmlContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error: %v \n", err)
	}
	//函数可以作为参数来传递
	forEachNode(doc, pre, post)

}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func pre(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func post(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth--
	}
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	//每次传入的stack都是上一次函数的返回的stack
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
func getLinksFromHTML(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = getLinksFromHTML(links, c)
	}
	return links
}

func getBlogContent() string {
	blogURL := "http://localhost:8080"
	resp, err := http.Get(blogURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch error: %v \n", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	//resp.Body是一个可读的服务器响应流
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reding %s error: %v \n", blogURL, err)
		os.Exit(1)
	}
	// fmt.Printf("%s", body)
	return string(body)
}
