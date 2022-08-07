package main

import (
	"testing"

	karmemgo "karmem.org/golang"

	"leki75/flbtest/schema/karmem"
)

func BenchmarkKarmemQuoteMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quoteBatch()
	}
}

func BenchmarkKarmemTradeMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tradeBatch()
	}
}

func BenchmarkKarmemQuoteRead(b *testing.B) {
	qb := quoteBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader := karmemgo.NewReader(qb)
		quotes := karmem.NewQuoteBatchViewer(reader, 0)
		for _, rawQuote := range quotes.Quotes(reader) {
			quote := rawQuote.Data(reader)
			_ = quote.Symbol(reader)
			_ = quote.Conditions(reader)
			_ = quote.AskExchange()
			_ = quote.AskPrice()
			_ = quote.AskSize()
			_ = quote.BidExchange()
			_ = quote.BidPrice()
			_ = quote.BidSize()
			_ = quote.Timestamp()
			_ = quote.Nbbo()
			_ = quote.Tape()
			_ = quote.ReceivedAt()
		}
	}
}

func BenchmarkKarmemTradeRead(b *testing.B) {
	qb := tradeBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader := karmemgo.NewReader(qb)
		trades := karmem.NewTradeBatchViewer(reader, 0)
		for _, rawTrade := range trades.Trades(reader) {
			trade := rawTrade.Data(reader)
			_ = trade.Symbol(reader)
			_ = trade.Conditions(reader)
			_ = trade.ID()
			_ = trade.Timestamp()
			_ = trade.ReceivedAt()
			_ = trade.Price()
			_ = trade.Volume()
			_ = trade.Exchange()
			_ = trade.Tape()
		}
	}
}

func BenchmarkKarmemQuoteUnmarshal(b *testing.B) {
	qb := quoteBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := &karmem.QuoteBatch{}
		batch.ReadAsRoot(karmemgo.NewReader(qb))
		for _, rawQuote := range batch.Quotes {
			_ = rawQuote.Data
		}
	}
}

func BenchmarkKarmemTradeUnmarshal(b *testing.B) {
	tb := tradeBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := &karmem.TradeBatch{}
		batch.ReadAsRoot(karmemgo.NewReader(tb))
		for _, rawTrade := range batch.Trades {
			_ = rawTrade.Data
		}
	}
}
