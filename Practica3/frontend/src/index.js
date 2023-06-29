import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Monitor from './components/Monitor/Monitor'
import reportWebVitals from './reportWebVitals';
import 'bootstrap/dist/css/bootstrap.css';

ReactDOM.render(
  <React.StrictMode>
    <Monitor URL={"https://c11f-45-191-245-250.ngrok.io"} />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
