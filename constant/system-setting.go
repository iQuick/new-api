package constant

var ServerAddress = "http://localhost:3000"
var WorkerUrl = ""
var WorkerValidKey = ""

var OutProxyUrl = ""
var OutProxyUsername = ""
var OutProxyPassword = ""

func EnableWorker() bool {
	return WorkerUrl != ""
}

func EnableOutProxy() bool {
	return OutProxyUrl != ""
}
