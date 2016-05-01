word2vec: word2vec.go
	go build word2vec.go
	chmod +x test.sh

clean: word2vec
	rm word2vec
