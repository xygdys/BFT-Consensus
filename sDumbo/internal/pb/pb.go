package pb //provable broadcast

import (
	"bytes"
	"context"
	"log"
	"sDumbo/internal/party"
	"sDumbo/pkg/core"
	"sDumbo/pkg/protobuf"
	"sync"

	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/tbls"
	"golang.org/x/crypto/sha3"
)

//Sender is run by the sender of a instance of provable broadcast
func Sender(ctx context.Context, p *party.HonestParty, ID []byte, value []byte, validation []byte) ([]byte, []byte, bool) {
	valueMessage := core.Encapsulation("Value", ID, p.PID, &protobuf.Value{
		Value:      value,
		Validation: validation,
	})

	p.Broadcast(valueMessage)

	sigs := [][]byte{}
	h := sha3.Sum512(value)
	var buf bytes.Buffer
	buf.Write([]byte("Echo"))
	buf.Write(ID)
	buf.Write(h[:])
	sm := buf.Bytes()

	for {
		select {
		case <-ctx.Done():
			return nil, nil, false
		case m := <-p.GetMessage("Echo", ID):
			payload := core.Decapsulation("Echo", m).(*protobuf.Echo)
			sigs = append(sigs, payload.Sigshare)
			if len(sigs) > int(2*p.F) {
				signature, err := tbls.Recover(pairing.NewSuiteBn256(), p.SigPK, sm, sigs, int(2*p.F+1), int(p.N))

				if err != nil {
					log.Fatalln(err)
				}

				return h[:], signature, true
			}
		}
	}

}

//Receiver is run by the receiver of a instance of provable broadcast
func Receiver(ctx context.Context, p *party.HonestParty, sender uint32, ID []byte, validator func(*party.HonestParty, []byte, []byte, []byte, *sync.Map, *sync.Map) error, hashVerifyMap *sync.Map, sigVerifyMap *sync.Map) ([]byte, []byte, bool) {
	select {
	case <-ctx.Done():
		return nil, nil, false
	case m := <-p.GetMessage("Value", ID):
		payload := (core.Decapsulation("Value", m)).(*protobuf.Value)
		if validator != nil {
			err := validator(p, ID, payload.Value, payload.Validation, hashVerifyMap, sigVerifyMap)
			if err != nil {
				log.Fatalln(err, "PB validator for", m.Sender)
				return nil, nil, false //sender is dishonest
			}
		}
		h := sha3.Sum512(payload.Value)
		var buf bytes.Buffer
		buf.Write([]byte("Echo"))
		buf.Write(ID)
		buf.Write(h[:])
		sm := buf.Bytes()
		sigShare, _ := tbls.Sign(pairing.NewSuiteBn256(), p.SigSK, sm) //sign("Echo"||ID||h)

		echoMessage := core.Encapsulation("Echo", ID, p.PID, &protobuf.Echo{
			Sigshare: sigShare,
		})
		p.Send(echoMessage, sender)

		return payload.Value, payload.Validation, true
	}
}
