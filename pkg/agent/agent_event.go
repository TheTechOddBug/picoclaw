// PicoClaw - Ultra-lightweight personal AI agent

package agent

import (
	"fmt"

	runtimeevents "github.com/sipeed/picoclaw/pkg/events"
	"github.com/sipeed/picoclaw/pkg/logger"
)

func (al *AgentLoop) newTurnEventScope(agentID, sessionKey string, turnCtx *TurnContext) turnEventScope {
	seq := al.turnSeq.Add(1)
	return turnEventScope{
		agentID:    agentID,
		sessionKey: sessionKey,
		turnID:     fmt.Sprintf("%s-turn-%d", agentID, seq),
		context:    cloneTurnContext(turnCtx),
	}
}

func (ts turnEventScope) meta(iteration int, source, tracePath string) HookMeta {
	return HookMeta{
		AgentID:     ts.agentID,
		TurnID:      ts.turnID,
		SessionKey:  ts.sessionKey,
		Iteration:   iteration,
		Source:      source,
		TracePath:   tracePath,
		turnContext: cloneTurnContext(ts.context),
	}
}

func (al *AgentLoop) emitEvent(kind runtimeevents.Kind, meta HookMeta, payload any) {
	clonedMeta := cloneHookMeta(meta)
	eventCtx := cloneTurnContext(clonedMeta.turnContext)
	evt := runtimeevents.Event{
		Kind:        kind,
		Source:      runtimeevents.Source{Component: "agent", Name: clonedMeta.AgentID},
		Scope:       runtimeScopeFromHookMeta(clonedMeta, eventCtx),
		Correlation: runtimeCorrelationFromHookMeta(clonedMeta),
		Severity:    runtimeSeverityForAgentEvent(kind, payload),
		Payload:     payload,
		Attrs:       runtimeAttrsFromHookMeta(clonedMeta),
	}

	if al == nil {
		return
	}

	al.logEvent(evt, clonedMeta, eventCtx)

	al.publishRuntimeEvent(evt)
}

func (al *AgentLoop) logEvent(evt runtimeevents.Event, meta HookMeta, eventCtx *TurnContext) {
	fields := map[string]any{
		"event_kind":  evt.Kind.String(),
		"agent_id":    meta.AgentID,
		"turn_id":     meta.TurnID,
		"session_key": meta.SessionKey,
		"iteration":   meta.Iteration,
	}

	if meta.TracePath != "" {
		fields["trace"] = meta.TracePath
	}
	if meta.Source != "" {
		fields["source"] = meta.Source
	}

	appendEventContextFields(fields, eventCtx)

	switch payload := evt.Payload.(type) {
	case TurnStartPayload:
		fields["user_len"] = len(payload.UserMessage)
		fields["media_count"] = payload.MediaCount
	case TurnEndPayload:
		fields["status"] = payload.Status
		fields["iterations_total"] = payload.Iterations
		fields["duration_ms"] = payload.Duration.Milliseconds()
		fields["final_len"] = payload.FinalContentLen
	case LLMRequestPayload:
		fields["model"] = payload.Model
		fields["messages"] = payload.MessagesCount
		fields["tools"] = payload.ToolsCount
		fields["max_tokens"] = payload.MaxTokens
	case LLMDeltaPayload:
		fields["content_delta_len"] = payload.ContentDeltaLen
		fields["reasoning_delta_len"] = payload.ReasoningDeltaLen
	case LLMResponsePayload:
		fields["content_len"] = payload.ContentLen
		fields["tool_calls"] = payload.ToolCalls
		fields["has_reasoning"] = payload.HasReasoning
	case LLMRetryPayload:
		fields["attempt"] = payload.Attempt
		fields["max_retries"] = payload.MaxRetries
		fields["reason"] = payload.Reason
		fields["error"] = payload.Error
		fields["backoff_ms"] = payload.Backoff.Milliseconds()
	case ContextCompressPayload:
		fields["reason"] = payload.Reason
		fields["dropped_messages"] = payload.DroppedMessages
		fields["remaining_messages"] = payload.RemainingMessages
	case SessionSummarizePayload:
		fields["summarized_messages"] = payload.SummarizedMessages
		fields["kept_messages"] = payload.KeptMessages
		fields["summary_len"] = payload.SummaryLen
		fields["omitted_oversized"] = payload.OmittedOversized
	case ToolExecStartPayload:
		fields["tool"] = payload.Tool
		fields["args_count"] = len(payload.Arguments)
	case ToolExecEndPayload:
		fields["tool"] = payload.Tool
		fields["duration_ms"] = payload.Duration.Milliseconds()
		fields["for_llm_len"] = payload.ForLLMLen
		fields["for_user_len"] = payload.ForUserLen
		fields["is_error"] = payload.IsError
		fields["async"] = payload.Async
	case ToolExecSkippedPayload:
		fields["tool"] = payload.Tool
		fields["reason"] = payload.Reason
	case SteeringInjectedPayload:
		fields["count"] = payload.Count
		fields["total_content_len"] = payload.TotalContentLen
	case FollowUpQueuedPayload:
		fields["source_tool"] = payload.SourceTool
		fields["content_len"] = payload.ContentLen
	case InterruptReceivedPayload:
		fields["interrupt_kind"] = payload.Kind
		fields["role"] = payload.Role
		fields["content_len"] = payload.ContentLen
		fields["queue_depth"] = payload.QueueDepth
		fields["hint_len"] = payload.HintLen
	case SubTurnSpawnPayload:
		fields["child_agent_id"] = payload.AgentID
		fields["label"] = payload.Label
	case SubTurnEndPayload:
		fields["child_agent_id"] = payload.AgentID
		fields["status"] = payload.Status
	case SubTurnResultDeliveredPayload:
		fields["target_channel"] = payload.TargetChannel
		fields["target_chat_id"] = payload.TargetChatID
		fields["content_len"] = payload.ContentLen
	case ErrorPayload:
		fields["stage"] = payload.Stage
		fields["error"] = payload.Message
	}

	logger.InfoCF("eventbus", fmt.Sprintf("Agent event: %s", evt.Kind.String()), fields)
}

// MountHook registers an in-process hook on the agent loop.
func (al *AgentLoop) MountHook(reg HookRegistration) error {
	if al == nil || al.hooks == nil {
		return fmt.Errorf("hook manager is not initialized")
	}
	return al.hooks.Mount(reg)
}

// UnmountHook removes a previously registered in-process hook.
func (al *AgentLoop) UnmountHook(name string) {
	if al == nil || al.hooks == nil {
		return
	}
	al.hooks.Unmount(name)
}

// RuntimeEvents returns the root runtime event channel.
func (al *AgentLoop) RuntimeEvents() runtimeevents.EventChannel {
	if al == nil || al.runtimeEvents == nil {
		return nil
	}
	return al.runtimeEvents.Channel()
}

// RuntimeEventStats returns runtime event bus counters.
func (al *AgentLoop) RuntimeEventStats() runtimeevents.Stats {
	if al == nil || al.runtimeEvents == nil {
		return runtimeevents.Stats{Closed: true}
	}
	return al.runtimeEvents.Stats()
}

// RuntimeEventBus returns the runtime event bus used by the agent loop.
func (al *AgentLoop) RuntimeEventBus() runtimeevents.Bus {
	if al == nil {
		return nil
	}
	return al.runtimeEvents
}
