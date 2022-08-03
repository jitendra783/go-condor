package db

import (
	"context"
	"database/sql"
	"errors"
	"testing/quick"
)

type Store interface{
	Querier 
	TransferTx(ctx context.Context,arg TransferTxParams)(TransferTxResult,err)

}
type SQLStore struct {
	*Queries
	db *sql.DB
}


func NewStore(db *sqlDB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
func (store *SQLStore) execTx(ctx context.Context,fn func (*Queries)error)error{
	tx,err != store.db.BeginTx(ctx,nil)
	if err != nil{
		return err
	}
	q := New(tx)
	err= fn(q)
	if err != nil{
		if rbErr:=tx.Rollback();rbErr!= nil{
			return fmt.Errorf("tx err: %v ,rb err : %v",err,rbErr)

		}
		return err

	}
	return tx.commit()
}

type TransferTxParams struct{
	FromAccountID int64 `json :"from_account_id"`
	ToAccountID int64 	` json: "to_account_id"`
	Amount int64 `josn: "amount"`
}
type  TransferTxResult struct{
 Transfer Transfer `josn:`
}
func (store * SQLStore) transferTx(ctx context.Context ,arg TransferTxParams)(TransferTxResults,error){
	var result TransferTxResult

	err :=store.db.execTx(ctx,func(q *Queries) error{
		var err error

		result.Transfer,err =q.CreateTransfer(ctx CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:		   arg.Amount,  
		})
		if err != nil{
			return err
		}
		return nil
	})
	return result,err
}
