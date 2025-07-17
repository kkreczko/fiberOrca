clean:
	rm ./release/*

build-net:
	@mkdir -p output
	@clang -I=./net/include/ ./net/main.cpp -o ./output/net-bin
