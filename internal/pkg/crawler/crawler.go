package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"code.sajari.com/docconv"
)

type Client struct {
	httpClient *http.Client
}

var (
	once   sync.Once
	client *Client
)

func NewClient(httpClient *http.Client) *Client {
	once.Do(func() {
		client = &Client{
			httpClient: httpClient,
		}
	})

	return client
}

func (c *Client) downloadPDF(uri string) error {
	resp, err := c.httpClient.Get(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if b, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else if err = ioutil.WriteFile("civil.pdf", b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (c *Client) ConvertPDF() error {
	return c.convertPDF()
}

func (c *Client) convertPDF() error {
	f, err := os.Open("./civil.pdf")
	if err != nil {
		return err
	}

	defer f.Close()

	// b := make([]byte, 1000)
	// // ch := make(chan []byte)
	// var n int
	// var e error
	// for e == nil {
	// 	n, e = f.Read(b)
	// 	fmt.Print(string(b))
	// }

	// if n != 0 && e == io.EOF {
	// 	n, e = f.Read(b)
	// 	fmt.Println(string(b))
	// 	return nil
	// } else if n == 0 && e == io.EOF {
	// 	fmt.Println(len(b))
	// 	fmt.Println(string(b))
	// 	return nil
	// } else {
	// 	panic(e)
	// }

	s, m, err := docconv.ConvertPDF(f)
	if err != nil {
		return err
	}
	fmt.Println(m)
	fmt.Println(s)

	return nil
}
