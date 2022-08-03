package api
import (

)

func TestAccountAPI(t *testing.T){
	account : = randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store:= mockdb.NewMockStore(ctrl)
	store.EXPECT().
		GetAccount(gomock.Any(),gomock.Eq(account.ID)).\
		times(1).
		return(account,nil)
	
	server:= NewServer(store)
	recorder:= httptest.NewRecorder()







	defer ctrl .mu.Unlock()
	
	if ctrl.finished{
		ctrl.T.Fetalf("Controller.Finish was called more than onece, iIt gas to be called exactly once")
	}
	ctrl.finished = true

	if err := recover (); err != nil{
		panic(err)

	}

	failures := ctrl.expectedCalls.Failures()
	for _,call := range failures{
		ctrl.T.Errorf("missing call(s) to %v ",call)

	}
	if len(failures) !=0 {
		ctrl.T.Fetalf(aborting test due to missing call(s)"")

	}
} 

func callerInfo(skip int )string {
	if_,file, line,ok := runtime.Caller(skip + 1); ok{
		return fmt.Sprintf("%s :%d" , file ,line)

	}
	return "unkown file"
}

func randomAccount() db.Account{
	return db.Account{
		ID: util.RandomInt(1,1000),
		Owner: util.RandomOwner(),
		Balance: util.RanodmMoney(),
		Currency: util.Currency(),
	}
}

func requireBodyMatchAccount()