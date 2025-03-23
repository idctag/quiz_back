package sqlc

import (
	"testing"

	"github.com/idctag/quiz_back/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomQuestion(t *testing.T) {
	phase := CreateRandomPhase(t)
	text := util.RandomString(20)
	img := util.RandomString(20)
	audio := util.RandomString(20)
	arg := CreateQuestionParams{
		PhaseID:          phase.ID,
		Text:             text,
		ImgUrl:           pgtype.Text{String: img},
		AudioUrl:         pgtype.Text{String: audio},
		IsMultipleChoice: true,
	}

	question, err := testQueries.CreateQuestion(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, question)

	require.Equal(t, question.PhaseID, phase.ID)
	require.Equal(t, question.Text, text)
	require.Equal(t, question.ImgUrl, img)
	require.Equal(t, question.AudioUrl, audio)
	require.Equal(t, question.IsMultipleChoice, arg.IsMultipleChoice)
}
