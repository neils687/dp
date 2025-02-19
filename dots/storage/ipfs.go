// Scry Info.  All rights reserved.
// license that can be found in the license file.

package storage

import (
    "github.com/ipfs/go-ipfs-api"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "strings"
)

const (
    IpfsTypeId = "6763f97f-dfd2-40eb-8925-b8a031aab461"
    IpfsLiveId = "6763f97f-dfd2-40eb-8925-b8a031aab461"
)

type Ipfs struct {
    sh *shell.Shell
}

func (c *Ipfs) Create(l dot.Line) error {

    return nil
}

//construct dot
func newIpfsDot() (dot.Dot, error) {
    d := &Ipfs{}
    return d, nil
}

//Data structure needed when generating newer component
func IpfsTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
        Meta: dot.Metadata{TypeId: IpfsTypeId,
            NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
                return newIpfsDot()
            }},
    }
}

func (c *Ipfs) Initialize(storageSrvAddr string) error {
    c.sh = shell.NewShell(storageSrvAddr)
    if c.sh == nil {
        return errors.New("Failed to create ipfs shell.")
    }

    return nil
}

func (c *Ipfs) Save(value []byte) (string, error) {
    if c.sh == nil {
        return "", errors.New("Ipfs api shell is nil")
    }

    return c.sh.Add(strings.NewReader(string(value)))
}

func (c *Ipfs) Get(key string, outDir string) error {
    if c.sh == nil {
        return errors.New("Get from ipfs failed, ipfs api shell is nil. ")
    }

    return c.sh.Get(key, outDir)
}
