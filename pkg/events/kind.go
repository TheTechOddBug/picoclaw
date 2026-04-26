package events

const (
	// KindAgentTurnStart is emitted when an agent turn starts.
	KindAgentTurnStart Kind = "agent.turn.start"
	// KindAgentTurnEnd is emitted when an agent turn ends.
	KindAgentTurnEnd Kind = "agent.turn.end"

	// KindAgentLLMRequest is emitted before an LLM request.
	KindAgentLLMRequest Kind = "agent.llm.request"
	// KindAgentLLMDelta is emitted for streaming LLM deltas.
	KindAgentLLMDelta Kind = "agent.llm.delta"
	// KindAgentLLMResponse is emitted after an LLM response.
	KindAgentLLMResponse Kind = "agent.llm.response"
	// KindAgentLLMRetry is emitted before retrying an LLM request.
	KindAgentLLMRetry Kind = "agent.llm.retry"

	// KindAgentContextCompress is emitted when agent context is compressed.
	KindAgentContextCompress Kind = "agent.context.compress"
	// KindAgentSessionSummarize is emitted when session summarization completes.
	KindAgentSessionSummarize Kind = "agent.session.summarize"

	// KindAgentToolExecStart is emitted before a tool executes.
	KindAgentToolExecStart Kind = "agent.tool.exec_start"
	// KindAgentToolExecEnd is emitted after a tool finishes.
	KindAgentToolExecEnd Kind = "agent.tool.exec_end"
	// KindAgentToolExecSkipped is emitted when a tool call is skipped.
	KindAgentToolExecSkipped Kind = "agent.tool.exec_skipped"

	// KindAgentSteeringInjected is emitted when steering is injected into context.
	KindAgentSteeringInjected Kind = "agent.steering.injected"
	// KindAgentFollowUpQueued is emitted when async follow-up input is queued.
	KindAgentFollowUpQueued Kind = "agent.follow_up.queued"
	// KindAgentInterruptReceived is emitted when a turn interrupt is accepted.
	KindAgentInterruptReceived Kind = "agent.interrupt.received"

	// KindAgentSubTurnSpawn is emitted when a sub-turn is spawned.
	KindAgentSubTurnSpawn Kind = "agent.subturn.spawn"
	// KindAgentSubTurnEnd is emitted when a sub-turn ends.
	KindAgentSubTurnEnd Kind = "agent.subturn.end"
	// KindAgentSubTurnResultDelivered is emitted when a sub-turn result is delivered.
	KindAgentSubTurnResultDelivered Kind = "agent.subturn.result_delivered"
	// KindAgentSubTurnOrphan is emitted when a sub-turn result cannot be delivered.
	KindAgentSubTurnOrphan Kind = "agent.subturn.orphan"
	// KindAgentError is emitted when agent execution reports an error.
	KindAgentError Kind = "agent.error"

	// KindChannelLifecycleStarted is emitted when a channel starts.
	KindChannelLifecycleStarted Kind = "channel.lifecycle.started"
	// KindChannelLifecycleStartFailed is emitted when a channel fails to start.
	KindChannelLifecycleStartFailed Kind = "channel.lifecycle.start_failed"
	// KindChannelMessageOutboundSent is emitted when an outbound channel message is sent.
	KindChannelMessageOutboundSent Kind = "channel.message.outbound_sent"
	// KindChannelMessageOutboundFailed is emitted when an outbound channel message fails.
	KindChannelMessageOutboundFailed Kind = "channel.message.outbound_failed"

	// KindGatewayReloadStarted is emitted when gateway reload starts.
	KindGatewayReloadStarted Kind = "gateway.reload.started"
	// KindGatewayReloadCompleted is emitted when gateway reload completes.
	KindGatewayReloadCompleted Kind = "gateway.reload.completed"
	// KindGatewayReloadFailed is emitted when gateway reload fails.
	KindGatewayReloadFailed Kind = "gateway.reload.failed"

	// KindMCPServerConnected is emitted when an MCP server connects.
	KindMCPServerConnected Kind = "mcp.server.connected"
	// KindMCPServerFailed is emitted when an MCP server fails.
	KindMCPServerFailed Kind = "mcp.server.failed"
)
