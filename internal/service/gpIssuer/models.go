package gpIssuer

import (
	"encoding/json"
	"fmt"
	"massimple.com/wallet-controller/internal/models"
	"reflect"
	"strconv"
	"strings"
	"time"
)

///////////////////////////////////////////////////////
/// 				Public models					///
///////////////////////////////////////////////////////
type GPError struct {
	Code    int
	Message string
}

func (e *GPError) Error() string {
	return fmt.Sprintf("GP Error(%d): %s", e.Code, e.Message)
}

///////////////////////////////////////////////////////
/// 				Default values					///
///////////////////////////////////////////////////////

func fillDefaults()  {
	recargaDefaults = recargaDefaultsType{
		moneda:      1,
		comprobante: 0000000000,
		formaDePago: 0,
		origen:      8,
		sucursal:    api.branch,
		observacion: vacio,
	}
	cuentaDefaults = cuentaDefaultsType{
		domicilioParticular: domicilio{
			Calle:        "Pje. Gral. Lorenzo Vintter",
			Altura:       "817",
			Departamento: vacio,
			CodigoPostal: vacio,
			Localidad:    vacio,
			Provincia:    3,
			Pais:         32,
			Comentario:   "Más Simple",
		},
		domicilioCorrespondencia: domicilio{
			Calle:        "Pje. Gral. Lorenzo Vintter",
			Altura:       "817",
			Departamento: vacio,
			CodigoPostal: vacio,
			Localidad:    vacio,
			Provincia:    3,
			Pais:         32,
			Comentario:   "Más Simple",
		},
		email: vacio,
		sexo:  sexo.masculino,
		telefonos: telefonos{
			Particular:      vacio,
			Celular:         vacio,
			Correspondencia: vacio,
			Laboral:         vacio,
			Otro:            vacio,
		},
		documentoOtro: "30716618834",
		embozado: embozado{
			Nombre:      vacio,
			CuartaLinea: vacio,
		},
		nacionalidad: 32,
		trackingId:   vacio,
		tipoTarjeta:  tarjetas.virtual,
	}
}
const vacio = "vacio"
type cuentaDefaultsType struct {
	domicilioParticular      domicilio
	domicilioCorrespondencia domicilio
	email                    string
	sexo                     string
	telefonos                telefonos
	documentoOtro            string
	embozado                 embozado
	nacionalidad             int
	trackingId               string
	tipoTarjeta              int
}
type recargaDefaultsType struct {
	moneda		int
	comprobante	int
	formaDePago	int
	origen		int
	sucursal	int
	observacion	string
}
var recargaDefaults recargaDefaultsType
var cuentaDefaults cuentaDefaultsType

///////////////////////////////////////////////////////
/// 				Private models					///
///////////////////////////////////////////////////////

type gpFecha time.Time

func (j *gpFecha) UnmarshalJSON(b []byte) error {
	// https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
	// This implementation of (j *gpFecha) UnmarshalJSON(b []byte) is used by json.Decode() to generate structs
	// from json strings. Here we are parsing the weired time format used by GP. ex: 2019-07-22T00:00:00
	s := strings.Trim(string(b), "\"") // from []byte to string
	// 2019-07-22T00:00:00
	dateTime := strings.Split(s, "T")        // separate between date and time ["2019-07-22","00:00:00"]
	dates := strings.Split(dateTime[0], "-") // separate date data to ["2019","07","22"]
	times := strings.Split(dateTime[1], ":") // separate time data to ["00","00","00"]
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])
	hour, _ := strconv.Atoi(times[0])
	minute, _ := strconv.Atoi(times[1])
	second, _ := strconv.Atoi(times[2])
	timeZone := time.FixedZone("Argentina Time", int((-3 * time.Hour).Seconds()))
	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, timeZone)
	*j = gpFecha(t)
	return nil
}

type tipoTarjeta int

func (j *tipoTarjeta) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"") // from []byte to string
	switch s {
	case "0":
		*j = tipoTarjeta(0)
	case "1":
		*j = tipoTarjeta(1)
	default:
		return &json.UnmarshalTypeError{
			Value:  s,
			Type:   reflect.TypeOf(s),
			Offset: 0,
			Struct: "",
			Field:  "tipoTarjeta",
		}
	}
	return nil
}

type respConsultaDeCuenta struct {
	IdCuenta                 int
	Producto                 producto
	Marca                    marca
	Tipo                     string
	Nombre                   string
	Apellido                 string
	Adicional                int
	Documento                documento
	DomicilioParticular      domicilio
	DomicilioCorrespondencia domicilio
	Telefonos                telefonos
	Email                    string
	Sexo                     string
	FechaNacimiento          gpFecha
	IdCuentaExterna          models.ID
	Estado                   string
	SucursalEmisora          int
	Tarjetas                 []tarjetaReferencia
}

type respConsultaDeTransaccion struct {
	Id           int
	Methodo      string
	Parametros   string
	RequestId    string
	FechaIngreso gpFecha
	FechaSalida  gpFecha
	Resultado    int
	Detalles     string
}

type marca struct {
	IdMarca     int
	Descripcion string
}

type movimiento struct {
	ID                   int
	Tipo                 int
	Fecha                gpFecha
	Descripcion          string
	Importe              importeOut
	TCC                  string
	MCC                  string
	DescripcionActividad string
	DescripcionPlan      string
	Observaciones        string
}

type respConsultaDeMovimientos struct {
	Desde      int
	Cantidad   int
	FechaDesde gpFecha
	FechaHasta gpFecha
	Total      int
	Resultado  []movimiento
}

type respConsultaDeDisponibleSaldo struct {
	DisponibleCompra        float64
	DisponibleAnticipo      float64
	DisponibleCompraDolar   float64
	DisponibleAnticipoDolar float64
	SaldoPesos              float64
	SaldoDolar              float64
}

type respConsultaDeCargas struct {
	Desde      int
	Cantidad   int
	FechaDesde gpFecha
	FechaHasta gpFecha
	Total      int
	Resultado  []carga
}

type carga struct {
	Fecha             gpFecha
	CodigoComprobante int
	Importe           importeOut
	NumeroComprobante int
}

type respCargaDeTarjeta struct {
	CodigoConfirmacion int
	GeneroCargos       bool
}

type importeIn struct {
	Monto  float64
	Moneda int
}

type importeOut struct {
	Monto  float64
	Moneda string
}

type recarga struct {
	Importe     importeIn
	Comprobante int
	FormaDePago int
	Origen      int
	Sucursal    int
	Observacion string
}

type producto struct {
	IdProducto  int
	Descripcion string
}

type tarjeta struct {
	NumeroTarjeta string
	Vencimiento   gpFecha
	Cvc           string
	TipoTarjeta   tipoTarjeta
}

type tarjetaReferencia struct {
	Referencia  string
	Estado      string
	Vencimiento gpFecha
}

type respAltaDeCuenta struct {
	IdCuenta        int
	Producto        producto
	IdCuentaExterna models.ID
	Tarjetas        []tarjeta
}

type documento struct {
	Tipo   string
	Numero string
}

type domicilio struct {
	Calle        string
	Altura       string
	Piso         string
	Departamento string
	CodigoPostal string
	Localidad    string
	Provincia    int
	Pais         int
	Comentario   string
}

type telefonos struct {
	Particular      string
	Celular         string
	Correspondencia string
	Laboral         string
	Otro            string
}

type cuentaGp struct {
	TipoProducto             int
	Nombre                   string
	Apellido                 string
	Documento                documento
	DomicilioParticular      domicilio
	DomicilioCorrespondencia domicilio
	Sexo                     string
	FechaNacimiento          string
	Telefonos                telefonos
	Email                    string
	IdCuentaExterna          models.ID
	SucursalEmisora          int
	GrupoAfinidad            int
	Embozado                 embozado
	Nacionalidad             int
	TrackingId               string
	TipoTarjeta              int
	DocumentoOtro            string
}

type embozado struct {
	Nombre      string
	CuartaLinea string
}

var sexo = struct {
	masculino string
	femenino  string
}{
	masculino: "M",
	femenino:  "F",
}

var tarjetas = struct {
	fisica  int
	virtual int
}{
	fisica:  0,
	virtual: 1,
}

var tipoDocumento = struct {
	nada  string
	DNI   string
	CI    string
	LE    string
	Pass  string
	LC    string
	RUT   string
	CUIL  string
	CUIT  string
	otros string
}{
	nada:  "Nada",
	DNI:   "DNI",
	CI:    "CI",
	LE:    "LE",
	Pass:  "Pass",
	LC:    "LC",
	RUT:   "RUT",
	CUIL:  "CUIL",
	CUIT:  "CUIT",
	otros: "Otros",
}

func cuentaDtoToDefaultCuenta(dto models.GpNewAccountInput) cuentaGp {
	return cuentaGp{
		TipoProducto: api.productNumber,
		Nombre:       dto.Name,
		Apellido:     dto.Lastname,
		Documento:                documento{
			Tipo:   tipoDocumento.DNI,
			Numero: dto.DocumentNumber,
		},
		DomicilioParticular:      cuentaDefaults.domicilioParticular,
		DomicilioCorrespondencia: cuentaDefaults.domicilioParticular,
		Sexo:                     cuentaDefaults.sexo,
		FechaNacimiento:          dto.BirthDate,
		Telefonos:                cuentaDefaults.telefonos,
		Email:                    cuentaDefaults.email,
		IdCuentaExterna:          dto.ExternalId,
		SucursalEmisora:          api.branch,
		GrupoAfinidad:            api.afinityGroup,
		Embozado:                 cuentaDefaults.embozado,
		Nacionalidad:             cuentaDefaults.nacionalidad,
		TrackingId:               cuentaDefaults.trackingId,
		TipoTarjeta:              cuentaDefaults.tipoTarjeta,
		DocumentoOtro:            cuentaDefaults.documentoOtro,
	}
}

func cuentaToNewAccountDto(cuenta respAltaDeCuenta) models.GpNewAccountOutput {
	cards := gpTarjetasToCardsDto(cuenta.Tarjetas)
	return models.GpNewAccountOutput{
		ID:         models.IDGP(cuenta.IdCuenta),
		ExternalId: cuenta.IdCuentaExterna,
		Cards:      cards,
	}
}

func gpTarjetasToCardsDto(tarjetas []tarjeta) []models.GPCard {
	var cards []models.GPCard
	for _, c := range tarjetas {
		cards = append(cards, models.GPCard{
			CardNumber: c.NumeroTarjeta,
			Cvc:        c.Cvc,
		})
	}
	return cards
}

func recargaDtoToDefaultRecarga(dto models.GpRecharge) recarga {
	return recarga{
		Importe:     importeIn{
			Monto:  dto.Amount,
			Moneda: recargaDefaults.moneda,
		},
		Comprobante: recargaDefaults.comprobante,
		FormaDePago: recargaDefaults.formaDePago,
		Origen:      recargaDefaults.origen,
		Sucursal:    recargaDefaults.sucursal,
		Observacion: recargaDefaults.observacion,
	}
}

func consultaDeMovimientosToAccountMovementsDto(resp respConsultaDeMovimientos) models.GPAccountMovements {
	var movements []models.GpMovement
	for _, m := range resp.Resultado {
		movements = append(movements, models.GpMovement{
			ID:           models.IDGP(m.ID),
			Type:         m.Tipo,
			Date:         time.Time(m.Fecha),
			Description:  m.Descripcion,
			Amount:       m.Importe.Monto,
			Observations: m.Observaciones,
		})
	}
	return models.GPAccountMovements{
		Amount:      resp.Cantidad,
		DateFrom:    time.Time(resp.FechaDesde),
		DateTo:      time.Time(resp.FechaHasta),
		TotalAmount: resp.Total,
		Movements:   movements,
	}
}
