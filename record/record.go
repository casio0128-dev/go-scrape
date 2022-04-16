package record

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sclevine/agouti"
	"go-scrape/browser"
	"net/http"
	"os"
	"strings"
	"time"
)

const url = "http://localhost:1323/rec"

func Run() {
	d := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
		//browser.IsHeadless(),
	}))

	if err := d.Start(); err != nil {
		panic(err)
	}
	defer func() {
		if err := d.Stop(); err != nil {
			panic(err)
		}
	}()

	page, err := d.NewPage(agouti.Browser("chrome"))
	if err != nil {
		panic(err)
	}

	beforeURL, err := page.URL()
	if err != nil {
		panic(err)
	}

	for {
		if d, err := time.ParseDuration("1s"); err == nil {
			<-time.Tick(d)
		} else {
			panic(err)
		}

		url, err := page.URL()
		if err != nil {
			panic(err)
		}

		if !strings.EqualFold(url, beforeURL) {
			fmt.Println(url, beforeURL)
			beforeURL = url
			if err := RunScripts(page); err != nil {
				panic(err)
			}
		}
	}
}

func RunScripts(page *agouti.Page) error {
	if err := page.RunScript(browser.MakeScript(browser.GetXPath(), browser.Post(), browser.SetEventListener(url)), nil, nil); err != nil {
		return err
	}
	return nil
}

type operationRequest struct {
	Action      string `json:"action"`
	Content     string `json:"content"`
	Target      string `json:"target"`
	CurrentHref string `json:"currentHref"`

	createdAt time.Time
}

type Recorder struct {
	Requests *[]operationRequest
}

func (r *Recorder) PushRequest(or *operationRequest) {
	or.createdAt = time.Now()
	*r.Requests = append(*r.Requests, *or)
}

func NewRecorder() *Recorder {
	return &Recorder{
		Requests: new([]operationRequest),
	}
}

func (r *Recorder) indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "index")
}

func (r *Recorder) recHandler(c echo.Context) error {
	or := new(operationRequest)
	if err := c.Bind(or); err != nil {
		return err
	}

	fmt.Println("act", or.Action)
	fmt.Println("xpath", or.Target)
	fmt.Println("content", or.Content)
	fmt.Println("currentHref", or.CurrentHref)

	r.PushRequest(or)

	return c.JSON(http.StatusOK, r.Requests)
}

func Recording() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	rec := NewRecorder()
	e.GET("/", rec.indexHandler)
	e.POST("/rec", rec.recHandler)

	defer func() {
		echoLog := e.Logger.Output()

		f, e := os.Create(time.Now().Format("server_20060102150405.log"))
	}

	if err := e.Start(":1323"); err != nil {
		panic(err)
	}
}
