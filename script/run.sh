#!/bin/bash

find img/actor001/stands/right/ -type f |
	./bin/tkimgutil scale -s 100 |
	./bin/tkimgutil trim -x 40 -y 320 |
  sort |
	./bin/tkimgutil paste
