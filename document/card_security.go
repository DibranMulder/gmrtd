package document

import (
	"fmt"
	"log/slog"
	"slices"

	cms "github.com/dibranmulder/gmrtd/cms"
	"github.com/dibranmulder/gmrtd/oid"
	"github.com/dibranmulder/gmrtd/utils"
)

type CardSecurity struct {
	RawData       []byte          `json:"rawData,omitempty"`
	SD            *cms.SignedData `json:"sd,omitempty"`
	SecurityInfos *SecurityInfos  `json:"securityInfos,omitempty"`
}

func NewCardSecurity(data []byte) (out *CardSecurity, err error) {
	if len(data) < 1 {
		return nil, nil
	}

	out = new(CardSecurity)

	out.RawData = slices.Clone(data)

	slog.Debug("NewCardSecurity", "bytes", utils.BytesToHex(out.RawData))

	// NB no root node for CardSecurity, so directly parse ASN1 SignedData

	{
		var sd *cms.SignedData
		var err error

		sd, err = cms.ParseSignedData(out.RawData)
		if err != nil {
			return nil, err
		}

		out.SD = sd

		// verify the content-type is as expected
		if !sd.Content.EContentType.Equal(oid.OidSecurityObject) {
			return nil, fmt.Errorf("incorrect ContentType (got:%s)", sd.Content.EContentType.String())
		}
		eContent := sd.Content.EContent

		if out.SecurityInfos, err = DecodeSecurityInfos(eContent); err != nil {
			return nil, err
		}
	}

	return out, nil
}
