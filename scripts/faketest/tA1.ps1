#!/usr/bin/env pwsh

$ErrorActionPreference = "Stop"
Set-Location $PSScriptRoot/../..

$tmp = [System.IO.Path]::GetTempPath()
New-Item -Path $tmp -Name xxxqqq -ItemType "directory" -Force

Write-Output "127.0.0.1 local1
1.1.1.1 some-service
" > $tmp/hosts1

$host.ui.RawUI.WindowTitle = "== A1 =="

./dist/client `
	--insecure -D --external-ip-nohttp `
	--hosts-file=$tmp/hosts1 `
	--netgroup=A `
	--server=127.0.0.1 `
	--external-ip=172.0.1.1 `
	--internal-ip=127.0.1.1 `
	--perfer-ip=111.1 `
	--hostname=peer-a1 `
	--title="test A 1" `
	--netgroup="A"
