package user

import (
	"bjungle-consenso/internal/env"

	"bjungle-consenso/internal/grpc/accounting_proto"
	"bjungle-consenso/internal/grpc/wallet_proto"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcMetadata "google.golang.org/grpc/metadata"
	"net/http"
)

type handlerUser struct {
	DB   *sqlx.DB
	TxID string
}

// GetAccountByWalletID godoc
// @Summary Método para obtener una cuenta por wallet id
// @Description Método para obtener una cuenta asociada a una wallet
// @tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization" default(Bearer <Add access token here>)
// @Param wallet path string true "Id de la wallet"
// @Success 200 {object} resAccount
// @Router /api/v1/user/accounting/{wallet} [get]
func (h *handlerUser) GetAccountByWalletID(c *fiber.Ctx) error {
	res := resAccount{Error: true}
	walletId := c.Params("wallet")
	e := env.NewConfiguration()

	connAuth, err := grpc.Dial(e.AuthService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		res.Code, res.Type, res.Msg = 22, 1, "error conectando con el servicio auth de blockchain"
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connAuth.Close()

	clientAccount := accounting_proto.NewAccountingServicesAccountingClient(connAuth)

	token := c.Get("Authorization")[7:]
	ctx := grpcMetadata.AppendToOutgoingContext(context.Background(), "authorization", token)

	resWsAccount, err := clientAccount.GetAccountingByWalletById(ctx, &accounting_proto.RequestGetAccountingByWalletId{Id: walletId})
	if err != nil {
		logger.Error.Printf("error conectando con el servicio account de blockchain: %s", err)
		res.Code, res.Type, res.Msg = 22, 1, "error conectando con el servicio account de blockchain"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resWsAccount == nil {
		logger.Error.Printf("error conectando con el servicio account de blockchain")
		res.Code, res.Type, res.Msg = 22, 1, "error conectando con el servicio account de blockchain"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resWsAccount.Error {
		logger.Error.Printf(resWsAccount.Msg)
		res.Code, res.Type, res.Msg = int(resWsAccount.Code), int(resWsAccount.Type), resWsAccount.Msg
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = Accounting{
		Id:       resWsAccount.Data.Id,
		IdWallet: resWsAccount.Data.IdWallet,
		Amount:   resWsAccount.Data.Amount,
	}
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

// GetFreezeMoney godoc
// @Summary Método para obtener la cantidad de acais congelados
// @Description Método para obtener la cantidad de acais congelados
// @tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization" default(Bearer <Add access token here>)
// @Param wallet path string true "Id de la wallet"
// @Success 200 {object} resFreezeMoney
// @Router /api/v1/user/freeze-money/{wallet} [get]
func (h *handlerUser) GetFreezeMoney(c *fiber.Ctx) error {

	res := resFreezeMoney{Error: true}
	walletId := c.Params("wallet")

	e := env.NewConfiguration()
	connAuth, err := grpc.Dial(e.AuthService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		res.Code, res.Type, res.Msg = 22, 1, "error conectando con el servicio auth de blockchain"
		return c.Status(http.StatusAccepted).JSON(res)
	}
	defer connAuth.Close()

	clientWallet := wallet_proto.NewWalletServicesWalletClient(connAuth)

	token := c.Get("Authorization")[7:]
	ctx := grpcMetadata.AppendToOutgoingContext(context.Background(), "authorization", token)

	resFrozen, err := clientWallet.GetFrozenMoney(ctx, &wallet_proto.RqGetFrozenMoney{WalletId: walletId})
	if err != nil {
		logger.Error.Printf("error conectando con el servicio account de blockchain: %s", err)
		res.Code, res.Type, res.Msg = 22, 1, "error conectando con el servicio account de blockchain"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resFrozen == nil {
		logger.Error.Printf("error conectando con el servicio account de blockchain")
		res.Code, res.Type, res.Msg = 22, 1, "error conectando con el servicio account de blockchain"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if resFrozen.Error {
		logger.Error.Printf(resFrozen.Msg)
		res.Code, res.Type, res.Msg = int(resFrozen.Code), int(resFrozen.Type), resFrozen.Msg
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resFrozen.Data
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
