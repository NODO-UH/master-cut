build64:
	go build -ldflags="-s -w" -o master-cut_x64.bin -i src/main.go

compress: build64
	upx --brute master-cut_x64.bin
