package main
import (
	
)

func main() {
    tr := &http.Transport{
        MaxIdleConns:    100,
        IdleConnTimeout: 3 * time.Second,
    }

    n := 5
    for i := 0; i < n; i++ {
        req, _ := http.NewRequest("POST", "https://www.baidu.com", nil)
        req.Header.Add("content-type", "application/json")
        client := &http.Client{
            Transport: tr,
            Timeout:   3 * time.Second,
        }
        resp, _ := client.Do(req)
        _, _ = ioutil.ReadAll(resp.Body)
        _ = resp.Body.Close()
    }
    time.Sleep(time.Second * 5)
    fmt.Printf("goroutine num is %d\n", runtime.NumGoroutine())
}