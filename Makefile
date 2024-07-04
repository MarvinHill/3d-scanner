frontend:
	cd frontend-scanner; npm install; npm run build
firmware:
	cd firmware-rpi/cmd; go get .; GOOS=linux GOARCH=arm GOARM=6 go build -o ./../../build/scanner
build: frontend firmware
clean:
	rm ./firmware-rpi/cmd/scanner
	rm ./firmware-rpi/cmd/resources/*