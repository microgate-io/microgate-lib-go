package db

import (
	context "context"

	apilog "github.com/microgate-io/microgate-lib-go/v1/log"
	grpc "google.golang.org/grpc"
)

func InTransactionDo(
	ctx context.Context,
	dbClient DatabaseServiceClient,
	work func(transaction_id string) error,
	opts ...grpc.CallOption) error {

	// begin
	bResp, err := dbClient.Begin(ctx, new(BeginRequest))
	if err != nil {
		return apilog.ErrorWithLog(ctx, err, "failed to begin db transaction", "db", "begin")
	}

	err = work(bResp.TransactionId)

	if err != nil {
		// rollback
		apilog.Debugw(ctx, "Rollback")
		_, _ = dbClient.Rollback(ctx, &RollbackRequest{TransactionId: bResp.TransactionId})
		return apilog.ErrorWithLog(ctx, err, "failed to execute transaction", "db", "mutate")
	} else {

		// commit
		_, err = dbClient.Commit(ctx, &CommitRequest{TransactionId: bResp.TransactionId})
		if err != nil {
			return apilog.ErrorWithLog(ctx, err, "failed to commit db transaction", "db", "commit")
		}
	}
	return nil
}
