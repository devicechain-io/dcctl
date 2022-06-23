/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package graphql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	dmgql "github.com/devicechain-io/dc-device-management/gqlclient"
	dmmodel "github.com/devicechain-io/dc-device-management/model"
	"github.com/spf13/cobra"
)

// Gets a GraphQL client based on command flags and other settings.
func GetDeviceManagementGraphQLClient(cmd *cobra.Command) graphql.Client {
	return GetGraphQLClientForCommand(cmd, "device-management")
}

// Assure a device type (check for existing or create new).
func AssureDeviceType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("device type", token)
	req := dmmodel.DeviceTypeCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceType.Token)
	}
}

// Assure a device (check for existing or create new).
func AssureDevice(ctx context.Context, cli graphql.Client, token string, deviceTypeToken string, name *string,
	description *string, metadata *string) {
	assure("device", token)
	req := dmmodel.DeviceCreateRequest{
		Token:           token,
		DeviceTypeToken: deviceTypeToken,
		Name:            name,
		Description:     description,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureDevice(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DevicesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateDevice.Token)
	}
}

// Assure a device relationship type (check for existing or create new).
func AssureDeviceRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("device relationship type", token)
	req := dmmodel.DeviceRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceRelationshipType.Token)
	}
}

// Assure a device relationship (check for existing or create new).
func AssureDeviceRelationship(ctx context.Context, cli graphql.Client, token string, source string,
	target string, relation string, metadata *string) {
	assure("device relationship", token)
	req := dmmodel.DeviceRelationshipCreateRequest{
		Token:            token,
		SourceDevice:     source,
		TargetDevice:     target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceRelationship.Token)
	}
}

// Assure a device group (check for existing or create new).
func AssureDeviceGroup(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("device group", token)
	req := dmmodel.DeviceGroupCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceGroupsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceGroup.Token)
	}
}

// Assure a device group relationship type (check for existing or create new).
func AssureDeviceGroupRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("device group relationship type", token)
	req := dmmodel.DeviceGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceGroupRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceGroupRelationshipType.Token)
	}
}

// Assure a device group relationship (check for existing or create new).
func AssureDeviceGroupRelationship(ctx context.Context, cli graphql.Client, token string, deviceGroup string,
	device string, relation string, metadata *string) {
	assure("device group relationship", token)
	req := dmmodel.DeviceGroupRelationshipCreateRequest{
		Token:            token,
		DeviceGroup:      deviceGroup,
		Device:           device,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureDeviceGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.DeviceGroupRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateDeviceGroupRelationship.Token)
	}
}

// Assure an asset type (check for existing or create new).
func AssureAssetType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("asset type", token)
	req := dmmodel.AssetTypeCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureAssetType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AssetTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAssetType.Token)
	}
}

// Assure an asset (check for existing or create new).
func AssureAsset(ctx context.Context, cli graphql.Client, token string, assetTypeToken string, name *string,
	description *string, metadata *string) {
	assure("asset", token)
	req := dmmodel.AssetCreateRequest{
		Token:          token,
		AssetTypeToken: assetTypeToken,
		Name:           name,
		Description:    description,
		Metadata:       metadata,
	}
	gresp, cresp, err := dmgql.AssureAsset(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AssetsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAsset.Token)
	}
}

// Assure an asset relationship type (check for existing or create new).
func AssureAssetRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("asset relationship type", token)
	req := dmmodel.AssetRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureAssetRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AssetRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAssetRelationshipType.Token)
	}
}

// Assure an asset relationship (check for existing or create new).
func AssureAssetRelationship(ctx context.Context, cli graphql.Client, token string, source string,
	target string, relation string, metadata *string) {
	assure("asset relationship", token)
	req := dmmodel.AssetRelationshipCreateRequest{
		Token:            token,
		SourceAsset:      source,
		TargetAsset:      target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureAssetRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AssetRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAssetRelationship.Token)
	}
}

// Assure an asset group (check for existing or create new).
func AssureAssetGroup(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("asset group", token)
	req := dmmodel.AssetGroupCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureAssetGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AssetGroupsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAssetGroup.Token)
	}
}

// Assure an asset group relationship type (check for existing or create new).
func AssureAssetGroupRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("asset group relationship type", token)
	req := dmmodel.AssetGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureAssetGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AssetGroupRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAssetGroupRelationshipType.Token)
	}
}

// Assure an asset group relationship (check for existing or create new).
func AssureAssetGroupRelationship(ctx context.Context, cli graphql.Client, token string, assetGroup string,
	asset string, relation string, metadata *string) {
	assure("asset group relationship", token)
	req := dmmodel.AssetGroupRelationshipCreateRequest{
		Token:            token,
		AssetGroup:       assetGroup,
		Asset:            asset,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureAssetGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AssetGroupRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAssetGroupRelationship.Token)
	}
}

// Assure an area type (check for existing or create new).
func AssureAreaType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("area type", token)
	req := dmmodel.AreaTypeCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureAreaType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AreaTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAreaType.Token)
	}
}

// Assure an area (check for existing or create new).
func AssureArea(ctx context.Context, cli graphql.Client, token string, areaTypeToken string, name *string,
	description *string, metadata *string) {
	assure("area", token)
	req := dmmodel.AreaCreateRequest{
		Token:         token,
		AreaTypeToken: areaTypeToken,
		Name:          name,
		Description:   description,
		Metadata:      metadata,
	}
	gresp, cresp, err := dmgql.AssureArea(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AreasByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateArea.Token)
	}
}

// Assure an area relationship type (check for existing or create new).
func AssureAreaRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("area relationship type", token)
	req := dmmodel.AreaRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureAreaRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AreaRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAreaRelationshipType.Token)
	}
}

// Assure an area relationship (check for existing or create new).
func AssureAreaRelationship(ctx context.Context, cli graphql.Client, token string, source string,
	target string, relation string, metadata *string) {
	assure("area relationship", token)
	req := dmmodel.AreaRelationshipCreateRequest{
		Token:            token,
		SourceArea:       source,
		TargetArea:       target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureAreaRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AreaRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAreaRelationship.Token)
	}
}

// Assure an area group (check for existing or create new).
func AssureAreaGroup(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("area group", token)
	req := dmmodel.AreaGroupCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureAreaGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AreaGroupsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAreaGroup.Token)
	}
}

// Assure an area group relationship type (check for existing or create new).
func AssureAreaGroupRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("area group relationship type", token)
	req := dmmodel.AreaGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureAreaGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AreaGroupRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAreaGroupRelationshipType.Token)
	}
}

// Assure an area group relationship (check for existing or create new).
func AssureAreaGroupRelationship(ctx context.Context, cli graphql.Client, token string, areaGroup string,
	area string, relation string, metadata *string) {
	assure("area group relationship", token)
	req := dmmodel.AreaGroupRelationshipCreateRequest{
		Token:            token,
		AreaGroup:        areaGroup,
		Area:             area,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureAreaGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.AreaGroupRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateAreaGroupRelationship.Token)
	}
}

// Assure a customer type (check for existing or create new).
func AssureCustomerType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("customer type", token)
	req := dmmodel.CustomerTypeCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureCustomerType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.CustomerTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateCustomerType.Token)
	}
}

// Assure a customer (check for existing or create new).
func AssureCustomer(ctx context.Context, cli graphql.Client, token string, customerTypeToken string, name *string,
	description *string, metadata *string) {
	assure("customer", token)
	req := dmmodel.CustomerCreateRequest{
		Token:             token,
		CustomerTypeToken: customerTypeToken,
		Name:              name,
		Description:       description,
		Metadata:          metadata,
	}
	gresp, cresp, err := dmgql.AssureCustomer(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.CustomersByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateCustomer.Token)
	}
}

// Assure a customer relationship type (check for existing or create new).
func AssureCustomerRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("customer relationship type", token)
	req := dmmodel.CustomerRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureCustomerRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.CustomerRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateCustomerRelationshipType.Token)
	}
}

// Assure a customer relationship (check for existing or create new).
func AssureCustomerRelationship(ctx context.Context, cli graphql.Client, token string, source string,
	target string, relation string, metadata *string) {
	assure("customer relationship", token)
	req := dmmodel.CustomerRelationshipCreateRequest{
		Token:            token,
		SourceCustomer:   source,
		TargetCustomer:   target,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureCustomerRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.CustomerRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateCustomerRelationship.Token)
	}
}

// Assure a customer group (check for existing or create new).
func AssureCustomerGroup(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, imageUrl *string, icon *string, backgroundColor *string, foregroundColor *string,
	borderColor *string, metadata *string) {
	assure("customer group", token)
	req := dmmodel.CustomerGroupCreateRequest{
		Token:           token,
		Name:            name,
		Description:     description,
		ImageUrl:        imageUrl,
		Icon:            icon,
		BackgroundColor: backgroundColor,
		ForegroundColor: foregroundColor,
		BorderColor:     borderColor,
		Metadata:        metadata,
	}
	gresp, cresp, err := dmgql.AssureCustomerGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.CustomerGroupsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateCustomerGroup.Token)
	}
}

// Assure a customer group relationship type (check for existing or create new).
func AssureCustomerGroupRelationshipType(ctx context.Context, cli graphql.Client, token string, name *string,
	description *string, metadata *string) {
	assure("customer group relationship type", token)
	req := dmmodel.CustomerGroupRelationshipTypeCreateRequest{
		Token:       token,
		Name:        name,
		Description: description,
		Metadata:    metadata,
	}
	gresp, cresp, err := dmgql.AssureCustomerGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.CustomerGroupRelationshipTypesByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateCustomerGroupRelationshipType.Token)
	}
}

// Assure a customer group relationship (check for existing or create new).
func AssureCustomerGroupRelationship(ctx context.Context, cli graphql.Client, token string, customerGroup string,
	customer string, relation string, metadata *string) {
	assure("customer group relationship", token)
	req := dmmodel.CustomerGroupRelationshipCreateRequest{
		Token:            token,
		CustomerGroup:    customerGroup,
		Customer:         customer,
		RelationshipType: relation,
		Metadata:         metadata,
	}
	gresp, cresp, err := dmgql.AssureCustomerGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if gresp != nil {
		found(gresp.CustomerGroupRelationshipsByToken[0].Token)
	}
	if cresp != nil {
		created(cresp.CreateCustomerGroupRelationship.Token)
	}
}
