package email

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//爬邮箱
//爬链接
//怕图片...

var (
	emailReg=`\w+@\w+\.\w+`
	//s?代表又或者没有s
	//+代表出1次或多次
	//\s\S代表各种字符
	//+?代表贪婪模式,到“停下
	linkReg=`href="(https?://[\s\S]+?)"`
	phoneReg=`1[3456789]\d\s?\d{4}\s?\d{4}`
	imgReg=`https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))"`
)

//错误处理
func HandleErr(err error,why string){
	if err!=nil{
		fmt.Println(err,why)
	}
}

//抽取 根据url获取内容
func GetPageStr(url string)(pageStr string){
	//1.确定目标
	resp, err := http.Get(url)
	defer resp.Body.Close()
	HandleErr(err,"http.Get url")
	//2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleErr(err,"ioutil.ReadAll")
	//字节转字符串
	pageStr=string(pageBytes)
	return
}

//爬邮箱
func GetUrl(url string){
	//1.确定目标
	pageStr := GetPageStr(url)
	//3.正则筛选想要的内容
	re := regexp.MustCompile(emailReg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	//4.获取到的内容进行处理
	for  _,result :=range results {
		fmt.Println(result)
	}
}

//爬链接
func GetLink(url string){
	//1.确定目标
	pageStr := GetPageStr(url)
	//3.正则筛选想要的内容
	re := regexp.MustCompile(linkReg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	//4.获取到的内容进行处理
	for  _,result :=range results {
		fmt.Println(result[1])
	}
}

func GetPhone(url string){
	//1.确定目标
	pageStr := GetPageStr(url)
	//3.正则筛选想要的内容
	re := regexp.MustCompile(phoneReg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	//4.获取到的内容进行处理
	for  _,result :=range results {
		fmt.Println(result)
	}
}

func GetImg(url string){
	//1.确定目标
	pageStr := GetPageStr(url)
	//3.正则筛选想要的内容
	re := regexp.MustCompile(imgReg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	//4.获取到的内容进行处理
	for  _,result :=range results {
		fmt.Println(result[0])
	}
}

func main(){
	//爬邮箱
	//GetUrl("https://tieba.baidu.com/f?kw=%D3%CA%CF%E4&fr=ala0&tpl=5&traceid=")

	//爬链接
	//GetLink("https://tieba.baidu.com/p/6573828404?fr=ala0&pstaala=2&tpl=5&fid=4628&red_tag=1501655536")

	//爬手机号
	//GetPhone("https://www.zhaohaowang.cn/")

	//爬图片
	GetImg("https://www.umei.cc/p/gaoqing/rihan/1.htm")
}

