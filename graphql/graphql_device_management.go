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
	resp, wascreated, err := dmgql.AssureDeviceType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device types by token.
func GetDeviceTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IDeviceType, error) {
	return dmgql.GetDeviceTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureDevice(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get devices by token.
func GetDevicesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IDevice, error) {
	return dmgql.GetDevicesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureDeviceRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device relationship types by token.
func GetDeviceRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IDeviceRelationshipType, error) {
	return dmgql.GetDeviceRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureDeviceRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device relationships by token.
func GetDeviceRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IDeviceRelationship, error) {
	return dmgql.GetDeviceRelationshipsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureDeviceGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device groups by token.
func GetDeviceGroupsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IDeviceGroup, error) {
	return dmgql.GetDeviceGroupsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureDeviceGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device group relationship types by token.
func GetDeviceGroupRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IDeviceGroupRelationshipType, error) {
	return dmgql.GetDeviceGroupRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureDeviceGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get device group relationships by token.
func GetDeviceGroupRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IDeviceGroupRelationship, error) {
	return dmgql.GetDeviceGroupRelationshipsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAssetType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset types by token.
func GetAssetTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAssetType, error) {
	return dmgql.GetAssetTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAsset(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get assets by token.
func GetAssetsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAsset, error) {
	return dmgql.GetAssetsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAssetRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset relationship types by token.
func GetAssetRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAssetRelationshipType, error) {
	return dmgql.GetAssetRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAssetRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset relationships by token.
func GetAssetRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAssetRelationship, error) {
	return dmgql.GetAssetRelationshipsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAssetGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset groups by token.
func GetAssetGroupsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAssetGroup, error) {
	return dmgql.GetAssetGroupsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAssetGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset group relationship types by token.
func GetAssetGroupRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAssetGroupRelationshipType, error) {
	return dmgql.GetAssetGroupRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAssetGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get asset group relationships by token.
func GetAssetGroupRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAssetGroupRelationship, error) {
	return dmgql.GetAssetGroupRelationshipsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAreaType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area types by token.
func GetAreaTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAreaType, error) {
	return dmgql.GetAreaTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureArea(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get areas by token.
func GetAreasByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IArea, error) {
	return dmgql.GetAreasByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAreaRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area relationship types by token.
func GetAreaRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAreaRelationshipType, error) {
	return dmgql.GetAreaRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAreaRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area relationships by token.
func GetAreaRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAreaRelationship, error) {
	return dmgql.GetAreaRelationshipsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAreaGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area groups by token.
func GetAreaGroupsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAreaGroup, error) {
	return dmgql.GetAreaGroupsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAreaGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area group relationship types by token.
func GetAreaGroupRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAreaGroupRelationshipType, error) {
	return dmgql.GetAreaGroupRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureAreaGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get area group relationships by token.
func GetAreaGroupRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.IAreaGroupRelationship, error) {
	return dmgql.GetAreaGroupRelationshipsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureCustomerType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer types by token.
func GetCustomerTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.ICustomerType, error) {
	return dmgql.GetCustomerTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureCustomer(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customers by token.
func GetCustomersByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.ICustomer, error) {
	return dmgql.GetCustomersByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureCustomerRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer relationship types by token.
func GetCustomerRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.ICustomerRelationshipType, error) {
	return dmgql.GetCustomerRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureCustomerRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer relationships by token.
func GetCustomerRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.ICustomerRelationship, error) {
	return dmgql.GetCustomerRelationshipsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureCustomerGroup(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer groups by token.
func GetCustomerGroupsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.ICustomerGroup, error) {
	return dmgql.GetCustomerGroupsByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureCustomerGroupRelationshipType(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer group relationship types by token.
func GetCustomerGroupRelationshipTypesByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.ICustomerGroupRelationshipType, error) {
	return dmgql.GetCustomerGroupRelationshipTypesByToken(ctx, cli, tokens)
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
	resp, wascreated, err := dmgql.AssureCustomerGroupRelationship(ctx, cli, req)
	if err != nil {
		panic(err)
	}
	if wascreated {
		created(resp.GetToken())
	} else {
		found(resp.GetToken())
	}
}

// Get customer group relationships by token.
func GetCustomerGroupRelationshipsByToken(ctx context.Context, cli graphql.Client, tokens []string) (map[string]dmgql.ICustomerGroupRelationship, error) {
	return dmgql.GetCustomerGroupRelationshipsByToken(ctx, cli, tokens)
}
