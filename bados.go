package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var running = true

func OnReady()  {
	fmt.Println("Executed")

	//icon, _ := ioutil.ReadFile("seed.png")

	//systray.SetIcon(icon)
	systray.SetTitle("AntiWPC")
	systray.SetTooltip("AntiWPC")
	close := systray.AddMenuItem("Quit", "Close the application")

	go func() {
		<-close.ClickedCh
		fmt.Println("Closing")
		systray.Quit()
	}()

	go func() {
		for running {
			WpcMonScan()
			WpcMonSvcScan()
			time.Sleep(time.Second * 1)
		}
	}()
}

func OnExit()  {
	running = false
}

func WpcMonScan()  {
	WPCMonRegex := regexp.MustCompile(`WpcMon\.exe( +)(.*) ([CS])`)

	out, err := exec.Command("tasklist").Output()
	if err != nil {
		fmt.Println("Error retrieving tasklist" + fmt.Sprintf("%s",err))
	}

	if WPCMonRegex.MatchString(string(out)) {
		PID := WPCMonRegex.FindStringSubmatch(string(out))[2]
		out2, err2 := exec.Command("taskkill","/T", "/F", "/PID", PID).Output()
		if err2 != nil {
			fmt.Println("Error killing WpcMon: " + fmt.Sprintf("%s",err2))
		}
		fmt.Println(string(out2))
	}
}

func WpcMonSvcScan()  {
	PIDRegex := regexp.MustCompile(`PID( +: )(.*)\n`)

	out3, err3 := exec.Command("C:\\Windows\\system32\\sc.exe", "queryex", "WpcMonSvc").Output()
	if err3 != nil {
		fmt.Println("Error querying WpcMonSvc" + fmt.Sprintf("%s",err3))
	}

	if PIDRegex.MatchString(string(out3)) {
		PID := PIDRegex.FindStringSubmatch(string(out3))[2]
		if PID != "0\r" {
			PID = strings.ReplaceAll(PID,"\r","")
			out4, err4 := exec.Command("taskkill", "/T", "/F", "/PID", PID).Output()
			if err4 != nil {
				fmt.Println("Error killing WpcMonSvc: " + fmt.Sprintf("%s",err4))
			}
			fmt.Println(string(out4))
		}
	}
}

func main() {
	time.Sleep(time.Second * 3)
	systray.Run(OnReady, OnExit)
}