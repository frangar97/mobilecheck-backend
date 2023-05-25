package service

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type TareaService interface {
	CrearTareaMovil(context.Context, model.CreateTareaModelMovil, int64) (model.TareaModelMovil, error)
	CrearTareaWeb(context.Context, model.CreateTareaModelWeb) (model.TareaModelWeb, error)
	ObtenerTareasWeb(context.Context, string, string) ([]model.TareaModelWeb, error)
	ObtenerTareasDelDia(context.Context, string, int64) ([]model.TareaModelMovil, error)
	ObtenerCantidadTareasUsuarioPorFecha(context.Context, string, string) ([]model.CantidadTareaPorUsuario, error)
	CompletarTarea(*gin.Context, model.CompletarTareaModel, int64) error
	CrearTareaMasivaWeb(context.Context, model.CreateTareaMasivaModelWeb) error
	CrearTareaMasivaExcelWeb(context.Context, model.CreateTareasExcelWeb) error
	VerificarTarea(context.Context, string, int64) (int, error)
	ValidarDataExcel(*gin.Context, int64, int64, int64, string) (model.ValidarTareasExcelWeb, error)
	ObtenerTareasHorasWeb(context.Context, model.ParamReportTareasHoras) ([]model.TareaHorasModelReporteWeb, error)
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

func (t *tareaServiceImpl) CrearTareaMasivaWeb(ctx context.Context, tareaCreate model.CreateTareaMasivaModelWeb) error {

	for _, fecha := range tareaCreate.Fechas {
		tareaModel := model.CreateTareaModelWeb{
			ClienteId:       tareaCreate.ClienteId,
			UsuarioId:       tareaCreate.UsuarioId,
			Meta:            tareaCreate.Meta,
			Fecha:           fecha,
			ImagenRequerida: tareaCreate.ImagenRequerida,
			TipoVisitaId:    tareaCreate.TipoVisitaId,
		}

		_, err := t.tareaRepository.CrearTareaWeb(ctx, tareaModel)

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *tareaServiceImpl) ObtenerTareasWeb(ctx context.Context, fechaInicio string, fechaFinal string) ([]model.TareaModelWeb, error) {
	return t.tareaRepository.ObtenerTareasWeb(ctx, fechaInicio, fechaFinal)
}

func (t *tareaServiceImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	return t.tareaRepository.ObtenerTareasDelDia(ctx, fecha, usuarioId)
}

func (t *tareaServiceImpl) ObtenerCantidadTareasUsuarioPorFecha(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadTareaPorUsuario, error) {
	return t.tareaRepository.ObtenerCantidadTareasUsuarioPorFecha(ctx, fechaInicio, fechaFin)
}

func (t *tareaServiceImpl) CompletarTarea(ctx *gin.Context, tarea model.CompletarTareaModel, usuarioId int64) error {
	var urlImagen string

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
	}

	visita := model.CreateVisitaModel{
		Comentario: tarea.Comentario,
		Latitud:    tarea.Latitud,
		Longitud:   tarea.Longitud,
		Fecha:      tarea.Fecha,
		ClienteId:  tarea.ClienteId,
		Meta:       tarea.Meta,
	}

	visitaId, err := t.visitaRepository.CrearVisita(ctx.Request.Context(), visita, urlImagen, usuarioId)

	if err != nil {
		return err
	}

	_, err = t.tareaRepository.CompletarTarea(ctx, tarea.TareaId, visitaId)

	if err != nil {
		return err
	}

	return nil
}

func (t *tareaServiceImpl) CrearTareaMasivaExcelWeb(ctx context.Context, tareaCreate model.CreateTareasExcelWeb) error {

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
			fmt.Println("------------------")
			fmt.Println(usuarioId)
			fmt.Println(fecha)

			var tarea model.TareaHorasModelReporteWeb
			result, err := t.tareaRepository.ObtenerTareasHorasWeb(ctx, usuarioId, fecha)
			if err != nil {
				return result, err
			}
			for _, report := range result {
				tarea = report
			}

			if tarea.Respomsable != "" {
				tareas = append(tareas, tarea)
			}
		}
	}
	return tareas, nil
}
