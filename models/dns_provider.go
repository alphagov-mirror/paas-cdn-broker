package models

import (
	"context"
	"crypto/rsa"
	"fmt"
	"gopkg.in/square/go-jose.v1"
	"time"

	"github.com/alphagov/paas-cdn-broker/config"
	"github.com/alphagov/paas-cdn-broker/utils"

	legoacme "github.com/alphagov/paas-cdn-broker/lego/acme"
	goacme "golang.org/x/crypto/acme"
)

func (acp *AcmeClientProvider) GetDNS01Client(user *utils.User, settings config.Settings) (legoacme.ClientInterface, error) {
	logSess := acp.logger.Session("get-dns-01-client")

	logSess.Info("create-goacme-client")
	key := user.GetPrivateKey().(*rsa.PrivateKey)
	client := goacme.Client{Key: key}

	ctx := context.Background()

	logSess.Info("create-goacme-account-struct")
	a := goacme.Account{
		Contact: []string {fmt.Sprintf("mailto:%s", user.Email)},
	}

	logSess.Info("fetch-goacme-account")
	account, err := client.GetReg(ctx, "this argument is ignored because the CA is RFC8555 compliant. The key on the client is used instead.")
	if err == goacme.ErrNoAccount {
		logSess.Info("goacme-account-not-found")
		logSess.Info("register-goacme-account")
		account, err = client.Register(ctx, &a, goacme.AcceptTOS)
		if err != nil {
			logSess.Error("register-goacme-account-error", err)
			return nil, err
		}
	} else if err != nil {
		logSess.Error("fetch-goacme-account", err)
		return nil, err
	}

	if user.GetRegistration() == nil {
		logSess.Info("create-user-registration-resource")
		user.Registration = &legoacme.RegistrationResource{
			Body:        legoacme.Registration{
				Resource:       account.URI,
				ID:             0,
				Key:            jose.JsonWebKey{Key:key},
				Contact:        nil,
				Agreement:      account.AgreedTerms,
				Authorizations: account.Authorizations,
				Certificates:   account.Certificates,
			},
			URI:         account.URI,
			NewAuthzURL: "https://acme-v01.api.letsencrypt.org/acme/new-authz",
			TosURL:      "",
		}
	}
	
	logSess.Info("create-legoacme-client")
	legoClient, err := legoacme.NewClient(settings.AcmeUrl, user, legoacme.RSA2048)
	if err != nil {
		logSess.Error("create-legoacme-client-error", err)
		return nil, err
	}

	logSess.Info("set-challenge-provider")
	err = legoClient.SetChallengeProvider(legoacme.DNS01, &DNSProvider{})
	if err != nil {
		logSess.Error("set-challenge-provider-error", err)
		return nil, err
	}
	legoClient.ExcludeChallenges([]legoacme.Challenge{legoacme.TLSSNI01, legoacme.HTTP01})

	logSess.Info("created")
	return legoClient, nil
}

type DNSProvider struct{}

func (p *DNSProvider) Present(domain, token, keyAuth string) error {
	return nil
}

func (p *DNSProvider) CleanUp(domain, token, keyAuth string) error {
	return nil
}

func (p *DNSProvider) Timeout() (time.Duration, time.Duration) {
	return 10 * time.Second, 2 * time.Second
}
