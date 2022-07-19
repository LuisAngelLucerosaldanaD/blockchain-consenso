package participants

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/grpc/wallet_proto"
	"bjungle-consenso/internal/helpers"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"bjungle-consenso/pkg/bc"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcMetadata "google.golang.org/grpc/metadata"
	"net/http"
)

type handlerParticipant struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerParticipant) RegisterParticipant(c *fiber.Ctx) error {
	e := env.NewConfiguration()
	res := responseRegisterParticipant{}
	request := requestRegisterParticipant{}
	err := c.BodyParser(&request)
	if err != nil {
		logger.Error.Printf("couldn't bind model requestRegisterParticipant: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if request.Amount <= 0 {
		logger.Error.Printf("el monto de acais debe ser mayo que 0")
		res.Code, res.Type, res.Msg = 22, 1, "el monto de acais debe ser mayo que 0"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if request.Amount < e.App.MinimumFee {
		logger.Error.Printf("el monto de acais superar a la cuata minima de ingreso")
		res.Code, res.Type, res.Msg = 22, 1, fmt.Sprintf("el monto de acais superar a la cuata minima de ingreso, monto minimo: %d", e.App.MinimumFee)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	u, err := helpers.GetUserContextV2(c)
	if err != nil {
		logger.Error.Printf("couldn't get user token: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvBk := bc.NewServerBk(h.DB, nil, h.TxID)

	connAuth, err := grpc.Dial(e.AuthService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connAuth.Close()

	clientWallet := wallet_proto.NewWalletServicesWalletClient(connAuth)

	token := c.Get("Authorization")[7:]
	ctx := grpcMetadata.AppendToOutgoingContext(context.Background(), "authorization", token)

	resWallet, err := clientWallet.GetWalletByUserId(ctx, &wallet_proto.RequestGetWalletByUserId{UserId: u.ID})
	if err != nil {
		logger.Error.Printf("couldn't get wallet by user id, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resWallet == nil {
		logger.Error.Printf("couldn't get wallet by user id, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resWallet.Error {
		logger.Error.Printf(resWallet.Msg)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if len(resWallet.Data) <= 0 {
		logger.Error.Printf("El usuario no tiene ninguna wallet registrada")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	participant, code, err := srvBk.SrvParticipants.GetParticipantsByWalletID(resWallet.Data[0].Id)
	if err != nil {
		logger.Error.Printf("couldn't get participant by wallet id, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if participant == nil {
		lotteryActive, code, err := srvBk.SrvLottery.GetLotteryActive()
		if err != nil {
			logger.Error.Printf("couldn't get active lottery, error: %s", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
			return c.Status(http.StatusAccepted).JSON(res)
		}

		_, code, err = srvBk.SrvParticipants.CreateParticipants(uuid.New().String(), lotteryActive.ID, resWallet.Data[0].Id, request.Amount, false, 21, false)
		if err != nil {
			logger.Error.Printf("couldn't get active lottery, error: %s", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
			return c.Status(http.StatusAccepted).JSON(res)
		}
		res.Data = "inscrito correctamente"
		res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
		res.Error = false
		return c.Status(http.StatusOK).JSON(res)
	}

	lottery, code, err := srvBk.SrvLottery.GetLotteryByID(participant.LotteryId)
	if err != nil {
		logger.Error.Printf("couldn't validate participant, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if lottery != nil && lottery.RegistrationEndDate == nil && lottery.LotteryEndDate == nil && lottery.ProcessEndDate == nil {
		logger.Error.Printf("couldn't registration user, error: %s", err)
		res.Code, res.Type, res.Msg = 22, 1, "El usuario ya se encuentra a un sorteo pendiente"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	lotteryActive, code, err := srvBk.SrvLottery.GetLotteryActive()
	if err != nil {
		logger.Error.Printf("couldn't get active lottery, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	_, code, err = srvBk.SrvParticipants.CreateParticipants(uuid.New().String(), lotteryActive.ID, resWallet.Data[0].Id, request.Amount, false, 21, false)
	if err != nil {
		logger.Error.Printf("couldn't get active lottery, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	resFrozen, err := clientWallet.FrozenMoney(ctx, &wallet_proto.RqFrozenMoney{
		WalletId:  resWallet.Data[0].Id,
		Amount:    request.Amount,
		LotteryId: lotteryActive.ID,
	})
	if err != nil {
		logger.Error.Printf("couldn't frozen money, error: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resFrozen == nil {
		logger.Error.Printf("couldn't frozen money")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resFrozen.Error {
		logger.Error.Printf(resFrozen.Msg)
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = "inscrito correctamente"
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
