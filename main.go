package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type Result struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Ping      struct {
		Jitter  float64 `json:"jitter"`
		Latency float64 `json:"latency"`
		Low     float64 `json:"low"`
		High    float64 `json:"high"`
	} `json:"ping"`
	Download struct {
		Bandwidth int64 `json:"bandwidth"`
		Bytes     int64 `json:"bytes"`
		Elapsed   int64 `json:"elapsed"`
		Latency   struct {
			Iqm    float64 `json:"iqm"`
			Low    float64 `json:"low"`
			High   float64 `json:"high"`
			Jitter float64 `json:"jitter"`
		} `json:"latency"`
	} `json:"download"`
	Upload struct {
		Bandwidth int64 `json:"bandwidth"`
		Bytes     int64 `json:"bytes"`
		Elapsed   int64 `json:"elapsed"`
		Latency   struct {
			Iqm    float64 `json:"iqm"`
			Low    float64 `json:"low"`
			High   float64 `json:"high"`
			Jitter float64 `json:"jitter"`
		} `json:"latency"`
	} `json:"upload"`
	PacketLoss int    `json:"packetLoss"`
	ISP        string `json:"isp"`
	Interface  struct {
		InternalIP string `json:"internalIp"`
		Name       string `json:"name"`
		MACAddr    string `json:"macAddr"`
		IsVPN      bool   `json:"isVpn"`
		ExternalIP string `json:"externalIp"`
	} `json:"interface"`
	Server struct {
		ID       int    `json:"id"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		Location string `json:"location"`
		Country  string `json:"country"`
		IP       string `json:"ip"`
	} `json:"server"`
	Result struct {
		ID        string `json:"id"`
		URL       string `json:"url"`
		Persisted bool   `json:"persisted"`
	} `json:"result"`
}

func main() {
	args := []string{}
	args = append(args, "-f")
	args = append(args, "json")

	cmd := exec.Command("speedtest", args...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	var r Result

	err = json.Unmarshal([]byte(out), &r)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := time.Parse(time.RFC3339, r.Timestamp)

	fmt.Println(r.Timestamp)
	if err != nil {
		fmt.Println(err)
		return
	}
	formatttedDate := t.Format("02/01/06")
	formatttedTime := t.Format("15:30")

	downloadBits := float64(r.Download.Bandwidth) * 8
	downloadMbps := downloadBits / 1000000
	uploadBits := float64(r.Upload.Bandwidth) * 8
	uploadMbps := uploadBits / 1000000
	notifyCmd := exec.Command("terminal-notifier", "-message", fmt.Sprintf("Download %.2f Mbps \nUpload %.2f Mbps", downloadMbps, uploadMbps), "-title", fmt.Sprintf("%s - %s - %s", formatttedDate, formatttedTime, r.Server.Name), "-sound", "Funky")
	_ = notifyCmd.Start()

}
