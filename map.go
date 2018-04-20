package main

import "fmt"

func main() {
  //声明map变量
	var country map[string]string
  //定义 map 变量
	country = make(map[string]string)

	country["Franch"] = "Paris"
	country["US"] = "Washton"
	country["Japan"] = "Tok"
  
  // Range 返回map 类型的 key
	for c:= range(country){
		fmt.Printf("The captical of %s is %s\n", c, country[c])
	}

  // map[]第二个参数，返回的是 是否 存在
	i, ok := country["Korean"]
	if ok {
		fmt.Printf("The captical of %s is %s\n", i, country[i])
	}else
	{
		fmt.Printf("The captical is not found")
	}
}

