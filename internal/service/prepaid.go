package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const url1 = "https://34.231.98.134/crediware"
const url2 = "https://34.231.98.134:1486/WebServiceLink.asmx"
const url3 = "https://34.231.98.134:44301/GP.IDS/connect/token"


const GPCLientId = "asdf" // todo env
const GPCLientPass = "asdf"


var auth = struct {
	AccessToken string `json:"Access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}{}

func GPInit() {
	getToken()
	log.Info("GP-Emisor connected")
}

func getToken() {
	values := url.Values{}
	values.Set("client_id", GPCLientId)
	values.Set("client_secret", GPCLientPass)
	values.Set("grant_type","client_credentials")
	values.Set("scope","PrepagaApi")

	resp, err := http.PostForm(url2 + "/connect/token", values)
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&auth)
	if err != nil {
		panic(err)
	}
}

func post(url string, body interface{}) *reqExecutor {
	return baseRequest("POST", url, body)
}

func put(url string, body interface{}) *reqExecutor {
	return baseRequest("PUT", url, body)
}

func get(url string) *reqExecutor {
	return baseRequest("GET", url, nil)
}

func baseRequest(method string, url string, body interface{}) *reqExecutor {
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		// marshal de json body
		bytesArr, _ := json.Marshal(body)
		// Create a new request using http
		req, err = http.NewRequest(method, url, bytes.NewBuffer(bytesArr))
	}

	if err == nil {
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + auth.AccessToken
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}
	return &reqExecutor{req: req, err: err}
}

type ReqExecutor interface {
	Execute(interface{}) error
}

type reqExecutor struct {
	req *http.Request
	err error
}

func (re *reqExecutor) Execute(resultPtr interface{}) error {
	if re.err != nil {
		return re.err
	}
	client := &http.Client{}
	resp, err := client.Do(re.req)
	if err != nil {
		return err
	}
	// todo caso 4XX y 5XX
	if resultPtr == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(&resultPtr)
}

func AltaDeCuenta(cuenta CuentaGp) (RespAltaDeTarjetaVirtual, error){
	prod := "NO IDEA PAILOR" // todo consultar
	resp := RespAltaDeTarjetaVirtual{}
	err := post(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas", cuenta).Execute(&resp)
	return resp, err
}

func ModificarDatosDeCuenta(cuenta CuentaGp) error{
	prod := "NO IDEA PAILOR" // todo consultar
	err := put(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas", cuenta).Execute(nil)
	return err
}

func ConsultaDeCuenta(cuenta IDGP) (RespConsultaDeCuenta, error){
	prod := "NO IDEA PAILOR" // todo consultar
	resp := RespConsultaDeCuenta{}
	err := get(fmt.Sprintf(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas/%s", cuenta)).Execute(&resp)
	return resp, err
}

func CargaDeTarjeta(cuenta IDGP, carga CargaDeTarjetaGP) (RespCargaDeTarjeta, error){
	prod := "NO IDEA PAILOR" // todo consultar
	resp := RespCargaDeTarjeta{}
	err := post(fmt.Sprintf(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas/%s/Transacciones/Cargas", cuenta), carga).Execute(&resp)
	return resp, err
}

/*
 * Retorna las cargas del mes.
 * @mes represnta el mes relativo al mes actual.
 * ej: mes=0 => mes actual
 *     mes=1 => mes anterior
 *     mes=12 => mes actual año anterior
*/
func ConsultaDeCargas(cuenta IDGP, mes int) (RespConsultaDeCargas, error){
	prod := "NO IDEA PAILOR" // todo consultar
	resp := RespConsultaDeCargas{}
	// gp solo retorna datos de hasta 30 días. entonces el "desde" podría ser calculado
	desde, hasta := calcularFechasDeMesEntero(mes)
	fmt.Print("desde", desde, "hasta", hasta)
	query := fmt.Sprintf("FechaDesde=%s&FechaHasta=%s", desde.Format("2017-08-01"), hasta.Format("2017-08-01"))
	err := get(fmt.Sprintf(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas/%s/Transacciones/Cargas?%s", cuenta, query)).Execute(&resp)
	return resp, err
}

/*
 * Calcula el time.Time del primer día del mes y del último día del mes usando @mes como reltivo al mes actual
 */
func calcularFechasDeMesEntero(mes int) (time.Time, time.Time) {
	t := time.Now()
	desde := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0,-mes,0)
	hasta := desde.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	return desde, hasta
}

/*
 * Retorna todos los movimientos del mes.
 * @mes represnta el mes relativo al mes actual.
 * ej: mes=0 => mes actual
 *     mes=1 => mes anterior
 *     mes=12 => mes actual año anterior
 */
func ConsultaDeMovimientos(cuenta IDGP, mes int) (RespConsultaDeMovimientos, error){
	prod := "NO IDEA PAILOR" // todo consultar
	resp := RespConsultaDeMovimientos{}
	desde, hasta := calcularFechasDeMesEntero(mes)
	fmt.Print("desde", desde, "hasta", hasta)
	query := fmt.Sprintf("FechaDesde=%s&FechaHasta=%s", desde.Format("2017-08-01"), hasta.Format("2017-08-01"))
	err := get(fmt.Sprintf(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas/%s/Transacciones/Cargas?%s", cuenta, query)).Execute(&resp)
	return resp, err
}

func ConsultaDeDisponible(cuenta IDGP) (RespConsultaDeDisponibleSaldo, error){
	prod := "NO IDEA PAILOR" // todo consultar
	resp := RespConsultaDeDisponibleSaldo{}
	err := get(fmt.Sprintf(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas/%s/Disponible", cuenta)).Execute(&resp)
	return resp, err
}

func ConsultaDeSaldo(cuenta IDGP) (RespConsultaDeDisponibleSaldo, error){
	prod := "NO IDEA PAILOR" // todo consultar
	resp := RespConsultaDeDisponibleSaldo{}
	err := get(fmt.Sprintf(url2 + "/GP.Prepagas/Api/Productos/" + prod + "/cuentas/%s/Disponible", cuenta)).Execute(&resp)
	return resp, err
}

func ConsultaDeTransaccion(transId IDGP) (RespConsultaDeTransaccion, error){
	resp := RespConsultaDeTransaccion{}
	err := get(fmt.Sprintf(url2 + "/GP.Prepagas/Api/Requests/%s", transId)).Execute(&resp)
	return resp, err
}


type ID string
type IDGP string

type RespConsultaDeCuenta struct {
	IdCuenta int
	Producto producto
	Marca marca
	Tipo string
	Nombre string
	Apellido string
	Adicional int
	Documento documento
	DomicilioParticular domicilio
	DomicilioCorrespondencia domicilio
	Telefonos telefonos
	Email string
	Sexo string
	FechaNacimiento time.Time
	IdCuentaExterna ID
	Estado string
	SucursalEmisora int
	Tarjetas []tarjetaReferencia
}

type RespConsultaDeTransaccion struct {
	Id int
	Methodo string
	Parametros string
	RequestId string
	FechaIngreso time.Time
	FechaSalida time.Time
	Resultado int
	Detalles string
}

type marca struct {
	IdMarca int
	Descripcion string
}
type movimiento struct {
	Id IDGP
	Tipo int
	Fecha time.Time
	Descripcion string
	Importe importe
	TCC string
	MCC string
	DescripcionActividad string
	DescripcionPlan string
	Observaciones string
}

type RespConsultaDeMovimientos struct {
	Desde int
	Cantidad int
	FechaDesde time.Time
	FechaHasta time.Time
	Total int
	Resultado []movimiento
}

type RespConsultaDeDisponibleSaldo struct {
	DisponibleCompra float64
	DisponibleAnticipo float64
	DisponibleCompraDolar float64
	DisponibleAnticipoDolar float64
	SaldoPesos float64
	SaldoDolar float64
}

type RespConsultaDeCargas struct {
	Desde int
	Cantidad int
	FechaDesde time.Time
	FechaHasta time.Time
	Total int
	Resultado []carga
}

type carga struct {
	Fecha time.Time
	CodigoComprobante int
	Importe importe
	NumeroComprobante int
}

type RespCargaDeTarjeta struct {
	CodigoConfirmacion int
	GeneroCargos bool
}

type importe struct {
	Monto float64
	Moneda string
}

type CargaDeTarjetaGP struct {
	Importe importe
	Comprobante int
	FormaDePago int
	Origen int
	Sucursal int
	Observacion string
}

type producto struct {
	IdProducto int
	Descripcion string
}

type tarjeta struct {
	NumeroTarjeta string
	Vencimiento time.Time
	Cvc string
	TipoTarjeta int
}

type tarjetaReferencia struct {
	Referencia string
	Estado string
	Vencimiento time.Time
}

type RespAltaDeTarjetaVirtual struct {
	IdCuenta int
	Producto producto
	IdCuentaExterna ID
	Tarjetas  []tarjeta
}

type documento struct {
	Tipo	string
	Numero	string
}

type domicilio struct {
	Calle	string
	Altura	string
	Departamento	string
	CodigoPostal	string
	Localidad	string
	Provincia	int
	Pais	int
	Comentario string
}

type telefonos struct {
	Particular	string
	Celular	string
	Correspondencia	string
	Laboral	string
	Otro	string
}

type CuentaGp struct {
	TipoProducto	int
	Nombre	string
	Apellido	string
	Documento documento
	DomicilioParticular domicilio
	DomicilioCorrespondencia domicilio
	Sexo string
	FechaNacimiento string
	Telefonos telefonos
	Email	string
	IdCuentaExterna	ID
	SucursalEmisora string
	GrupoAfinidad	int
	Embozado	Embozado
	Nacionalidad	int
	TrackingId	string
	TipoTarjeta	int
	DocumentoOtro string
}

type Embozado struct {
	Nombre string
	CuartaLinea	string
}

