if [ ! -e pg100.txt ]; then
  wget http://www.gutenberg.org/cache/epub/100/pg100.txt
  sed -i e 124368,124787d pg100.txt
  sed -i e 1,173d pg100.txt
fi
./word2vec -t pg100.txt
