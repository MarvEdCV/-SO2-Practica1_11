import React from 'react';
import './Monitor.css';
import Header from '../Header/Header';
import {
    BrowserRouter as Router,
    Switch,
    Route
} from "react-router-dom";
import Homepage from '../Homepage/Homepage';
import RAM from '../RAM/RAM';
import HOME from '../Home/Home';
export default class Monitor extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div>
                <Header />

                <Router>
                    <Switch>
                        <Route exact path="/">
                            <HOME URL={this.props.URL} />
                        </Route>
                        <Route exact path="/RAM_Monitor">
                            <RAM URL={this.props.URL} />
                        </Route>
                        <Route exact path="/CPU">
                            <Homepage URL={this.props.URL} />
                        </Route>
                    </Switch>
                </Router>
            </div>
        );
    }
}

