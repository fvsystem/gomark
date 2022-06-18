package adapter

type ReportMaker interface {
	MakeReport(interface{}) error
}
