package main

import "fmt"

const fal = "falsePATH"
func handle(str string) string {
	if len(str)<4 {		//正常路径至少4个字符，比如 C:\1
		return fal
	}
	var res string	//路径的结果
	var tmp string	//保存暂时的英文路径名
	var final string  //获取最后的文件名
	var n int // .的个数
	final=str[len(str)-4:]
	if final!=".txt" {
		return fal
	}
	str=str[:len(str)-4]
	if str[len(str)-1]=='\\'||str[len(str)-1]=='.' {
		return fal
	}
	str+="\\"	//保证路径最后为\
	//必须是类似C:\,如果符合，进入逻辑
	if str[0]<='Z'&&str[0]>='A'&&str[1]==':'&&str[2]=='\\' {
		var k int
		res=string([]byte(str[:3]))
		for i:=3;i< len(str);i++{	//以 \ 为一个周期
			if str[i]=='\\'{
				if str[i-1]=='\\'{	// C:\dd\\
					return fal
				}
				if n==2 {	// C:\dd\..\
					for k= len(res)-2;k>=0;k--{		//回溯到上层目录
						if res[k]=='\\'{
							break
						}
					}
					if k==0 {	//只有一种情况能进入，即 C:\..\
						return fal
					}
					res=res[:k+1]
				}else if n==0{
					res+=tmp
					res+="\\"
				}
				tmp=""
				n=0
			}else if str[i]=='.'{
				if (str[i-1]=='.'||str[i-1]=='\\')&&n<2{	//可能的情况 \ \.
					n++
				}else {
					return fal
				}
			}else  {	//路径名的字母
				if n==0{	//如果是字母，必须上一个\后面没有出现.，即全是字母
					tmp+=string(str[i])
				}else{
					return fal
				}
			}
		}
		res=res[:len(res)-1]	//去掉末尾的\
		res+=final
		return res
	}
	return "falsePATH"
}
func main()  {
	fmt.Println(handle(`C:\1\2\..\3\.\..\..txt`))
}