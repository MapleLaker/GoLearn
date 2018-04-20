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
