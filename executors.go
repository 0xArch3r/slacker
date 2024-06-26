package slacker

import "github.com/slack-go/slack/socketmode"

func executeCommand(ctx *CommandContext, handler CommandHandler, middlewares ...CommandMiddlewareHandler) {
	if handler == nil {
		return
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	handler(ctx)
}

func executeInteraction(event *socketmode.Event, ctx *InteractionContext, handler InteractionHandler, middlewares ...InteractionMiddlewareHandler) {
	if handler == nil {
		return
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	handler(*event, ctx)
}

func executeSuggestion(socketEvent socketmode.Event, ctx *InteractionContext, handler SuggestionHandler, middlewares ...SuggestionMiddlewareHandler) {
	if handler == nil {
		return
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	handler(socketEvent, ctx)
}

func executeJob(ctx *JobContext, handler JobHandler, middlewares ...JobMiddlewareHandler) func() {
	if handler == nil {
		return func() {}
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return func() {
		handler(ctx)
	}
}
