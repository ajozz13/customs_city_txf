@echo off
set GO_ENV=production

REM usage cc.exe PATH_TO_FILE T or F    true is for ECCF F for CFS
cc.exe %1 %2
