package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"scientific-workflow-provenance/internal/artifact"
	"scientific-workflow-provenance/internal/domain/provenance"
	"scientific-workflow-provenance/internal/domain/sample"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/internal/planner"
	"scientific-workflow-provenance/internal/repository"
	"scientific-workflow-provenance/internal/scheduler"
	"scientific-workflow-provenance/internal/worker"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	workflows *repository.WorkflowRepo
	runs      *repository.RunRepo
	samples   *sample.Store
	artifacts *artifact.Store
	prov      *provenance.Graph
	planner   planner.Planner
	queue     *scheduler.Queue
	worker    worker.Adapter
}

func NewServer() *Server {
	return &Server{workflows: repository.NewWorkflowRepo(), runs: repository.NewRunRepo(), samples: sample.NewStore(), artifacts: artifact.NewStore(), prov: provenance.New(), queue: scheduler.New(8), worker: worker.Simulator{}}
}
func (s *Server) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.health)
	mux.HandleFunc("/readyz", s.health)
	mux.HandleFunc("/api/v1/workflows", s.workflowsHandler)
	mux.HandleFunc("/api/v1/workflows/", s.workflowAction)
	mux.HandleFunc("/api/v1/samples", s.samplesHandler)
	mux.HandleFunc("/api/v1/samples/", s.sampleAction)
	mux.HandleFunc("/api/v1/runs", s.runsHandler)
	mux.HandleFunc("/api/v1/runs/", s.runAction)
	mux.HandleFunc("/api/v1/artifacts", s.artifactsHandler)
	mux.HandleFunc("/api/v1/artifacts/", s.artifactAction)
	mux.HandleFunc("/api/v1/provenance/", s.provenanceAction)
	return logging(mux)
}
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-ID", strconv.FormatInt(time.Now().UnixNano(), 10))
		next.ServeHTTP(w, r)
	})
}
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	write(w, 200, map[string]any{"status": "ok", "time": time.Now().UTC()})
}
func decode(r *http.Request, v any) error { return json.NewDecoder(r.Body).Decode(v) }
func write(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeRun(w http.ResponseWriter, status int, run workflow.Run) {
	write(w, status, run)
}
func errWrite(w http.ResponseWriter, status int, err error) {
	write(w, status, map[string]any{"error": map[string]any{"code": http.StatusText(status), "message": err.Error()}})
}
func (s *Server) workflowsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		write(w, 200, s.workflows.List())
	case http.MethodPost:
		var d workflow.Definition
		if err := decode(r, &d); err != nil {
			errWrite(w, 400, err)
			return
		}
		if d.CreatedAt.IsZero() {
			d.CreatedAt = time.Now().UTC()
		}
		if err := s.workflows.Put(d); err != nil {
			errWrite(w, 422, err)
			return
		}
		d, _ = s.workflows.Get(d.ID, d.Version)
		write(w, 201, d)
	default:
		errWrite(w, 405, context.Canceled)
	}
}
func (s *Server) workflowAction(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 4 {
		errWrite(w, 404, context.Canceled)
		return
	}
	id := parts[3]
	if len(parts) == 4 && r.Method == http.MethodGet {
		v, _ := strconv.Atoi(r.URL.Query().Get("version"))
		d, err := s.workflows.Get(id, v)
		if err != nil {
			errWrite(w, 404, err)
			return
		}
		write(w, 200, d)
		return
	}
	if len(parts) >= 5 && parts[4] == "validate" {
		d, err := s.workflows.Get(id, 0)
		if err != nil {
			errWrite(w, 404, err)
			return
		}
		err = d.Validate()
		if err != nil {
			errWrite(w, 422, err)
			return
		}
		write(w, 200, map[string]any{"valid": true, "digest": d.ComputeDigest()})
		return
	}
	if len(parts) >= 5 && parts[4] == "plan" {
		d, err := s.workflows.Get(id, 0)
		if err != nil {
			errWrite(w, 404, err)
			return
		}
		var p map[string]string
		_ = decode(r, &p)
		pl, err := s.planner.Build(d, p)
		if err != nil {
			errWrite(w, 422, err)
			return
		}
		write(w, 200, pl)
		return
	}
	errWrite(w, 404, context.Canceled)
}
func (s *Server) samplesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		write(w, 200, s.samples.List())
		return
	}
	if r.Method != http.MethodPost {
		errWrite(w, 405, context.Canceled)
		return
	}
	var v sample.Sample
	if err := decode(r, &v); err != nil {
		errWrite(w, 400, err)
		return
	}
	if err := s.samples.Put(v); err != nil {
		errWrite(w, 422, err)
		return
	}
	s.prov.Add(provenance.Record{ID: "sample:" + v.ID, Kind: provenance.Entity, Type: "sample", Attributes: map[string]any{"name": v.Name, "batch": v.Batch}})
	write(w, 201, v)
}
func (s *Server) sampleAction(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 4 {
		errWrite(w, 404, context.Canceled)
		return
	}
	id := parts[3]
	if len(parts) == 4 {
		v, err := s.samples.Get(id)
		if err != nil {
			errWrite(w, 404, err)
			return
		}
		write(w, 200, v)
		return
	}
	switch parts[4] {
	case "lineage":
		write(w, 200, s.prov.Lineage("sample:"+id))
	case "freeze":
		if err := s.samples.Freeze(id); err != nil {
			errWrite(w, 409, err)
			return
		}
		write(w, 200, map[string]string{"status": "frozen"})
	case "destroy":
		if err := s.samples.Destroy(id); err != nil {
			errWrite(w, 409, err)
			return
		}
		write(w, 200, map[string]string{"status": "destroyed"})
	default:
		errWrite(w, 404, context.Canceled)
	}
}
func (s *Server) runsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		write(w, 200, s.runs.List())
		return
	}
	if r.Method != http.MethodPost {
		errWrite(w, 405, context.Canceled)
		return
	}
	var in struct {
		ID, WorkflowID string
		Version        int
		Parameters     map[string]string
	}
	if err := decode(r, &in); err != nil {
		errWrite(w, 400, err)
		return
	}
	d, err := s.workflows.Get(in.WorkflowID, in.Version)
	if err != nil {
		errWrite(w, 404, err)
		return
	}
	pl, err := s.planner.Build(d, in.Parameters)
	if err != nil {
		errWrite(w, 422, err)
		return
	}
	run := workflow.Run{ID: in.ID, WorkflowID: d.ID, Version: d.Version, Status: workflow.Planned, NodeStates: map[string]workflow.RunStatus{}, Attempts: map[string]int{}, Parameters: in.Parameters, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()}
	if run.ID == "" {
		run.ID = "run-" + strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	for _, st := range pl.Steps {
		run.NodeStates[st.ID] = workflow.Queued
	}
	s.runs.Put(run)
	write(w, 201, run)
}
func (s *Server) runAction(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 4 {
		errWrite(w, 404, context.Canceled)
		return
	}
	id := parts[3]
	run, ok := s.runs.Get(id)
	if !ok {
		errWrite(w, 404, context.Canceled)
		return
	}
	if len(parts) == 4 {
		writeRun(w, 200, run)
		return
	}
	switch parts[4] {
	case "cancel":
		run.Status = workflow.Cancelled
		run.UpdatedAt = time.Now().UTC()
		s.runs.Put(run)
		write(w, 200, run)
	case "provenance":
		write(w, 200, s.prov.Export())
	default:
		errWrite(w, 404, context.Canceled)
	}
}
func (s *Server) artifactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		write(w, 200, s.artifacts.List())
		return
	}
	if r.Method != http.MethodPost {
		errWrite(w, 405, context.Canceled)
		return
	}
	var a artifact.Artifact
	if err := decode(r, &a); err != nil {
		errWrite(w, 400, err)
		return
	}
	if err := s.artifacts.Put(a); err != nil {
		errWrite(w, 422, err)
		return
	}
	s.prov.Add(provenance.Record{ID: "artifact:" + a.Digest, Kind: provenance.Entity, Type: "artifact", Attributes: map[string]any{"schema": a.Schema, "unit": a.Unit}})
	write(w, 201, a)
}
func (s *Server) artifactAction(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 4 {
		errWrite(w, 404, context.Canceled)
		return
	}
	d := parts[3]
	a, ok := s.artifacts.Get(d)
	if !ok {
		errWrite(w, 404, context.Canceled)
		return
	}
	if len(parts) == 4 {
		write(w, 200, a)
		return
	}
	if parts[4] == "invalidate" {
		a.Status = artifact.Invalidated
		_ = s.artifacts.Invalidate(d, "manual invalidation")
		write(w, 200, a)
		return
	}
	if parts[4] == "verify" {
		write(w, 200, map[string]any{"digest": d, "valid": a.Status == artifact.Published})
		return
	}
	errWrite(w, 404, context.Canceled)
}
func (s *Server) provenanceAction(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 4 {
		errWrite(w, 404, context.Canceled)
		return
	}
	write(w, 200, s.prov.Lineage(parts[3]))
}

var _ = log.Printf
