package sign

import (
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/msg"
	"bjungle-consenso/pkg/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerSign struct {
	DB   *sqlx.DB
	TxID string
}

// createSign godoc
// @Summary Método para crear una firma del cuerpo de la transacción
// @Description Método para crear una firma del cuerpo de la transacción
// @tags Sign
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization" default(Bearer <Add access token here>)
// @Param ReqSign body ReqSign true "Datos para registrar la firma"
// @Success 200 {object} ResSign
// @Router /api/v1/sign/create [post]
func (h *handlerSign) createSign(c *fiber.Ctx) error {
	res := ResSign{Error: true}
	req := ReqSign{}
	err := c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("No se pudo parsear el cuerpo de la petición: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if req.Key == "" || len(req.Key) != 32 {
		logger.Error.Printf("La key es requerida")
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, nil, h.TxID)

	access, code, err := srvCfg.SrvBLionAccess.CreateBLionAccess(uuid.New().String(), req.Key, "created")
	if err != nil {
		logger.Error.Printf("No se pudo crear la key, error: ", err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if access == nil {
		logger.Error.Printf("No se pudo crear la key")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = access.ID
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	return c.Status(http.StatusOK).JSON(res)
}

// exportSign godoc
// @Summary Método para exportar firmas por id
// @Description Método para exportar firmas por id
// @tags Sign
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization" default(Bearer <Add access token here>)
// @Param ReqExportSign body ReqExportSign true "Ids de los registros"
// @Success 200 {object} ResExportSign
// @Router /api/v1/sign/export [post]
func (h *handlerSign) exportSign(c *fiber.Ctx) error {
	res := ResExportSign{Error: true}
	req := ReqExportSign{}
	err := c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("No se pudo parsear el cuerpo de la petición: %s", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if len(req.Id) == 0 {
		logger.Error.Printf("Los ids son requeridos")
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, nil, h.TxID)

	access, err := srvCfg.SrvBLionAccess.GetBLionAccessByIds(req.Id)
	if err != nil {
		logger.Error.Printf("No se pudo obtener los registros, error: ", err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if access == nil {
		logger.Error.Printf("No se pudo obtener los registros")
		res.Code, res.Type, res.Msg = msg.GetByCode(22, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	var keys []Sign
	for _, lionAccess := range access {
		keys = append(keys, Sign{
			Id:  lionAccess.ID,
			Key: lionAccess.Key,
		})

		_, err = srvCfg.SrvBLionAccess.DeleteBLionAccess(lionAccess.ID)
		if err != nil {
			logger.Error.Printf("No se pudo eliminar el registro ", lionAccess.ID, ", error: ", err.Error())
			continue
		}
	}

	res.Data = keys
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	return c.Status(http.StatusOK).JSON(res)
}
