package main

// func mymain() {
// 	client := new(http.Client)
// 	req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/test", bytes.NewBuffer([]byte("我是客户端"))) // post请求体 reader类型
// 	// 写个请求头
// 	req.Header["test"] = []string{"test1"}
// 	// 执行客户端请求
// 	res, _ := client.Do(req)
// 	// 读取reder,转为字节
// 	b, _ := io.ReadAll(res.Body)
// 	fmt.Println(string(b))
// }
