package encode

import (
	"github.com/0987363/go-geobuf/pkg/geojson"
	"github.com/0987363/go-geobuf/proto"
)

func EncodeFeatureCollection(collection *geojson.FeatureCollection, opts *EncodingConfig) (*proto.Data_FeatureCollection, error) {
	features := make([]*proto.Data_Feature, len(collection.Features))

	for i, feature := range collection.Features {
		encoded, err := EncodeFeature(feature, opts)
		if err != nil {
			return nil, err
		}
		features[i] = encoded
	}

	return &proto.Data_FeatureCollection{
		Features: features,
	}, nil
}
