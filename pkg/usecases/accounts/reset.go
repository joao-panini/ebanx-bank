package accounts

func (accountUseCase *accountUseCase) ResetAccountStates() error {
	err := accountUseCase.accountStore.ResetAccountStates()
	if err != nil {
		return err
	}
	return nil
}
