'use strict';

var e = React.createElement;

class HgmnLetter extends React.Component {

    constructor(props){
        super(props)
    }

    render(){
        
        console.log(this.props.answered)
        let letter_class = this.props.answered ? "letter answered" : "letter unanswered";

        return e(
            "div",
            {className: letter_class},
            this.props.value
        )
    }
}

class HgmnWord extends React.Component {

    constructor(props){
        super(props)
    }

    render(){

        var allletters = []

        for(let i = 0; i<this.props.answer_length; i++){
            if(i in this.props.answer_letters){
                var letter_answered = true
                var letter_value = this.props.answer_letters[i]
            } else {
                var letter_answered = false
                var letter_value = null
            }
            allletters.push(e(HgmnLetter, {answered: letter_answered, index: 0, value: letter_value}))
        }

        return e(
            "div",
            {className: "word"},
            allletters
        )
    }
}



// const hgmnword = document.querySelector('#hgmnword');
// ReactDOM.render(e(HgmnWord), hgmnword);


