package main

import (
	"testing"

	karmemgo "karmem.org/golang"

	"leki75/flbtest/schema/karmem"
)

func TestKarmemQuoteBatch(t *testing.T) {
	reader := karmemgo.NewReader(quoteBatch())
	batch := karmem.NewQuoteBatchViewer(reader, 0).Quotes(reader)

	if len(batch) != 2 {
		t.Error("invalid number of quotes")
	}

	{ // quote1
		aapl := [11]byte{'A', 'A', 'P', 'L'}
		quote := batch[0]
		if string(quote.Symbol()) != string(aapl[:]) {
			t.Error("symbol name mismatch on quote 1", string(quote.Symbol()))
		}
		if quote.BidPrice() != 100.0 {
			t.Error("size mismatch on quote 1", quote.AskSize())
		}
		if quote.AskSize() != 200 {
			t.Error("size mismatch on quote 1", quote.AskSize())
		}
	}
	{ // quote2
		tsla := [11]byte{'T', 'S', 'L', 'A'}
		quote := batch[1]
		if string(quote.Symbol()) != string(tsla[:]) {
			t.Error("symbol name mismatch on quote 2", string(quote.Symbol()))
		}
	}
}

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
		for _, quote := range quotes.Quotes(reader) {
			_ = quote.Symbol()
			_ = quote.Conditions()
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
	tb := tradeBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader := karmemgo.NewReader(tb)
		trades := karmem.NewTradeBatchViewer(reader, 0)
		for _, trade := range trades.Trades(reader) {
			_ = trade.Symbol()
			_ = trade.Conditions()
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
		for _, quote := range batch.Quotes {
			_ = quote
		}
	}
}

func BenchmarkKarmemTradeUnmarshal(b *testing.B) {
	tb := tradeBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := &karmem.TradeBatch{}
		batch.ReadAsRoot(karmemgo.NewReader(tb))
		for _, trade := range batch.Trades {
			_ = trade
		}
	}
}
