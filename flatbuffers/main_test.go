package main

import (
	"testing"

	"leki75/flbtest/schema/flatbuf"
)

func TestFlatbuffersQuoteBatch(t *testing.T) {
	batch := flatbuf.GetRootAsQuoteBatch(quoteBatch(), 0)
	{ // quote1
		quote := new(flatbuf.Quote)
		if !batch.Quotes(quote, 0) {
			t.Error("failed to get quote 1")
		}
		if string(quote.Symbol()) != "TSLA" {
			t.Error("symbol name mismatch on quote 1", string(quote.Symbol()))
		}
	}
	{ // quote2
		quote := new(flatbuf.Quote)
		if !batch.Quotes(quote, 1) {
			t.Error("failed to get quote 2")
		}
		if string(quote.Symbol()) != "AAPL" {
			t.Error("symbol name mismatch on quote 2", string(quote.Symbol()))
		}
	}
}

func TestFlatbuffersRawQuoteBatch(t *testing.T) {
	batch := flatbuf.GetRootAsRawQuoteBatch(rawQuoteBatch(), 0)
	{ // quote1
		rawQuote := new(flatbuf.RawQuote)
		if !batch.RawQuotes(rawQuote, 0) {
			t.Error("failed to get quote 1")
		}
		quote := flatbuf.GetRootAsQuote(rawQuote.DataBytes(), 0)
		if string(quote.Symbol()) != "TSLA" {
			t.Error("symbol name mismatch on quote 1", string(quote.Symbol()))
		}
	}
	{ // quote2
		rawQuote := new(flatbuf.RawQuote)
		if !batch.RawQuotes(rawQuote, 1) {
			t.Error("failed to get quote 2")
		}
		quote := flatbuf.GetRootAsQuote(rawQuote.DataBytes(), 0)
		if string(quote.Symbol()) != "AAPL" {
			t.Error("symbol name mismatch on quote 2", string(quote.Symbol()))
		}
	}
}

func BenchmarkFlatbuffersQuoteMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quoteBatch()
	}
}

func BenchmarkFlatbuffersQuoteUnmarshal(b *testing.B) {
	quote := new(flatbuf.Quote)
	qb := quoteBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := flatbuf.GetRootAsQuoteBatch(qb, 0)

		batch.Quotes(quote, 0)
		_ = quote.Symbol()
		_ = quote.AskExchange()
		_ = quote.AskPrice()
		_ = quote.AskSize()
		_ = quote.BidExchange()
		_ = quote.BidPrice()
		_ = quote.BidSize()
		_ = quote.Timestamp()
		_ = quote.ConditionsBytes()
		_ = quote.Nbbo()
		_ = quote.Tape()
		_ = quote.ReceivedAt()

		batch.Quotes(quote, 1)
		_ = quote.Symbol()
		_ = quote.AskExchange()
		_ = quote.AskPrice()
		_ = quote.AskSize()
		_ = quote.BidExchange()
		_ = quote.BidPrice()
		_ = quote.BidSize()
		_ = quote.Timestamp()
		_ = quote.ConditionsBytes()
		_ = quote.Nbbo()
		_ = quote.Tape()
		_ = quote.ReceivedAt()
	}
}

func BenchmarkFlatbuffersRawQuoteMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rawQuoteBatch()
	}
}

func BenchmarkFlatbuffersRawQuoteUnmarshal(b *testing.B) {
	var quote *flatbuf.Quote
	rawQuote := new(flatbuf.RawQuote)
	qb := rawQuoteBatch()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := flatbuf.GetRootAsRawQuoteBatch(qb, 0)

		batch.RawQuotes(rawQuote, 0)
		quote = flatbuf.GetRootAsQuote(rawQuote.DataBytes(), 0)
		_ = quote.Symbol()
		_ = quote.AskExchange()
		_ = quote.AskPrice()
		_ = quote.AskSize()
		_ = quote.BidExchange()
		_ = quote.BidPrice()
		_ = quote.BidSize()
		_ = quote.Timestamp()
		_ = quote.ConditionsBytes()
		_ = quote.Nbbo()
		_ = quote.Tape()
		_ = quote.ReceivedAt()

		batch.RawQuotes(rawQuote, 1)
		quote = flatbuf.GetRootAsQuote(rawQuote.DataBytes(), 0)
		_ = quote.Symbol()
		_ = quote.AskExchange()
		_ = quote.AskPrice()
		_ = quote.AskSize()
		_ = quote.BidExchange()
		_ = quote.BidPrice()
		_ = quote.BidSize()
		_ = quote.Timestamp()
		_ = quote.ConditionsBytes()
		_ = quote.Nbbo()
		_ = quote.Tape()
		_ = quote.ReceivedAt()
	}
}
