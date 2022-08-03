package db
import "context"
const createAccount=`-- name: createAccount: one
insert into accounts(
	owner,
	balance,
	currency
	)values(
		$1,$2,$3
	) RETRUNING id,owner,balance,currency,created_at 
	`

	type createAccountParams struct{
		owner string `json:"owner" `
		balance int64 `json:"balance `
		currency string `json:"currency" `
	}


	func (q *Queries) createAccount(ctx context,arg createAccountParams)(Account,error){
		row:=q.db.QueryContext(ctx,createAccount,arg.owner,arg.balance,arg.currency)
		var i Account
		err:=row.Scan(
		&i.ID,
		&i.owner,
		&i.balance,
		&i.currency,
		&i.CreatedAt,)
		return i,err
	}

	const getAccount =`-- name: GetAccount: one 
	select id ,owner,balance,currency,created_at FROM accounts
	WHERE id =$1 limit 1`

	func(q *Queries) GetAccount(ctx context.Context, id int64)(Account,error){
		row:= q.db.QueryRowContext(ctx,getAccount,id)
		var i Account
		err:= row.Scan(
			&i.ID,
			&i.owner,
			&i.balance,
			&i.currency,
			&i.CreatedAt,

		)
		retrun i ,err
	}


	const listAccounts=`--name: ListAccounts : many
	SELECT id,owner,balance,currency,created_at FROM accounts
	ORDER BY id
	LIMIT $1
	OFFSET $2
	`


	type ListAccountsparams struct{
		limit int32 `json:"limit"`
		offset int32 `json:"offset"`
	}


	func (q *Queries) ListAccounts(ctx context.context,arg ListAccountsparams)([]Account,error) {
		rows,err:= q.db.QueryContext(ctx,listlistAccounts,arg.Limit,arg.offset)
		if err!=nil{
			return nil,err}		
		defer rows.Close()
		var items [] Account
		for rows.Next(){
			var i Account
			if err := row.Scan(
				&i.ID,
				&i.Owner,
				&i.Balance,
				&i.Currency,
				&i.CreatedAt,
			);
			err != nil{
				return nil,err
		}
		items= append (items,i)
	}
	if err := rows.Close();
	err!=nil{
		return nil,err
	}
	if err :=rows.Err(); err!=nil{
		return nil ,err
	}
	return items,nil 
}