package usecase

import (
	"context"
	"github.com/quadgod/email-service-go/internal/app/config"
	"github.com/quadgod/email-service-go/internal/app/db/repository"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type IUnlockEmailsUseCase interface {
	Start()
}

type UnlockEmailsUseCase struct {
	emailRepository *repository.EmailRepository
	config          *config.Config
}

func NewUnlockEmailsUseCase(
	emailRepository repository.EmailRepository,
	config config.Config,
) IUnlockEmailsUseCase {
	return &UnlockEmailsUseCase{
		emailRepository: &emailRepository,
		config:          &config,
	}
}

func unlockEmails(allDone *sync.WaitGroup, emailRepository repository.EmailRepository) {
	defer allDone.Done()
	unlockedCount, err := emailRepository.Unlock(context.TODO())
	if err != nil {
		log.Errorf("Unlock emails error: %s", err.Error())
	} else {
		log.Infof("Unlocked emails %d", unlockedCount)
	}
}

func (instance *UnlockEmailsUseCase) Start() {
	var allDone sync.WaitGroup
	for {
		allDone.Add(1)
		go unlockEmails(&allDone, *instance.emailRepository)
		allDone.Wait()
		unlockAfterSeconds := (*instance.config).GetUnlockEmailsAfterSec()
		time.Sleep(time.Duration(unlockAfterSeconds) * time.Second)
	}
}
