'use string';

var e = React.createElement;

class Addwords_container extends React.Component {
    constructor(props){
        super(props)
        this.state = {
            wordlist: {
                0: {"word": "brooklyn", "message": "Cooking raw with the brooklyn boy"},
                1: {"word": "desk", "message": "Dele Eriksen Son Kane"}
            }
        }
    }

    render(){

        var words_cards = []

        for(const [key, value] of Object.entries(this.state.wordlist)){
            var card_num = key
            var card_word = value["word"]
            var card_message = value["message"]
            
            words_cards.push(e(Addwords, {num: card_num, word: card_word, message: card_message}))
        }

        return e(
            "div",
            {className: "addwords_container"},
            [
                words_cards,
                e(Addanother, {})
            ]
        )
    }
}

class Addwords extends React.Component{
    
    constructor(props){
        super(props)
        this.state = {}

    }


    render () {
        
        return e(
            "div",
            {className: "addwords card"},
            [
                e(Addword, {word: this.props.word}),
                e(Addmessage, {message: this.props.message})
            ]
        )
    }
}



class Addword extends React.Component{

    constructor(props){
        super(props)
        this.state = {
            value: this.props.word
        }
    }

    handleKeyPress = (e) => {

        const re = /^[A-Za-z]+$/;
        if (e.target.value === "" || re.test(e.target.value)){
          this.setState({ value: e.target.value });
        }
    }

    render() {
        return e(
            "div",
            {className: "addword_box large"},
            e(
                "input",
                {type: "text", className: "addword_cursor large", placeholder: "Enter word here", value: this.state.value, onChange: this.handleKeyPress}
            )
        )
    }
}

class Addmessage extends React.Component{

    constructor(props){
        super(props)
        this.state = {
            value: this.props.message
        }
    }

    handleKeyPress = (e) => {
        this.setState({ value: e.target.value });
    }

    render() {
        return e(
            "div",
            {className: "addword_box small"},
            e(
                "input",
                {type: "text", className: "addword_cursor small", placeholder: "Enter message to display on completion here", value: this.state.value, onChange: this.handleKeyPress}
            )
        )
    }
}

class Addanother extends React.Component {

    constructor(props){
        super(props)
        this.state = {}
    }

    render() {
        return e(
            "div",
            {className: "addanother card"},
            " + Click to add another"
        )
    }
}

const addwords_container = document.querySelector('#addwords_container');
ReactDOM.render(e(Addwords_container, {}), addwords_container);