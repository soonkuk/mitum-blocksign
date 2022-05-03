package extension

import (
	"github.com/spikeekips/mitum/base"
	"github.com/spikeekips/mitum/util"
	"github.com/spikeekips/mitum/util/encoder"
	"github.com/spikeekips/mitum/util/valuehash"
)

func (fact *CreateContractAccountsFact) unpack(
	enc encoder.Encoder,
	h valuehash.Hash,
	tk []byte,
	bSender base.AddressDecoder,
	bits []byte,
) error {
	sender, err := bSender.Encode(enc)
	if err != nil {
		return err
	}

	hits, err := enc.DecodeSlice(bits)
	if err != nil {
		return err
	}

	its := make([]CreateContractAccountsItem, len(hits))
	for i := range hits {
		j, ok := hits[i].(CreateContractAccountsItem)
		if !ok {
			return util.WrongTypeError.Errorf("expected CreateContractAccountsItem, not %T", hits[i])
		}

		its[i] = j
	}

	fact.h = h
	fact.token = tk
	fact.sender = sender
	fact.items = its

	return nil
}
