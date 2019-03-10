package server

import (
	"fmt"
	"github.com/chvck/meal-planner/proto/service"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"log"
	"net/http"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
)

// Run is the entry point for running the server
func Run(cfg *Info) (*http.Server, error) {
	dataStore, err := NewDataStore(
		cfg.DbServer,
		cfg.DbPort,
		cfg.DbName,
		cfg.DbUsername,
		cfg.DbPassword,
	)
	if err != nil {
		return nil, err
	}

	//cont := NewController(dataStore, cfg.AuthKey)
	//
	//handler := NewHandler(cont)

	//r := routes(handler, cfg.AuthKey)

	address := fmt.Sprintf("%v:%v", cfg.Hostname, cfg.HTTPPort)

	//srv := &http.Server{Addr: address, Handler: r}

	grpcServer := grpc.NewServer()
	service.RegisterMealPlannerServiceServer(grpcServer, &MealPlannerService{
		datastore: dataStore,
		authKey:   cfg.AuthKey,
	})

	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	fmt.Println(grpcweb.ListGRPCResources(grpcServer))

	httpServer := &http.Server{Addr: address, Handler: http.HandlerFunc(handler)}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Printf("Error running server: %s", err)
			return
		}
	}()

	return httpServer, nil
}
