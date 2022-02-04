'use strict';

var answer_word = "doggo";

var e = React.createElement;

class Hangman extends React.Component{

    constructor(props){
        super(props)
        this.state={
            answer_length: api_lengthcheck(),
            answer_letters: {},
            guessed_letters: [],
            complete: false
        }
    }

    check_letter = (letter) => {
        this.setState(prevState => ({guessed_letters: [...prevState.guessed_letters, letter]}))
        var result = api_lettercheck(letter)

        // Create an object of the new answer letters revealed
        var answer_letters_toadd = {}
        for(const [key, value] of Object.entries(result)) {
            answer_letters_toadd = {...answer_letters_toadd, [key]: value}
        }

        // Combine that object with the existing state object (note: can't change state in-place)
        const answer_letters_new = {...this.state.answer_letters, ...answer_letters_toadd}
        this.setState({answer_letters: answer_letters_new})

        if(Object.keys(this.state.answer_letters).length == this.state.answer_length){
            this.setState({complete: true})
            console.log("Complete - well done!")
        }
    }

    render(){

        return e(
            "div",
            {className: "hangman"},
            [
                e(HgmnWord, {answer_length: this.state.answer_length, answer_letters: this.state.answer_letters }),
                e(KeyBoard, {guessed_letters: this.state.guessed_letters, onHandleClick: this.check_letter }) // Sending function down to keyboard
            ]
        )
    }
}



const hangman = document.querySelector('#hangman');
ReactDOM.render(e(Hangman, {answer_word: answer_word}), hangman);



function api_lengthcheck(){
    return answer_word.length
}
function api_lettercheck(letter){

    var answer_word_upper = answer_word.toUpperCase();

    var answer_list = {}

    for(let i=0; i<answer_word_upper.length; i++){
        if(answer_word_upper[i] == letter){
            answer_list[i] = letter
        }
    }

    return answer_list

}