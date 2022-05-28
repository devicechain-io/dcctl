/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package cmd

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	gql "github.com/devicechain-io/dcctl/graphql"
	"github.com/spf13/cobra"
)

const (
	DATASET_CONSTRUCTION = "Construction"
)

// Create common command for creating DeviceChain resources
var constructionDataCmd = &cobra.Command{
	Use:   "construction",
	Short: "Bootstrap construction sample data",
	Long:  `Bootstraps system microservices with sample data for construction use case`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return bootstrapConstructionData(context.Background(), cmd)
	},
	SilenceUsage: true,
}

// Bootstraps system microservices with construction sample dataset.
func bootstrapConstructionData(ctx context.Context, cmd *cobra.Command) error {
	title(DATASET_CONSTRUCTION)
	gqlcli := gql.GetDeviceManagementGraphQLClient(cmd)

	header("Device Types", DATASET_CONSTRUCTION)
	bootstrapDeviceTypes(ctx, gqlcli)

	header("Devices", DATASET_CONSTRUCTION)
	bootstrapDevices(ctx, gqlcli)

	header("Device Relationship Types", DATASET_CONSTRUCTION)
	bootstrapDeviceRelationshipTypes(ctx, gqlcli)

	header("Device Relationships", DATASET_CONSTRUCTION)
	bootstrapDeviceRelationships(ctx, gqlcli)

	header("Device Groups", DATASET_CONSTRUCTION)
	bootstrapDeviceGroups(ctx, gqlcli)

	header("Device Group Relationship Types", DATASET_CONSTRUCTION)
	bootstrapDeviceGroupRelationshipTypes(ctx, gqlcli)

	header("Device Group Relationships", DATASET_CONSTRUCTION)
	bootstrapDeviceGroupRelationships(ctx, gqlcli)

	footer(DATASET_CONSTRUCTION)
	return nil
}

// Bootstrap device types.
func bootstrapDeviceTypes(ctx context.Context, client graphql.Client) error {
	// Cat D1
	gql.AssureDeviceType(ctx, client, "catd1", s("Cat D1"),
		unspace(`The new Cat® D1 delivers superior performance and the broadest choice of technology features to 
		help you get the most from your dozer. Nimble and responsive, it has power for dozing and ﬁnesse for 
		grading. Fully hydrostatic transmission gives you seamless acceleration, so you can get the job done quickly. 
		The load sensing system automatically optimizes ground speed based on load, for the greatest productivity and 
		fuel efficiency.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"engineModel": "Cat C3.6",
			"powerNet": "80 HP",
			"operatingWeight": "17855 lb"
		}`))
	return nil
}

// Bootstrap devices.
func bootstrapDevices(ctx context.Context, client graphql.Client) error {
	// Cat D1 SDK7GV3WXZ3FBXZ
	gql.AssureDevice(ctx, client, "SDK7GV3WXZ3FBXZ", "catd1", s("Cat D1 VIN:SDK7GV3WXZ3FBXZ"),
		unspace(`This is a Cat D1 with VIN SDK7GV3WXZ3FBXZ`),
		unspace(`
		{
			"vin": "SDK7GV3WXZ3FBXZ",
			"owner": "CatCorp",
			"purchaseDate": "2022/01/01"
		}`))
	// Cat D1 WDVM4L7YPRM7HU2
	gql.AssureDevice(ctx, client, "WDVM4L7YPRM7HU2", "catd1", s("Cat D1 VIN:WDVM4L7YPRM7HU2"),
		unspace(`This is a Cat D1 with VIN WDVM4L7YPRM7HU2`),
		unspace(`
		{
			"vin": "WDVM4L7YPRM7HU2",
			"owner": "CatCorp",
			"purchaseDate": "2022/02/01"
		}`))
	return nil
}

// Bootstrap device relationship types.
func bootstrapDeviceRelationshipTypes(ctx context.Context, client graphql.Client) error {
	// Tracks location of
	gql.AssureDeviceRelationshipType(ctx, client, "tracksLocationOf", s("Tracks location of"),
		unspace(`The source device tracks the location of the target device`),
		unspace(`
		{
			"accuracy": "1 meter"
		}`))
	// Tracks temperature of
	gql.AssureDeviceRelationshipType(ctx, client, "tracksTempOf", s("Tracks temperature of"),
		unspace(`The source device tracks the temperature of the target device`),
		unspace(`
		{
			"accuracy": "1 degree C"
		}`))
	return nil
}

// Bootstrap device relationships.
func bootstrapDeviceRelationships(ctx context.Context, client graphql.Client) error {
	// SDK7GV3WXZ3FBXZ tracksLocationOf WDVM4L7YPRM7HU2
	gql.AssureDeviceRelationship(ctx, client, "SDK7GV3WXZ3FBXZ-tracksLocationOf-WDVM4L7YPRM7HU2",
		"SDK7GV3WXZ3FBXZ", "WDVM4L7YPRM7HU2", "tracksLocationOf",
		unspace(`
		{
			"accuracy": "1 meter"
		}`))
	return nil
}

// Bootstrap device groups.
func bootstrapDeviceGroups(ctx context.Context, client graphql.Client) error {
	// Small Dozers
	gql.AssureDeviceGroup(ctx, client, "smalldoz", s("Small Dozers"),
		unspace(`Under 105 hp, the Cat® small dozers are designed to optimize speed, transportability, maneuverability, 
		versatility and finish grading accuracy. These crawler dozers are ideal for residential construction performing 
		such tasks as clearing and grading lots, sloping the sides of roads, back-filling, and final grade work for 
		landscaping and driveway construction.`),
		s("https://devicechain.s3.amazonaws.com/datasets/construction/catd1.jpg"), nil, nil, nil, nil,
		unspace(`
		{
			"maxWeight": "20000 lb"
		}`))
	return nil
}

// Bootstrap device group relationship types.
func bootstrapDeviceGroupRelationshipTypes(ctx context.Context, client graphql.Client) error {
	// Tracks location of
	gql.AssureDeviceGroupRelationshipType(ctx, client, "contains", s("Contains"),
		unspace(`The group contains the target device`), nil)
	return nil
}

// Bootstrap device group relationships.
func bootstrapDeviceGroupRelationships(ctx context.Context, client graphql.Client) error {
	// smalldoz contains SDK7GV3WXZ3FBXZ
	gql.AssureDeviceGroupRelationship(ctx, client, "smalldoz-contains-SDK7GV3WXZ3FBXZ",
		"smalldoz", "SDK7GV3WXZ3FBXZ", "contains", nil)
	// smalldoz contains WDVM4L7YPRM7HU2
	gql.AssureDeviceGroupRelationship(ctx, client, "smalldoz-contains-WDVM4L7YPRM7HU2",
		"smalldoz", "WDVM4L7YPRM7HU2", "contains", nil)
	return nil
}

func init() {
	bootstrapCmd.AddCommand(constructionDataCmd)
}
