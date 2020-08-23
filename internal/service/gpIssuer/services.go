package gpIssuer

import (
	"fmt"
	"massimple.com/wallet-controller/internal/models"
	"time"
)

///////////////////////////////////////////////////////
/// 			Implemented methods					///
///////////////////////////////////////////////////////
func AltaDeCuenta(cuentaDto models.GpNewAccountInput) (models.GpNewAccountOutput, error){
	resp := respAltaDeCuenta{}
	cuenta := cuentaDtoToDefaultCuenta(cuentaDto)
	err := post(productBaseUrl + "/Cuentas", cuenta).execute(&resp)
	newAcc := cuentaToNewAccountDto(resp)
	return newAcc, err
}

func CargaDeTarjeta(cuentaId models.ID, recargaDto models.GpRecharge) error {
	recarga := recargaDtoToDefaultRecarga(recargaDto)
	err := post(fmt.Sprintf(productBaseUrl + "/Cuentas/%s/Transacciones/Cargas", cuentaId), recarga).execute(nil)
	return err
}

/*
 * Calcula el time.Time del primer día del mes y del último día del mes usando @mes como reltivo al mes actual
 */
func generarQueryDeFechas(mes int) string {
	t := time.Now()
	desde := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0,-mes,0)
	hasta := desde.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	//desdeString := fmt.Sprintf("%d-%d-%d", desde.Year(), desde.Month(), desde.Day())
	//hastaString := fmt.Sprintf("%d-%d-%d", hasta.Year(), hasta.Month(), hasta.Day())
	return fmt.Sprintf("FechaDesde=%s&FechaHasta=%s", desde.Format("2006-01-02"), hasta.Format("2006-01-02"))
}

/*
 * Retorna todos los movimientos del mes.
 * @mes represnta el mes relativo al mes actual.
 * ej: mes=0 => mes actual
 *     mes=1 => mes anterior
 *     mes=12 => mes actual año anterior
 */
func ConsultaDeMovimientos(cuentaId models.ID, mes int) (models.GPAccountMovements, error){
	resp := respConsultaDeMovimientos{}
	query := generarQueryDeFechas(mes)
	err := get(fmt.Sprintf(productBaseUrl + "/Cuentas/%s/Transacciones/Cargas?%s", cuentaId, query)).execute(&resp)
	movements := consultaDeMovimientosToAccountMovementsDto(resp)
	return movements, err
}

func ConsultaDeDisponileYSaldo(cuentaId models.ID) (models.GPAvailable, error) {
	disponible, err := consultaDeDisponible(cuentaId)
	if err != nil {
		return models.GPAvailable{}, err
	}
	saldo, err := consultaDeSaldo(cuentaId)
	if err != nil {
		return models.GPAvailable{}, err
	}
	return models.GPAvailable{
		LocalAvailableBuy:      disponible.DisponibleCompra,
		LocalAvailableAdvance:  disponible.DisponibleAnticipo,
		DollarAvailableBuy:     disponible.DisponibleCompraDolar,
		DollarAvailableAdvance: disponible.DisponibleAnticipoDolar,
		LocalBalance:           saldo.SaldoPesos,
		DollarBalance:          saldo.SaldoDolar,
	}, nil
}

func consultaDeDisponible(cuentaId models.ID) (respConsultaDeDisponibleSaldo, error){
	resp := respConsultaDeDisponibleSaldo{}
	err := get(fmt.Sprintf(productBaseUrl + "/Cuentas/%s/Disponible", cuentaId)).execute(&resp)
	return resp, err
}

func consultaDeSaldo(cuentaId models.ID) (respConsultaDeDisponibleSaldo, error){
	resp := respConsultaDeDisponibleSaldo{}
	err := get(fmt.Sprintf(productBaseUrl + "/Cuentas/%s/Disponible", cuentaId)).execute(&resp)
	return resp, err
}

///////////////////////////////////////////////////////
/// 			Non-implemented methods				///
///////////////////////////////////////////////////////

func ConsultaDeTransaccion(transIdId models.ID) (respConsultaDeTransaccion, error){
	/*resp := respConsultaDeTransaccion{}
	err := get(fmt.Sprintf(api.baseUrl + "/Requests/%s", transId)).execute(&resp)*/
	return respConsultaDeTransaccion{}, models.NotImplementedError("ConsultaDeTransaccion not implemented")
}

/*
 * Retorna las cargas del mes.
 * @mes represnta el mes relativo al mes actual.
 * ej: mes=0 => mes actual
 *     mes=1 => mes anterior
 *     mes=12 => mes actual año anterior
 */
func ConsultaDeCargas(cuentaId models.ID, mes int) (respConsultaDeCargas, error){
	/*resp := respConsultaDeCargas{}
	// gp solo retorna datos de hasta 30 días. entonces el "desde" podría ser calculado
	desde, hasta := generarQueryDeFechas(mes)
	fmt.Print("desde", desde, "hasta", hasta)
	query := fmt.Sprintf("FechaDesde=%s&FechaHasta=%s", desde.Format("2017-08-01"), hasta.Format("2017-08-01"))
	err := get(fmt.Sprintf(productBaseUrl + "/Cuentas/%s/Transacciones/Cargas?%s", cuenta, query)).execute(&resp)*/
	return respConsultaDeCargas{}, models.NotImplementedError("Consulta de cargas not implemented")
}

func ModificarDatosDeCuenta(cuenta cuentaGp) error {
	//err := put(productBaseUrl + "/Cuentas", cuenta).execute(nil)
	return models.NotImplementedError("CModificarDatosDeCuenta not implemented")
}

func ConsultaDeCuenta(cuenta models.ID) (respConsultaDeCuenta, error){
	/*resp := respConsultaDeCuenta{}
	err := get(fmt.Sprintf(productBaseUrl + "/Cuentas/%s", cuenta)).execute(&resp)
	acc := cuentaToNewAccountDto(resp)
	return acc, err*/
	return respConsultaDeCuenta{}, models.NotImplementedError("Consulta de Cuenta not implemented")
}
