'use string';

var e = React.createElement;

class Addwords extends React.Component{
    
    constructor(props){
        super(props)
        this.state = {}


    }


    render () {
        
        return e(
            "div",
            {className: "addwords"},
            [
                e(Addword, {}),
                e(Addmessage, {})
            ]
        )
    }
}



class Addword extends React.Component{

    constructor(props){
        super(props)
        this.state = {
            value: ""
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
                {type: "text", className: "addword_cursor large", placeholder: "Word", value: this.state.value, onChange: this.handleKeyPress}
            )
        )
    }
}

class Addmessage extends React.Component{

    constructor(props){
        super(props)
        this.state = {
            value: ""
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
                {type: "text", className: "addword_cursor small", placeholder: "Message to display on completion", value: this.state.value, onChange: this.handleKeyPress}
            )
        )
    }
}

const addwords = document.querySelector('#addwords');
ReactDOM.render(e(Addwords, {}), addwords);