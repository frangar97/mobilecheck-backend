package service

import (
	"context"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type TareaService interface {
	CrearTareaMovil(context.Context, model.CreateTareaModelMovil, int64) (model.TareaModelMovil, error)
	CrearTareaWeb(context.Context, model.CreateTareaModelWeb) (model.TareaModelWeb, error)
	ObtenerTareasWeb(context.Context, string, string, int64) ([]model.TareaModelWeb, error)
	ObtenerTareasDelDia(context.Context, string, int64) ([]model.TareaModelMovil, error)
	ObtenerCantidadTareasUsuarioPorFecha(context.Context, string, string) ([]model.CantidadTareaPorUsuario, error)
	CompletarTarea(*gin.Context, model.CompletarTareaModel, int64) error
	CrearTareaMasivaWeb(context.Context, model.CreateTareaMasivaModelWeb, int64, time.Time) error
	CrearTareaMasivaExcelWeb(context.Context, model.CreateTareasExcelWeb, int64, time.Time) error
	VerificarTarea(context.Context, string, int64) (int, error)
	ValidarDataExcel(*gin.Context, int64, int64, int64, string) (model.ValidarTareasExcelWeb, error)
	ObtenerTareasHorasWeb(context.Context, model.ParamReportTareasHoras) ([]model.TareaHorasModelReporteWeb, error)
	EliminarTareas(context.Context, []int64) (int, error)
	ObtenerTareasPorAprobar(context.Context, string, int64) ([]model.AprobarTareas, error)
	AprobarTarea(context.Context, model.CreateAprobarTarea) (bool, error)
	CantidadTareasPendientesAprobar(string) (int64, error)
}

type tareaServiceImpl struct {
	tareaRepository      repository.TareaRepository
	visitaRepository     repository.VisitaRepository
	clienteRepository    repository.ClienteRepository
	usuarioRepository    repository.UsuarioRepository
	tipoVisitaRepository repository.TipoVisitaRepository
}

func newTareaService(tareaRepository repository.TareaRepository, visitaRepository repository.VisitaRepository, clienteRepository repository.ClienteRepository, usuarioRepository repository.UsuarioRepository, tipoVisitaRepository repository.TipoVisitaRepository) *tareaServiceImpl {
	return &tareaServiceImpl{
		tareaRepository:      tareaRepository,
		visitaRepository:     visitaRepository,
		clienteRepository:    clienteRepository,
		usuarioRepository:    usuarioRepository,
		tipoVisitaRepository: tipoVisitaRepository,
	}
}

func (t *tareaServiceImpl) CrearTareaMovil(ctx context.Context, tareaCreate model.CreateTareaModelMovil, usuarioId int64) (model.TareaModelMovil, error) {
	idGenerado, err := t.tareaRepository.CrearTareaMovil(ctx, tareaCreate, usuarioId)
	if err != nil {
		return model.TareaModelMovil{}, err
	}

	return t.tareaRepository.ObtenerTareaPorIdMovil(ctx, idGenerado)
}

func (t *tareaServiceImpl) CrearTareaWeb(ctx context.Context, tareaCreate model.CreateTareaModelWeb) (model.TareaModelWeb, error) {
	idGenerado, err := t.tareaRepository.CrearTareaWeb(ctx, tareaCreate)
	if err != nil {
		return model.TareaModelWeb{}, err
	}

	return t.tareaRepository.ObtenerTareaPorIdWeb(ctx, idGenerado)
}

func (t *tareaServiceImpl) CrearTareaMasivaWeb(ctx context.Context, tareaCreate model.CreateTareaMasivaModelWeb, usuarioCrea int64, fechacrea time.Time) error {

	for _, fecha := range tareaCreate.Fechas {
		tareaModel := model.CreateTareaModelWeb{
			ClienteId:       tareaCreate.ClienteId,
			UsuarioId:       tareaCreate.UsuarioId,
			Meta:            tareaCreate.Meta,
			Fecha:           fecha,
			ImagenRequerida: tareaCreate.ImagenRequerida,
			TipoVisitaId:    tareaCreate.TipoVisitaId,
			MetaLinea:       tareaCreate.MetaLinea,
			MetaSubLinea:    tareaCreate.MetaSubLinea,
			UsuarioCrea:     usuarioCrea,
			FechaCrea:       fechacrea,
		}

		_, err := t.tareaRepository.CrearTareaWeb(ctx, tareaModel)

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *tareaServiceImpl) ObtenerTareasWeb(ctx context.Context, fechaInicio string, fechaFinal string, paisId int64) ([]model.TareaModelWeb, error) {
	return t.tareaRepository.ObtenerTareasWeb(ctx, fechaInicio, fechaFinal, paisId)
}

func (t *tareaServiceImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	return t.tareaRepository.ObtenerTareasDelDia(ctx, fecha, usuarioId)
}

func (t *tareaServiceImpl) ObtenerCantidadTareasUsuarioPorFecha(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadTareaPorUsuario, error) {
	return t.tareaRepository.ObtenerCantidadTareasUsuarioPorFecha(ctx, fechaInicio, fechaFin)
}

func (t *tareaServiceImpl) CompletarTarea(ctx *gin.Context, tarea model.CompletarTareaModel, usuarioId int64) error {
	var urlImagen string

	visita := model.CreateVisitaModel{
		Comentario:   tarea.Comentario,
		Latitud:      tarea.Latitud,
		Longitud:     tarea.Longitud,
		Fecha:        tarea.Fecha,
		ClienteId:    tarea.ClienteId,
		Meta:         tarea.Meta,
		MetaLinea:    tarea.MetaLinea,
		MetaSubLinea: tarea.MetaSubLinea,
		Ip:           tarea.Ip,
	}

	visitaId, err := t.visitaRepository.CrearVisita(ctx.Request.Context(), visita, urlImagen, usuarioId)

	if err != nil {
		return err
	}

	var completada bool

	if *tarea.NecesitaAprobacion {
		completada = false
	} else {
		completada = true
	}

	println(completada)

	_, err = t.tareaRepository.CompletarTarea(ctx, tarea.TareaId, visitaId, completada, *tarea.NecesitaAprobacion)

	if err != nil {
		return err
	}

	if *tarea.ImagenRequerida {
		formfile, _, err := ctx.Request.FormFile("imagen")

		if err != nil {
			return err
		}

		cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDNAME"), os.Getenv("APIKEY"), os.Getenv("APISECRET"))
		resp, err := cld.Upload.Upload(ctx, formfile, uploader.UploadParams{Folder: os.Getenv("CLOUDFOLDER")})

		if err != nil {
			return err
		}

		urlImagen = resp.SecureURL

		_, err = t.visitaRepository.ActualizarVisitaImagen(ctx, visitaId, urlImagen)

		if err != nil {
			return err
		}

	}

	return nil
}

func (t *tareaServiceImpl) CrearTareaMasivaExcelWeb(ctx context.Context, tareaCreate model.CreateTareasExcelWeb, usuarioCrea int64, fechaCrea time.Time) error {

	for _, tarea := range tareaCreate.Tareas {
		tareaModel := model.CreateTareaModelWeb{
			ClienteId:       tarea.ClienteId,
			UsuarioId:       tarea.UsuarioId,
			Meta:            tarea.Meta,
			Fecha:           tarea.Fecha,
			ImagenRequerida: tarea.ImagenRequerida,
			TipoVisitaId:    tarea.TipoVisitaId,
			MetaLinea:       tarea.MetaLinea,
			MetaSubLinea:    tarea.MetaSubLinea,
			UsuarioCrea:     usuarioCrea,
			FechaCrea:       fechaCrea,
		}

		_, err := t.tareaRepository.CrearTareaWeb(ctx, tareaModel)

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *tareaServiceImpl) VerificarTarea(ctx context.Context, fecha string, usuarioId int64) (int, error) {
	return t.tareaRepository.VerificarTarea(ctx, fecha, usuarioId)
}

func (t *tareaServiceImpl) ValidarDataExcel(ctx *gin.Context, clienteId int64, usuarioId int64, tipoVisitaId int64, fecha string) (model.ValidarTareasExcelWeb, error) {

	clienteResult, err := t.clienteRepository.ObtenerClientePorId(ctx, clienteId)

	responsableResult, err := t.usuarioRepository.ObtenerUsuarioPorId(ctx, usuarioId)

	tipovisitaResult, err := t.tipoVisitaRepository.ObtenerTipoVisitaPorId(ctx, tipoVisitaId)

	tareaExiste, err := t.VerificarTarea(ctx, fecha, usuarioId)

	tarea := ""
	cliente := ""
	responsable := ""
	tipovisita := ""

	if tareaExiste > 0 {
		tarea = "La impulsadora(Responsable Id) ya tiene asignada una tarea para esta fecha y hora"
	}

	if !clienteResult {
		cliente = "Cliente(Cliente Id) no existe"
	}

	if !responsableResult {
		responsable = "Impulsadora(Responsable Id) no existe"
	}

	if !tipovisitaResult {
		tipovisita = "Tipo de visita(Tipo Visita Id) no existe"
	}

	validar := false

	if !clienteResult || !responsableResult || !tipovisitaResult || tarea != "" {
		validar = true
	}

	result := model.ValidarTareasExcelWeb{
		Cliente:     cliente,
		Responsable: responsable,
		TipoVisita:  tipovisita,
		Tarea:       tarea,
		Error:       validar,
	}

	return result, err
}

func (t *tareaServiceImpl) ObtenerTareasHorasWeb(ctx context.Context, parametros model.ParamReportTareasHoras) ([]model.TareaHorasModelReporteWeb, error) {
	tareas := []model.TareaHorasModelReporteWeb{}
	for _, usuarioId := range parametros.UsuarioId {
		for _, fecha := range parametros.Fechas {

			result, err := t.tareaRepository.ObtenerTareasHorasWeb(ctx, usuarioId, fecha)
			if err != nil {
				return result, err
			}
			for _, report := range result {
				tareas = append(tareas, report)
			}
		}
	}
	return tareas, nil
}

func (t *tareaServiceImpl) EliminarTareas(ctx context.Context, tareasId []int64) (int, error) {

	rowsDeleted := 0

	for _, tareaId := range tareasId {

		result, err := t.tareaRepository.EliminarTarea(ctx, tareaId)
		if err != nil {
			return int(result), err
		}
		rowsDeleted = rowsDeleted + int(result)
	}

	return rowsDeleted, nil
}

func (t *tareaServiceImpl) ObtenerTareasPorAprobar(ctx context.Context, fecha string, paisId int64) ([]model.AprobarTareas, error) {
	return t.tareaRepository.ObtenerTareasPorAprobar(ctx, fecha, paisId)
}

func (t *tareaServiceImpl) AprobarTarea(ctx context.Context, tarea model.CreateAprobarTarea) (bool, error) {
	return t.tareaRepository.AprobarTarea(ctx, tarea)
}

func (t *tareaServiceImpl) CantidadTareasPendientesAprobar(fecha string) (int64, error) {
	return t.tareaRepository.CantidadTareasPendientesAprobar(fecha)
}
