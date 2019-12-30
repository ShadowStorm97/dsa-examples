package main

import "fmt"

func main() {
	str1 := "xjybz"
	str2 := "djebru"
	ret := getLCS(str1,str2)
	fmt.Println(ret)
}


func getLCS(str1,str2 string) (ret string){
	ret = LCS([]rune(str1),[]rune(str2),len([]rune(str1)),len([]rune(str2)))
	return
}

func LCS(str1,str2 []rune,len1,len2 int) (ret string) {

	//递归基
	if 0 == len1 || 0 == len2 {
		return
	}

	//如果末尾能匹配上
	if str1[len1-1] == str2[len2-1] {
		//分解为n-1的子问题
		temp1 := str1[0:len1-1]
		temp2 := str2[0:len2-1]
		ret = LCS(temp1,temp2,len1-1,len2-1) + string(str1[len1-1])
	}else{
		//如果末尾不能匹配上,分解为两个不同的递归方向

		//第一个方向为 str1中最后一位将不能再提供任何贡献 移除掉
		ret1 := LCS(str1[0:len1-1],str2,len1-1,len2)

		//第二个方向为 str2中最后一位将不能再提供任何贡献 移除掉
		ret2 := LCS(str1,str2[0:len2-1],len1,len2-1)

		//取两个方向中 最长的那个
		if len(ret1) > len(ret2){
			ret = ret1
		}else{
			ret = ret2
		}
	}
	return
}

