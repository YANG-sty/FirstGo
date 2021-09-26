package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	var ascii = string("我喜欢java")
	fmt.Println([]byte(ascii)) //byte类型 输出的是字节
	fmt.Println([]rune(ascii)) //rune类型 输出的是字符

	fmt.Println(len(ascii))
	fmt.Println(len([]rune(ascii)))            //输出的是字符的长度
	fmt.Println(utf8.RuneCountInString(ascii)) //输出字符的长度

	fmt.Println(string([]byte(ascii))) //byte类型的数据转为字符串输出
	fmt.Println(string([]rune(ascii))) //rune类型的数据转为字符串输出

	fmt.Println("-----------------char int float bool--------------------")

	fmt.Println("----------char int---------")
	fmt.Println(strconv.Itoa('a'))  //char 转 字符串
	atoi, err := strconv.Atoi("97") //字符串 转 char
	fmt.Println(atoi, err)          //字符串 转 int

	/**
	参数1：要转换的字符串
	参数2：转换位int后的进制数，2进制，8进制，10进制， 16进制 等
	参数3：转换的int后，该int值所占用的位数， 8位<=127, 16位<=32767 等
	*/
	parseInt, err := strconv.ParseInt("127", 10, 8)        // string 转 int
	parseInt16, err16 := strconv.ParseInt("32767", 10, 16) // string 转 int
	fmt.Println(parseInt, err)
	fmt.Println(parseInt16, err16)

	fmt.Println("----------float---------")
	/**
	参数1：转换的浮点类型数据
	参数2：f 表示无指数
	参数3：小数位精确的位数
	参数4：转换的浮点类型数据的大小，32位，64位
	*/
	fmt.Println(strconv.FormatFloat(3.1415927, 'f', 10, 64)) //float 转 string
	float, err := strconv.ParseFloat("3.1415927000", 64)     // string 转 float
	fmt.Println(float, err)

	fmt.Println("----------bool---------")

	fmt.Println(strconv.FormatBool(true)) // bool 转 string
	parseBool, err := strconv.ParseBool("false")
	fmt.Println(parseBool, err) // string 转 bool

	fmt.Println("-----------------stringUtils--------------------")
	fmt.Println("------------字符串比较----------")
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(strings.Compare("1", "1"))
	fmt.Println(strings.Compare("3", "2"))

	fmt.Println("------------包含----------")

	fmt.Println(strings.Contains("我喜欢java", "java"))  //是否包含 指定字符串的内容
	fmt.Println(strings.ContainsAny("我喜欢java", "av")) //是否包含 任意 char类型字符串
	fmt.Println(strings.ContainsRune("我喜欢java", 'a')) //是否包含rune类型的字符
	fmt.Println(strings.Count("我喜欢java", "a"))        //计算字符串出现的紫薯

	fmt.Println("------------相等----------")

	fmt.Println(strings.EqualFold("我喜欢java", "Java")) //两字符串是否相同，不区分大小写

	fmt.Println("-----------空白符----------")
	//空白符 \n换行 \tTab符号 \r回车 \f换页
	fmt.Printf("%#v\n", strings.Fields("a b c d\ne f\tg h\ri J\fk "))

	fmt.Println("------------出现位置----------")
	//开头 结尾
	fmt.Println(strings.HasPrefix("我喜欢java", "我")) //是否以 该字符串 开头
	fmt.Println(strings.HasSuffix("我喜欢java", "a")) //是否以 该字符串 结尾

	//字符串出现的位置，索引的下标从0开始，该函数对中文不友好，是按照 字节 进行匹配
	fmt.Println(strings.Index("I love java", "love")) // 元素最初出现的位置， 正常返回
	fmt.Println(strings.Index("我喜欢java", "喜欢"))       //返回异常，应该返回1，结果返回的是3
	fmt.Println(strings.Index("我喜欢java", "我"))
	fmt.Println(strings.LastIndex("I love java", "a")) //元素最后出现的位置

	fmt.Println("------------连接----------")
	//连接
	join := strings.Join([]string{"1", "2", "3", "4", "5", "6", "7"}, "_")
	fmt.Println(join) //字符串通过，_ 进行连接

	fmt.Println("------------分割----------")
	//分割
	fmt.Println(strings.Split(join, "_"))     //字符串所有的关键字进行分割
	fmt.Println(strings.SplitN(join, "_", 3)) //将元素分割成3个

	fmt.Println("------------分割----------")
	fmt.Println(strings.Repeat("&", 6)) //指定字符，使其连续重复出现

	fmt.Println("------------替换----------")
	fmt.Println(strings.Replace("I love java, java is very good", "java", "go", -1))              //全部替换
	fmt.Println(strings.ReplaceAll("I love java, java is very good", "java", "go"))               //全部替换
	fmt.Println(strings.Replace("I love java, java is very good, java is life", "java", "go", 2)) //替换前两个

	fmt.Println("------------首字母大写----------")
	fmt.Println(strings.Title("i love java"))                  //首字母大写
	fmt.Println(strings.ToLower("I Love JAVA"))                //全部小写
	fmt.Println(strings.ToUpper("i love java, javaisvergood")) //全部大写

	fmt.Println("------------trim-------------")
	fmt.Println(strings.Trim("abcdefabc", "abdc")) //ef
	fmt.Println(strings.Trim("abcdefabc", "ac"))   //bcdefab
	fmt.Println(strings.Trim("abcdefabc", "ab"))   //cdefabc
	fmt.Println(strings.Trim("abcdefabc", "a"))    //bcdefabc
	fmt.Println("---------------------")
	fmt.Println(strings.TrimLeft("ababeabcd", "abd")) //eabcd //左边字符串出现在子字符中则替换
	fmt.Println(strings.TrimRight("ababcdb", "abd"))  //ababc //右边字符串出现在子字符中则替换

	fmt.Println("---------------------")
	fmt.Println(strings.TrimPrefix("ababeabcd", "abd")) //ababeabcd //字符串当成一个整体替换
	fmt.Println(strings.TrimSuffix("ababeabcd", "abd")) //ababeabcd //字符串当成一个整体替换
	fmt.Println(strings.TrimPrefix("ababeabcd", "ab"))  //abeabcd //字符串当成一个整体替换
	fmt.Println(strings.TrimSuffix("ababeabcd", "cd"))  //ababeab //字符串当成一个整体替换

	fmt.Println("---------------------")
	fmt.Println(strings.TrimSpace("ababcd ")) //ababcd //去除字符串两边的空格

}
