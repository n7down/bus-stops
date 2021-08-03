import './App.css';
import React from 'react'
import moment from 'moment'

class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            currentTime: "12:03AM",
            stop1Route1: null,
            stop1Route2: null,
            stop1Route3: null,
            stop1NextRoute1: null,
            stop1NextRoute2: null,
            stop1NextRoute3: null,

            stop2Route1: null,
            stop2Route2: null,
            stop2Route3: null,
            stop2NextRoute1: null,
            stop2NextRoute2: null,
            stop2NextRoute3: null,
        };
    }

    componentDidMount() {
        const requestOptions = {
            method: 'POST',
            body: JSON.stringify({ stop: 1, time: "12:03AM" })
        };
        fetch('http://localhost:8080/api/v1/bus/routes', requestOptions)
            .then(response => response.json())
            .then(data => this.setState({ stop1Route1: data.route1, stop1Route2: data.route2, stop1Route3: data.route3, stop1NextRoute1: data.nextRoute1, stop1NextRoute2: data.nextRoute2, stop1NextRoute3: data.nextRoute3 }));

        const requestOptions2 = {
            method: 'POST',
            body: JSON.stringify({ stop: 2, time: "12:03AM" })
        };
        fetch('http://localhost:8080/api/v1/bus/routes', requestOptions2)
            .then(response => response.json())
            .then(data => this.setState({ stop2Route1: data.route1, stop2Route2: data.route2, stop2Route3: data.route3, stop2NextRoute1: data.nextRoute1, stop2NextRoute2: data.nextRoute2, stop2NextRoute3: data.nextRoute3 }));
    }

    render() {
        const { currentTime } = this.state;
        const { stop1Route1 } = this.state;
        const { stop1Route2 } = this.state;
        const { stop1Route3 } = this.state;
        const { stop1NextRoute1 } = this.state;
        const { stop1NextRoute2 } = this.state;
        const { stop1NextRoute3 } = this.state;
        const { stop2Route1 } = this.state;
        const { stop2Route2 } = this.state;
        const { stop2Route3 } = this.state;
        const { stop2NextRoute1 } = this.state;
        const { stop2NextRoute2 } = this.state;
        const { stop2NextRoute3 } = this.state;

        var startDate = moment(currentTime, "HH:mmA");
        var timeEnd = moment(stop1Route1, "HH:mmA");
        var diff = timeEnd.diff(startDate);
        var diffDuration = moment.duration(diff);
        var stop1Route1Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop1NextRoute1, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop1NextRoute1Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop1Route2, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop1Route2Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop1NextRoute2, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop1NextRoute2Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop1Route3, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop1Route3Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop1NextRoute3, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop1NextRoute3Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop2Route1, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop2Route1Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop2NextRoute1, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop2NextRoute1Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop2Route2, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop2Route2Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop2NextRoute2, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop2NextRoute2Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop2Route3, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop2Route3Minutes = parseInt(diffDuration.asMinutes())%60;

        timeEnd = moment(stop2NextRoute3, "HH:mmA");
        diff = timeEnd.diff(startDate);
        diffDuration = moment.duration(diff);
        var stop2NextRoute3Minutes = parseInt(diffDuration.asMinutes())%60;
        console.log(stop2NextRoute3Minutes)
        console.log(isNaN(stop2NextRoute3Minutes))

        return (
            <div className="card text-center m-3">
                <h4 className="card-header">{currentTime}</h4>
                <div className="card-body">
                    <p>Stop 1</p>
                    <p>Route 1 in {stop1Route1Minutes} mins and {stop1NextRoute1Minutes} mins</p>
                    <p>Route 2 in {stop1Route2Minutes} mins and {stop1NextRoute2Minutes} mins</p>
                    <p>Route 3 in {stop1Route3Minutes} mins and {stop1NextRoute3Minutes} mins</p>
                    <p>Stop 2</p>
                    <p>Route 1 in {stop2Route1Minutes} mins and {stop2NextRoute1Minutes} mins</p>
                    <p>Route 2 in {stop2Route2Minutes} mins and {stop2NextRoute2Minutes} mins</p>
                    <p>Route 3 in {stop2Route3Minutes} mins and {stop2NextRoute3Minutes} mins</p>
                </div>
            </div>
        );
    }
}

class Timer extends React.Component {
  constructor(props) {
    super(props);
    // console.log(String(this.props.secs))
    this.state = { time: {}, seconds: this.props.secs };
    this.timer = 0;
    this.startTimer = this.startTimer.bind(this);
    this.countDown = this.countDown.bind(this);
  }

  secondsToTime(secs){
    let hours = Math.floor(secs / (60 * 60));

    let divisor_for_minutes = secs % (60 * 60);
    let minutes = Math.floor(divisor_for_minutes / 60);

    let divisor_for_seconds = divisor_for_minutes % 60;
    let seconds = Math.ceil(divisor_for_seconds);

    let obj = {
      "h": hours,
      "m": minutes,
      "s": seconds
    };
    return obj;
  }

  componentDidMount() {
    let timeLeftVar = this.secondsToTime(this.state.seconds);
    this.setState({ time: timeLeftVar });
    this.startTimer()
  }

  startTimer() {
    if (this.timer == 0 && this.state.seconds > 0) {
      this.timer = setInterval(this.countDown, 1000);
    }
  }

  countDown() {
    // Remove one second, set state so a re-render happens.
    let seconds = this.state.seconds - 1;
    this.setState({
      time: this.secondsToTime(seconds),
      seconds: seconds,
    });
    
    // Check if we're at zero.
    if (seconds == 0) { 
      clearInterval(this.timer);
    }
  }

  render() {
    return(
      <div>
        m: {this.state.time.m} s: {this.state.time.s}
      </div>
    );
  }
}

export default App;
