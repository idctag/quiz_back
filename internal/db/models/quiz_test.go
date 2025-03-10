package db

import (
	"context"
	"testing"
	"time"

	"github.com/idctag/quiz_back/util"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/require"
)

func createRandomQuiz(t *testing.T) Quiz {
	name := util.RandomName()

	quiz, err := testQueries.CreateQuiz(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, quiz)

	require.Equal(t, name, quiz.Name)

	require.NotZero(t, quiz.ID)
	require.NotZero(t, quiz.CreatedAt)

	return quiz
}

func TestCreateQuiz(t *testing.T) {
	createRandomQuiz(t)
}

func TestGetQuiz(t *testing.T) {
	quiz1 := createRandomQuiz(t)
	quiz2, err := testQueries.GetQuiz(context.Background(), quiz1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, quiz2)

	require.Equal(t, quiz1.ID, quiz2.ID)
	require.Equal(t, quiz1.Name, quiz2.Name)
	require.WithinDuration(t, quiz1.CreatedAt.Time, quiz2.CreatedAt.Time, time.Second)
}

func TestUpdateQuiz(t *testing.T) {
	quiz1 := createRandomQuiz(t)

	arg := UpdateQuizParams{
		ID:   quiz1.ID,
		Name: util.RandomName(),
	}

	quiz2, err := testQueries.UpdateQuiz(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, quiz2)

	require.Equal(t, quiz2.ID, arg.ID)
	require.Equal(t, quiz2.Name, arg.Name)
	require.WithinDuration(t, quiz1.CreatedAt.Time, quiz2.CreatedAt.Time, time.Second)
}

func TestDeleteQuiz(t *testing.T) {
	quiz1 := createRandomQuiz(t)
	err := testQueries.DeleteQuiz(context.Background(), quiz1.ID)
	require.NoError(t, err)

	quiz2, err := testQueries.GetQuiz(context.Background(), quiz1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, quiz2)
}

func TestListQuic(t *testing.T) {
	for range 10 {
		createRandomQuiz(t)
	}

	arg := ListQuizzesParams{
		Limit:  5,
		Offset: 5,
	}

	quizzes, err := testQueries.ListQuizzes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, quizzes, 5)

	for _, quiz := range quizzes {
		require.NotEmpty(t, quiz)
	}
}
