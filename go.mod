module ntsc.ac.cn/st-pcie-sync

go 1.18

replace ntsc.ac.cn/ta-registry v0.0.0 => ../ta-registry

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	ntsc.ac.cn/tas/tas-commons v0.0.0
)

require (
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
)

replace ntsc.ac.cn/tas/tas-commons v0.0.0 => ../tas-commons
