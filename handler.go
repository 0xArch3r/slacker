package slacker

import "github.com/slack-go/slack/socketmode"

// CommandMiddlewareHandler represents the command middleware handler function
type CommandMiddlewareHandler func(CommandHandler) CommandHandler

// CommandHandler represents the command handler function
type CommandHandler func(*CommandContext)

// InteractionMiddlewareHandler represents the interaction middleware handler function
type InteractionMiddlewareHandler func(InteractionHandler) InteractionHandler

// SuggestionMiddlewareHandler represents the suggestion middleware handler function
type SuggestionMiddlewareHandler func(SuggestionHandler) SuggestionHandler

// InteractionHandler represents the interaction handler function
type InteractionHandler func(sm socketmode.Event, ic *InteractionContext)

// SuggestionHandler represents the interaction handler function for block_suggestion
type SuggestionHandler func(sm socketmode.Event, ic *InteractionContext)

// JobMiddlewareHandler represents the job middleware handler function
type JobMiddlewareHandler func(JobHandler) JobHandler

// JobHandler represents the job handler function
type JobHandler func(*JobContext)
