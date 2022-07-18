package bc

import (
	"bjungle-consenso/internal/models"
	"bjungle-consenso/pkg/bc/lottery_table"
	"bjungle-consenso/pkg/bc/miner_response"
	"bjungle-consenso/pkg/bc/participants_table"
	"bjungle-consenso/pkg/bc/reward_table"
	"bjungle-consenso/pkg/bc/validator_votes"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvLottery        lottery_table.PortsServerLotteryTable
	SrvParticipants   participants_table.PortsServerParticipantsTable
	SrvReward         reward_table.PortsServerRewardTable
	SrvValidatorsVote validator_votes.PortsServerValidatorVotes
	SrvMinerResponse  miner_response.PortsServerMinerResponse
}

func NewServerBk(db *sqlx.DB, user *models.User, txID string) *Server {

	repoLottery := lottery_table.FactoryStorage(db, user, txID)
	srvLottery := lottery_table.NewLotteryTableService(repoLottery, user, txID)

	repoParticipants := participants_table.FactoryStorage(db, user, txID)
	srvParticipants := participants_table.NewParticipantsTableService(repoParticipants, user, txID)

	repoReward := reward_table.FactoryStorage(db, user, txID)
	srvReward := reward_table.NewRewardTableService(repoReward, user, txID)

	repoValidatorsVote := validator_votes.FactoryStorage(db, user, txID)
	srvValidatorsVote := validator_votes.NewValidatorVotesService(repoValidatorsVote, user, txID)

	repoMinerResponse := miner_response.FactoryStorage(db, user, txID)
	srvMinerResponse := miner_response.NewMinerResponseService(repoMinerResponse, user, txID)

	return &Server{
		SrvLottery:        srvLottery,
		SrvParticipants:   srvParticipants,
		SrvReward:         srvReward,
		SrvValidatorsVote: srvValidatorsVote,
		SrvMinerResponse:  srvMinerResponse,
	}
}
