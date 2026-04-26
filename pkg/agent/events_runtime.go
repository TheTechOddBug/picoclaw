package agent

import (
	"context"
	"time"

	runtimeevents "github.com/sipeed/picoclaw/pkg/events"
)

const runtimeEventPublishTimeout = 100 * time.Millisecond

func (al *AgentLoop) publishRuntimeEvent(evt Event) {
	if al == nil || al.runtimeEvents == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), runtimeEventPublishTimeout)
	defer cancel()

	al.runtimeEvents.Publish(ctx, runtimeevents.Event{
		Kind:        runtimeKindForAgentEvent(evt.Kind),
		Source:      runtimeevents.Source{Component: "agent", Name: evt.Meta.AgentID},
		Scope:       runtimeScopeFromAgentEvent(evt),
		Correlation: runtimeCorrelationFromAgentEvent(evt),
		Severity:    runtimeSeverityForAgentEvent(evt),
		Payload:     evt.Payload,
		Attrs:       runtimeAttrsFromAgentEvent(evt),
	})
}

func runtimeScopeFromAgentEvent(evt Event) runtimeevents.Scope {
	scope := runtimeevents.Scope{
		AgentID:    evt.Meta.AgentID,
		SessionKey: evt.Meta.SessionKey,
		TurnID:     evt.Meta.TurnID,
	}

	if evt.Context == nil || evt.Context.Inbound == nil {
		return scope
	}

	inbound := evt.Context.Inbound
	scope.Channel = inbound.Channel
	scope.Account = inbound.Account
	scope.ChatID = inbound.ChatID
	scope.TopicID = inbound.TopicID
	scope.SpaceID = inbound.SpaceID
	scope.SpaceType = inbound.SpaceType
	scope.ChatType = inbound.ChatType
	scope.SenderID = inbound.SenderID
	scope.MessageID = inbound.MessageID
	return scope
}

func runtimeCorrelationFromAgentEvent(evt Event) runtimeevents.Correlation {
	return runtimeevents.Correlation{
		TraceID:      evt.Meta.TracePath,
		ParentTurnID: evt.Meta.ParentTurnID,
	}
}

func runtimeSeverityForAgentEvent(evt Event) runtimeevents.Severity {
	switch evt.Kind {
	case EventKindError, EventKindSubTurnOrphan:
		return runtimeevents.SeverityError
	case EventKindLLMRetry, EventKindContextCompress, EventKindToolExecSkipped:
		return runtimeevents.SeverityWarn
	case EventKindTurnEnd:
		payload, ok := evt.Payload.(TurnEndPayload)
		if !ok {
			return runtimeevents.SeverityInfo
		}
		switch payload.Status {
		case TurnEndStatusError:
			return runtimeevents.SeverityError
		case TurnEndStatusAborted:
			return runtimeevents.SeverityWarn
		default:
			return runtimeevents.SeverityInfo
		}
	case EventKindToolExecEnd:
		payload, ok := evt.Payload.(ToolExecEndPayload)
		if ok && payload.IsError {
			return runtimeevents.SeverityWarn
		}
		return runtimeevents.SeverityInfo
	default:
		return runtimeevents.SeverityInfo
	}
}

func runtimeAttrsFromAgentEvent(evt Event) map[string]any {
	attrs := make(map[string]any, 2)
	if evt.Meta.Source != "" {
		attrs["agent_source"] = evt.Meta.Source
	}
	if evt.Meta.Iteration != 0 {
		attrs["iteration"] = evt.Meta.Iteration
	}
	if len(attrs) == 0 {
		return nil
	}
	return attrs
}
