# scripts/build.ps1 — build ghostline for windows
Write-Host "[*] tidying modules..." -ForegroundColor Cyan
go mod tidy

Write-Host "[*] building ghostline.exe..." -ForegroundColor Cyan
go build -ldflags "-s -w" -o ghostline.exe .

Write-Host "[*] done. run: .\ghostline.exe" -ForegroundColor Green