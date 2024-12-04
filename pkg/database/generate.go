package database

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i TxManeger -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Handler -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Transactor -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Pinger -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i SqlExecer -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Client -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i NamedExecer -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i QueryExecer -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i DB -o ./mocks/ -s "_minimock.go"
