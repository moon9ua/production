package response

import (
	"fmt"
	"main/internal/core/data"
	"main/internal/core/pb"
	"main/internal/pkg/logger"
	"time"
)

type ChartData struct {
	StartDate       time.Time `time_format:"2006-01-02T15:04:05.000Z" json:"start_date"`
	ProfitRateData  []float32 `json:"profit_rate_data"`
	ProfitKospiData []float32 `json:"profit_kospi_data"`
}

type QuantResponse struct {
	QuantID             uint      `json:"-"`
	CumulativeReturn    float64   `json:"cumulative_return"`
	AnnualAverageReturn float64   `json:"annual_average_return"`
	WinningPercentage   float64   `json:"winning_percentage"`
	MaxLossRate         float64   `json:"max_loss_rate"`
	HoldingsCount       int32     `json:"holdings_count"`
	ChartData           ChartData `json:"chart_data"`
}

func NewQuantResultFromPB(pb *pb.QuantResult) *QuantResponse {
	return &QuantResponse{
		CumulativeReturn:    pb.CumulativeReturn,
		AnnualAverageReturn: pb.AnnualAverageReturn,
		WinningPercentage:   pb.WinningPercentage,
		MaxLossRate:         pb.MaxLossRate,
		HoldingsCount:       pb.HoldingsCount,
		ChartData: ChartData{
			StartDate:      pb.ChartData.StartDate.AsTime(),
			ProfitRateData: pb.ChartData.ProfitRateData,
		},
	}
}

func (qr *QuantResponse) AddKospiData() error {
	var kospiVal []float32

	dataLen := len(qr.ChartData.ProfitRateData)
	startTime := qr.ChartData.StartDate
	idx := data.KospiData.Date[startTime]
	if idx < dataLen {
		return fmt.Errorf("error in AddKospiData got wrong data: 'idx':%d, 'dataLen':%d", idx, dataLen)
	}

	for i := idx; i >= idx-dataLen; i-- {
		if i < 0 {
			logger.Logger.Errorf("error in AddKospiData: got negative index: 'i': %d", idx)
			return fmt.Errorf("error in AddKospiData: got negative index: 'i': %d", idx)
		}
		kospiVal = append(kospiVal, data.KospiData.IndexVal[i])
	}
	qr.ChartData.ProfitKospiData = kospiVal

	return nil
}