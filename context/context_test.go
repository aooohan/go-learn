package context

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dst)
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// 处理完业务之后 必须手动调用func
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func TestDeadline(t *testing.T) {
	ctx, cancel := context.WithDeadline(
		context.Background(),
		time.Now().Add(50*time.Millisecond),
	)
	defer cancel()
	select {
	case <-time.After(1 * time.Millisecond):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting...")
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("over")
}

func TestServer(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := rand.Intn(2)
		if n == 0 {
			time.Sleep(10 * time.Second)
			fmt.Fprintf(w, "slow response")
			return
		}
		fmt.Fprintf(w, "quick response")
	})

	http.ListenAndServe(":8000", nil)
}

type respData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *respData, 1)
	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000/", nil)
	if err != nil {
		fmt.Printf("new request failed!%v\n", err)
		return
	}
	//request = request.WithContext(ctx)
	go func() {
		resp, err := client.Do(request)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}
func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	doCall(ctx)
}
