package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDb_Insert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/file.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("path/to/file.mp4", "Pending", video)

	require.Nil(t, err)
	require.NotEmpty(t, job.ID)
	require.Equal(t, job.Video.ID, video.ID)
}

func TestJobRepositoryDb_Update(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/file.mp4"
	video.CreatedAt = time.Now()

	videoRepo := repositories.VideoRepositoryDb{Db: db}
	videoRepo.Insert(video)

	newJob, err := domain.NewJob("path/to/file.mp4", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(newJob)

	newJob.Status = "Completed"

	repoJob.Update(newJob)

	job, err := repoJob.Find(newJob.ID)
	require.NotEmpty(t, job.ID)
	require.Nil(t, err)
	require.Equal(t, job.Status, newJob.Status)
}
