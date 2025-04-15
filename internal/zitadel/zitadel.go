package zitadel

import (
	"context"
	"github.com/materials-resources/store-api/app"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/client"
	managementClient "github.com/zitadel/zitadel-go/v3/pkg/client/management"
	"github.com/zitadel/zitadel-go/v3/pkg/client/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel"
	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/management"
)

type Client struct {
	management *managementClient.Client
}

func NewZitadelClient(a *app.App) (*Client, error) {
	ctx := context.Background()

	cli, err := managementClient.NewClient(ctx, a.Config.Zitadel.Issuer, a.Config.Zitadel.ApiUrl, []string{
		oidc.ScopeOpenID, client.ScopeZitadelAPI(),
	}, zitadel.WithJWTProfileTokenSource(middleware.JWTProfileFromPath(ctx, a.Config.Zitadel.JwtPath)), zitadel.WithOrgID(a.Config.Zitadel.OrgId))

	if err != nil {
		return nil, err
	}

	return &Client{
		management: cli,
	}, nil

}

func (c *Client) ChangeUserBranchId(ctx context.Context, userId, branchId string) error {

	_, err := c.management.SetUserMetadata(ctx, &management.SetUserMetadataRequest{
		Id:    userId,
		Key:   "branch_id",
		Value: []byte(branchId),
	})

	return err
}
