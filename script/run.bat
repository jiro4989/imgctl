@echo off

setlocal exe=tkimgutil.exe

%exe% generate | ^
	%exe% scale -s 50 | ^
	%exe% trim -x 100 -y 290 | ^
	sort | ^
	%exe% paste
