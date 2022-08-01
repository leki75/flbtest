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
