package workerutil

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/andrewpillar/djinn/config"
	"github.com/andrewpillar/djinn/errors"
	"github.com/andrewpillar/djinn/worker"
)

func ParseFlags(args []string) (string, string, bool) {
	var (
		config  string
		driver  string
		version bool
	)

	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	fs.StringVar(&config, "config", "djinn-worker.toml", "the config file to use")
	fs.StringVar(&driver, "driver", "djinn-driver.toml", "the driver config to use")
	fs.BoolVar(&version, "version", false, "show the version and exit")
	fs.Parse(args[1:])

	return config, driver, version
}

func Init(workerPath, driverPath string) (*worker.Worker, config.Worker, func(), error) {
	var cfg config.Worker

	f1, err := os.Open(workerPath)

	if err != nil {
		return nil, cfg, nil, errors.Err(err)
	}

	defer f1.Close()

	cfg, err = config.DecodeWorker(f1)

	if err != nil {
		return nil, cfg, nil, errors.Err(err)
	}

	pidfile := cfg.Pidfile()

	db := cfg.DB()
	redis := cfg.Redis()
	smtp, postmaster := cfg.SMTP()

	log := cfg.Log()

	close_ := func() {
		db.Close()
		redis.Close()
		smtp.Close()
		log.Close()

		if pidfile != nil {
			if err := os.RemoveAll(pidfile.Name()); err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
				os.Exit(1)
			}
		}
	}

	f2, err := os.Open(driverPath)

	if err != nil {
		return nil, cfg, nil, errors.Err(err)
	}

	defer f2.Close()

	drivers, driverCfg, err := config.DecodeDriver(f2)

	if err != nil {
		return nil, cfg, nil, errors.Err(err)
	}

	return &worker.Worker{
		DB:          db,
		Redis:       redis,
		SMTP:        smtp,
		Admin:       postmaster,
		Block:       cfg.BlockCipher(),
		Log:         log,
		Queue:       cfg.Queue(),
		Parallelism: cfg.Parallelism(),
		Timeout:     cfg.Timeout(),
		Config:      driverCfg,
		Drivers:     drivers,
		Providers:   cfg.Providers(),
		Placer:      cfg.Objects(),
		Collector:   cfg.Artifacts(),
	}, cfg, close_, nil
}

func Start(ctx context.Context, w *worker.Worker) {
	go func() {
		if err := w.Run(ctx); err != nil {
			w.Log.Error.Println(errors.Cause(err))
		}
	}()
}
