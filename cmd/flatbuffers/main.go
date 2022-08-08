package main

import (
	"fmt"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"

	"leki75/flbtest/schema/flatbuf"
)

var now = uint64(time.Now().UnixNano())

func main() {
	fmt.Println("FlatbuffersQuoteBatch:", len(quoteBatch()))
	fmt.Println("FlatbuffersRawQuoteBatch:", len(rawQuoteBatch()))
}

func quoteBatch() []byte {
	var quote1, quote2 flatbuffers.UOffsetT
	builder := flatbuffers.NewBuilder(256)

	{ // quote1
		sym := builder.CreateString("AAPL")
		cond := builder.CreateByteVector([]byte{'C', 'D'})
		flatbuf.QuoteStart(builder)
		flatbuf.QuoteAddSymbol(builder, sym)
		flatbuf.QuoteAddBidExchange(builder, 'A')
		flatbuf.QuoteAddBidPrice(builder, 100.0)
		flatbuf.QuoteAddBidSize(builder, 100)
		flatbuf.QuoteAddAskExchange(builder, 'B')
		flatbuf.QuoteAddAskPrice(builder, 200.0)
		flatbuf.QuoteAddAskSize(builder, 200)
		flatbuf.QuoteAddTimestamp(builder, now)
		flatbuf.QuoteAddConditions(builder, cond)
		flatbuf.QuoteAddNbbo(builder, true)
		flatbuf.QuoteAddTape(builder, 'E')
		flatbuf.QuoteAddReceivedAt(builder, now)
		quote1 = flatbuf.QuoteEnd(builder)
	}
	{ // quote2
		sym := builder.CreateString("TSLA")
		cond := builder.CreateByteVector([]byte{'H', 'I'})
		flatbuf.QuoteStart(builder)
		flatbuf.QuoteAddSymbol(builder, sym)
		flatbuf.QuoteAddBidExchange(builder, 'F')
		flatbuf.QuoteAddBidPrice(builder, 300.0)
		flatbuf.QuoteAddBidSize(builder, 300)
		flatbuf.QuoteAddAskExchange(builder, 'G')
		flatbuf.QuoteAddAskPrice(builder, 400.0)
		flatbuf.QuoteAddAskSize(builder, 400)
		flatbuf.QuoteAddTimestamp(builder, now)
		flatbuf.QuoteAddConditions(builder, cond)
		flatbuf.QuoteAddNbbo(builder, true)
		flatbuf.QuoteAddTape(builder, 'J')
		flatbuf.QuoteAddReceivedAt(builder, now)
		quote2 = flatbuf.QuoteEnd(builder)
	}

	flatbuf.QuoteBatchStartQuotesVector(builder, 2)
	builder.PrependUOffsetT(quote1)
	builder.PrependUOffsetT(quote2)
	quotes := builder.EndVector(2)

	flatbuf.QuoteBatchStart(builder)
	flatbuf.QuoteBatchAddQuotes(builder, quotes)
	batch := flatbuf.QuoteBatchEnd(builder)

	builder.Finish(batch)
	return builder.FinishedBytes()
}

func rawQuoteBatch() []byte {
	var quote1, quote2 []byte

	{ // quote1
		builder := flatbuffers.NewBuilder(128)
		sym := builder.CreateString("AAPL")
		cond := builder.CreateByteVector([]byte{'C', 'D'})
		flatbuf.QuoteStart(builder)
		flatbuf.QuoteAddSymbol(builder, sym)
		flatbuf.QuoteAddBidExchange(builder, 'A')
		flatbuf.QuoteAddBidPrice(builder, 100.0)
		flatbuf.QuoteAddBidSize(builder, 100)
		flatbuf.QuoteAddAskExchange(builder, 'B')
		flatbuf.QuoteAddAskPrice(builder, 200.0)
		flatbuf.QuoteAddAskSize(builder, 200)
		flatbuf.QuoteAddTimestamp(builder, now)
		flatbuf.QuoteAddConditions(builder, cond)
		flatbuf.QuoteAddNbbo(builder, true)
		flatbuf.QuoteAddTape(builder, 'E')
		flatbuf.QuoteAddReceivedAt(builder, now)
		builder.Finish(flatbuf.QuoteEnd(builder))
		quote1 = builder.FinishedBytes()
	}
	{ // quote2
		builder := flatbuffers.NewBuilder(128)
		sym := builder.CreateString("TSLA")
		cond := builder.CreateByteVector([]byte{'H', 'I'})
		flatbuf.QuoteStart(builder)
		flatbuf.QuoteAddSymbol(builder, sym)
		flatbuf.QuoteAddBidExchange(builder, 'F')
		flatbuf.QuoteAddBidPrice(builder, 300.0)
		flatbuf.QuoteAddBidSize(builder, 300)
		flatbuf.QuoteAddAskExchange(builder, 'G')
		flatbuf.QuoteAddAskPrice(builder, 400.0)
		flatbuf.QuoteAddAskSize(builder, 400)
		flatbuf.QuoteAddTimestamp(builder, now)
		flatbuf.QuoteAddConditions(builder, cond)
		flatbuf.QuoteAddNbbo(builder, true)
		flatbuf.QuoteAddTape(builder, 'J')
		flatbuf.QuoteAddReceivedAt(builder, now)
		builder.Finish(flatbuf.QuoteEnd(builder))
		quote2 = builder.FinishedBytes()
	}

	builder := flatbuffers.NewBuilder(256)
	rawQuotes := make([]flatbuffers.UOffsetT, 0, 2)
	{ // quote1
		data := builder.CreateByteVector(quote1)
		flatbuf.RawQuoteStart(builder)
		flatbuf.RawQuoteAddData(builder, data)
		rawQuotes = append(rawQuotes, flatbuf.RawQuoteEnd(builder))
	}
	{ // quote2
		data := builder.CreateByteVector(quote2)
		flatbuf.RawQuoteStart(builder)
		flatbuf.RawQuoteAddData(builder, data)
		rawQuotes = append(rawQuotes, flatbuf.RawQuoteEnd(builder))
	}

	flatbuf.RawQuoteBatchStartRawQuotesVector(builder, 2)
	builder.PrependUOffsetT(rawQuotes[0])
	builder.PrependUOffsetT(rawQuotes[1])
	quotes := builder.EndVector(len(rawQuotes))

	flatbuf.RawQuoteBatchStart(builder)
	flatbuf.RawQuoteBatchAddRawQuotes(builder, quotes)
	batch := flatbuf.RawQuoteBatchEnd(builder)

	builder.Finish(batch)
	return builder.FinishedBytes()
}
