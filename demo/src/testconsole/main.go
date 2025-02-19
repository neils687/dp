// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/dp/demo/src/sdk"
    "github.com/scryinfo/dp/demo/src/sdk/core/chainoperations"
    "github.com/scryinfo/dp/demo/src/sdk/core/ethereum/events"
    "github.com/scryinfo/dp/demo/src/sdk/scryclient"
    cif "github.com/scryinfo/dp/demo/src/sdk/scryclient/chaininterfacewrapper"
    "github.com/scryinfo/dp/dots/service"
    "math/big"
    "time"
)

var (
    publishId                        = ""
    txId                    *big.Int = big.NewInt(0)
    metaDataIdEncWithSeller []byte
    metaDataIdEncWithBuyer  []byte
    protocolContractAddr                           = "0x3c4d26e916d79fc3fc925027a79612012462f691"
    tokenContractAddr                              = "0x5c29f42d640ee25f080cdc648641e8e358459ddc"
    clientPassword                                 = "888888"
    suAddress                                      = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8"
    suPassword                                     = "111111"
    seller                  *scryclient.ScryClient = nil
    buyer                   *scryclient.ScryClient = nil
    verifier1               *scryclient.ScryClient = nil
    verifier2               *scryclient.ScryClient = nil
    verifier3               *scryclient.ScryClient = nil
    arbitrator              *scryclient.ScryClient = nil
    sleepTime               time.Duration          = 20000000000
    startVerify             bool                   = false
    startTestFromBlock      bool                   = false
)

func main() {
    //note: asServiceAddr is the host of key management service which is outside
    err := sdk.Init(
        "http://127.0.0.1:8545/",
        "192.168.1.14:48080",
        protocolContractAddr,
        tokenContractAddr,
        "/ip4/127.0.0.1/tcp/5001",
        "./log/sdk.log",
        "testconsole")
    if err != nil {
        fmt.Println("failed to initialize sdk, error:", err)
        return
    }

    initClients()

    testAccounts()

    testTxWithoutVerify()

    testTxWithVerify()

    time.Sleep(100000000000000)
}

func testAccounts() {
    fmt.Println("Start testing accounts...")

    ac, err := service.GetAMIns().CreateAccount("12345")
    if err != nil {
        fmt.Println("failed to create interface, error:", err)
    }

    rv, err := service.GetAMIns().AuthAccount(ac.Address, "12345")
    if err != nil {
        fmt.Println("failed to authenticate interface, error:", err)
    }

    if rv {
        fmt.Println("account_ authentication passed")
    } else {
        fmt.Println("account_ authentication not passed")
    }

}

func testTxWithoutVerify() {
    fmt.Println("Start testing tx without verification...")

    startVerify = false

    SellerPublishData(false)

}

func testTxWithVerify() {
    fmt.Println("Start testing tx with verification...Note: please restart the test chain before running this case")

    startVerify = true
    startTestFromBlock = false

    subscribeAllEvents()

    VerifierApproveTransfer(verifier1)

    time.Sleep(sleepTime)

    VerifierApproveTransfer(verifier2)

    time.Sleep(sleepTime)

    VerifierApproveTransfer(verifier3)

    time.Sleep(sleepTime)

    RegisterAsVerifier(verifier1)

    time.Sleep(sleepTime)

    RegisterAsVerifier(verifier2)

    time.Sleep(sleepTime)

    RegisterAsVerifier(verifier3)

    time.Sleep(sleepTime)

    SellerPublishData(true)
}

func testFromBlock() {
    startTestFromBlock = true
    seller.SubscribeEvent("DataPublish", onPublish)
    sdk.StartScan(1)
}

func subscribeAllEvents() {
    seller.SubscribeEvent("DataPublish", onPublish)
    seller.SubscribeEvent("Buy", onPurchase)

    verifier1.SubscribeEvent("Approval", onApprovalVerifierTransfer)
    verifier1.SubscribeEvent("RegisterVerifier", OnRegisterVerifier)

    verifier2.SubscribeEvent("Approval", onApprovalVerifierTransfer)
    verifier2.SubscribeEvent("RegisterVerifier", OnRegisterVerifier)

    verifier3.SubscribeEvent("Approval", onApprovalVerifierTransfer)
    verifier3.SubscribeEvent("RegisterVerifier", OnRegisterVerifier)

    if startVerify {
        verifier1.SubscribeEvent("TransactionCreate", onVerifier1Chosen)
        verifier2.SubscribeEvent("TransactionCreate", onVerifier2Chosen)
        verifier3.SubscribeEvent("TransactionCreate", onVerifier3Chosen)
    }

    buyer.SubscribeEvent("Vote", onVote)
    buyer.SubscribeEvent("Approval", onApprovalBuyerTransfer)
    buyer.SubscribeEvent("TransactionCreate", onTransactionCreate)
    buyer.SubscribeEvent("ReadyForDownload", onReadyForDownload)
    buyer.SubscribeEvent("TransactionClose", onClose)
    buyer.SubscribeEvent("VerifierDisable", onVerifierDisable)
}

func unsubscribeAllEvents() {
    seller.UnSubscribeEvent("DataPublish")
    seller.UnSubscribeEvent("Buy")

    verifier1.UnSubscribeEvent("Approval")
    verifier1.UnSubscribeEvent("RegisterVerifier")

    verifier2.UnSubscribeEvent("Approval")
    verifier2.UnSubscribeEvent("RegisterVerifier")

    verifier3.UnSubscribeEvent("Approval")
    verifier3.UnSubscribeEvent("RegisterVerifier")

    if startVerify {
        verifier1.UnSubscribeEvent("TransactionCreate")
        verifier2.UnSubscribeEvent("TransactionCreate")
        verifier3.UnSubscribeEvent("TransactionCreate")
    }

    buyer.UnSubscribeEvent("Vote")
    buyer.UnSubscribeEvent("Approval")
    buyer.UnSubscribeEvent("TransactionCreate")
    buyer.UnSubscribeEvent("ReadyForDownload")
    buyer.UnSubscribeEvent("TransactionClose")
    buyer.UnSubscribeEvent("VerifierDisable")
}

func initClients() {
    var err error
    seller, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    buyer, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    verifier1, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    verifier2, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    verifier3, err = CreateClientWithToken(big.NewInt(10000000), big.NewInt(10000000))
    if err != nil {
        fmt.Println("failed to init clients, error:", err)
        panic(err)
    }

    time.Sleep(sleepTime)
}

func CreateClientWithToken(token *big.Int, eth *big.Int) (*scryclient.ScryClient, error) {
    client, err := scryclient.CreateScryClient(clientPassword)
    if err != nil {
        fmt.Println("failed to create user, error:", err)
        return nil, err
    }

    _, err = cif.TransferEth(common.HexToAddress(suAddress),
        suPassword,
        common.HexToAddress(client.Account.Address),
        eth)
    if err != nil {
        fmt.Println("failed to transfer token, error:", err)
        return nil, err
    }

    txParam := chainoperations.TransactParams{common.HexToAddress(suAddress), suPassword, big.NewInt(0), false}
    err = cif.TransferTokens(&txParam, common.HexToAddress(client.Account.Address), token)
    if err != nil {
        fmt.Println("failed to transfer token, error:", err)
        return nil, err
    }

    return client, nil
}

func SellerPublishData(supportVerify bool) {
    //publish data
    metaData := []byte("QmcHXkMXwgvZP56tsUJNtcfedojHkqrDsgkC4fbsBM1zre")
    proofData := []string{"QmNrTHX545s7hGfbEVrJuCiMqKVwUWJ4cPwXrAPv3GW5pm", "Qmb7csVP7wGco16XHVNqqRXUE7vQMjuA24QRypnZdkeQMD"}
    despData := "QmQXqZdEwXnWgKpfJmUVaACuVar9R7vpBxtZgddQMTa2UN"

    txParam := chainoperations.TransactParams{common.HexToAddress(seller.Account.Address), clientPassword, big.NewInt(0), false}
    cif.Publish(&txParam, big.NewInt(1000), metaData, proofData, 2, despData, supportVerify)
}

func VerifierApproveTransfer(verifier *scryclient.ScryClient) {
    txParam := chainoperations.TransactParams{common.HexToAddress(verifier.Account.Address), clientPassword, big.NewInt(0), false}
    err := cif.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(10000))
    if err != nil {
        fmt.Println("VerifierApproveTransfer", err)
    }
}

func RegisterAsVerifier(verifier *scryclient.ScryClient) {
    txParam := chainoperations.TransactParams{common.HexToAddress(verifier.Account.Address), clientPassword, big.NewInt(0), false}
    err := cif.RegisterAsVerifier(&txParam)
    if err != nil {
        fmt.Println("RegisterAsVerifier", err)
    }
}

func Vote(verifier *scryclient.ScryClient) {
    txParam := chainoperations.TransactParams{common.HexToAddress(verifier.Account.Address), clientPassword, big.NewInt(0), false}
    err := cif.Vote(&txParam, txId, true, "This could be real from "+verifier.Account.Address)
    if err != nil {
        fmt.Println("Vote:", err)
    }
}

func CreditsToVerifier(to common.Address) {
    txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), clientPassword, big.NewInt(0), false}
    err := cif.CreditsToVerifier(&txParam, txId, 1, 5)
    if err != nil {
        fmt.Println("CreditsToVerifier:", err)
    }
}

func BuyerApproveTransfer() {
    txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), clientPassword, big.NewInt(0), false}
    err := cif.ApproveTransfer(&txParam, common.HexToAddress(protocolContractAddr), big.NewInt(1600))
    if err != nil {
        fmt.Println("BuyerApproveTransfer:", err)
    }
}

func PrepareToBuy(publishId string, startVerify bool) {
    txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), clientPassword,
        big.NewInt(0), false}
    err := cif.PrepareToBuy(&txParam, publishId, startVerify)
    if err != nil {
        fmt.Println("failed to prepareToBuy, error:", err)
    }
}

func Buy(txId *big.Int) {
    txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address), clientPassword, big.NewInt(0), false}
    err := cif.BuyData(&txParam, txId)
    if err != nil {
        fmt.Println("failed to buyData, error:", err)
    }
}

func SubmitMetaDataIdEncWithBuyer(txId *big.Int) {
    txParam := chainoperations.TransactParams{common.HexToAddress(seller.Account.Address), clientPassword, big.NewInt(0), false}
    err := cif.SubmitMetaDataIdEncWithBuyer(&txParam, txId, metaDataIdEncWithBuyer)
    if err != nil {
        fmt.Println("failed to SubmitMetaDataIdEncWithBuyer, error:", err)
    }
}

func ConfirmDataTruth(txId *big.Int) {
    txParam := chainoperations.TransactParams{common.HexToAddress(buyer.Account.Address),
        clientPassword, big.NewInt(0), false}
    err := cif.ConfirmDataTruth(&txParam, txId, true)
    if err != nil {
        fmt.Println("failed to ConfirmDataTruth, error:", err)
    }
}

func onPurchase(event events.Event) bool {
    fmt.Println("onPurchase:", event)
    metaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
    fmt.Println("Node: EncID. ", metaDataIdEncWithSeller)
    metaDataIdEncWithBuyer = make([]byte, len(metaDataIdEncWithSeller))
    copy(metaDataIdEncWithBuyer, metaDataIdEncWithSeller)

    SubmitMetaDataIdEncWithBuyer(txId)
    return true
}

func onReadyForDownload(event events.Event) bool {
    fmt.Println("onReadyForDownload:", event)
    metaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)

    ConfirmDataTruth(txId)

    return true
}

func onClose(event events.Event) bool {
    fmt.Println("onClose:", event)

    fmt.Println("Testing Tx end")

    unsubscribeAllEvents()

    fmt.Println("Start testing from block")

    testFromBlock()

    return true
}

func onApprovalBuyerTransfer(event events.Event) bool {
    fmt.Println("onApprovalBuyerTransfer:", event)

    PrepareToBuy(publishId, startVerify)
    return true
}

func onApprovalVerifierTransfer(event events.Event) bool {
    fmt.Println("onApprovalVerifierTransfer:", event)

    return true
}

func onTransactionCreate(event events.Event) bool {
    fmt.Println("onTransactionCreated:", event)
    txId = event.Data.Get("transactionId").(*big.Int)
    if !startVerify {
        Buy(txId)
    }

    return true
}

func onVerifier1Chosen(event events.Event) bool {
    fmt.Println("onVerifier1Chosen:", event)

    txId = event.Data.Get("transactionId").(*big.Int)
    Vote(verifier1)
    return true
}

func onVerifier2Chosen(event events.Event) bool {
    fmt.Println("onVerifier2Chosen:", event)

    txId = event.Data.Get("transactionId").(*big.Int)
    Vote(verifier2)
    return true
}

func onVerifier3Chosen(event events.Event) bool {
    fmt.Println("onVerifier3Chosen:", event)

    txId = event.Data.Get("transactionId").(*big.Int)
    Vote(verifier3)
    return true
}

func onPublish(event events.Event) bool {
    fmt.Println("onpublish: ", event)

    publishId = event.Data.Get("publishId").(string)

    if !startTestFromBlock {
        BuyerApproveTransfer()
    }

    return true
}

func OnRegisterVerifier(event events.Event) bool {
    fmt.Println("OnRegisterVerifier: ", event)

    return true
}

func onVote(event events.Event) bool {
    fmt.Println("onVote: ", event)

    Buy(txId)

    return true
}

func onVerifierDisable(event events.Event) bool {
    fmt.Println("onVerifierDisable: ", event)

    return true
}
