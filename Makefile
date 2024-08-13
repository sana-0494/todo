build:
	CGO_ENABLED=0 go build -o build/todo .

migrate_up: build 
	build/todo server migrate up -c ./cfg.yaml

migrate_down: build 
	build/todo server migrate down -c ./cfg.yaml
    
start: build
	build/todo server start -c ./cfg.yaml
    