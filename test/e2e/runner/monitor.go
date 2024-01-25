package main

import (
	"context"

	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

func LogMetrics(ctx context.Context, testnet *e2e.Testnet, portals map[uint64]Portal, network netconf.Network) error {
	// Just pick the first node for now.
	if err := MonitorCProvider(ctx, testnet.Nodes[0], network); err != nil {
		return errors.Wrap(err, "monitoring cchain provider")
	}

	if err := MonitorCursors(ctx, portals, network); err != nil {
		return errors.Wrap(err, "monitoring cursors")
	}

	return nil
}

func MonitorCursors(ctx context.Context, portals map[uint64]Portal, network netconf.Network) error {
	for _, dest := range network.Chains {
		for _, src := range network.Chains {
			if src.ID == dest.ID {
				continue
			}

			offset, err := portals[dest.ID].Contract.InXStreamOffset(nil, src.ID)
			if err != nil {
				return errors.Wrap(err, "getting inXStreamOffset")
			}

			log.Info(ctx, "Submitted cross chain messages",
				"src", src.Name,
				"dest", dest.Name,
				"count", offset,
			)
		}
	}

	return nil
}

func MonitorCProvider(ctx context.Context, node *e2e.Node, network netconf.Network) error {
	client, err := node.Client()
	if err != nil {
		return errors.Wrap(err, "getting client")
	}

	cprov := cprovider.NewABCIProvider(client)

	for _, chain := range network.Chains {
		aggs, err := cprov.ApprovedFrom(ctx, chain.ID, 0)
		if err != nil {
			return errors.Wrap(err, "getting approved attestations")
		}

		log.Info(ctx, "Halo approved attestations", "chain", chain.Name, "count", len(aggs))
	}

	return nil
}