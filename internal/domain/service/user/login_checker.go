package user

type LoginChecker interface {
	LoginCheck() (bool, error)
}
