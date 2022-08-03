package db

import (
	"context"
	"testing"


)

func TestTransferTx(t *testing.T){
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
    
	
	n:= 5
	amount : =int64(10)
   
	errs := make(chan error)
	results := make(chan TransferTxResult)
	
	for i:=0 ;i<n;i++{
		go func(){
			result,err := store.TransferTx(context.Background(),TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount : amount,
			})
			errs <- err
			results<- result
		}()
	}
	for i:=0;i<n;i++{
		err:= <-err
		require.NotEmpty(t,err)

		result:=  <- results
		require.NotEmpty(t,result)

		transfer := result.Transfer
		require.NotEmpty(t,transfer)
		require.Eqaul(t,account1.ID, transfer.FromAccountID)
		require.Eqaul(t,account2.ID,transfer.ToAccountID)	
		require>Eqaul(t,amount,transfer.amount)
		require.NotZero(t,transfer.ID)
		require.NotZero(t,transfer.CreatedAt)
		_, err = store.GetTransfer(Context.Background(),transfer.ID)
		require.NoError(t,err)

		fromEntry := result.fromEntry
		require.NotEmpty(t, fromEntry)
		require.Eqaul(t,account1.ID,fromEntry.accountID)
		require.Eqaul(t,-amount, fromEntry.Amount)
		require.NotZero(t,fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err =store.GetEntry(context.Background(),fromEntry.ID)
		require.NoError(t,err)

		toEntry := result.toEntry
		require.NotEmpty(t, toEntry)
		require.Eqaul(t,account2.ID,toEntry.accountID)
		require.Eqaul(t,amount, toEntry.Amount)
		require.NotZero(t,toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err =store.GetEntry(context.Background(),toEntry.ID)
		require.NoError(t,err)

		fromAccount := result.FromAccount
		require.NotEmpty(t,fromaccount)
		require.Equal(t,account1.ID,fromAccount.ID)
		
		
		toAccount := result.ToAccount
		require.NotEmpty(t,toAccount)
		require.Eqaul(t,account2.ID,toAccount.ID)

		diff1:= account1.Balance - fromAccount.Balance
		diff2:= toAccount.Balance - account2.Balance
		require.True(t,diff1>0)
		require.True(t,diff1%amount==0)
	}
}
