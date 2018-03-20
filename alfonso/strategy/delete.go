package strategy

import (
	kelp "github.com/lightyeario/kelp/support"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// DeleteStrategy is a strategy to delete the orders for a given currency pair in an account on one side of the orderbook
type DeleteStrategy struct {
	txButler   *kelp.TxButler
	assetBase  *horizon.Asset
	assetQuote *horizon.Asset
}

// ensure it implements Strategy
var _ Strategy = &DeleteStrategy{}

// MakeDeleteStrategy is a factory method for DeleteStrategy
func MakeDeleteStrategy(
	txButler *kelp.TxButler,
	assetBase *horizon.Asset,
	assetQuote *horizon.Asset,
) Strategy {
	return &DeleteStrategy{
		txButler:   txButler,
		assetBase:  assetBase,
		assetQuote: assetQuote,
	}
}

// PruneExistingOffers impl
func (s *DeleteStrategy) PruneExistingOffers(buyingAOffers []horizon.Offer, sellingAOffers []horizon.Offer) ([]horizon.Offer, []horizon.Offer) {
	offers := selectOfferSide(buyingAOffers, sellingAOffers)
	for i := 0; i < len(offers); i++ {
		s.txButler.DeleteOffer(offers[i])
	}

	offers = []horizon.Offer{}
	if buyingAOffers != nil {
		return offers, sellingAOffers
	}
	return buyingAOffers, offers
}

// PreUpdate impl
func (s *DeleteStrategy) PreUpdate(
	maxAssetBase float64,
	maxAssetQuote float64,
	buyingAOffers []horizon.Offer,
	sellingAOffers []horizon.Offer,
) error {
	return nil
}

// UpdateWithOps impl
func (s *DeleteStrategy) UpdateWithOps(
	buyingAOffers []horizon.Offer,
	sellingAOffers []horizon.Offer,
) ([]build.TransactionMutator, error) {
	return []build.TransactionMutator{}, nil
}

// PostUpdate impl
func (s *DeleteStrategy) PostUpdate() error {
	return nil
}