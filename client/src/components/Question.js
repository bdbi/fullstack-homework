import { React, useState } from 'react'

const Question = (props) => {
    const { question, answerListener, peviousAnswer } = props
    let q
    switch (question.__typename) {
        case "ChoiceQuestion":
            q = <ChoiceQuestion
                previousAnswer={!peviousAnswer ? null : peviousAnswer.selectedOption}
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
    const [answerText, setAnswerText] = useState(previousAnswer)
    const handleAnswerText = (e) => {
        setAnswerText(e.target.value)
        answerListener({ questionID: questionID, text: e.target.value })
    }
    return (<div>
        <p>{text}</p>
        <div>
            <textarea onChange={handleAnswerText} value={answerText} />
        </div>
    </div>)
}

const ChoiceQuestion = (props) => {
    const { text, options, questionID, answerListener, previousAnswer } = props
    const [selectedOption, setSelectedOption] = useState(previousAnswer)
    const handleOptionSelection = (e) => {
        setSelectedOption(e.target.value)
        answerListener({ questionID: questionID, selectedOption: e.target.value })
    }
    return (<div>
        <p>{text}</p>
        <div>
            {options.map((o, i) => {
                return <div key={i}>
                    <input
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
        </div>
    </div>)
}

export default Question