package bc

import (
	"bjungle-consenso/internal/models"
	"bjungle-consenso/pkg/bc/lotteries"
	"bjungle-consenso/pkg/bc/miner_response"
	"bjungle-consenso/pkg/bc/participants"
	"bjungle-consenso/pkg/bc/penalty_participants"
	"bjungle-consenso/pkg/bc/rewards"
	"bjungle-consenso/pkg/bc/validator_votes"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvLottery        lotteries.PortsServerLottery
	SrvParticipants   participants.PortsServerParticipants
	SrvPenalty        penalty_participants.PortsServerPenaltyParticipants
	SrvReward         rewards.PortsServerRewards
	SrvValidatorsVote validator_votes.PortsServerValidatorVotes
	SrvMinerResponse  miner_response.PortsServerMinerResponse
}

func NewServerBk(db *sqlx.DB, user *models.User, txID string) *Server {

	repoLottery := lotteries.FactoryStorage(db, user, txID)
	srvLottery := lotteries.NewLotteryService(repoLottery, user, txID)

	repoParticipants := participants.FactoryStorage(db, user, txID)
	srvParticipants := participants.NewParticipantsService(repoParticipants, user, txID)

	repoReward := rewards.FactoryStorage(db, user, txID)
	srvReward := rewards.NewRewardsService(repoReward, user, txID)

	repoValidatorsVote := validator_votes.FactoryStorage(db, user, txID)
	srvValidatorsVote := validator_votes.NewValidatorVotesService(repoValidatorsVote, user, txID)

	repoMinerResponse := miner_response.FactoryStorage(db, user, txID)
	srvMinerResponse := miner_response.NewMinerResponseService(repoMinerResponse, user, txID)

	repoPenalty := penalty_participants.FactoryStorage(db, user, txID)
	srvPenalty := penalty_participants.NewPenaltyParticipantsService(repoPenalty, user, txID)

	return &Server{
		SrvLottery:        srvLottery,
		SrvParticipants:   srvParticipants,
		SrvReward:         srvReward,
		SrvValidatorsVote: srvValidatorsVote,
		SrvMinerResponse:  srvMinerResponse,
		SrvPenalty:        srvPenalty,
	}
}
