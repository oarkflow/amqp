package grabbit

import (
	"context"
)

// ChanUsageParameters embeds [PublisherUsageOptions] and [ConsumerUsageOptions].
// It is a private member of the ChannelOptions and cen be passed
// via [WithChannelUsageParams].
type ChanUsageParameters struct {
	PublisherUsageOptions
	ConsumerUsageOptions
}

// ChannelOptions represents the options for configuring a channel.
type ChannelOptions struct {
	notifier          chan Event              // feedback channel
	name              string                  // tag for this channel
	delayer           DelayProvider           // how much to wait between re-attempts
	cbDown            CallbackWhenDown        // callback on conn lost
	cbUp              CallbackWhenUp          // callback when conn recovered
	cbReconnect       CallbackWhenRecovering  // callback when recovering
	cbNotifyPublish   CallbackNotifyPublish   // publish notification handler
	cbNotifyReturn    CallbackNotifyReturn    // returned message notification handler
	cbProcessMessages CallbackProcessMessages // user defined message processing routine
	topology          []*TopologyOptions      // the _whole_ infrastructure involved as array of queues and exchanges
	implParams        ChanUsageParameters     // implementation trigger for publishers or consumers
	ctx               context.Context         // cancellation context
	cancelCtx         context.CancelFunc      // aborts the reconnect loop
}

// OnChannelDown returns a function that sets the callback function to be called when the channel is down.
//
// down - The callback function to be called when the channel is down.
// options - The ChannelOptions object to be modified.
func OnChannelDown(down CallbackWhenDown) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.cbDown = down
	}
}

// OnChannelUp returns a function that sets the callback function
// to be executed when the channel is up.
//
// up: the callback function to be executed when the channel is up.
// options: the ChannelOptions to be modified.
//
// returns: a function that modifies the ChannelOptions by setting the
// callback function to be executed when the channel is up.
func OnChannelUp(up CallbackWhenUp) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.cbUp = up
	}
}

// OnChannelRecovering generates a function that sets the callback function to be called when recovering from an error in the ChannelOptions.
//
// Parameters:
//   - recover: a CallbackWhenRecovering function that will be called when recovering from an error in the ChannelOptions.
//
// Returns:
//   - A function that takes a pointer to ChannelOptions and sets the cbReconnect field to the provided recover function.
func OnChannelRecovering(recover CallbackWhenRecovering) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.cbReconnect = recover
	}
}

// WithChannelCtx creates a function that sets the context of a ChannelOptions struct.
//
// It takes a context.Context as a parameter and returns a function that takes a pointer to a ChannelOptions struct.
// The returned function sets the ctx field of the ChannelOptions struct to the provided context.
func WithChannelCtx(ctx context.Context) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.ctx = ctx
	}
}

// WithChannelDelay returns a function that sets the "delayer" field of the ChannelOptions struct to the given DelayProvider.
//
// Parameters:
// - delayer: The DelayProvider that will be set as the "delayer" field of ChannelOptions.
//
// Return type: A function that takes a pointer to a ChannelOptions struct as its parameter.
func WithChannelDelay(delayer DelayProvider) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.delayer = delayer
	}
}

// WithChannelName creates a function that sets the name field of the ChannelOptions struct.
//
// It takes a string parameter 'name' and returns a function that takes a pointer to the ChannelOptions struct as a parameter.
func WithChannelName(name string) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.name = name
	}
}

// WithChannelNotification provides an application defined
// [Event] receiver to handle various alerts about the channel status.
func WithChannelNotification(ch chan Event) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.notifier = ch
	}
}

// WithChannelTopology returns a function that sets the topology options for a channel.
//
// The function takes a slice of TopologyOptions as a parameter, which specifies the desired topology for the channel.
// It returns a function that takes a pointer to a ChannelOptions struct as a parameter.
// The function sets the topology field of the ChannelOptions struct to the provided topology slice.
func WithChannelTopology(topology []*TopologyOptions) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.topology = topology
	}
}

// OnPublishSuccess returns a function that sets the callback function
// for notifying the publish event in the ChannelOptions.
//
// It takes a single parameter:
// - publishNotifier: the callback function for notifying the publish event.
//
// It returns a function that takes a pointer to ChannelOptions as a parameter.
func OnPublishSuccess(publishNotifier CallbackNotifyPublish) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.cbNotifyPublish = publishNotifier
	}
}

// OnPublishFailure generates a function that sets the returnNotifier
// callback for a ChannelOptions struct.
//
// It takes a returnNotifier parameter of type CallbackNotifyReturn which represents
// a function that will be called when a return value is received.
//
// The generated function takes an options parameter of type *ChannelOptions and sets
// the cbNotifyReturn field to the provided returnNotifier.
func OnPublishFailure(returnNotifier CallbackNotifyReturn) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.cbNotifyReturn = returnNotifier
	}
}

// WithChannelProcessor is a function that returns a function which sets the callback
// process messages for the ChannelOptions struct.
//
// The parameter `proc` is a CallbackProcessMessages function that will be assigned to the
// `cbProcessMessages` field of the `ChannelOptions` struct.
//
// The return type of the returned function is `func(options *ChannelOptions)`.
func WithChannelProcessor(proc CallbackProcessMessages) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.cbProcessMessages = proc
	}
}

// WithChannelUsageParams returns a function that sets the implementation parameters of the ChannelOptions struct.
//
// It takes a parameter of type ChanUsageParameters and returns a function that takes a pointer to a ChannelOptions struct.
func WithChannelUsageParams(params ChanUsageParameters) func(options *ChannelOptions) {
	return func(options *ChannelOptions) {
		options.implParams = params

	}
}
