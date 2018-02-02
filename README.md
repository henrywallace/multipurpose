# multipurpose

Simple swiss-army knife go tool `mp` for various tasks while working the shell.

`mp durfmt DURATION` - parse a time duration and output a time.Duration parsed
version with each part (e.g. seconds) truncated to %.2f precision.

`curl http://theinter.net/book.txt | mp counter | head -10` - count the
top 10 words via strings.Fields on input.

TODO:
  - switch git branches
  - show git branch status
  - add option to remove stop words in counter
  - add option to do ngrams in counter
