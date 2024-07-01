DEST_IP = 192.168.188.59
build:
	cd frontend-scanner; npm install; npm run build
	cd firmware-rpi/cmd; GOOS=linux GOARCH=arm GOARM=6 go build -o scanner
deploy:
	scp ./firmware-rpi/cmd/scanner marvin@${DEST_IP}:/tmp
doAll: build deploy
clean:
	rm ./firmware-rpi/cmd/scanner
	rm ./firmware-rpi/cmd/resources/*