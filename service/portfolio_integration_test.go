// +build integration

package service

import (
	"testing"

	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/dao"
	"github.com/jeremyhahn/tradebot/mapper"
	"github.com/stretchr/testify/assert"
)

func TestPortfolioService_Build(t *testing.T) {
	ctx := NewIntegrationTestContext()
	marketcapService := NewMarketCapService(ctx.GetLogger())

	pluginDAO := dao.NewPluginDAO(ctx)
	userDAO := dao.NewUserDAO(ctx)
	userMapper := mapper.NewUserMapper()
	userExchangeMapper := mapper.NewUserExchangeMapper()
	ethereumService, err := NewEthereumService(ctx, userDAO, userMapper)
	assert.Nil(t, err)
	userService := NewUserService(ctx, userDAO, pluginDAO, marketcapService, ethereumService, userMapper, userExchangeMapper)

	service := NewPortfolioService(ctx, marketcapService, userService, ethereumService)
	currencyPair := &common.CurrencyPair{
		Base:          "BTC",
		Quote:         "USD",
		LocalCurrency: "USD"}
	portfolio, err := service.Build(ctx.GetUser(), currencyPair)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), portfolio.GetUser().GetId())
	assert.Equal(t, true, len(portfolio.GetExchanges()) > 0)
	assert.Equal(t, true, len(portfolio.GetWallets()) > 0)
	assert.Equal(t, true, portfolio.GetNetWorth() > 0)
	CleanupIntegrationTest()
}

func TestPortfolioService_Stream(t *testing.T) {
	ctx := NewIntegrationTestContext()
	marketcapService := NewMarketCapService(ctx.GetLogger())

	pluginDAO := dao.NewPluginDAO(ctx)
	userDAO := dao.NewUserDAO(ctx)
	userMapper := mapper.NewUserMapper()
	userExchangeMapper := mapper.NewUserExchangeMapper()
	ethereumService, err := NewEthereumService(ctx, userDAO, userMapper)
	assert.Nil(t, err)
	userService := NewUserService(ctx, userDAO, pluginDAO, marketcapService, ethereumService, userMapper, userExchangeMapper)

	service := NewPortfolioService(ctx, marketcapService, userService, ethereumService)
	currencyPair := &common.CurrencyPair{
		Base:          "BTC",
		Quote:         "USD",
		LocalCurrency: "USD"}

	portfolioChan, err := service.Stream(ctx.GetUser(), currencyPair)
	assert.Nil(t, err)

	portfolio := <-portfolioChan
	service.Stop(ctx.GetUser())

	assert.Equal(t, uint(1), portfolio.GetUser().GetId())
	assert.Equal(t, true, len(portfolio.GetExchanges()) > 0)
	assert.Equal(t, true, len(portfolio.GetWallets()) > 0)
	assert.Equal(t, true, portfolio.GetNetWorth() > 0)
	CleanupIntegrationTest()
}
