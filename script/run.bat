@echo off

find img/actor001/stands/right/ -type f |
	./bin/tkimgutil.exe scale -s 100 |
	./bin/tkimgutil.exe trim -x 40 -y 320 |
  sort |
	./bin/tkimgutil.exe paste
