package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var p = fmt.Println

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	p(true)
	bolB, _ := json.Marshal(true)
	p(string(bolB))

	p(1)
	intB, _ := json.Marshal(1)
	p(string(intB))

	p(2.34)
	fltB, _ := json.Marshal(2.34)
	p(string(fltB))

	p("gopher")
	strB, _ := json.Marshal("gopher")
	p(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	p(slcD)
	slcB, _ := json.Marshal(slcD)
	p(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	p(mapD)
	mapB, _ := json.Marshal(mapD)
	p(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	p(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	p(string(res2B))

	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p(dat)

	//value是interface{},通过a.(type)的形式进行类型断言，返回value和err
	num := dat["num"].(float64)
	p(num)

	strs := dat["strs"].([]interface{})
	p(strs)
	fmt.Printf("%T\n", strs)
	str1 := strs[0].(string)
	p(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	p(res)
	p(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
