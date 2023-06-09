import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Monitor from './components/Monitor/Monitor'
import reportWebVitals from './reportWebVitals';
import 'bootstrap/dist/css/bootstrap.css';

ReactDOM.render(
  <React.StrictMode>
    <Monitor URL={"http://34.125.205.134:8080"} />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
