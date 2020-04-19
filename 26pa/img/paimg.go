package main

import (
	"code.oldbody.com/studygolang/mylearn/26pa/email"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//并发爬图片

var (
	imgReg=`https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))"`
)


//并发爬思路：
//1.初始化数据管道
//2.爬虫协程：26个协程向管道中添加图片链接
//3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
//4.下载协程：从管道读取链接并下载

var (
	//存放图片链接的数据管道
	chanImgUrls chan string
	waitGroup sync.WaitGroup
	//用于监控协程
	chanTask chan string
)



//任务统计携程
func CheckOK(){
	defer  waitGroup.Done()
	var count int
	for{
		url:=<-chanTask
		fmt.Printf("%s完成了爬取任务\n",url)
		count++
		if count==7{
			close(chanImgUrls)
			break
		}
	}
}


//爬图片链接到管道
//url 是传的整页链接
func getImgUrls(url string){
	defer waitGroup.Done()
	urls:=getImgs(url)
	//遍历切片里的所有链接，存入数据管道
	for _, value := range urls {
		chanImgUrls<-value
	}
	//标识当前协程完成
	//每完成一个任务，写一条数据，用于监控携程知道已经完成了几个任务
	chanTask<-url
}

//获取当前页的图片链接
func getImgs(url string)(urls[]string){
	//1.确定目标
	//2.获取内容
	pageStr := email.GetPageStr(url)
	//3.正则筛选想要的内容
	re := regexp.MustCompile(imgReg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("找到%d条结果\n",len(results))
	//4.获取到的内容进行处理
	for  _,result :=range results {
		//fmt.Println(result[0])
		urls=append(urls,result[0])
	}
	return
}

//下载图片
func DownLoadFile(url string,filename string)(ok bool){
	resp, err := http.Get(url)
	email.HandleErr(err,"http.get utl")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	email.HandleErr(err,"resp.Body")
	filename="../images/"+filename
	fmt.Println("aaaa:"+filename)
	//写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err!=nil{
		ok=false
	}
	ok=true
	return
}


func DownLoadImg(){
	defer waitGroup.Done()
	for url:= range chanImgUrls{
		filename := getFilenameFormURL(url)
		if ok := DownLoadFile(url, filename);ok{
			fmt.Printf("%s下载成功\n",url)
		}else {
			fmt.Printf("%s下载失败\n",url)
		}
	}
}

//截取URL名字
func getFilenameFormURL(url string)(filename string){
	//返回最后一个/的位置
	lastindex := strings.LastIndex(url, "/")
	filename=url[lastindex+1:]
	//时间戳解决重名
	timePrefix:=strconv.Itoa(int(time.Now().Unix()))
	filename=timePrefix+"_"+filename
	return
}


func main(){
	//1.初始化管道
	chanImgUrls=make(chan string,100000)
	chanTask=make(chan string,7)
	//2.爬虫协程
	for i:=1;i<8 ;i++  {
		waitGroup.Add(1)
		go getImgUrls(fmt.Sprintf("https://www.umei.cc/p/gaoqing/rihan/%d.htm",i))
	}
	//3.任务统计携程，统计7个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()

	//4.下载携程，从管道中读取链接并下载
	for i:=0;i<5 ;i++  {
		waitGroup.Add(1)
		go DownLoadImg()
	}
	waitGroup.Wait()

}


