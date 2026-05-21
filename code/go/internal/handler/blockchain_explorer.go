package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
)

type blockchainHandler struct {
	svc service.Service
}

var BlockChainHandler = new(blockchainHandler)

func (c *blockchainHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "blockchain_explorer_index.html").WriteTpl(gin.H{
		"BlockHeight":        156085,
		"TotalTransactions":  2090,
		"OrganizationCount":  3,
		"PeerCount":          6,
		"RecentTransactions": []map[string]string{{"ID": "af31c...2191b", "Time": "2025-01-22 23:17:12"}, {"ID": "181c6...278a2", "Time": "2025-01-21 19:01:03"}, {"ID": "181d0...ce5ad", "Time": "2025-01-20 16:32:03"}, {"ID": "a41cd...103ef", "Time": "2025-01-20 14:01:03"}, {"ID": "1f92d...7dd23", "Time": "2025-01-20 9:01:03"}, {"ID": "f6502...f0982", "Time": "2025-01-19 23:01:03"}, {"ID": "eeffa...08b92", "Time": "2025-01-19 21:01:03"}, {"ID": "ed0b1...8be48", "Time": "2025-01-19 14:01:03"}, {"ID": "9557f...acbe9", "Time": "2025-01-19 10:04:03"}, {"ID": "17343...0c2fa", "Time": "2025-01-19 9:01:03"}},                                                                                             // 简化示例
		"LatestBlocks":       []map[string]string{{"Number": "156085", "Transactions": "1", "Node": "org1.peer1"}, {"Number": "156084", "Transactions": "2", "Node": "org2.peer1"}, {"Number": "156083", "Transactions": "1", "Node": "org3.peer1"}, {"Number": "156082", "Transactions": "1", "Node": "org3.peer1"}, {"Number": "156081", "Transactions": "1", "Node": "org1.peer1"}, {"Number": "156080", "Transactions": "1", "Node": "org1.peer1"}, {"Number": "156079", "Transactions": "1", "Node": "org3.peer1"}, {"Number": "156078", "Transactions": "4", "Node": "org3.peer1"}, {"Number": "156077", "Transactions": "2", "Node": "org3.peer1"}, {"Number": "156076", "Transactions": "1", "Node": "org3.peer1"}}, // 简化示例
		"TransactionData":    []map[string]string{{"time": "2025-01-24T01:34:56Z", "count": "10"}, {"time": "2025-01-24T04:12:34Z", "count": "15"}, {"time": "2025-01-24T21:33:00Z", "count": "3"}, {"time": "2025-01-24T23:11:44Z", "count": "25"}, {"time": "2025-01-24T23:51:00Z", "count": "11"}, {"time": "2025-01-24T23:53:30Z", "count": "7"}, {"time": "2025-01-24T23:56:15Z", "count": "19"}, {"time": "2025-01-24T23:58:45Z", "count": "4"}, {"time": "2025-01-25T00:01:15Z", "count": "2"}, {"time": "2025-01-24T23:59:00Z", "count": "14"}, {"time": "2025-01-24T23:59:30Z", "count": "6"}},
	})
}
