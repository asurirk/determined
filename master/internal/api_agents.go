package internal

import (
	"context"
	"fmt"

	"github.com/determined-ai/determined/master/pkg/proto/apiv1"
)

func (a *apiServer) GetAgents(
	_ context.Context, req *apiv1.GetAgentsRequest) (resp *apiv1.GetAgentsResponse, err error) {
	err = a.actorRequest("/agents", req, &resp)
	return resp, err
}

func (a *apiServer) GetAgent(
	_ context.Context, req *apiv1.GetAgentRequest) (resp *apiv1.GetAgentResponse, err error) {
	err = a.actorRequest(fmt.Sprintf("/agents/%s", req.AgentId), req, &resp)
	return resp, err
}

func (a *apiServer) GetSlots(
	_ context.Context, req *apiv1.GetSlotsRequest) (resp *apiv1.GetSlotsResponse, err error) {
	err = a.actorRequest(fmt.Sprintf("/agents/%s", req.AgentId), req, &resp)
	return resp, err
}

func (a *apiServer) GetSlot(
	_ context.Context, req *apiv1.GetSlotRequest) (resp *apiv1.GetSlotResponse, err error) {
	err = a.actorRequest(fmt.Sprintf("/agents/%s/slots/%s", req.AgentId, req.SlotId), req, &resp)
	return resp, err
}

func (a *apiServer) EnableAgent(
	_ context.Context, req *apiv1.EnableAgentRequest) (resp *apiv1.EnableAgentResponse, err error) {
	err = a.actorRequest(fmt.Sprintf("/agents/%s", req.AgentId), req, &resp)
	return resp, err
}

func (a *apiServer) DisableAgent(
	_ context.Context, req *apiv1.DisableAgentRequest) (resp *apiv1.DisableAgentResponse, err error) {
	err = a.actorRequest(fmt.Sprintf("/agents/%s", req.AgentId), req, &resp)
	return resp, err
}

func (a *apiServer) EnableSlot(
	_ context.Context, req *apiv1.EnableSlotRequest) (resp *apiv1.EnableSlotResponse, err error) {
	err = a.actorRequest(fmt.Sprintf("/agents/%s/slots/%s", req.AgentId, req.SlotId), req, &resp)
	return resp, err
}

func (a *apiServer) DisableSlot(
	_ context.Context, req *apiv1.DisableSlotRequest) (resp *apiv1.DisableSlotResponse, err error) {
	err = a.actorRequest(fmt.Sprintf("/agents/%s/slots/%s", req.AgentId, req.SlotId), req, &resp)
	return resp, err
}