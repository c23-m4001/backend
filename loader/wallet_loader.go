package loader

import (
	"capstone/constant"
	"capstone/model"
	"capstone/repository"
	"capstone/util"
	"context"

	"github.com/graph-gophers/dataloader"
)

type WalletLoader struct {
	loaderById               dataloader.Loader
	loaderByUserId           dataloader.Loader
	loaderHaveWalletByUserId dataloader.Loader
}

func (l *WalletLoader) loadById(id string) (*model.Wallet, error) {
	thunk := l.loaderById.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Wallet), nil
}

func (l *WalletLoader) loadByUserId(userId string) ([]model.Wallet, error) {
	thunk := l.loaderByUserId.Load(context.TODO(), dataloader.StringKey(userId))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.Wallet), nil
}

func (l *WalletLoader) loadHaveWalletByUserId(userId string) (*bool, error) {
	thunk := l.loaderHaveWalletByUserId.Load(context.TODO(), dataloader.StringKey(userId))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*bool), nil
}

func (l *WalletLoader) UserFn(user *model.User) func() error {
	return func() error {
		wallets, err := l.loadByUserId(user.Id)
		if err != nil {
			return err
		}

		user.Wallets = wallets

		return nil
	}
}

func (l *WalletLoader) TransactionFn(transaction *model.Transaction) func() error {
	return func() error {
		wallet, err := l.loadById(transaction.WalletId)
		if err != nil {
			return err
		}

		transaction.Wallet = wallet

		return nil
	}
}

func NewWalletloader(
	walletRepository repository.WalletRepository,
) *WalletLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		wallets, err := walletRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		walletsById := map[string]*model.Wallet{}
		for _, wallet := range wallets {
			walletsById[wallet.Id] = &wallet
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var wallet *model.Wallet
			if v, ok := walletsById[k.String()]; ok {
				wallet = v
			}

			result := &dataloader.Result{Data: wallet, Error: nil}

			if wallet == nil {
				result.Error = constant.ErrNoData
			}

			results[idx] = result
		}
		return results
	}

	batchUserIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		wallets, err := walletRepository.FetchByUserIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		walletsByUserId := map[string][]model.Wallet{}
		for _, wallet := range wallets {
			walletsByUserId[wallet.UserId] = append(walletsByUserId[wallet.UserId], wallet)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var wallets []model.Wallet
			if v, ok := walletsByUserId[k.String()]; ok {
				wallets = v
			}

			result := &dataloader.Result{Data: wallets, Error: nil}

			results[idx] = result
		}
		return results
	}

	haveWalletBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		haveWalletByUserId, err := walletRepository.IsExistByUserIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var haveWallet *bool
			if v, ok := haveWalletByUserId[k.String()]; ok {
				haveWallet = &v
			} else {
				haveWallet = util.BoolP(false)
			}

			result := &dataloader.Result{Data: haveWallet, Error: nil}

			if haveWallet == nil {
				result.Error = constant.ErrNoData
			}

			results[idx] = result
		}
		return results
	}

	return &WalletLoader{
		loaderById:               newDataloader(batchFn),
		loaderByUserId:           newDataloader(batchUserIdFn),
		loaderHaveWalletByUserId: newDataloader(haveWalletBatchFn),
	}
}
