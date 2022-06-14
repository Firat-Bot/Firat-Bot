token:
	export BOT_TOKEN=OTg1NjQ3MTExMTEyNjI2Mjg2.G-_GYC.BIXd3wpbuIxzYfmc9hXF39xtyH586GK1wAEogc

run:
	go run main.go -t $BOT_TOKEN

task:
	task build

build:
	./bin/Firat-Bot  -t  $ BOT_TOKEN

.PHONY: token run task build