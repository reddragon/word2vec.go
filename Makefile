word2vec: word2vec.go
	go build word2vec.go

clean: word2vec
	rm word2vec
