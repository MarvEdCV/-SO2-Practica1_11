import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Monitor from './components/Monitor/Monitor'
import reportWebVitals from './reportWebVitals';
import 'bootstrap/dist/css/bootstrap.css';

ReactDOM.render(
  <React.StrictMode>
    <Monitor URL={"https://918c-2803-d100-98a8-341e-ac5f-cfc6-c56b-a689.ngrok.io"} />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
