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
        let attempts = this.props.guessed_letters.length

        let time_to_next = "tomorrow";

        if(this.props.nextwordtime != null){
            var timenew = new Date(
                this.props.nextwordtime["year"], 
                this.props.nextwordtime["month"]-1, // Month is indexed at 0 
                this.props.nextwordtime["day"], 
                this.props.nextwordtime["hour"],
                this.props.nextwordtime["minute"],
                this.props.nextwordtime["second"]
            );

            var timenow = new Date();

            var timediff_mins = (timenew.getTime() - timenow.getTime()) / 60000 // Converting milliseconds to minutes
            var timediff_hours = Math.floor(timediff_mins / 60)
            var timediff_remainingmins = Math.floor(timediff_mins % 60)

            time_to_next = "in " + timediff_hours.toString() + " hours and " + timediff_remainingmins.toString() + " minutes!"
            }

        return e(
            "div",
            {className: finishbox_class, onClick: this.handleClick},
            [
                e("p", {}, this.props.message),
                e("p", {}, "You did it in " + attempts.toString() + " guesses! Next word available " + time_to_next)
            ]
        )
    }
}