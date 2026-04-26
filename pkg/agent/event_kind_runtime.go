package agent

import runtimeevents "github.com/sipeed/picoclaw/pkg/events"

func runtimeKindForAgentEvent(kind EventKind) runtimeevents.Kind {
	switch kind {
	case EventKindTurnStart:
		return runtimeevents.KindAgentTurnStart
	case EventKindTurnEnd:
		return runtimeevents.KindAgentTurnEnd
	case EventKindLLMRequest:
		return runtimeevents.KindAgentLLMRequest
	case EventKindLLMDelta:
		return runtimeevents.KindAgentLLMDelta
	case EventKindLLMResponse:
		return runtimeevents.KindAgentLLMResponse
	case EventKindLLMRetry:
		return runtimeevents.KindAgentLLMRetry
	case EventKindContextCompress:
		return runtimeevents.KindAgentContextCompress
	case EventKindSessionSummarize:
		return runtimeevents.KindAgentSessionSummarize
	case EventKindToolExecStart:
		return runtimeevents.KindAgentToolExecStart
	case EventKindToolExecEnd:
		return runtimeevents.KindAgentToolExecEnd
	case EventKindToolExecSkipped:
		return runtimeevents.KindAgentToolExecSkipped
	case EventKindSteeringInjected:
		return runtimeevents.KindAgentSteeringInjected
	case EventKindFollowUpQueued:
		return runtimeevents.KindAgentFollowUpQueued
	case EventKindInterruptReceived:
		return runtimeevents.KindAgentInterruptReceived
	case EventKindSubTurnSpawn:
		return runtimeevents.KindAgentSubTurnSpawn
	case EventKindSubTurnEnd:
		return runtimeevents.KindAgentSubTurnEnd
	case EventKindSubTurnResultDelivered:
		return runtimeevents.KindAgentSubTurnResultDelivered
	case EventKindSubTurnOrphan:
		return runtimeevents.KindAgentSubTurnOrphan
	case EventKindError:
		return runtimeevents.KindAgentError
	default:
		return runtimeevents.Kind("agent." + kind.String())
	}
}
