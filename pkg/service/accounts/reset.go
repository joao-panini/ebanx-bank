package accounts

func (s *accountService) ResetAccountStates() error {
	err := s.accStore.ResetAccountStates()
	if err != nil {
		return err
	}
	return nil
}
