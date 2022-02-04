'use strict';

var answer_word = "doggo";

var e = React.createElement;

class Hangman extends React.Component{

    constructor(props){
        super(props)
        this.state={
            // answer_length: api_lengthcheck().then(response => {return response}),
            answer_length: null,
            answer_letters: {},
            guessed_letters: [],
            complete: false
        }
    }

    componentDidMount(){
        this.api_lengthcheck();
    }

    api_lengthcheck = () => {
        return fetch("/hangman/api/wrdlen/abc123")
            .then(response => response.json())
            .then(data => this.setState({answer_length: data}))
            .catch(error => console.log(error));
    }
    
    api_lettercheck = (letter) => {  // Anonymous function so we can call this.State
        return fetch("/hangman/api/ltrchk/abc123?letter=" + letter)
            .then(response => response.json())
            .then(data => {
                console.log(data)

                // Adding letter to list of guessed letters
                this.setState(prevState => ({guessed_letters: [...prevState.guessed_letters, letter]}))

                // Create an object of the new answer letters revealed
                var answer_letters_toadd = {}
                for(const [key, value] of Object.entries(data)) {
                    answer_letters_toadd = {...answer_letters_toadd, [key]: value}
                }

                // Combine that object with the existing state object (note: can't change state in-place)
                const answer_letters_new = {...this.state.answer_letters, ...answer_letters_toadd}
                this.setState(
                    {answer_letters: answer_letters_new},
                    function() { // Callback function so it ensures state is updated before running
                        if(Object.keys(this.state.answer_letters).length == this.state.answer_length){
                            this.setState({complete: true})
                        }
                    })

            })
    }

    reset_hangman = () => {
        this.api_lengthcheck()
        this.setState({answer_letters: {}, guessed_letters: [], complete: false})
    }

    render(){

        return e(
            "div",
            {className: "hangman"},
            [
                e(HgmnWord, {answer_length: this.state.answer_length, answer_letters: this.state.answer_letters }),
                e(KeyBoard, {guessed_letters: this.state.guessed_letters, onHandleClick: this.api_lettercheck }), // Sending function down to keyboard
                e(Finish, {complete: this.state.complete, onHandleClick: this.reset_hangman})
            ]
        )
    }
}



const hangman = document.querySelector('#hangman');
ReactDOM.render(e(Hangman, {answer_word: answer_word}), hangman);



