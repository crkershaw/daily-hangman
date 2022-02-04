'use strict'

var e = React.createElement;

class Finish extends React.Component {
    constructor(props){
        super(props)

        this.handleClick = this.handleClick.bind(this);
    }

    handleClick(){
        this.props.onHandleClick()
    }

    render(){
        
        let finishbox_class = this.props.complete ? "finishbox" : "finishbox hidden"

        return e(
            "div",
            {className: finishbox_class, onClick: this.handleClick},
            "You did it! Click here to try another one!"
        )
    }
}