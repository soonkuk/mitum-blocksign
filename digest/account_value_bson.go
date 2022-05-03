package digest // nolint: dupl, revive

import (
	"github.com/spikeekips/mitum/base"
	bsonenc "github.com/spikeekips/mitum/util/encoder/bson"
	"go.mongodb.org/mongo-driver/bson"
)

func (va AccountValue) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(bsonenc.MergeBSONM(
		bsonenc.NewHintedDoc(va.Hint()),
		bson.M{
			"ac":              va.ac,
			"balance":         va.balance,
			"owner":           va.owner,
			"documents":       va.document,
			"height":          va.height,
			"previous_height": va.previousHeight,
		},
	))
}

type AccountValueBSONUnpacker struct {
	AC bson.Raw            `bson:"ac"`
	BL bson.Raw            `bson:"balance"`
	OW base.AddressDecoder `bson:"owner"`
	CD bson.Raw            `bson:"documents"`
	HT base.Height         `bson:"height"`
	PT base.Height         `bson:"previous_height"`
}

func (va *AccountValue) UnpackBSON(b []byte, enc *bsonenc.Encoder) error {
	var uva AccountValueBSONUnpacker
	if err := enc.Unmarshal(b, &uva); err != nil {
		return err
	}

	return va.unpack(enc, uva.AC, uva.BL, uva.OW, uva.CD, uva.HT, uva.PT)
}
