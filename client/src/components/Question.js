import { React, useState } from 'react'
import { gql, useMutation } from "urql";

const successStyle = { backgroundColor: 'green', color: 'white' }

const Question = (props) => {
    const { question, answerListener, peviousAnswer } = props
    let q
    switch (question.__typename) {
        case "ChoiceQuestion":
            q = <ChoiceQuestion
                previousAnswer={!peviousAnswer ? null : peviousAnswer.optionID}
                answerListener={answerListener}
                questionID={question.id}
                text={question.body}
                options={question.options} />
            break
        case "TextQuestion":
            q = <TextQuestion
                answerListener={answerListener}
                questionID={question.id}
                text={question.body}
                previousAnswer={!peviousAnswer ? "" : peviousAnswer.text}
            />
            break
        default:
            q = <p>unsupported question</p>
    }
    return q
}

const TextQuestion = (props) => {
    const { text, questionID, answerListener, previousAnswer } = props
    const [state, executeMutation] = useMutation(SUBMIT_ANSWER)
    const [answerText, setAnswerText] = useState(previousAnswer)
    const [dirty, setDirty] = useState(!!previousAnswer)
    const [submittedOK, setSubmittedOK] = useState(null)

    const handleAnswerText = (e) => {
        answerListener({ questionID: questionID, text: e.target.value })
        setAnswerText(e.target.value)
        setDirty(true)
    }
    const handleSubmitAnswer = async () => {
        try {
            const answer = { questionID: questionID, text: answerText }
            let res = await executeMutation({ answer })
            if (!res.error) {
                console.log("submitted answer:", res.data.submitAnswer)
                setSubmittedOK(true)
            }
            else console.log(res.error)
        } catch (e) {
            console.error(e)
        }
    }

    return (<div>
        <p style={successStyle}>{submittedOK && "answer correctly sent"}</p>
        <p>{text}</p>
        <div>
            <textarea
                disabled={submittedOK}
                onChange={handleAnswerText}
                value={answerText} />
            <br />
            <button type="button" disabled={!dirty || state.fetching || submittedOK}
                onClick={handleSubmitAnswer}>Submit answer</button>
        </div>
    </div>)
}

const ChoiceQuestion = (props) => {
    const { text, options, questionID, answerListener, previousAnswer } = props
    const [state, executeMutation] = useMutation(SUBMIT_ANSWER)
    const [selectedOption, setSelectedOption] = useState(previousAnswer)
    const [dirty, setDirty] = useState(!!previousAnswer)
    const [submittedOK, setSubmittedOK] = useState(null)


    const handleOptionSelection = (e) => {
        answerListener({ questionID: questionID, optionID: e.target.value })
        setSelectedOption(e.target.value)
        setDirty(true)
    }
    const handleSubmitAnswer = async () => {
        try {
            const answer = { questionID: questionID, optionID: selectedOption }
            let res = await executeMutation({ answer })
            if (!res.error) {
                console.log("submitted answer:", res.data.submitAnswer)
                setSubmittedOK(true)
            }
            else console.log(res.error)
        } catch (e) {
            console.error(e)
        }
    }
    return (<div>
        <p style={successStyle}>{submittedOK && "answer correctly sent"}</p>
        <p>{text}</p>
        <div>
            {options.map((o, i) => {
                return <div key={i}>
                    <input
                        disabled={submittedOK}
                        checked={o.id === selectedOption}
                        name={questionID}
                        type="radio"
                        id={i}
                        value={o.id}
                        onChange={handleOptionSelection}
                    />
                    <label htmlFor={i}>{o.body}</label><br />
                </div>
            })}
            <button type="button" disabled={!dirty || state.fetching || submittedOK} onClick={handleSubmitAnswer}>Submit answer</button>
        </div>
    </div>)
}

export default Question

const SUBMIT_ANSWER = gql`
mutation SubmitAnswer($answer: AnswerInput!){
  submitAnswer(answer: $answer){
    id
  }
}
`