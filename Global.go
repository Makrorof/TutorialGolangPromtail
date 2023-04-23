package main

import (
	"fmt"
	"log"

	"github.com/ic2hrmk/promtail"
)

const (
	lokiServerPort = 3100
	lokiServerIP   = "localhost"
	//InstanceID yerine baska bir degisken adi kullanilabilir ServerID gibi grafana'da label degeri olarak gozukuyor.
	lokiInstanceId = "my-Tutorial-Promtail-Client"
)

var Log = newPromtailClient()

func newPromtailClient() promtail.Client {
	//InstanceID yerine baska bir degisken adi kullanilabilir ServerID gibi grafana'da label degeri olarak gozukuyor.
	identifiers := map[string]string{
		"instanceId": lokiInstanceId,
	}

	promtailClient, err := promtail.NewJSONv1Client(lokiServerIP+":"+fmt.Sprint(lokiServerPort), identifiers)
	if err != nil {
		log.Fatal(err)
	}

	return promtailClient
	//defer promtailClient.Close()
}

// Close promtail client
func CloseLog() {
	Log.Close()
}
