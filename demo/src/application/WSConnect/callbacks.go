// Scry Info.  All rights reserved.
// license that can be found in the license file.

package WSConnect

import (
    "encoding/json"
    "github.com/btcsuite/btcutil/base58"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/scryinfo/dp/demo/src/application/definition"
    "github.com/scryinfo/dp/demo/src/sdk/core/ethereum/events"
    "github.com/scryinfo/dp/demo/src/sdk/util/storage/ipfsaccess"
    rlog "github.com/sirupsen/logrus"
    "io/ioutil"
    "math/big"
    "os"
    "strconv"
)

const (
    IPFSOutDir      = "D:/desktop"
    EventSendFailed = " event send failed. "
)

func onPublish(event events.Event) bool {
    go func() {
        var op definition.OnPublish
        if op, err = getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
            rlog.Error(errors.Wrap(err, "onPublish: get publish data details failed. "))
        }
        op.Block = event.BlockNumber
        op.Price = event.Data.Get("price").(*big.Int).String()
        op.PublishID = event.Data.Get("publishId").(string)
        op.SupportVerify = event.Data.Get("supportVerify").(bool)

        if err = sendMessage("onPublish", op); err != nil {
            rlog.Error(errors.Wrap(err, "onPublish"+EventSendFailed))
        }
    }()
    return true
}

func getPubDataDetails(ipfsID string) (definition.OnPublish, error) {
    defer func() {
        if er := recover(); er != nil {
            rlog.Error("onPublish.callback: get publish data details failed. ", er)
        }
    }()

    var detailsData definition.OnPublish
    for {
        if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID); err != nil {
            break
        }

        oldFileName := IPFSOutDir + "/" + ipfsID
        newFileName := oldFileName + ".txt"
        if err = os.Rename(oldFileName, newFileName); err != nil {
            break
        }

        var details []byte
        if details, err = ioutil.ReadFile(newFileName); err != nil {
            break
        }
        if err = json.Unmarshal(details, &detailsData); err != nil {
            break
        }

        if err = os.Remove(newFileName); err != nil {
            rlog.Debug("Node - onPublish.callback: delete details file failed. ", err)
        }
        break
    }

    return detailsData, err
}

func onApprove(_ events.Event) bool {
    go func() {
        // -
    }()
    return true
}

func onVerifiersChosen(event events.Event) bool {
    go func() {
        var (
            ovc        definition.OnVerifiersChosen
            extensions []string
        )
        ovc.PublishID = event.Data.Get("publishId").(string)
        if err = sendMessage("onProofFilesExtensions", ovc.PublishID); err != nil {
            rlog.Error(errors.Wrap(err, "onProofFilesExtensions"+EventSendFailed))
        }

        ovc.Block = event.BlockNumber
        ovc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
        ovc.TxState = setTxState(event.Data.Get("state").(uint8))

        extensions = <-channel
        if ovc.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
            rlog.Error(errors.Wrap(err, "Node - onVC.callback: get and rename proof files failed. "))
        }
        if err = sendMessage("onVerifiersChosen", ovc); err != nil {
            rlog.Error(errors.Wrap(err, "onVerifiersChosen"+EventSendFailed))
        }
    }()
    return true
}

func onTransactionCreate(event events.Event) bool {
    go func() {
        var (
            otc        definition.OnTransactionCreate
            extensions []string
        )
        otc.PublishID = event.Data.Get("publishId").(string)
        if err = sendMessage("onProofFilesExtensions", otc.PublishID); err != nil {
            rlog.Error(errors.Wrap(err, "onProofFilesExtensions"+EventSendFailed))
        }

        otc.Block = event.BlockNumber
        otc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
        otc.Buyer = event.Data.Get("users").([]common.Address)[0].String()
        otc.StartVerify = event.Data.Get("needVerify").(bool)
        otc.TxState = setTxState(event.Data.Get("state").(uint8))

        extensions = <-channel
        if otc.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
            rlog.Error(errors.Wrap(err, "Node - onTC.callback: get and rename proof files failed. "))
        }
        if err = sendMessage("onTransactionCreate", otc); err != nil {
            rlog.Error(errors.Wrap(err, "onTransactionCreate"+EventSendFailed))
        }
    }()
    return true
}

// Get proof files from proofIDs.
// ipfsGet -> modify file name, how can I get extension?
func getAndRenameProofFiles(ipfsIDs [][32]byte, extensions []string) ([]string, error) {
    defer func() {
        if er := recover(); er != nil {
            rlog.Error(er, "in callback: get and rename proof files failed. ")
        }
    }()
    var (
        proofs                   = make([]string, len(ipfsIDs))
        oldFileName, newFileName string
    )
    if len(ipfsIDs) != len(extensions) {
        return nil, errors.New("Invalid IPFS IDs or extensions. ")
    }
    for i := 0; i < len(ipfsIDs); i++ {
        ipfsID := ipfsBytes32ToHash(ipfsIDs[i])
        if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID); err != nil {
            return nil, errors.Wrap(err, "Node - callback: IPFS get failed. ")
        }
        oldFileName = IPFSOutDir + "/" + ipfsID
        newFileName = oldFileName + extensions[i]
        if err = os.Rename(oldFileName, newFileName); err != nil {
            return nil, errors.Wrap(err, "Node - callback: rename proof file failed. ")
        }
        proofs[i] = newFileName
    }

    return proofs, nil
}
func ipfsBytes32ToHash(ipfsb [32]byte) string {
    byte34 := make([]byte, 34)
    // if ipfs change encrypt algorithm, byte will change together.
    copy(byte34[:2], []byte{byte(18), byte(32)})
    copy(byte34[2:], ipfsb[:])

    return base58.Encode(byte34)
}

func onPurchase(event events.Event) bool {
    go func() {
        var op definition.OnPurchase
        op.Block = event.BlockNumber
        op.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
        op.MetaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
        op.PublishID = event.Data.Get("publishId").(string)
        op.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
        op.TxState = setTxState(event.Data.Get("state").(uint8))

        // temp
        op.Buyer = event.Data.Get("buyer").(common.Address).String()

        if err = sendMessage("onPurchase", op); err != nil {
            rlog.Error(errors.Wrap(err, "onPurchase"+EventSendFailed))
        }
    }()
    return true
}

func onReadyForDownload(event events.Event) bool {
    go func() {
        var orfd definition.OnReadyForDownload
        orfd.Block = event.BlockNumber
        orfd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
        orfd.MetaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
        orfd.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
        orfd.TxState = setTxState(event.Data.Get("state").(uint8))

        if err = sendMessage("onReadyForDownload", orfd); err != nil {
            rlog.Error(errors.Wrap(err, "onReadyForDownload"+EventSendFailed))
        }
    }()
    return true
}

func onClose(event events.Event) bool {
    go func() {
        var oc definition.OnClose
        oc.Block = event.BlockNumber
        oc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
        oc.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
        oc.TxState = setTxState(event.Data.Get("state").(uint8))

        if err = sendMessage("onClose", oc); err != nil {
            rlog.Error(errors.Wrap(err, "onClose"+EventSendFailed))
        }
    }()

    return true
}

func onRegisterAsVerifier(event events.Event) bool {
    go func() {
        var orav definition.OnRegisterAsVerifier
        orav.Block = event.BlockNumber

        if err = sendMessage("onRegisterVerifier", orav); err != nil {
            rlog.Error(errors.Wrap(err, "onRegisterVerifier"+EventSendFailed))
        }
    }()

    return true
}

func onVote(event events.Event) bool {
    go func() {
        var (
            ov      definition.OnVote
            judge   bool
            comment string
        )
        ov.Block = event.BlockNumber
        ov.VerifierIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
        judge = event.Data.Get("judge").(bool)
        comment = event.Data.Get("comments").(string)
        ov.VerifierResponse = setJudge(judge) + ", " + comment
        ov.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
        ov.TxState = setTxState(event.Data.Get("state").(uint8))

        if err = sendMessage("onVote", ov); err != nil {
            rlog.Error(errors.Wrap(err, "onVote"+EventSendFailed))
        }
    }()

    return true
}

func onVerifierDisable(event events.Event) bool {
    go func() {
        var ovd definition.OnVerifierDisable
        ovd.Block = event.BlockNumber

        if err = sendMessage("onVerifierDisable", ovd); err != nil {
            rlog.Error(errors.Wrap(err, "onVerifierDisable"+EventSendFailed))
        }
    }()

    return true
}

func setTxState(state byte) string {
    var str string
    switch state {
    case 1:
        str = "Created"
    case 2:
        str = "Voted"
    case 3:
        str = "Buying"
    case 4:
        str = "ReadyForDownload"
    case 5:
        str = "Closed"
    default:
        str = "Unknow TxState!"
    }
    return str
}

func setJudge(judge bool) string {
    var str = "Not suggest to buy"
    if judge {
        str = "Suggest to buy"
    }
    return str
}
