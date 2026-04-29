.PHONY: dev
dev:
	go run main.go

.PHONE: edgeone-dev
edgeone-dev:
	edgeone pages dev

.PHONE: edgeone-deploy
edgeone-deploy:
	edgeone pages deploy
