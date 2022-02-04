'use strict';

var e = React.createElement;

class KeyButton extends React.Component {

    constructor(props){
        super(props);
        this.state= {clicked: false};
        this.handleClick = this.handleClick.bind(this);
        
    };

    handleClick(){
        this.setState({clicked: !this.state.clicked})
        this.props.onHandleClick(this.props.letter) // Sends action up to KeyRow
    };

    render() {
        let btn_class = this.state.clicked ? "key clicked" : "key unclicked";

        // React.createElement(component, props, ...children)
        return e(
            "button", 
            { onClick: this.handleClick, className: btn_class }, 
            this.props.letter
        );
    }
}


class KeyRow extends React.Component {

    constructor(props){
        super(props);

        this.handleClick = this.handleClick.bind(this)
    }

    handleClick = (letter) => {
        this.props.onHandleClick(letter) // Sends action up pto Keyboard
    }


    render() {
        var rowkeys = []

        for(let i = 0; i<this.props.keys.length; i++){
            rowkeys.push(e(KeyButton, {letter: this.props.keys[i], onHandleClick: this.handleClick}))
        }

        return e(
            "div",
            {className: "keyrow"},
            rowkeys
        )
    }
}

class KeyBoard extends React.Component {

    constructor(props){
        super(props);
    }

    handleClick = (letter) => {
        this.props.onHandleClick(letter) // Sends action up to hangman
    }

    render() {

        var row1_keys = ["Q","W","E","R","T","Y","U","I","O","P"]
        var row2_keys = ["A","S","D","F","G","H","J","K","L"]
        var row3_keys = ["Z","X","C","V","B","N","M"]        

        return e(
            "div",
            {className: "keyboard"},
            [
                e(KeyRow, {keys: row1_keys, onHandleClick: this.handleClick}),
                e(KeyRow, {keys: row2_keys, onHandleClick: this.handleClick}),
                e(KeyRow, {keys: row3_keys, onHandleClick: this.handleClick})
            ]
        )
    }
}
