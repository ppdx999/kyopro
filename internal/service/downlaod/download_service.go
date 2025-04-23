package downlaod

/*
DownloadServiceは問題のテストケースをダウンロードします。
*/
type DownloadService interface {
	Download() error
}

type DownloadServiceImpl struct {
	loader CurrentProblemLoader
	getter TestCasesGetter
	saver  TestCaseSaver
}

func NewDownloadServiceImpl(
	loader CurrentProblemLoader,
	getter TestCasesGetter,
	saver TestCaseSaver,
) *DownloadServiceImpl {
	return &DownloadServiceImpl{
		loader: loader,
		getter: getter,
		saver:  saver,
	}
}

func (d *DownloadServiceImpl) Download() error {
	p, err := d.loader.LoadCurrentProblem()
	if err != nil {
		return err
	}

	testCases, err := d.getter.GetTestCases(
		p.Contest.ID,
		p.ID,
	)
	if err != nil {
		return err
	}

	for _, t := range testCases {
		if err := d.saver.SaveTestCase(t); err != nil {
			return err
		}
	}

	return nil
}
