all:
	gobuf easyapi_toy/services/service1/service1.go | gobuf-go > easyapi_toy/services/service1/service1.gb.go