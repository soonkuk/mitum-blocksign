package blocksign // nolint:dupl

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/soonkuk/mitum-data/currency"
	"github.com/spikeekips/mitum/base"
	bsonenc "github.com/spikeekips/mitum/util/encoder/bson"
)

func (it BaseTransferDocumentsItem) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bsonenc.MergeBSONM(bsonenc.NewHintedDoc(it.Hint()),
			bson.M{
				"documentid": it.docId,
				"owner":      it.owner,
				"receiver":   it.receiver,
				"currency":   it.cid,
			}),
	)
}

type BaseTransferDocumentsItemBSONUnpacker struct {
	DM currency.Big        `bson:"documentid"`
	OW base.AddressDecoder `bson:"owner"`
	RC base.AddressDecoder `bson:"receiver"`
	CI string              `bson:"currency"`
}

func (it *BaseTransferDocumentsItem) UnpackBSON(b []byte, enc *bsonenc.Encoder) error {
	var ht bsonenc.HintedHead
	if err := enc.Unmarshal(b, &ht); err != nil {
		return err
	}

	var uit BaseTransferDocumentsItemBSONUnpacker
	if err := enc.Unmarshal(b, &uit); err != nil {
		return err
	}

	return it.unpack(enc, ht.H, uit.DM, uit.OW, uit.RC, uit.CI)
}