build:
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/fancontrol fancontrol.go
	
clean:
	rm -rf build/