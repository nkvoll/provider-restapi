/*
Copyright 2023 Njal Karevoll.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "restapi"
	modulePath     = "github.com/nkvoll/provider-restapi"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("restapi.k8s.karevoll.no"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		func(p *ujconfig.Provider) {
			p.AddResourceConfigurator("restapi_object", func(r *ujconfig.Resource) {
				r.ShortGroup = "object"
			})
		},
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
