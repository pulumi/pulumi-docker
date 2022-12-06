package provider

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/mapper"
)

func (img *Image) unmarshalPropertyMap(pm resource.PropertyMap) error {
	data := pm.Mappable()

	if err := new(Build).unmarshalPropertyValue(pm["build"], &img.rawBuild); err != nil {
		return err
	}

	// Read everything else except Build using the standard mapper.
	type imageArgsSubset struct {
		ImageName string   `pulumi:"imageName"`
		Registry  Registry `pulumi:"registry,optional"`
		SkipPush  bool     `pulumi:"skipPush,optional"`
	}
	var i imageArgsSubset
	if err := mapper.New(&mapper.Opts{IgnoreUnrecognized: true}).Decode(data, &i); err != nil {
		return err
	}
	img.ImageName = i.ImageName
	img.Registry = i.Registry
	img.SkipPush = i.SkipPush
	return nil
}

func (*Build) unmarshalPropertyValue(pv resource.PropertyValue, buildOrString *interface{}) error {
	switch {
	case pv.IsNull():
		return nil
	case pv.IsString():
		*buildOrString = pv.StringValue()
		return nil
	case pv.IsObject():
		build := Build{}
		pm := pv.ObjectValue()

		// Read CacheFrom.
		if err := new(CacheFrom).unmarshalPropertyValue(pm["cacheFrom"], &build.CacheFrom); err != nil {
			return err
		}

		// Read everything else except CacheFrom using the standard mapper.
		type buildSubset struct {
			Context      string             `pulumi:"context,optional"`
			Dockerfile   string             `pulumi:"dockerfile,optional"`
			Env          map[string]string  `pulumi:"env,optional"`
			Args         map[string]*string `pulumi:"args,optional"`
			ExtraOptions []string           `pulumi:"extraOptions,optional"`
			Target       string             `pulumi:"target,optional"`
		}
		var b buildSubset
		if err := mapper.New(&mapper.Opts{IgnoreUnrecognized: true}).Decode(pm.Mappable(), &b); err != nil {
			return err
		}

		build.Context = b.Context
		build.Dockerfile = b.Dockerfile
		build.Env = b.Env
		build.Args = b.Args
		build.ExtraOptions = b.ExtraOptions
		build.Target = b.Target

		*buildOrString = build
		return nil
	default:
		return fmt.Errorf("Cannot recognize Build from: %v", pv)
	}
}

func (*CacheFrom) unmarshalPropertyValue(pv resource.PropertyValue, boolOrCacheFrom *interface{}) error {
	switch {
	case pv.IsNull():
		return nil
	case pv.IsBool():
		*boolOrCacheFrom = pv.BoolValue()
		return nil
	case pv.IsObject():
		cf := CacheFrom{}
		err := mapper.Map(pv.ObjectValue().Mappable(), &cf)
		*boolOrCacheFrom = cf
		return err
	default:
		return fmt.Errorf("Cannot recognize CacheFrom from: %v", pv)
	}
}
