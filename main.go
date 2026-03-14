package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	_ "github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	_ "github.com/yuin/goldmark/parser"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type BlogPageData struct {
	BlogTitle string
	Date      string
	Content   template.HTML
}

type Blog struct {
	Title    string
	Date     string
	Filename string
}

type BlogList struct {
	Blogs []Blog
}

var BlogsList BlogList

func compBlogs(i int, j int) bool {
	t1, err1 := time.Parse("02 January, 2006", BlogsList.Blogs[i].Date)

	t2, err2 := time.Parse("02 January, 2006", BlogsList.Blogs[j].Date)
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	return t2.Before(t1)
}

func main() {
	blogs, err := os.ReadDir("./blogs/")
	check(err)
	for _, blog := range blogs {
		generateBlog(filepath.Join("./blogs/", blog.Name()))
		BlogsList.Blogs[len(BlogsList.Blogs)-1].Filename = strings.Join([]string{strings.TrimSuffix(blog.Name(), filepath.Ext(blog.Name())), ".html"}, "")
	}
	sort.Slice(BlogsList.Blogs, compBlogs)
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	f, err := os.Create("serve/index.html")
	check(err)
	tmpl.Execute(f, BlogsList)
}

func generateBlog(blogDest string) {
	tmpl := template.Must(template.ParseFiles("./templates/blog.html"))
	_, file := filepath.Split(blogDest)
	file = strings.TrimSuffix(file, filepath.Ext(blogDest))
	file = fmt.Sprintf("serve/blogs/%s.html", file)
	f, err := os.Create(file)
	check(err)
	tmpl.Execute(f, parseBlogFile(blogDest))
}

func parseBlogFile(path string) BlogPageData {
	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	out, err := os.ReadFile(path)
	check(err)
	ctx := parser.NewContext()
	//fmt.Println(string(out))
	var buf bytes.Buffer
	err = md.Convert(out, &buf, parser.WithContext(ctx))
	metaData := meta.Get(ctx)
	//dat := map[string]string(metaData)
	var blogData BlogPageData = BlogPageData{
		BlogTitle: metaData["Title"].(string),
		Date:      metaData["Date"].(string),
		Content:   template.HTML(buf.String()),
	}
	BlogsList.Blogs = append(BlogsList.Blogs, Blog{Title: metaData["Title"].(string), Date: metaData["Date"].(string)})
	return blogData
}
