'use strict';

var e = React.createElement;

class KeyButton extends React.Component {

    constructor(props){
        super(props);
        this.state= {clicked: false};
        this.changeClicked = this.changeClicked.bind(this);
    };

    changeClicked(){
        this.setState({clicked: !this.state.clicked})

    };

    render() {
        let btn_class = this.state.clicked ? "key clicked" : "key unclicked";

        // React.createElement(component, props, ...children)
        return e(
            "button", 
            { onClick: this.changeClicked, className: btn_class }, 
            this.props.letter
        );
    }
}


class KeyRow extends React.Component {

    constructor(props){
        super(props);
    }

    render() {
        return e(
            "div",
            {className: "keyrow"},
            this.props.rowkeys
        )
    }
}

class KeyBoard extends React.Component {

    constructor(props){
        super(props);
    }

    render() {
        return e(
            "div",
            {className: "keyboard"},
            allrows
        )
    }
}


var row1_keys = ["Q","W","E","R","T","Y","U","I","O","P"]
var row2_keys = ["A","S","D","F","G","H","J","K","L"]
var row3_keys = ["Z","X","C","V","B","N","M"]

var row1 = []
var row2 = []
var row3 = []

for(let i = 0; i<row1_keys.length; i++){
    row1.push(e(KeyButton, {letter: row1_keys[i]}))
}

for(let i = 0; i<row2_keys.length; i++){
    row2.push(e(KeyButton, {letter: row2_keys[i]}))
}

for(let i = 0; i<row3_keys.length; i++){
    row3.push(e(KeyButton, {letter: row3_keys[i]}))
}

var allrows = [
    e(KeyRow, {rowkeys: row1}),
    e(KeyRow, {rowkeys: row2}),
    e(KeyRow, {rowkeys: row3})
]


const keyboard = document.querySelector('#keyboard');
ReactDOM.render(e(KeyBoard), keyboard);