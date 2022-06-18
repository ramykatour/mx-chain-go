package transactionAPI

import "errors"

// ErrTransactionNotFound signals that a transaction was not found
var ErrTransactionNotFound = errors.New("transaction not found")

// ErrCannotRetrieveTransaction signals that a transaction cannot be retrieved
var ErrCannotRetrieveTransaction = errors.New("transaction cannot be retrieved")

// ErrNilStatusComputer signals that the status computer is nil
var ErrNilStatusComputer = errors.New("nil transaction status computer")

// ErrNilAPITransactionProcessorArg signals that a nil arguments structure has been provided
var ErrNilAPITransactionProcessorArg = errors.New("nil api transaction processor arg")

// ErrNilFeeComputer signals that the fee computer is nil
var ErrNilFeeComputer = errors.New("nil fee computer")

// ErrNilLogsFacade signals that the logs facade is nil
var ErrNilLogsFacade = errors.New("nil logs facade")

var errCannotLoadReceipts = errors.New("cannot load receipt(s)")
var errCannotContractResults = errors.New("cannot load contract result(s)")
