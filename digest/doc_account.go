package digest

import (
	"github.com/pkg/errors"
	"github.com/protoconNet/mitum-document/document"
	"github.com/protoconNet/mitum-document/extension"
	"github.com/spikeekips/mitum-currency/currency"
	"github.com/spikeekips/mitum/base"
	"github.com/spikeekips/mitum/base/state"
	mongodbstorage "github.com/spikeekips/mitum/storage/mongodb"
	"github.com/spikeekips/mitum/util/encoder"
	bsonenc "github.com/spikeekips/mitum/util/encoder/bson"
)

type AccountDoc struct {
	mongodbstorage.BaseDoc
	address string
	height  base.Height
	pubs    []string
}

func NewAccountDoc(rs AccountValue, enc encoder.Encoder) (AccountDoc, error) {
	b, err := mongodbstorage.NewBaseDoc(nil, rs, enc)
	if err != nil {
		return AccountDoc{}, err
	}

	var pubs []string
	if keys := rs.Account().Keys(); keys != nil {
		ks := keys.Keys()
		pubs = make([]string, len(ks))
		for i := range ks {
			k := ks[i].Key()
			pubs[i] = k.String()
		}
	}

	address := rs.ac.Address()
	return AccountDoc{
		BaseDoc: b,
		address: address.String(),
		height:  rs.height,
		pubs:    pubs,
	}, nil
}

func (doc AccountDoc) MarshalBSON() ([]byte, error) {
	m, err := doc.BaseDoc.M()
	if err != nil {
		return nil, err
	}

	m["address"] = doc.address
	m["height"] = doc.height
	m["pubs"] = doc.pubs

	return bsonenc.Marshal(m)
}

type BalanceDoc struct {
	mongodbstorage.BaseDoc
	st state.State
	am currency.Amount
}

// NewBalanceDoc gets the State of Amount
func NewBalanceDoc(st state.State, enc encoder.Encoder) (BalanceDoc, error) {
	am, err := currency.StateBalanceValue(st)
	if err != nil {
		return BalanceDoc{}, errors.Wrap(err, "balanceDoc needs Amount state")
	}

	b, err := mongodbstorage.NewBaseDoc(nil, st, enc)
	if err != nil {
		return BalanceDoc{}, err
	}

	return BalanceDoc{
		BaseDoc: b,
		st:      st,
		am:      am,
	}, nil
}

func (doc BalanceDoc) MarshalBSON() ([]byte, error) {
	m, err := doc.BaseDoc.M()
	if err != nil {
		return nil, err
	}

	address := doc.st.Key()[:len(doc.st.Key())-len(currency.StateKeyBalanceSuffix)-len(doc.am.Currency())-1]
	m["address"] = address
	m["currency"] = doc.am.Currency().String()
	m["height"] = doc.st.Height()

	return bsonenc.Marshal(m)
}

type DocumentsDoc struct {
	mongodbstorage.BaseDoc
	st state.State
}

// NewDocumentsDoc gets the State of Document Inventory
func NewDocumentsDoc(st state.State, enc encoder.Encoder) (DocumentsDoc, error) {
	b, err := mongodbstorage.NewBaseDoc(nil, st, enc)
	if err != nil {
		return DocumentsDoc{}, err
	}
	return DocumentsDoc{
		BaseDoc: b,
		st:      st,
	}, nil
}

func (doc DocumentsDoc) MarshalBSON() ([]byte, error) {
	m, err := doc.BaseDoc.M()
	if err != nil {
		return nil, err
	}
	address := doc.st.Key()[:len(doc.st.Key())-len(document.StateKeyDocumentsSuffix)]
	m["address"] = address
	m["height"] = doc.st.Height()

	return bsonenc.Marshal(m)
}

type ContractAccountOwnerDoc struct {
	mongodbstorage.BaseDoc
	st state.State
}

// NewContractAccountOwnerDoc gets the State of contract account owner
func NewContractAccountOwnerDoc(st state.State, enc encoder.Encoder) (ContractAccountOwnerDoc, error) {
	b, err := mongodbstorage.NewBaseDoc(nil, st, enc)
	if err != nil {
		return ContractAccountOwnerDoc{}, err
	}
	return ContractAccountOwnerDoc{
		BaseDoc: b,
		st:      st,
	}, nil
}

func (doc ContractAccountOwnerDoc) MarshalBSON() ([]byte, error) {
	m, err := doc.BaseDoc.M()
	if err != nil {
		return nil, err
	}
	address := doc.st.Key()[:len(doc.st.Key())-len(extension.StateKeyContractAccountOwnerSuffix)]
	m["address"] = address
	m["height"] = doc.st.Height()

	return bsonenc.Marshal(m)
}
