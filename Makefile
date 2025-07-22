clean:
	@rm ./output/*

build-net-dbg:
	@mkdir -p output
	@clang -I=./net/include/ ./net/main.c -o ./output/net-bin -g
