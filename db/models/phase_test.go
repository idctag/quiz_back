package sqlc

import (
	"context"
	"testing"

	"github.com/idctag/quiz_back/util"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/require"
)

func CreateRandomPhase(t *testing.T) Phase {
	quiz := CreateRandomQuiz(t)
	name := util.RandomString(10)
	arg := CreatePhaseParams{
		Name:   name,
		QuizID: quiz.ID,
	}
	phase, err := testQueries.CreatePhase(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, phase)

	require.Equal(t, arg.Name, phase.Name)
	require.Equal(t, arg.QuizID, phase.QuizID)

	return phase
}

func TestCreatePhase(t *testing.T) {
	CreateRandomPhase(t)
}

func TestGetPhase(t *testing.T) {
	phase1 := CreateRandomPhase(t)
	phase2, err := testQueries.GetPhase(context.Background(), phase1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, phase2)

	require.Equal(t, phase1.ID, phase2.ID)
	require.Equal(t, phase1.Name, phase2.Name)
	require.Equal(t, phase1.QuizID, phase2.QuizID)
}

func TestUpdatePhase(t *testing.T) {
	quiz1 := CreateRandomPhase(t)
	name := util.RandomName()
	arg := UpdatePhaseParams{
		ID:   quiz1.ID,
		Name: name,
	}

	quiz2, err := testQueries.UpdatePhase(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, quiz2)

	require.Equal(t, quiz2.ID, quiz1.ID)
	require.Equal(t, quiz2.Name, arg.Name)
}

func TestDeletePhase(t *testing.T) {
	phase1 := CreateRandomPhase(t)
	err := testQueries.DeletePhase(context.Background(), phase1.ID)
	require.NoError(t, err)

	phase2, err := testQueries.GetPhase(ctx, phase1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, phase2)
}

func TestListPhase(t *testing.T) {
	for range 10 {
		CreateRandomPhase(t)
	}

	arg := ListPhasesParams{
		Limit:  5,
		Offset: 5,
	}
	phases, err := testQueries.ListPhases(ctx, arg)
	require.NoError(t, err)
	require.Len(t, phases, 5)

	for _, phase := range phases {
		require.NotEmpty(t, phase)
	}
}
