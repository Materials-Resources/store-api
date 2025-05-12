package zitadel

import (
	"context"
	"github.com/materials-resources/store-api/app"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/client"
	managementClient "github.com/zitadel/zitadel-go/v3/pkg/client/management"
	"github.com/zitadel/zitadel-go/v3/pkg/client/middleware"
	userClient "github.com/zitadel/zitadel-go/v3/pkg/client/user/v2"
	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel"
	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/management"
	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/user/v2"
)

type Client struct {
	management *managementClient.Client
	user       *userClient.Client
}

func NewZitadelClient(a *app.App) (*Client, error) {
	ctx := context.Background()

	cli, err := managementClient.NewClient(ctx, a.Config.Zitadel.Issuer, a.Config.Zitadel.ApiUrl, []string{
		oidc.ScopeOpenID, client.ScopeZitadelAPI(),
	}, zitadel.WithJWTProfileTokenSource(middleware.JWTProfileFromPath(ctx, a.Config.Zitadel.JwtPath)))

	if err != nil {
		return nil, err
	}

	userCl, err := userClient.NewClient(ctx, a.Config.Zitadel.Issuer, a.Config.Zitadel.ApiUrl, []string{
		oidc.ScopeOpenID, client.ScopeZitadelAPI(),
	}, zitadel.WithJWTProfileTokenSource(middleware.JWTProfileFromPath(ctx, a.Config.Zitadel.JwtPath)), zitadel.WithOrgID(a.Config.Zitadel.OrgId))
	if err != nil {
		return nil, err
	}
	return &Client{
		management: cli,
		user:       userCl,
	}, nil

}

func (c *Client) ChangeUserBranchId(ctx context.Context, userId, branchId string) error {
	// Get user details
	userResp, err := c.user.GetUserByID(ctx, &user.GetUserByIDRequest{
		UserId: userId,
	})
	fmt.Println(err)
	if err != nil {
		return err
	}

	_, err = c.management.SetUserMetadata(middleware.SetOrgID(ctx, userResp.GetDetails().GetResourceOwner()), &management.SetUserMetadataRequest{
		Id:    userId,
		Key:   "branch_id",
		Value: []byte(branchId),
	})

	return err
}

// IsUserActive checks if a user is active in Zitadel
func (c *Client) IsUserActive(ctx context.Context, userId string) (bool, error) {
	resp, err := c.user.GetUserByID(ctx, &user.GetUserByIDRequest{
		UserId: userId,
	})
	if err != nil {
		return false, err
	}

	// Check if user state is active (1)
	// Zitadel user states: 0=unspecified, 1=active, 2=inactive
	return resp.User.State == user.UserState_USER_STATE_ACTIVE, nil
}
