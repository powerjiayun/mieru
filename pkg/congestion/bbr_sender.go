// Copyright (C) 2024  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package congestion

import "time"

type bbrMode int

const (
	// Startup phase of the connection.
	modeStartUp bbrMode = iota

	// After achieving the highest possible bandwidth during the startup, lower
	// the pacing rate in order to drain the queue.
	modeDrain

	// Cruising mode.
	modeProbeBW

	// Temporarily slow down sending in order to empty the buffer and measure
	// the real minimum RTT.
	modeProbeRTT
)

// Indicates how the congestion control limits the amount of bytes in flight.
type bbrRecoveryState int

const (
	// Do not limit.
	stateNotInRecovery bbrRecoveryState = iota

	// Allow 1 extra outstanding byte for each byte acknowledged.
	stateConservation

	// Allow 1.5 extra outstanding bytes for each byte acknowledged.
	stateMediumGrowth

	// Allow 2 extra outstanding bytes for each byte acknowledged (slow start).
	stateGrowth
)

type BBRSender struct {
	rttStats *RTTStats

	// unackedPackets *QuicUnackedPacketMap

	// Current BBR running mode.
	mode bbrMode

	// Bandwidth sampler provides BBR with the bandwidth measurements at
	// individual points.
	sampler BandwidthSamplerInterface

	// The number of the round trips that have occurred during the connection.
	roundTripCount uint64

	// The packet number of the most recently sent packet.
	lastSentPacket int64

	// Acknowledgement of any packet after currentRoundTripEnd will cause
	// the round trip counter to advance.
	currentRoundTripEnd int64

	// Tracks the maximum bandwidth over the multiple recent round-trips.
	maxBandwidth WindowedFilter[uint64]

	// Tracks the maximum number of bytes acked faster than the sending rate.
	maxAckHeight WindowedFilter[uint64]

	// The time this aggregation started and the number of bytes acked during it.
	aggregationEpochStartTime time.Time
	aggregationEpochBytes     int64

	// The number of bytes acknowledged since the last time bytes in flight
	// dropped below the target window.
	bytesAckedSinceQueueDrained int64

	// The muliplier for calculating the max amount of extra CWND to add to
	// compensate for ack aggregation.
	maxAggregationBytesMultiplier float64

	// Minimum RTT estimate. Automatically expires within 10 seconds
	// (and triggers PROBE_RTT mode) if no new value is sampled
	// during that period.
	minRTT time.Duration

	// The time at which the current value of minRTT was assigned.
	minRTTTimestamp time.Time

	// The maximum allowed number of bytes in flight.
	congestionWindow int64

	// The initial value of the congestionWindow.
	initialCongestionWindow int64

	// The largest value the congestionWindow can achieve.
	maxCongestionWindow int64

	// The smallest value the congestionWindow can achieve.
	minCongestionWindow int64

	// The current pacing rate of the connection.
	pacingRate int64

	// The gain currently applied to the pacing rate.
	pacingGain float64

	// The gain currently applied to the congestion window.
	congestionWindowGain float64

	// The gain used for the congestion window during PROBE_BW.
	congestionWindowGainConstant float64

	// The coefficient by which mean RTT variance is added to the congestion window.
	rttVarianceWeight float64

	// The number of RTTs to stay in STARTUP mode. Defaults to 3.
	numStartupRTTs uint64

	// If true, exit startup if 1 RTT has passed with no bandwidth increase and
	// the connection is in recovery.
	exitStartupOnLoss bool

	// Number of round-trips in PROBE_BW mode, used for determining the current
	// pacing gain cycle.
	cycleCurrentOffset int

	// The time at which the last pacing gain cycle was started.
	lastCycleStart time.Time

	// Indicates whether the connection has reached the full bandwidth mode.
	isAtFullBandwidth bool

	// Number of rounds during which there was no significant bandwidth increase.
	roundsWithoutBandwidthGain uint64

	// The bandwidth compared to which the increase is measured.
	bandwidthAtLastRound int64

	// Set to true upon exiting quiescence.
	exitingQuiescence bool

	// Time at which PROBE_RTT has to be exited. Setting it to zero indicates
	// that the time is yet unknown as the number of packets in flight has not
	// reached the required value.
	exitProbeRTTAt time.Time

	// Indicates whether a round-trip has passed since PROBE_RTT became active.
	probeRTTRoundPassed bool

	// Indicates whether the most recent bandwidth sample was marked as
	// app-limited.
	lastSampleIsAppLimited bool

	// Current state of recovery.
	recoveryState bbrRecoveryState

	// Receiving acknowledgement of a packet after endRecoveryAt will cause
	// BBR to exit the recovery mode. A value above zero indicates at least one
	// loss has been detected, so it must not be set back to zero.
	endRecoveryAt int64

	// A window used to limit the number of bytes in flight during loss recovery.
	recoveryWindow int64

	// When true, recovery is rate based rather than congestion window based.
	rateBasedRecovery bool

	// When true, pace at 1.5x and disable packet conservation in STARTUP.
	slowerStartup bool

	// When true, disables packet conservation in STARTUP.
	rateBasedStartup bool

	// Used as the initial packet conservation mode when first entering recovery.
	initialConservationInStartup bbrRecoveryState

	// If true, will not exit low gain mode until bytesInFlight drops below BDP
	// or it's time for high gain mode.
	fullyDrainQueue bool

	// If true, use a CWND of 0.75*BDP during probe_rtt instead of 4 packets.
	probeRTTBasedOnBDP bool

	// If true, skip PROBE_RTT and update the timestamp of the existing minRTT to
	// now if minRTT over the last cycle is within 12.5% of the current minRTT.
	// Even if the minRTT is 12.5% too low, the 25% gain cycling and 2x CWND gain
	// should overcome an overly small minRTT.
	probeRTTSkippedIfSimilarRTT bool

	// If true, disable PROBE_RTT entirely as long as the connection was recently
	// app limited.
	probeRTTDisabledIfAppLimited bool

	appLimitedSinceLastProbeRTT bool
	minRTTSinceLastProbeRTT     time.Duration
}
