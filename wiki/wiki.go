//https://go.dev/doc/articles/wiki/ 来自此官方教程
package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

//正则验证
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // 标题是第二个子表达式。
}

//定义标题与正文内容
type Page struct {
	Title string
	Body  []byte
}

//定义页面的保存功能
func (p *Page) save() error {
	fileName := "data/" + p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

//加载页面
func loadPage(title string) (*Page, error) {
	fileName := "data/" + title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

//模板处理
//func renderTemplate(w http.ResponseWriter, tmp string, p *Page) {
//	t, _ := template.ParseFiles(tmp)
//	t.Execute(w, p)
//
//	//err := templates.ExecuteTemplate(w, tmp, p)
//	//if err != nil {
//	//	http.Error(w, err.Error(), http.StatusInternalServerError)
//	//	return
//	//}
//}

//模板缓存
var templates = template.Must(template.ParseFiles("template/edit.html", "template/view.html"))

//模板处理
//这里注意 使用缓存后 tmp只需输入文件名 而不是路径
func renderTemplate(w http.ResponseWriter, tmp string, p *Page) {
	err := templates.ExecuteTemplate(w, tmp, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//查看页面
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p, err := loadPage(title)

	//如果没有该页面则编辑该页面
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view.html", p)
}

//编辑页面
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p, err := loadPage(title)

	//该页面不存在的情况
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit.html", p)
}

//保存页面的处理
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body := r.FormValue("body")
	p := &Page{title, []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+p.Title, http.StatusFound)
}

//引入闭包处理

func main() {
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
