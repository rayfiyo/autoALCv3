// ノードの数を返す

package cmd

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func NodeCount(ctx context.Context, sel interface{}, opts ...chromedp.QueryOption) (int, error) {
	nodes := []*cdp.Node{}
	if err := chromedp.Run(ctx, chromedp.Nodes(sel, &nodes, opts...)); err != nil {
		return -1, xerrors.Errorf("Failed to get nodes count: %w", err)
	}
	return len(nodes), nil
}
