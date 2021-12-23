package skill

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"gitlab.com/hieuxeko19991/job4e_be/services/autocomplete"

	"go.uber.org/zap"
)

type ISkillService interface {
	CreateSkill(ctx context.Context, skill *models.RequestCreateSkill) error
	UpdateSkill(ctx context.Context, skill *models.RequestCreateSkill, skill_id int64) error
	GetListSkill(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSkill, error)
	GetAll(ctx context.Context, name string) ([]models.Skill, error)
}

type ISkillDatabase interface {
	Create(ctx context.Context, skill *models.Skill) error
	Update(ctx context.Context, skill *models.RequestUpdateSkill, skillID int64) error
	Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSkill, error)
	GetAll(ctx context.Context, name string) ([]models.Skill, error)
}

type Skill struct {
	SkillGorm     ISkillDatabase
	CandidateTrie *autocomplete.Trie
	RecruiterTrie *autocomplete.Trie
	JobTrie       *autocomplete.Trie
	Logger        *zap.Logger
}

func NewSkill(skillGorm *SkillGorm, canTrie *autocomplete.Trie, recTrie *autocomplete.Trie, jobTrie *autocomplete.Trie, logger *zap.Logger) *Skill {
	return &Skill{
		SkillGorm:     skillGorm,
		CandidateTrie: canTrie,
		RecruiterTrie: recTrie,
		JobTrie:       jobTrie,
		Logger:        logger,
	}
}

/*Create Skill*/
func (s *Skill) CreateSkill(ctx context.Context, skill *models.RequestCreateSkill) error {
	skillModels := &models.Skill{
		Name:        skill.Name,
		Description: skill.Description,
		Questions:   skill.Questions,
		Icon:        skill.Icon,
		Status:      skill.Status}
	err := s.SkillGorm.Create(ctx, skillModels)
	if err != nil {
		return err
	}
	//s.CandidateTrie.Insert(skillModels.Name)
	//s.RecruiterTrie.Insert(skillModels.Name)
	//s.JobTrie.Insert(skillModels.Name)
	return nil
}

/*Update Skill*/
func (s *Skill) UpdateSkill(ctx context.Context, skill *models.RequestCreateSkill, skill_id int64) error {
	skillModels := &models.RequestUpdateSkill{
		Name:        skill.Name,
		Description: skill.Description,
		Questions:   skill.Questions,
		Icon:        skill.Icon,
		Status:      skill.Status}
	err := s.SkillGorm.Update(ctx, skillModels, skill_id)
	if err != nil {
		return err
	}
	//s.CandidateTrie.Insert(skillModels.Name)
	//s.RecruiterTrie.Insert(skillModels.Name)
	//s.JobTrie.Insert(skillModels.Name)
	return nil
}

func (s *Skill) GetListSkill(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSkill, error) {
	acc, err := s.SkillGorm.Get(ctx, name, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *Skill) GetAll(ctx context.Context, name string) ([]models.Skill, error) {
	acc, err := s.SkillGorm.GetAll(ctx, name)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
