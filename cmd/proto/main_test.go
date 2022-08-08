package main

import (
	"testing"

	"leki75/flbtest/schema/proto"
)

func TestProtoQuoteBatch(t *testing.T) {
	qb := quoteBatch()
	batch := &proto.QuoteBatch{}
	if err := batch.UnmarshalVT(qb); err != nil {
		panic(err)
	}

	if len(batch.Quotes) != 2 {
		t.Error("invalid number of quotes")
	}

	{ // quote1
		quote := batch.Quotes[0]
		if string(quote.Symbol) != "AAPL" {
			t.Error("symbol name mismatch on quote 1", string(quote.Symbol))
		}
		if quote.BidPrice != 100.0 {
			t.Error("size mismatch on quote 1", quote.BidPrice)
		}
		if quote.AskSize != 200 {
			t.Error("size mismatch on quote 1", quote.AskSize)
		}
	}
	{ // quote2
		quote := batch.Quotes[1]
		if string(quote.Symbol) != "TSLA" {
			t.Error("symbol name mismatch on quote 2", quote.Symbol)
		}
	}
}

func BenchmarkProtoQuoteMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quoteBatch()
	}
}

func BenchmarkProtoTradeMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tradeBatch()
	}
}

func BenchmarkProtoQuoteUnmarshal(b *testing.B) {
	qb := quoteBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := &proto.QuoteBatch{}
		if err := batch.UnmarshalVT(qb); err != nil {
			panic(err)
		}
		for _, quote := range batch.Quotes {
			_ = quote
		}
	}
}

func BenchmarkProtoTradeUnmarshal(b *testing.B) {
	tb := tradeBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := &proto.TradeBatch{}
		if err := batch.UnmarshalVT(tb); err != nil {
			panic(err)
		}
		for _, trade := range batch.Trades {
			_ = trade
		}
	}
}
