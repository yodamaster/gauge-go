package messageprocessors

import (
	m "github.com/manuviswam/gauge-go/gauge_messages"
	t "github.com/manuviswam/gauge-go/testsuit"
)

type ScenarioExecutionStartingRequestProcessor struct{}

func (r *ScenarioExecutionStartingRequestProcessor) Process(msg *m.Message, context t.GaugeContext) *m.Message {
	tags := msg.GetScenarioExecutionStartingRequest().GetCurrentExecutionInfo().GetCurrentScenario().GetTags()
	hooks := context.GetHooks(t.BEFORESCENARIO, tags)

	executionTime, err := executeHooks(hooks)
	return createResponseMessage(msg.MessageId, executionTime, err)
}
