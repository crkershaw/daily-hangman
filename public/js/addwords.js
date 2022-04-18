'use string';

var e = React.createElement;

class Addwords_container extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            wordlist: {
                0: {"word": "brooklyn", "message": "Cooking raw with the brooklyn boy"},
                1: {"word": "desk", "message": "Dele Eriksen Son Kane"}
            }
        };
    }

    addToWordlist = (id, word, message) => {

        // Create an object of the new wordlist
        var wordlist_toadd = {}

        wordlist_toadd = {...wordlist_toadd , [id]: {"word": word, "message": message}}
        
        // Combine that object with the existing state object (note: can't change state in-place)
        const wordlist_new = {...this.state.wordlist, ...wordlist_toadd}

        this.setState({wordlist: wordlist_new})


    }

    submitWordlist = () => {

        var new_words = {}
        
        var num = 0

        for(let i in this.state.wordlist){
            if(this.state.wordlist[i]["word"] != ""){
                new_words[num] = {"word": this.state.wordlist[i]["word"], "message": this.state.wordlist[i]["message"]}    
                num += 1
            }
        }
    }


    render(){

        var words_cards = []

        for(const [key, value] of Object.entries(this.state.wordlist)){
            var card_num = key
            var card_word = value["word"]
            var card_message = value["message"]
            
            words_cards.push(e(Addwords, 
                {
                    num: card_num, 
                    word: card_word, 
                    message: card_message,
                    onHandleInput: this.addToWordlist
                }))
        }

        return e(
            "div",
            {className: "addwords_container"},
            [
                words_cards,
                e(Addanother, 
                    {onHandleClick: this.addToWordlist,
                    wordlist: this.state.wordlist}
                ),
                e(Submit, 
                    {onHandleClick: this.submitWordlist}
                )
            ]
        )
    }
}

class Addwords extends React.Component{
    
    constructor(props){
        super(props)
        this.state = {
            word: this.props.word,
            message: this.props.message     
        }
    }

    handleInput = (type, text) => {
        if(type == "word"){
            this.setState({word: text}, () => {this.props.onHandleInput(this.props.num, this.state.word, this.state.message)})

        } else {
            this.setState({message: text}, () => {this.props.onHandleInput(this.props.num, this.state.word, this.state.message)})
        }
    }

    render () {
        
        return e(
            "div",
            {className: "addwords card"},
            [
                e(Addword, {word: this.props.word, onHandleInput: this.handleInput}),
                e(Addmessage, {message: this.props.message, onHandleInput: this.handleInput})
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
          this.setState({ value: e.target.value }, () => {this.props.onHandleInput("word", this.state.value)})
        }
    }

    render() {
        return e(
            "div",
            {className: "addword_box large"},
            e(
                "input",
                {type: "text", 
                className: "addword_cursor large", 
                placeholder: "Enter word here", 
                value: this.state.value, 
                onChange: this.handleKeyPress}
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
        this.setState({ value: e.target.value }, () => {this.props.onHandleInput("message", this.state.value)});
    }

    render() {
        return e(
            "div",
            {className: "addword_box small"},
            e(
                "input",
                {type: "text", 
                className: "addword_cursor small", 
                placeholder: "Enter message to display on completion here", 
                value: this.state.value, 
                onChange: this.handleKeyPress}
            )
        )
    }

}

class Addanother extends React.Component {

    constructor(props){
        super(props)
        this.state = {}
        this.handleClick = this.handleClick.bind(this)
    }

    handleClick = () => {
        var max = 0

        for(let i in this.props.wordlist){
           if(i > max){
               max = i
           } 
        }
        this.props.onHandleClick(parseInt(max) + 1, "", "")

    }

    render() {
        return e(
            "div",
            {className: "addanother card", onClick: this.handleClick},
            " + Click to add another"
        )
    }
}

class Submit extends React.Component {

    constructor(props){
        super(props)
        this.state = {}
        this.handleClick = this.handleClick.bind(this)
    }

    handleClick = () => {
        this.props.onHandleClick()
    }

    render() {
        return e(
            "div",
            {className: "submit card", onClick: this.handleClick},
            "Click to submit words and messages"
        )
    }
}

const addwords_container = document.querySelector('#addwords_container');
ReactDOM.render(e(Addwords_container, {}), addwords_container);