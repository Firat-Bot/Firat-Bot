token:
	export BOT_TOKEN=OTg1NjQ3MTExMTEyNjI2Mjg2.GtQ4Sw.dmlBANk_KXFi-MMRvbtpqCOQ6tG1iHb8t5IKLY

run:
	go run main.go -t $BOT_TOKEN

task:
	task build

build:
	./bin/Firat-Bot  -t  $BOT_TOKEN

.PHONY: token run task build