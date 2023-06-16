package validator

import (
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"bjungle-consenso/pkg/bc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerValidators struct {
	DB   *sqlx.DB
	TxID string
}

// RegisterVoteValidator godoc
// @Summary Método para registrar el voto de un validador
// @Description Método para registrar el voto de un validador
// @tags Validator
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization" default(Bearer <Add access token here>)
// @Param rqRegisterVote body rqRegisterVote true "Datos para registrar el voto"
// @Success 200 {object} resRegisterVote
// @Router /api/v1/validators/register [post]
func (h *handlerValidators) RegisterVoteValidator(c *fiber.Ctx) error {
	res := resRegisterVote{Error: true}
	request := rqRegisterVote{}
	err := c.BodyParser(&request)
	if err != nil {
		logger.Error.Printf("couldn't bind model rqRegisterVote: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	srvBc := bc.NewServerBk(h.DB, nil, h.TxID)

	lottery, code, err := srvBc.SrvLottery.GetLotteryActiveForMined()
	if err != nil {
		logger.Error.Printf("couldn't get lottery for mined: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	participant, code, err := srvBc.SrvParticipants.GetParticipantsByWalletIDAndLotteryID(request.WalletID, lottery.ID)
	if err != nil {
		logger.Error.Printf("couldn't get participant by id: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if participant == nil {
		logger.Error.Printf("Usted no esta registrado como participante para esta lotería")
		res.Code, res.Type, res.Msg = 22, 1, "Usted no esta registrado como participante para esta lotería"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if !participant.Accepted {
		logger.Error.Printf("Usted no ha sido seleccionado para como participante de esta lotería")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if participant.TypeCharge != 24 {
		logger.Error.Printf("solo los validadores pueden registrar su voto con respecto al hash propuesto por el minero que hallo el hash")
		res.Code, res.Type, res.Msg = 22, 1, "solo los validadores pueden registrar su voto con respecto al hash propuesto por el minero que hallo el hash"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	hashMined, code, err := srvBc.SrvMinerResponse.GetMinerResponseRegister(lottery.ID)
	if err != nil {
		logger.Error.Printf("couldn't get hash mined: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if hashMined == nil {
		logger.Error.Printf("El hash aun no ha sido resuelto por ningún minero")
		res.Code, res.Type, res.Msg = 22, 1, "El hash aun no ha sido resuelto por ningún minero"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	vote := hashMined.Hash == request.Hash

	votePtr, _, err := srvBc.SrvValidatorsVote.GetValidatorVotesByParticipantID(participant.ID, hashMined.LotteryId)
	if err != nil {
		logger.Error.Printf("No se pudo validar el voto del participante, error: %V", err)
		res.Code, res.Type, res.Msg = 22, 1, "No se pudo validar el voto del participante"
		return c.Status(http.StatusAccepted).JSON(res)
	}
	if votePtr != nil {
		logger.Error.Printf("El participante ya ha registrado su voto para esta lotería")
		res.Code, res.Type, res.Msg = 22, 1, "El participante ya ha registrado su voto para esta lotería"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	_, code, err = srvBc.SrvValidatorsVote.CreateValidatorVotes(uuid.New().String(), hashMined.LotteryId, participant.ID, request.Hash, vote)
	if err != nil {
		logger.Error.Printf("couldn't create vote of validator: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = "Se ha registrado correctamente el voto"
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
