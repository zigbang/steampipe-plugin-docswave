install:
	go build -o  ~/.steampipe/plugins/hub.steampipe.io/plugins/zigbang/docswave@latest/steampipe-plugin-docswave.plugin *.go

local:
	go build -o  ~/.steampipe/plugins/local/docswave/docswave.plugin *.go
