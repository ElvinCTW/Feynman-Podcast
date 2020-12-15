package crawler

import (
	"net/http"
)

var (
	c = NewClient(&http.Client{Timeout: 5})
)

//
//func Test_GetCivilFromGov(t *testing.T) {
//	t.Fail()
//	if err := c.GetCivilFromGov("https://wwwq.moex.gov.tw/exam/wHandExamQandA_File.ashx?t=Q&code=109090&c=301&s=0101&q=1"); err != nil {
//		fmt.Println(err)
//	}
//}

//func Test_ConvertPDFToDOCX(t *testing.T) {
//	t.Fail()
//	if err := c.ConvertPDFToDOCX("./civil.pdf"); err != nil {
//		fmt.Println(err)
//	}
//}

//func Test_DownloadPDF(t *testing.T) {
//	t.Fail()
//	if err := c.DownloadPDF("https://wwwq.moex.gov.tw/exam/wHandExamQandA_File.ashx?t=Q&code=109090&c=301&s=0101&q=1"); err != nil {
//		fmt.Println(err)
//	}
//}
