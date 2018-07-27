package main

import (
	"encoding/csv"
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/whiteshtef/clockwork"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var isServer bool                // is the instance a server or the client?
var serverPort int               // port on which to run the server
var serverData map[string]string // pairs of instanceName:IPAddress values
var instanceName string
var destination string

func init() {
	flag.BoolVar(&isServer, "server", false, "Specify if the instance is a server or client")
	flag.IntVar(&serverPort, "server_port", 8000, "Port on which to run the server (if server flag is set to true)")
	hostName, _ := os.Hostname()
	flag.StringVar(&instanceName, "name", hostName, "The name of the client (will appear in logs on the server)")
	flag.StringVar(&destination, "destination", "0.0.0.0", "If running in client mode, specify the server adress and port")
	flag.Parse()
}

func Ping() {
	glog.Info("Reporting name ", instanceName, " to ", destination)
	resp, err := http.Get("http://" + destination + "/report/" + instanceName)
	if err != nil {
		glog.Error("Error:", err)
		return
	}
	if resp.StatusCode != 200 {
		glog.Error("Server returned non 200 status code:", resp)
	}
	glog.Info("Report successful")

}

func HandleReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	ip := strings.Split(r.RemoteAddr, ":")[0]
	serverData[name] = ip

	file, err := os.Create("reports.csv")
	if err != nil {
		glog.Warning("Cannot create reports file:", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for k, v := range serverData {
		var row []string
		row = append(row, k)
		row = append(row, v)
		writer.Write(row)
	}
	glog.Info("Processed report from ", name, " at ", ip)
}

func main() {
	if isServer {
		glog.Info("Starting in server mode. Listening at port ", strconv.Itoa(serverPort))
		serverData = make(map[string]string)
		router := mux.NewRouter()
		router.HandleFunc("/report/{name}", HandleReport).Methods("GET")
		glog.Fatal(http.ListenAndServe(":"+strconv.Itoa(serverPort), router))
	} else {
		glog.Info("Starting in client mode.")
		glog.Info("Reporting name ", instanceName, " to ", destination, " every 10 seconds.")
		sched := clockwork.NewScheduler()
		sched.Schedule().Every(10).Seconds().Do(Ping)
		sched.Run()
	}
}
