package main

import (
	"testing"

	"leki75/flbtest/schema/flatbuf"
)

func TestQuoteBatch(t *testing.T) {
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

func TestRawQuoteBatch(t *testing.T) {
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

func BenchmarkQuoteMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quoteBatch()
	}
}

func BenchmarkQuoteMarshalUnmarshal(b *testing.B) {
	quote := new(flatbuf.Quote)
	for i := 0; i < b.N; i++ {
		batch := flatbuf.GetRootAsQuoteBatch(quoteBatch(), 0)
		batch.Quotes(quote, 0)
		_ = quote.Symbol()
		batch.Quotes(quote, 1)
		_ = quote.Symbol()
	}
}

func BenchmarkRawQuoteMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quoteBatch()
	}
}

func BenchmarkRawQuoteMarhsalUnmarshal(b *testing.B) {
	var quote *flatbuf.Quote
	rawQuote := new(flatbuf.RawQuote)
	for i := 0; i < b.N; i++ {
		batch := flatbuf.GetRootAsRawQuoteBatch(rawQuoteBatch(), 0)
		batch.RawQuotes(rawQuote, 0)
		quote = flatbuf.GetRootAsQuote(rawQuote.DataBytes(), 0)
		_ = quote.Symbol()
		batch.RawQuotes(rawQuote, 1)
		quote = flatbuf.GetRootAsQuote(rawQuote.DataBytes(), 0)
		_ = quote.Symbol()
	}
}
