package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/fabulous-tech/go-redmine"
)

func main() {
	var (
		endpoint string
		apikey   string
	)
	flag.StringVar(&endpoint, "e", "", "Redmine Endpoint URL")
	flag.StringVar(&apikey, "k", "", "Redmine API Key")

	flag.Parse()
	c := redmine.NewClient(endpoint, apikey)
	pj, err := c.GetProjects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(pj) == 0 {
		fmt.Println("This Redmine Account is No Assign Projects.")
		os.Exit(1)
	}

	fmt.Println("プロジェクトIDを入力してください。")
	for _, v := range pj {
		fmt.Printf("%v : %v\n", v.ID, v.Name)
	}

	var pid int
	fmt.Scan(&pid)

	issues, err := c.GetIssuesByFilter(&redmine.IssueFilter{ProjectId: strconv.Itoa(pid)})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var iid int = 0
	if len(issues) != 0 {
		fmt.Println("IssueIDを入力してください。")
		for _, v := range issues {
			fmt.Printf("%v : %v\n", v.Id, v.Subject)
		}
		fmt.Scan(&iid)
	}

	var at int
	activities, err := c.GetTimeEntryActivities()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("作業分類IDを入力してください。")
	for _, v := range activities {
		fmt.Printf("%v : %v\n", v.Id, v.Name)
	}
	fmt.Scan(&at)

	var s string
	fmt.Println("作業日を入力してください。（format: yyyy-mm-dd） \n")
	fmt.Scan(&s)

	var h float32
	fmt.Println("作業時間を数値で入力してください\n")
	fmt.Scan(&h)

	req := &redmine.TimeEntryRequest{
		ProjectId:  pid,
		IssueId:    iid,
		SpentOn:    s,
		Hours:      h,
		ActivityId: at,
		Comments:   "",
	}
	fmt.Println(req)
	_, err = c.CreateTimeEntry(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
