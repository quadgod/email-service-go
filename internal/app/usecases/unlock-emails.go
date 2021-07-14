package usecases

import (
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/repos"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type IUnlockEmailsUseCase interface {
	StartUnlocking()
}

type UnlockEmailsUseCase struct {
	emailRepository *repos.IEmailRepository
	config          *config.IConfig
}

func unlockEmails(allDone *sync.WaitGroup, emailRepository repos.IEmailRepository) {
	defer allDone.Done()
	unlockedCount, err := emailRepository.UnlockEmails()
	if err != nil {
		log.Errorf("Unlock emails error: %s", err.Error())
	} else {
		log.Infof("Unlocked emails %d", unlockedCount)
	}
}

func (instance *UnlockEmailsUseCase) StartUnlocking() {
	var allDone sync.WaitGroup
	for {
		allDone.Add(1)
		go unlockEmails(&allDone, *instance.emailRepository)
		allDone.Wait()
		unlockAfterSeconds := (*instance.config).GetUnlockEmailsAfterSec()
		time.Sleep(time.Duration(unlockAfterSeconds) * time.Second)
	}
}
