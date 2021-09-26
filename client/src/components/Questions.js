import { React, useState } from 'react'
import { gql, useQuery } from "urql";
import Question from './Question'


export function Questions() {
  const [{ data, fetching, error }] = useQuery({ query });
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0)
  const [answers, setAnswers] = useState({})
  // storing anwsers to display the previous value when using next and previous buttons

  if (fetching) return "Loading...";
  if (error) return `Error: ${error}`;

  data.questions.sort(sortByWeight)
  data.questions.forEach(q => {
    if (q.options !== undefined) q.options.sort(sortByWeight)
  })
  if (data.questions.length === 0) return <>No questions</>

  const showNextQuestion = () => {
    if (currentQuestionIndex < data.questions.length - 1) {
      setCurrentQuestionIndex(i => i + 1)
    }
  }

  const showPreviousQuestion = () => {
    if (currentQuestionIndex > 0) {
      setCurrentQuestionIndex(i => i - 1)
    }
  }

  // I prefer to use a closure to ensure the questionID is not changed downstream
  const handleAnswerUpdate = (questionID) => {
    const storeAnswer = (answer) => {
      setAnswers(answers => {
        answers[questionID] = answer
        return answers
      })
    }
    return storeAnswer
  }

  return (
    <>
      {
        <Question peviousAnswer={answers[data.questions[currentQuestionIndex].id]} 
          answerListener={handleAnswerUpdate(data.questions[currentQuestionIndex].id)} 
          question={data.questions[currentQuestionIndex]} />
      }
      {
        currentQuestionIndex > 0 &&
        <button type="button" onClick={showPreviousQuestion}>Previous</button>
      }
      {
        currentQuestionIndex < data.questions.length - 1 && 
          <button type="button" onClick={showNextQuestion}>Next</button>
      }
    </>
  );
}

const sortByWeight = (a, b) => {
  return a.weight - b.weight
}

const query = gql`
  query {
    questions {
      __typename
      ... on ChoiceQuestion {
        id
        body
        weight
        options {
          id
          body
          weight
        }
      }
      ... on TextQuestion {
        id
        body
        weight
      }
    }
  }
`;
