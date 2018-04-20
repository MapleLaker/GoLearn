range 遍历容器。
第一个参数返回index, 第二个参数返回value

eg: 对于map, index 就是 key, value 就是map[index]
	var country map[string]string
	country = make(map[string]string)

	country["Franch"] = "Paris"
	country["US"] = "Washton"
	country["Japan"] = "Tok"

	for c,d := range(country){
		fmt.Printf("The captical of %s is %s\n", c, d)
	}
	
eg:   对于数组，index 就是序号，value 就是数值
	var numbers = [6]int{1,2,3,4,5,6}
	for a,b:= range numbers{
		fmt.Printf("The %d is %d\n",a, b)
	}
