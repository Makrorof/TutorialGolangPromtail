package main

import (
	"time"

	"github.com/ic2hrmk/promtail"
)

func main() {
	defer CloseLog()

	//log.Println("OK")
	//
	//Log.Infof("I am here for ya %s", "buddy")
	//
	//// Some staff...
	//
	//Log.Logf(promtail.Info, "Still here")
	//
	//Log.Debugf("I am here for ya %s", "buddy")
	//Log.Infof("I am here for ya %s", "buddy")
	//Log.Warnf("I am here for ya %s", "buddy")
	//Log.Errorf("I am here for ya %s", "buddy")
	//Log.Fatalf("I am here for ya %s", "buddy")
	//Log.Panicf("I am here for ya %s", "buddy")
	//
	//customLabels := map[string]string{
	//	"somethingSpecial": "right-here",
	//}
	//Log.LogfWithLabels(promtail.Info, customLabels, "Still here")
	//
	//log.Println("End")

	//Key-value gibi bir array verirken keylerin bosluksuz olduguna emin olun hata verebilir.
	customLabels := map[string]string{
		"TamamlananAsinler": "ASIN1,ASIN2,ASIN3,ASIN4,ASIN5,ASIN6,ASIN7,ASIN1,ASIN2,ASIN3,ASIN4,ASIN5,ASIN6,ASIN7",
		"HataliAsinler":     "ASIN8,ASIN9,ASIN10,ASIN11,ASIN12,ASIN13,ASIN14,ASIN15,ASIN16,",
	}
	Log.LogfWithLabels(promtail.Info, customLabels, "Tarama tamamalandi")

	for {
		time.Sleep(time.Second)
	}
}
