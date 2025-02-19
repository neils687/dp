Write-Output "Scry Info.  All rights reserved."
Write-Output "license that can be found in the license file."

Set-Location $PSScriptRoot

if (Test-Path ".\chain\*") {
    Remove-Item ".\chain\*" -Recurse
}

geth --datadir "chain" init genesis.json
geth --datadir "chain" --rpc --rpcapi "db,eth,net,web3" --nodiscover console
