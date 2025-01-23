package main

import (
	// "flag"
	// "fmt"
	"fmt"
	"net/http"

	// "os"
	// "strings"
	// "sync"
	// "time"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS, 
		middleware.Logging,
	)

	server := http.Server{
		Addr: ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server listening")
	server.ListenAndServe()
}








// самый простой пример создания сервера

// func hello(w http.ResponseWriter, req *http.Request)  {
// 	fmt.Println("Hello")
// }

// func main() {
// 	router := http.NewServeMux()
// 	router.HandleFunc("/hello", hello)

// 	server := http.Server{
// 		Addr: ":8081",
// 		Handler: router,
// 	}

// 	server.ListenAndServe()
// }


// проверка файла на правильность запросов

// func ping(url string, respCh chan int, errCh chan error)  {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		errCh <- err
// 		return
// 	}
// 	respCh <- resp.StatusCode
// }

// func main() {
// 	path := flag.String("file", "url.txt", "path to URL file")
// 	flag.Parse()
// 	file, err := os.ReadFile(*path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	urlSlice := strings.Split(string(file), "\n")
// 	respCh := make(chan int)
// 	errCh := make(chan error)
// 	for _, url := range urlSlice {
// 		go ping(url, respCh, errCh)
// 	}
// 	for range urlSlice {
// 		select {
// 		case err := <-errCh:
// 			fmt.Println(err)
// 		case res := <-respCh:
// 			fmt.Println(res)
// 		}
// 	}
// }



// сумма элементов в массиву

// func main() {
// 	ch := make(chan int, 3)
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

// 	for i := 0; i < 3; i++ {
// 		x1 := 4*i
// 		x2 := 4 + x1
// 		slice := arr[x1:x2]
// 		go sum(ch, slice)
// 		// go func(slice []int) {
// 		// 	sum(ch, slice)
// 		// 	wg.Done()
// 		// }(slice)
// 	}

// 	total := 0
// 	// for sum := range ch {
// 	// 	total += sum
// 	// }
// 	for i := 0; i < 3; i++ {
// 		total += <-ch
// 	}

// 	fmt.Println(total)

// }

// func sum(ch chan int, arr []int)  {
// 	summa := 0
// 	for _, el := range arr {
// 		summa += el
// 	}
// 	ch <- summa
// }
		


// func main() {
// 	code := make (chan int)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func () {
// 			getHttpCode(code)
// 			wg.Done()
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(code)
// 	}()

// 	for res := range code {
// 		fmt.Printf("Code: %d", res)
// 	}
// }

// func getHttpCode(codeCh chan int)  {
// 	resp, err := http.Get("https://google.com")

// 	if err != nil {
// 		fmt.Printf("Error %s", err.Error())
// 	}
// 	codeCh <- resp.StatusCode
// }


// func main() {
// 	t:= time.Now()
// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func()  {
// 			getHttpCode()
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println(time.Since(t))
// }

// func getHttpCode()  {
// 	resp, err := http.Get("https://google.com")

// 	if err != nil {
// 		fmt.Printf("Error %s", err.Error())
// 	}
// 	fmt.Printf("Code: %d", resp.StatusCode)
// }