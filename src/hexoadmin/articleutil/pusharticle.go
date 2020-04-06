package articleutil

import (
	"encoding/json"
	"fmt"
	"gowww/hexoadmin/param"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestInfo struct {
	Url string
	Content string
	Data map[string]string //post要传输的数据，必须key value必须都是string
	DataInterface map[string]interface{}
}
type NewPage struct {
	Title       string        `json:"title"`
	Author      string        `json:"author"`
	Content     string        `json:"_content"`
	Source      string        `json:"source"`
	Raw         string        `json:"raw"`
	Slug        string        `json:"slug"`
	Published   bool          `json:"published"`
	Date        time.Time     `json:"date"`
	Updated     time.Time     `json:"updated"`
	Comments    bool          `json:"comments"`
	Layout      string        `json:"layout"`
	Photos      []interface{} `json:"photos"`
	Link        string        `json:"link"`
	ID          string        `json:"_id"`
	Path        string        `json:"path"`
	Permalink   string        `json:"permalink"`
	FullSource  string        `json:"full_source"`
	AssetDir    string        `json:"asset_dir"`
	Tags        []interface{} `json:"tags"`
	Categories  []interface{} `json:"categories"`
	IsDraft     bool          `json:"isDraft"`
	IsDiscarded bool          `json:"isDiscarded"`
}

type SavePage struct {
	Post struct {
		Title       string        `json:"title"`
		Author      string        `json:"author"`
		Date        time.Time     `json:"date"`
		Content     string        `json:"_content"`
		Source      string        `json:"source"`
		Raw         string        `json:"raw"`
		Slug        string        `json:"slug"`
		Published   bool          `json:"published"`
		Updated     time.Time     `json:"updated"`
		ID          string        `json:"_id"`
		Comments    bool          `json:"comments"`
		Layout      string        `json:"layout"`
		Photos      []interface{} `json:"photos"`
		Link        string        `json:"link"`
		Path        string        `json:"path"`
		Permalink   string        `json:"permalink"`
		FullSource  string        `json:"full_source"`
		AssetDir    string        `json:"asset_dir"`
		Tags        []interface{} `json:"tags"`
		Categories  []interface{} `json:"categories"`
		IsDraft     bool          `json:"isDraft"`
		IsDiscarded bool          `json:"isDiscarded"`
	} `json:"post"`
	TagsCategoriesAndMetadata struct {
		Categories struct {
			Ck8Nz5Rk7001Vjcdm9Prevx35 string `json:"ck8nz5rk7001vjcdm9prevx35"`
		} `json:"categories"`
		Tags struct {
			Ck8Nz5Rku001Wjcdm65N825Dl string `json:"ck8nz5rku001wjcdm65n825dl"`
			Ck8Nz5Rkw001Zjcdma49Idhtx string `json:"ck8nz5rkw001zjcdma49idhtx"`
		} `json:"tags"`
		Metadata []interface{} `json:"metadata"`
	} `json:"tagsCategoriesAndMetadata"`
}
//适用于 application/x-www-form-urlencoded
func (this RequestInfo) postUrlEncoded( )([]byte,error){
	client := &http.Client{}
	this.Content = strings.Replace(this.Content, "\n", "\\n", -1)
	fmt.Println("-----33302->",this.Content)

	req,err := http.NewRequest("POST",this.Url,strings.NewReader(this.Content))
	if err != nil{
		fmt.Println("----req,err := http.NewRequest(\"POST\",this.Url,strings.NewReader(this.Content))->:",err,this.Url)
		return nil,err
	}
	//伪装头部
	req.Header.Set("Accept","*/*")
	req.Header.Add("Accept-Encoding","gzip, deflate, br")
	req.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Connection","keep-alive")
	req.Header.Add("Content-Length","15")
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Cookie","M_distinctid=16f2b92bd2e723-0417a106c9c8c-5f302f45-ff000-16f2b92bd2f878; _ga=GA1.1.2097300613.1584883320; __cfduid=d656753c2d962575c9c82843c051a09451585574191; __gads=ID=333855f8a06f03f4:T=1585912296:S=ALNI_MbkQ7KUzZS7AxY8-tFUi7QNR1o8sg; sm_anonymous_id=2ea93a02-d78d-446d-acdb-22e8c62398df; CNZZDATA1278144647=636958000-1584883305-%7C1586144257; __atuvc=13%7C13%2C31%7C14%2C20%7C15; __atuvs=5e8ab1ca6ef49f37000; connect.sid=s%3A6UiNZWaBVrkxmCyqrnoPgP63CN1au66r.aq9b4JYUFnPbX59WbxYqPQSOOUZNP1Omfx4bdUAY6l0; Hm_lvt_6aa3f11e071bb2936fd0275e7fd3a182=1586137220,1586137235,1586137238,1586148192; Hm_lpvt_6aa3f11e071bb2936fd0275e7fd3a182=1586148192; _ga_M7TMLNPPFB=GS1.1.1586144262.25.1.1586148194.0")
	req.Header.Add("Host","zshipu.com")
	req.Header.Add("Origin","https://zshipu.com")
	req.Header.Add("Referer","https://zshipu.com/admintlg/")
	req.Header.Add("Sec-Fetch-Dest","empty")
	req.Header.Add("Sec-Fetch-Mode","cors")
	req.Header.Add("Sec-Fetch-Site","same-origin")
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4090.0 Safari/537.36 Edg/83.0.467.0")
	//提交请求
	resp,err := client.Do(req)
	defer resp.Body.Close()
	if err != nil{
		return nil,err
	}
	//读取返回值
	result,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	return result,nil
}




func pushPage(someOne NewPage) {
	bdpush := &RequestInfo{
		Url:     "https://zshipu.com/admintlg/api/posts/" + someOne.ID + "/publish",
		Content: "",
	}
	bdpush.postUrlEncoded()
}

// 可以通过修改底层url.QueryEscape代码获得更高的效率，很简单
func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}


func Foo(src string, dist string) {
	r := url.QueryEscape(src)
	r = strings.Replace(r, "+", "%20", -1)
	if r != dist {
		fmt.Printf("ensrc:%s\ngo:%s\njs:%s\n\n", src, r, dist)
	}

	r2,_ := url.QueryUnescape(dist)

	if r2 != src {
		fmt.Printf("desrc:%s\ngo:%s\njs:%s\n\n", src, r, dist)
	}
}

type ContentStruct struct {
	Content string `json:"_content"`
}
func saveArticle(someOne NewPage,content string) {

	con:=&ContentStruct{
		Content:content,
	}
	constr, err := json.Marshal(con);
	if err != nil {
		fmt.Println("content to string",err)
		return
	}
	content = string(constr)

	bdsave := &RequestInfo{
		Url:     "https://zshipu.com/admintlg/api/posts/" + someOne.ID,
		Content: content,
	}


	_, errsave := bdsave.postUrlEncoded()
	if errsave != nil{
		fmt.Println("resultsave, errsave := bdsave.postUrlEncoded() ",errsave)
	}


}

func createPage(title string) NewPage {
	bd := &RequestInfo{
		Url:     "https://zshipu.com/admintlg/api/posts/new",
		Content: "{\"title\":\""+title+"\"}",
	}
	result, err := bd.postUrlEncoded()
	if err != nil {
		fmt.Println(err)
	}
	var someOne NewPage
	if err := json.Unmarshal([]byte(string(result)), &someOne); err == nil {
		fmt.Println(someOne.ID)
	} else {
		fmt.Println(err)
	}
	return someOne
}

func PostArticle(param *param.ReqUrlParam) {

	////创建一篇文章
	someOne := createPage(param.Title)

	////保存文章
	saveArticle(someOne,param.Content)

	////发布文章
	pushPage(someOne)

}