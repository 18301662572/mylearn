package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//验证身份证号后四位,只获取男性
//http://mid.weixingmap.com/
func main() {
	str := "11010819871023" //北京海淀区
	//str := "32090019840816" //江苏盐城
	result := getRand(str)
	//for _, v := range result {
	//	fmt.Println(":" + v)
	//}
	for index, v := range result {
		fmt.Printf("%d:"+v+"\n", index)
	}
}

//获取身份证号后四位 验证不重复
func getRand(codeid string) (reslutArr []string) {
	var int_Arr []string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(codeid)
		r_num := r.Intn(10000)
		var isExit bool = false
		result := strconv.Itoa(r_num)
		res_len := len(result)
		if res_len <= 1 {
			for _, v := range int_Arr {
				if len(v) == 1 {
					isExit = true
					break
				}
			}
			//十位数用0补足，所以肯定为偶数，为女,筛除女性
			if 0%2 == 0 {
				isExit = true
			}
			//if 0%2 > 0 { //为男,筛除男性
			//	isExit = true
			//}
		} else {
			for _, v := range int_Arr {
				if len(v) >= 1 {
					if strings.EqualFold(v[0:len(v)-1], result[0:res_len-1]) {
						isExit = true
						break
					}
				}
			}
			if r_len, err := strconv.Atoi(result[:res_len-1]); err == nil {
				if r_len%2 == 0 { //肯定为偶数,为女,筛除女性
					isExit = true
				}
				//if r_len%2 > 0 { //为男,筛除男性
				//	isExit = true
				//}
			}
		}
		if !isExit {
			int_Arr = append(int_Arr, strconv.Itoa(r_num))
			if len(result) < 4 {
				cl_num := 4 - len(result)
				num := ""
				for j := 0; j < cl_num; j++ {
					num = num + "0"
				}
				buffer.WriteString(num)
			}
			buffer.WriteString(result)
			reslutArr = append(reslutArr, buffer.String())
		}
	}
	return
}
