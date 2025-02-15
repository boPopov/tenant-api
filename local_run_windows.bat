@echo off
cd /d %~dp0
for /f "delims=" %%x in ('type .\env\api\.env') do set %%x
go run ./api/main.go