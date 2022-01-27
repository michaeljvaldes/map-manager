build:
	make prebuild
	go build -o bin/map-manager/map-manager.exe cmd/main.go
	make postbuild


prebuild:
	rm -rf bin/*
	mkdir bin/map-manager
	cp configs/config.yml bin/map-manager/config.yml
	
postbuild:
	cd bin && tar -czf map-manager.tar.gz map-manager