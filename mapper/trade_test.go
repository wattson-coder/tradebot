package mapper

import (
	"testing"
	"time"

	"github.com/jeremyhahn/tradebot/dto"
	"github.com/stretchr/testify/assert"
)

func TestTradeMapper(t *testing.T) {
	mapper := NewTradeMapper()
	dto := &dto.TradeDTO{
		Id:        1,
		ChartId:   1,
		UserId:    1,
		Base:      "BTC",
		Quote:     "USD",
		Exchange:  "Test",
		Date:      time.Now(),
		Type:      "buy",
		Price:     10000.0,
		Amount:    2.5,
		ChartData: "{}"}

	entity := mapper.MapTradeDtoToEntity(dto)
	assert.NotNil(t, entity)
	assert.Equal(t, dto.GetId(), entity.GetId())
	assert.Equal(t, dto.GetChartId(), entity.GetChartId())
	assert.Equal(t, dto.GetUserId(), entity.GetUserId())
	assert.Equal(t, dto.GetBase(), entity.GetBase())
	assert.Equal(t, dto.GetQuote(), entity.GetQuote())
	assert.Equal(t, dto.GetExchange(), entity.GetExchangeName())
	assert.Equal(t, dto.GetDate(), entity.GetDate())
	assert.Equal(t, dto.GetType(), entity.GetType())
	assert.Equal(t, dto.GetPrice(), entity.GetPrice())
	assert.Equal(t, dto.GetAmount(), entity.GetAmount())
	assert.Equal(t, dto.GetChartData(), entity.GetChartData())

	mappedDTO := mapper.MapTradeEntityToDto(entity)
	assert.NotNil(t, entity)
	assert.Equal(t, entity.GetId(), mappedDTO.GetId())
	assert.Equal(t, entity.GetChartId(), mappedDTO.GetChartId())
	assert.Equal(t, entity.GetUserId(), mappedDTO.GetUserId())
	assert.Equal(t, entity.GetBase(), mappedDTO.GetBase())
	assert.Equal(t, entity.GetQuote(), mappedDTO.GetQuote())
	assert.Equal(t, entity.GetExchangeName(), mappedDTO.GetExchange())
	assert.Equal(t, entity.GetDate(), mappedDTO.GetDate())
	assert.Equal(t, entity.GetType(), mappedDTO.GetType())
	assert.Equal(t, entity.GetPrice(), mappedDTO.GetPrice())
	assert.Equal(t, entity.GetAmount(), mappedDTO.GetAmount())
	assert.Equal(t, entity.GetChartData(), mappedDTO.GetChartData())

}
